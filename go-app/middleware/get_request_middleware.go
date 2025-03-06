package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/common"
	"github.com/prateek96/app/schema"
)

func GetMiddleware(d *common.Dependency) gin.HandlerFunc {
	timeout := time.Duration(1 * time.Second)
	return func(c *gin.Context) {
		respChan := make(chan schema.GetRequestResponse)

		queryParams := schema.GetRequestQueryParams{}
		if err := c.BindQuery(&queryParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request, missing required query params!",
			})
			return
		}

		getRequestContextKeeper := schema.GetRequestContextKeeper{
			QueryParams: queryParams,
			RespChan:    respChan,
		}

		ctx, cancel := context.WithTimeout(context.WithValue(
			c.Request.Context(), schema.GetRequestCtx, getRequestContextKeeper),
			timeout,
		)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		finalStatus, finalResponse := HandleGetRequestResponse(c, respChan)
		c.AbortWithStatusJSON(finalStatus, finalResponse)
	}
}
