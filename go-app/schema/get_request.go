package schema

type GetRequestQueryParams struct {
	ReqId  *string `form:"request_id" binding:"required"`
	UserId *string `form:"user_id" binding:"required"`
}

type GetRequestResponse struct {
	UserExists *bool `json:"user_exists_get_api" binding:"required"`
}
