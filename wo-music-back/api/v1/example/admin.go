package example

import (
	"github.com/gin-gonic/gin"
)

type AdminApi struct {
}

func (e *AdminApi) CreateAdmin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
