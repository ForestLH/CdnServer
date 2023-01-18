package main

import (
	"CdnServer/config"
	_ "CdnServer/docs"
	"CdnServer/middleware"
	"CdnServer/src"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Microseconds()
		hlog.CtxTracef(c, "status=%d cost=%d(us) method=%s full_path=%s client_ip=%s host=%s",
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host())
	}
}
func main() {

	h := GetServer()
	hlog.SetLevel(config.LogLevel)
	h.Spin()
	//h.POST("/douyin/publish/action/", func(c context.Context, ctx *app.RequestContext) {
	//	valid := &validation.Validation{}
	//	mp4, err := ctx.FormFile("data")
	//	if err != nil || mp4 == nil {
	//		hlog.Info("failed to get data")
	//		ctx.AbortWithStatus(http.StatusNotFound)
	//		return
	//	}
	//	destFileName := ctx.PostForm("title")
	//
	//	valid.Required(destFileName, "video file name").Message("video file name cannot be null")
	//	hlog.Info("title:", destFileName)
	//	mp4File, err := mp4.Open()
	//	mp4DestFile, _ := os.OpenFile("/tmp/video.mp4", os.O_CREATE|os.O_WRONLY, 0755)
	//	defer mp4DestFile.Close()
	//	defer mp4File.Close()
	//	io.Copy(mp4DestFile, mp4File)
	//	if err != nil {
	//
	//	}
	//	for valid.HasErrors() {
	//		errMap := make(map[string]string, 3)
	//		for _, err := range valid.Errors {
	//			errMap[err.Key] = err.Message
	//		}
	//		ctx.JSON(http.StatusOK, errMap)
	//	}
	//})
	//
	//h.Spin()
}

func GetServer() *server.Hertz {
	h := server.New(server.WithMaxRequestBodySize(config.RequestBodyMaxSize),
		server.WithHostPorts(config.CdnServerHost))
	h.Use(AccessLog(), middleware.GlobalErrorHandler, recovery.Recovery())
	src.InitRoute(h)
	return h
}
func setPref() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}
