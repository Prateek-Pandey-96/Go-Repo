package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/common"
	"github.com/prateek96/app/schema"
)

func GetController(d *common.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		contextData, _ := ctx.Value(schema.GetRequestCtx).(schema.GetRequestContextKeeper)

		respChan := contextData.RespChan
		val := true
		go func() {
			response := schema.GetRequestResponse{UserExists: &val}
			respChan <- response
		}()
	}
}
