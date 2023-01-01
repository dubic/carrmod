package api

import (
	"carrmod/backend/services"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, response any, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func Handle(c *gin.Context, errs []string) bool {
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msgs": errs,
		})
		return false
	}
	return true
}

// Checks if a value is set for header 'Authorization'
func Authentication(c *gin.Context) {
	var authorization = c.GetHeader("Authorization")
	if authorization == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("not authenticated"))
		return
	}

	authSplit := strings.Split(authorization, " ")
	token := strings.Trim(authSplit[1], " ")

	if email, err := services.ParseJwt(token); err != nil {
		log.Printf("could not parse jwt [%s]: %s", token, err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid authentication token"))
		return
	} else {
		c.Set("email", email)
	}
	c.Next()
}
