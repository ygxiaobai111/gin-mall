package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建商品

func CreateProduct(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	CreateProductService := service.ProductService{}
	if err := ctx.ShouldBind(&CreateProductService); err == nil {
		res := CreateProductService.Create(ctx.Request.Context(), claim.ID, files)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}

func ListProduct(ctx *gin.Context) {

	listProductService := service.ProductService{}
	if err := ctx.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}

func SearchProduct(ctx *gin.Context) {
	searchProductService := service.ProductService{}
	if err := ctx.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}

func ShowProduct(ctx *gin.Context) {
	showProductService := service.ProductService{}
	if err := ctx.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
