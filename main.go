package main

import (
	"CdnServer/config"
	_ "CdnServer/docs"
	"CdnServer/middleware"
	"CdnServer/src"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
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
	f, _ := os.OpenFile(config.LogDir+"cdnserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	hlog.SetLogger(GetLogger(f))
	h := GetServer()
	h.Spin()
}

func GetServer() *server.Hertz {
	h := server.New(server.WithMaxRequestBodySize(config.RequestBodyMaxSize),
		server.WithHostPorts(config.CdnServerHost))
	h.Use(AccessLog(), middleware.GlobalErrorHandler, recovery.Recovery())
	src.InitRoute(h)
	return h
}

func GetLogger(output io.Writer) *hertzZerolog.Logger {
	return hertzZerolog.New(
		hertzZerolog.WithOutput(output),         // allows to specify output
		hertzZerolog.WithLevel(config.LogLevel), // option with log level
		hertzZerolog.WithTimestamp(),            // option with timestamp
		hertzZerolog.WithCaller(),               // option with caller
	)
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
