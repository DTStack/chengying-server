package apibase

import (
	"github.com/kataras/iris/context"
	"net/http"
)

const (
	UNKNOWN_ERR = iota + 100
	API_PARAM_ERR
	DB_MODEL_ERR
	ACCESS_DENIED_ERR
)

type ApiResult struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type EmptyResult struct{}

func Feedback(ctx context.Context, result interface{}) {
	if err, ok := result.(error); ok {
		if IsApiParameterErrors(err) {
			errs, _ := err.(*ApiParameterErrors)
			data := map[string]string{}
			for pname, err := range errs.errors {
				data[pname] = err.Error()
			}
			ctx.JSON(&ApiResult{
				Code: API_PARAM_ERR,
				Msg:  "Invalid parameter(s)",
				Data: data,
			})
		} else if IsDBModelError(err) {
			e, _ := err.(*DBModelError)
			ctx.JSON(&ApiResult{
				Code: DB_MODEL_ERR,
				Msg:  "DB Model error",
				Data: e.err.Error(),
			})
		} else if IsAccessDeniedError(err) {
			e, _ := err.(*AccessDeniedError)
			ctx.JSON(&ApiResult{
				Code: ACCESS_DENIED_ERR,
				Msg:  "Access Denied",
				Data: e.Err.Error(),
			})
		} else {
			ctx.JSON(&ApiResult{
				Code: UNKNOWN_ERR,
				Msg:  err.Error(),
			})
		}
	} else if _, ok := result.(EmptyResult); ok {
		ctx.StatusCode(http.StatusOK)
		ctx.Done()
	} else {
		ctx.JSON(&ApiResult{
			Code: 0,
			Msg:  "ok",
			Data: result,
		})
	}
	ctx.StatusCode(http.StatusOK)
	ctx.Done()
}
