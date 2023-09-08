package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(ctx *gin.Context) {

	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
func UpdateAddress(ctx *gin.Context) {

	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Update(ctx.Request.Context(), ctx.Param("id"), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
func ListAddress(ctx *gin.Context) {
	_, _ = util.ParseToken(ctx.GetHeader("Authorization"))
	addressService := service.AddressService{}
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}

func DeleteAddress(ctx *gin.Context) {

	addressService := service.AddressService{}
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Delete(ctx.Request.Context(), ctx.Param("id"), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}

func ShowAddress(ctx *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&addressService); err == nil {
		res := addressService.Show(ctx.Request.Context(), ctx.Param("id"), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
