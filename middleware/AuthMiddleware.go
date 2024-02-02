package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/utils"
)

var SECRET_KEY = os.Getenv("CLERK_SECRET")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := clerk.NewClient(SECRET_KEY)
		if err != nil {
			fmt.Println("Failed to create Clerk client: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// get session token from Authorization header
		sessionToken := utils.GetAuthorizationHeader(c)

		// Check if sessionToken is empty after trimming the prefix
		if sessionToken == "" {
			// If sessionToken is empty, try to get the _session cookie
			sessionToken = utils.GetClerkSessionCookie(c)
		}

		// verify the session
		sessClaims, err := client.VerifyToken(sessionToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// get the user and attach it to the context
		user, err := client.Users().Read(sessClaims.Claims.Subject)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
