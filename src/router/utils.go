package router

import (
	"BitginHomework/model"
	"context"

	"github.com/gin-gonic/gin"
)

// set internal faild message to gin context
func internalFaild(c *gin.Context, errStr string) {
	c.JSON(500, gin.H{
		"status":  500,
		"message": errStr,
	})
}

// get context.Context from gin context
func getContext(c *gin.Context) *context.Context {
	ctxInterface, ctxExist := c.Get("context")
	if ctxExist {
		return ctxInterface.(*context.Context)
	}

	ctx := context.Background()
	return &ctx
}

// get user from gin context
func getUser(c *gin.Context) *model.User {
	userInterface, userExist := c.Get("user")
	if userExist {
		return userInterface.(*model.User)
	}

	return nil
}
