package schema

type PostRequestQueryParams struct {
	ReqId  *string `form:"request_id" binding:"required"`
	UserId *string `form:"user_id" binding:"required"`
}

type PostRequestBody struct {
	Name *string `json:"name" binding:"required"`
	City *string `json:"city" binding:"required"`
}

type PostRequestResponse struct {
	UserExists *bool `json:"user_exists_post_api" binding:"required"`
}
