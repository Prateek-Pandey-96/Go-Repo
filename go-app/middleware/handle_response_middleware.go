package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prateek96/app/schema"
)

func HandleRequestResponse[T any](reqContext *gin.Context, respChan <-chan T, createEmptyResponse func() T) (int, T) {
	ctx := reqContext.Request.Context()

	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return http.StatusRequestTimeout, createEmptyResponse()
			}
			return http.StatusInternalServerError, createEmptyResponse()
		case response := <-respChan:
			return http.StatusOK, response
		}
	}
}

func HandleGetRequestResponse(reqContext *gin.Context, respChan <-chan schema.GetRequestResponse) (int, schema.GetRequestResponse) {
	return HandleRequestResponse(reqContext, respChan, func() schema.GetRequestResponse {
		return schema.GetRequestResponse{UserExists: nil}
	})
}

func HandlePostRequestResponse(reqContext *gin.Context, respChan <-chan schema.PostRequestResponse) (int, schema.PostRequestResponse) {
	return HandleRequestResponse(reqContext, respChan, func() schema.PostRequestResponse {
		return schema.PostRequestResponse{UserExists: nil}
	})
}
