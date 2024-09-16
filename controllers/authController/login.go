package authcontroller

import (
	"app/middleware/serializers"
	"app/middleware/validators"
	authservice "app/services/authService"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
