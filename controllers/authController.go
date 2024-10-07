package controllers

import (
	"app/middleware/serializers"
	"app/middleware/validators"
	authservice "app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var req serializers.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "heloooo "})
		return
	}

	if err := validators.ValidationCheckSignup(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authservice.SignupService(c, req)

}

func Login(c *gin.Context) {
	var loginInput serializers.Logininput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidationChecklogin(loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authservice.LoginService(c, loginInput)
}
