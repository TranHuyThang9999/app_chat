package entities

import (
	"mime/multipart"
	"websocket_p4/core/adapter/domain"
)

type User struct {
	UserName string                `form:"user_name"`
	Age      int                   `form:"age"`
	Address  string                `form:"address"`
	Email    string                `form:"email"`
	File     *multipart.FileHeader `form:"file"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type UserRespAdd struct {
	Result    Result `json:"result"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}
type UserRespFindByUserName struct {
	Result    Result        `json:"result"`
	CreatedAt int64         `json:"created_at"`
	User      *domain.Users `json:"user"`
}
