package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckEndpoints() bool {
	client := http.Client{
		Timeout: 4 * time.Second,
	}

	endpoints := []string{
		"http://localhost:8080/auth/login",
		"http://localhost:8080/auth/signup",
		"http://localhost:8080/users",
		"http://localhost:8080/users/user",
		"http://localhost:8080/users/user/following",
		"http://localhost:8080/users/user/followers",
		"http://localhost:8080/users/user/update",
		"http://localhost:8080/users/user/delete",
		"http://localhost:8080/users/user/follow",
		"http://localhost:8080/users/user/unfollow",
		"http://localhost:8080/users/user/updatepassword",
	}

	for _, endpoint := range endpoints {
		_, err := client.Get(endpoint)
		if err != nil {
			return false
		}
	}
	return true
}

func Healthcheck(c *gin.Context) {
	if CheckEndpoints() {
		c.JSON(http.StatusOK, gin.H{"status": "Healthy"})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Unhealthy"})
	}
}
