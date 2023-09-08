package v1

import (
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(ctx *gin.Context) {
	var Category service.CategoryService

	if err := ctx.ShouldBind(&Category); err == nil {
		res := Category.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)

	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
