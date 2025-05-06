package gin_utils

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	gwm_app "github.com/slhmy/go-webmods/app"
)

type ErrorResponse struct {
	BindErrors    []*gin.Error `json:"bindErrors,omitempty"`
	RenderErrors  []*gin.Error `json:"renderErrors,omitempty"`
	PublicErrors  []*gin.Error `json:"publicErrors,omitempty"`
	PrivateErrors []*gin.Error `json:"privateErrors,omitempty"`
}

func ErrorHandler(ginCtx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			slog.ErrorContext(ginCtx, string(stack),
				slog.Any("error", r),
				slog.Bool("panic", true),
			)
			_ = ginCtx.Error(fmt.Errorf("panic: %v", r))
		}
		if len(ginCtx.Errors) == 0 {
			return
		}

		resp := ErrorResponse{
			BindErrors:    ginCtx.Errors.ByType(gin.ErrorTypeBind),
			RenderErrors:  ginCtx.Errors.ByType(gin.ErrorTypeRender),
			PublicErrors:  ginCtx.Errors.ByType(gin.ErrorTypePublic),
			PrivateErrors: ginCtx.Errors.ByType(gin.ErrorTypePrivate),
		}

		httpStatus := http.StatusInternalServerError
		if len(resp.BindErrors) > 0 {
			httpStatus = http.StatusBadRequest
		}
		lastPublicErr := ginCtx.Errors.ByType(gin.ErrorTypePublic).Last()
		if lastPublicErr != nil {
			if status, ok := lastPublicErr.Meta.(map[string]any)[gwm_app.ErrorMetaFieldHttpStatus].(int); ok {
				httpStatus = status
			}
		}

		ginCtx.Negotiate(httpStatus, gin.Negotiate{Offered: []string{"application/json"}, Data: resp})
		ginCtx.Abort()
	}()
	ginCtx.Next()
}
