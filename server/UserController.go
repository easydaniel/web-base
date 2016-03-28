package main

import (
	"github.com/gin-gonic/gin"
)

// UserControllerGET definition
func UserControllerGET(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"Name": c.Param("name"),
	})
}
