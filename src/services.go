package src

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func helpServices(c context.Context, ctx *app.RequestContext) {

	ctx.JSON(http.StatusOK, utils.H{
		"code":     200,
		"api docs": "/swagger/index.html",
	})
}

// @Summary 存入视频数据
// @Description 存入视频数据
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /video/:name [post]
func videoServices(c context.Context, ctx *app.RequestContext) {

}

// @Summary 存入图片数据
// @Description 存入图片数据
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /images/:name [post]
func imagesServices(c context.Context, ctx *app.RequestContext) {

}
