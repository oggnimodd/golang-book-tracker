package utils

import (
	"errors"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
)

func GetClerkSessionCookie(c *gin.Context) string {
	// get session cookie
	sessionCookie, err := c.Cookie("__session")

	// return an empty string in case of an error or if the cookie doesn't exist
	if err != nil || sessionCookie == "" {
		return ""
	}

	return sessionCookie
}

func GetAuthorizationHeader(c *gin.Context) string {
	// get authorization header
	authorizationHeader := c.GetHeader("Authorization")
	authorizationHeader = strings.TrimPrefix(authorizationHeader, "Bearer ")
	// return an empty string in case of an error or if the header doesn't exist
	if authorizationHeader == "" {
		return ""
	}
	return authorizationHeader
}

func GetUserFromContext(c *gin.Context) (*clerk.User, error) {
	val, exists := c.Get("user")
	if !exists {
		return nil, errors.New("invalid user type in context")
	}

	user, ok := val.(*clerk.User)
	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return user, nil
}
