package schema

type ctx_key string

var GetRequestCtx ctx_key = "GetRequestCtx"
var PostRequestCtx ctx_key = "PostRequestCtx"

type GetRequestContextKeeper struct {
	QueryParams GetRequestQueryParams
	RespChan    chan GetRequestResponse
}

type PostRequestContextKeeper struct {
	QueryParams PostRequestQueryParams
	PostBody    PostRequestBody
	RespChan    chan PostRequestResponse
}
