package gwm_app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ErrorMetaFieldHttpStatus = "httpStatus"

type Error struct {
	Err  error
	Meta map[string]any
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) GinError() *gin.Error {
	return &gin.Error{
		Err:  e.Err,
		Type: gin.ErrorTypePublic,
		Meta: e.Meta,
	}
}

func NewError(err error, httpStatus int, meta map[string]any) Error {
	if meta == nil {
		meta = make(map[string]any)
	}
	meta[ErrorMetaFieldHttpStatus] = httpStatus
	return Error{
		Err:  err,
		Meta: meta,
	}
}

var ErrUnauthorized = NewError(errors.New("unauthorized"), http.StatusUnauthorized, nil)

func NewNotFoundError(target, by string) Error {
	meta := map[string]any{}
	if target != "" {
		meta["target"] = target
	}
	if by != "" {
		meta["by"] = by
	}
	return NewError(errors.New("not found"), http.StatusNotFound, meta)
}

func NewForbiddenError(reason string) Error {
	meta := map[string]any{
		"reason": reason,
	}
	return NewError(errors.New("forbidden"), http.StatusForbidden, meta)
}

func NewBadRequestError(reason string) Error {
	meta := map[string]any{
		"reason": reason,
	}
	return NewError(errors.New("bad request"), http.StatusBadRequest, meta)
}
