package config

import (
	"carrmod/backend/api"
	"carrmod/backend/domain/models"
	"carrmod/backend/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func Logging() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// connect to db
func Database() {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Panicln("Error reading database uri. plase specify DB_URI")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}
	dbClient = client
	log.Println("Successfully connected to mongo database")
}

func disconnectDatabase() {
	dbClient.Disconnect(context.TODO())
	log.Println("Disconnected from mongo DB")
}

// register controllers here
func Web() {
	router := gin.Default()
	//services
	userRepo := models.NewUserRepo(dbClient.Database("carrmod").Collection("users"))
	userService := services.NewUserService(userRepo)
	//controllers
	api.RegisterUserController(router, userService)

	//start server
	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	term := make(chan os.Signal)
	go func() {
		signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
		<-term
		teardown()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	log.Println("starting server...")
	var err = server.ListenAndServe()
	if err != nil {
		log.Println("error occurred", err)
	}
}

func teardown() {
	log.Println("closing connections")
	disconnectDatabase()
}
