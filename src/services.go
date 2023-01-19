package src

import (
	"CdnServer/config"
	"context"
	"github.com/astaxie/beego/validation"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
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
	valid := &validation.Validation{}
	mp4, err := ctx.FormFile("data")
	if err != nil || mp4 == nil {
		hlog.Error("failed to get data")
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	destFileName := ctx.Param("name")

	valid.Required(destFileName, "video file name").Message("video file name cannot be null")
	hlog.Info("receive mp4 file:", destFileName)
	mp4File, err := mp4.Open()
	// TODO(lee) : 错误处理得换一个好一点儿的方式，这样写太傻了
	if err != nil {
		hlog.Error("failed to open mp4File")
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	mp4DestFile, err := os.OpenFile(config.ResourceDir+destFileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		hlog.Error("failed to create new mp4File:", mp4DestFile.Name())
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	_, err = io.Copy(mp4DestFile, mp4File)
	if err != nil {
		hlog.Error("failed to Copy to mp4File:", mp4DestFile.Name())
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	if mp4File.Close() != nil || mp4DestFile.Close() != nil {
		hlog.Error("failed to close mp4File:", mp4DestFile.Name())
		_ = ctx.Error(errors.WithStack(errors.New("failed to close mp4File")))
		return
	}
	ctx.JSON(http.StatusOK, utils.H{
		"code":     "200",
		"filename": destFileName,
	})
}

// @Summary 存入图片数据
// @Description 存入图片数据
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /images/:name [post]
func imagesServices(c context.Context, ctx *app.RequestContext) {
	valid := &validation.Validation{}
	image, err := ctx.FormFile("data")
	if err != nil || image == nil {
		hlog.Error("failed to get data")
		_ = ctx.Error(errors.WithStack(err))
		return
	}

	destFileName := ctx.Param("name")

	valid.Required(destFileName, "image file name").Message("image file name cannot be null")
	hlog.Info("receive image file:", destFileName)
	imageFile, err := image.Open()
	// TODO(lee) : 错误处理得换一个好一点儿的方式，这样写太傻了
	if err != nil {
		hlog.Error("failed to open imageFile")
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	imgDestFile, err := os.OpenFile(config.ResourceDir+destFileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		hlog.Error("failed to create new imageFile:", imgDestFile.Name())
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	_, err = io.Copy(imgDestFile, imageFile)
	if err != nil {
		hlog.Error("failed to copy to imageFile:", imgDestFile.Name())
		_ = ctx.Error(errors.WithStack(err))
		return
	}
	if imageFile.Close() != nil || imgDestFile.Close() != nil {
		hlog.Error("failed to close imageFile:", imgDestFile.Name())
		_ = ctx.Error(errors.WithStack(errors.New("failed to close imageFile")))
		return
	}
	ctx.JSON(http.StatusOK, utils.H{
		"code":     "200",
		"filename": destFileName,
	})
}
