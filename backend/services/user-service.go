package services

import (
	"carrmod/backend/domain/dto"
	"carrmod/backend/domain/models"
	"fmt"
	"log"
	"time"
)

type UserService struct {
	userRepo       *models.UserRepo
	sessionManager *SessionManager
}

func NewUserService(userRepo *models.UserRepo, sMan *SessionManager) *UserService {
	return &UserService{userRepo, sMan}
}

// Creates an account. emails are unique.
// If an account with the same email exists, a soft error is returned.
func (svc *UserService) CreateAccount(userCreationRequest dto.UserCreationRequest) (dto.UserCreationResponse, error) {
	hashedPassword := HashPassword(userCreationRequest.Password, userCreationRequest.Email)
	//check user exists
	exists, err := svc.userRepo.UserExists(userCreationRequest.Email)
	if err != nil {
		return dto.UserCreationResponse{}, err
	}
	if exists {
		log.Println("User with email exists: ", userCreationRequest.Email)
		return dto.UserCreationResponse{Created: false, Msg: fmt.Sprintf("User with email exists - %s", userCreationRequest.Email)}, err
	}
	//create user
	user := models.NewUser(userCreationRequest, hashedPassword)
	err = svc.userRepo.SaveNewUser(user)
	log.Println("created and saved new user: ", user.Print())
	log.Println("Send mail")
	return dto.UserCreationResponse{VerificationMailSent: false, Created: true, Email: user.Email}, err
}

func (svc *UserService) UserLoginResponse(user models.User) dto.LoginResponse {
	return dto.LoginResponse{Verified: user.EmailVerified, Authenticated: true, Account: dto.Account{
		Name: user.Name, Email: user.Email, CreatedAt: user.Base.CreatedAt, UpdatedAt: user.Base.UpdatedAt,
	}}
}

func (svc *UserService) UserToAccount(user models.User) dto.Account {
	return dto.Account{
		Name: user.Name, Email: user.Email, CreatedAt: user.Base.CreatedAt, UpdatedAt: user.Base.UpdatedAt,
	}
}

// Login with email and password of previously created account
func (svc *UserService) Login(loginRequest dto.LoginRequest) dto.LoginResponse {
	loginResponse := svc.Authenticate(loginRequest)
	if !loginResponse.Authenticated {
		return loginResponse
	}

	//if email is not verified return code
	if !loginResponse.Verified {
		log.Println("account email not verified - ", loginResponse.Account.Email)
		loginResponse.Msg = "Account not verified"
		return loginResponse
	}

	loginResponse = svc.CreateUserSession(loginResponse)
	return loginResponse
}

func (svc *UserService) Authenticate(loginRequest dto.LoginRequest) dto.LoginResponse {
	// hash password and authenticate
	hashedPassword := HashPassword(loginRequest.Password, loginRequest.Email)
	user, err := svc.userRepo.FindUserByEmailAndPassword(loginRequest.Email, hashedPassword)
	if err != nil {
		log.Printf("login error - %s", err)
		return dto.LoginResponse{Authenticated: false, Msg: "Wrong credentials"}
	}
	loginResponse := svc.UserLoginResponse(user)
	log.Println("successful authentication - ", loginResponse.Account.Email)
	return loginResponse
}

// create a session and token using an autheticated login response
func (svc *UserService) CreateUserSession(loginResponse dto.LoginResponse) dto.LoginResponse {
	token, jwtErr := GenerateJwt(loginResponse.Account.Email)
	if jwtErr != nil {
		log.Printf("JWT creating error for %s: %s", loginResponse.Account.Email, jwtErr)
		return dto.LoginResponse{Msg: "Error logging in"}
	}
	loginResponse.Token = token
	//Create mobile session in db
	session := models.Session{User: loginResponse.Account.Email, AccessedAt: time.Now(), ExpiresAt: time.Now().Add(time.Hour * 2160)}
	if err := svc.sessionManager.NewSession(session); err != nil {
		log.Printf("JWT creating error for %s: %s", loginResponse.Account.Email, jwtErr)
	}
	// return token
	loginResponse.LoginTime = time.Now()
	loginResponse.Msg = "Successful"
	return loginResponse
}

func (svc *UserService) Logout(email string) {
	sessions := svc.sessionManager.RemoveSession(email)
	log.Printf("logged [%s] out of %d sessions", email, sessions)
}

func (svc *UserService) LoadProfile(email string) (dto.Account, error) {
	user, err := svc.userRepo.FindUserByEmail(email)
	if err != nil {
		return dto.Account{}, fmt.Errorf("sorry, error occurred loading profile %s", email)
	}
	return svc.UserToAccount(user), nil
}
