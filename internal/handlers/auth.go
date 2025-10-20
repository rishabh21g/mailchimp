package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabh21g/mailchimp/internal/database"
	"github.com/rishabh21g/mailchimp/internal/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterHandler(c *gin.Context) {
	var user types.User            // create a user
	err := c.ShouldBindJSON(&user) // binding with request the req must be type of user
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": "error",
			"message": "Invalid inputs",
			"error":   err.Error(),
		})
		return

	}
	// Now we check user exist or not in our database
	var existingUser types.User
	err = database.DB.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": "failed",
			"message": "User already exists",
		})
		return

	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "failed",
			"message": "Internal Server Error",
		})
		return

	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "failed",
			"message": "Error while hashing password",
		})
		return
	}
	user.Password = string(hashedPassword)
	// create user in database
	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "failed",
			"message": "Error while creating user",
		})
		return

	}
	user.Password = "" 
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created",
		"user":    user.Email,
	})

}
