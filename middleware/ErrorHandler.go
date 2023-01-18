package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func GlobalErrorHandler(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)

	if len(c.Errors) == 0 {
		return
	}
	hertzErrors := c.Errors[0]
	// 获取errors包装的err
	err := hertzErrors.Unwrap()
	hlog.CtxErrorf(ctx, "%+v", err)
	err = errors.Unwrap(err)
	c.JSON(http.StatusBadRequest, utils.H{
		"code":    "400",
		"message": err.Error(),
	})
}
