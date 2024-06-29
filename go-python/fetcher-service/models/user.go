package models

type UserReq struct {
	Username string `json:"username" binding:"required"`
	Token    string `json:"token" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Depth    int    `json:"depth" binding:"required"`
}
