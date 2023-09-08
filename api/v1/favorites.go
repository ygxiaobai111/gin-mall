package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavorites(ctx *gin.Context) {

	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	CreateProductService := service.FavoriteService{}
	if err := ctx.ShouldBind(&CreateProductService); err == nil {
		res := CreateProductService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
func ListFavorites(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	listFavoritesService := service.FavoriteService{}
	if err := ctx.ShouldBind(&listFavoritesService); err == nil {
		res := listFavoritesService.Show(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}

func DeleteFavorites(ctx *gin.Context) {

	deleteFavoritesService := service.FavoriteService{}
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&deleteFavoritesService); err == nil {
		res := deleteFavoritesService.Delete(ctx.Request.Context(), ctx.Param("id"), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
