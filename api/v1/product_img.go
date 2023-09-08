package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProductImg(ctx *gin.Context) {
	listCarousel := service.ListProductImg{}
	if err := ctx.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("err: ", err)
	}
}
