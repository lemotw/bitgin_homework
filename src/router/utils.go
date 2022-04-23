package router

import (
	"context"

	"github.com/gin-gonic/gin"
)

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
