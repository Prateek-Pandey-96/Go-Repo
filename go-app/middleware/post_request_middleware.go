package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/common"
	"github.com/prateek96/app/schema"
)

func PostMiddleware(d *common.Dependency) gin.HandlerFunc {
	timeout := time.Duration(1 * time.Second)
	return func(c *gin.Context) {
		respChan := make(chan schema.PostRequestResponse)

		queryParams := schema.PostRequestQueryParams{}
		if err := c.BindQuery(&queryParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request, missing required query params!",
			})
			return
		}

		postBody := schema.PostRequestBody{}
		if err := c.BindJSON(&postBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request, malformed post body!",
			})
			return
		}

		postRequestContextKeeper := schema.PostRequestContextKeeper{
			QueryParams: queryParams,
			PostBody:    postBody,
			RespChan:    respChan,
		}

		ctx, cancel := context.WithTimeout(context.WithValue(
			c.Request.Context(), schema.PostRequestCtx, postRequestContextKeeper),
			timeout,
		)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		finalStatus, finalResponse := HandlePostRequestResponse(c, respChan)
		c.AbortWithStatusJSON(finalStatus, finalResponse)
	}
}
