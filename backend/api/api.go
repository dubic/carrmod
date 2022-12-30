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
