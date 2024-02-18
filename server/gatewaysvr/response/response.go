package response

import (
	"gatewaysvr/log"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	successCode = 0
	errCode     = 1
)

type response struct {
	StatusCode int32
	StatusMsg  string
}

// Response 这里的v一般是是pb文件中的Resp结构体
func Response(ctx *gin.Context, httpStatus int, v interface{}) {
	ctx.JSON(httpStatus, v)
}

func Success(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		Response(ctx, 200, response{successCode, msg})
	} else {
		//setResp(ctx, successCode, msg, v)
		Response(ctx, 200, v)
	}
}
func Fail(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		Response(ctx, 200, response{errCode, msg})
	} else {
		setResp(ctx, errCode, msg, v)
		Response(ctx, 200, v)
		ctx.Abort()
	}
}
func setResp(ctx *gin.Context, StatusCode int64, StatusMsg string, v interface{}) {
	getValue := reflect.ValueOf(v)
	field := getValue.Elem().FieldByName("StatusMsg")
	if field.CanSet() {
		field.SetString(StatusMsg)
	} else {
		log.Debug("can't set StatusMsg")
	}
	fieldCode := getValue.Elem().FieldByName("StatusCode")
	if fieldCode.CanSet() {
		fieldCode.SetInt(StatusCode)
	} else {
		log.Debug("cant set StatusMsg")
	}
}
