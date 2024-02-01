package controller

import (
	"net/http"
	"websocket_p4/common/log"
	"websocket_p4/core/entities"
	"websocket_p4/core/usecase"

	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	user *usecase.UseCaseUse
	// *baseController
}

func NewControllerEmployees(user *usecase.UseCaseUse) *ControllerUser {
	return &ControllerUser{
		user: user,
	}
}
func (o *ControllerUser) AddAcount(ctx *gin.Context) {

	var req entities.User

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil && err != http.ErrMissingFile && err != http.ErrNotMultipart {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Không thể tải ảnh lên.",
		})
		return
	}
	log.Infof("req : ", req.Email)
	req.File = file
	resp, err := o.user.AddAccount(ctx, &req)

	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (o *ControllerUser) GetUserByUserName(ctx *gin.Context) {
	user_name := ctx.Query("user_name")
	resp, err := o.user.FindByUserName(ctx, user_name)

	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
