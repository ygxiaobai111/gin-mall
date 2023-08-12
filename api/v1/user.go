package v1

import (
	"gin-mall/service"
	"gin-mall/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	if err := ctx.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func UploadAvatar(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(ctx.Request.Context(), claims.ID, file, fileSize)
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
