package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

func WithContext(c *gin.Context) {
	ctx := context.Background()
	c.Set("context", &ctx)

	c.Next()
}

func getContext(c *gin.Context) *context.Context {
	var ctx context.Context
	ctxInterface, ctxExist := c.Get("context")
	if ctxExist {
		ctx = *ctxInterface.(*context.Context)
	} else {
		ctx = context.Background()
	}
	return &ctx
}
