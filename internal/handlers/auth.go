package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func LoginHandler(c*gin.Context){
	 c.JSON(http.StatusOK , gin.H{
		"message" : "Login route working fine",
	})
}


func RegisterHandler(c *gin.Context){
	c.JSON(http.StatusOK , gin.H{
		"message":"Register route working fine",
	})
}