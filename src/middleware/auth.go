package middleware

import (
	"BitginHomework/service/auth"

	"github.com/gin-gonic/gin"
)

func WithUser(c *gin.Context) {
	ctx := getContext(c)
	tokenStr := c.GetHeader("token")

	// valid token and get user
	user, err := auth.ValidToken(*ctx, tokenStr)
	if err != nil {
		c.JSON(403, gin.H{
			"status":  403,
			"message": "auth not allow",
		})
		return
	}

	c.Set("user", user)
	c.Next()
}
