package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, response any, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusOK, response)
}
