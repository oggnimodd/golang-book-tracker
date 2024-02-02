package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
