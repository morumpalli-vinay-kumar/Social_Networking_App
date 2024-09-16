package authcontroller

import (
	"app/middleware/serializers"
	"app/middleware/validators"
	authservice "app/services/authService"
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
