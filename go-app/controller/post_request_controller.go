package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/common"
	"github.com/prateek96/app/schema"
)

func PostController(d *common.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		contextData, _ := ctx.Value(schema.PostRequestCtx).(schema.PostRequestContextKeeper)

		respChan := contextData.RespChan
		val := true
		go func() {
			response := schema.PostRequestResponse{UserExists: &val}
			respChan <- response
		}()
	}
}
