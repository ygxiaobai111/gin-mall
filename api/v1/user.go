package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

	}
}

func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	if err := ctx.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

	}
}

func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

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
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

	}
}

func SendEmail(ctx *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

	}
}

func ValidEmail(ctx *gin.Context) {
	var validEmail service.ValidEmailService

	if err := ctx.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

	}
}

func ShowMoney(ctx *gin.Context) {
	var showMoney service.ShowMoneyService

	if err := ctx.ShouldBind(&showMoney); err == nil {
		claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
		res := showMoney.Show(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)

	}
}
