package serverutils

import (
	"github.com/gin-gonic/gin"

	"github.com/shengchaohua/red-packet-backend/base/constants"
)

func wrapResponse(errcode constants.Errcode, errmsg string, response interface{}) gin.H {
	return gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
		"data":    response,
	}
}

func BadRequest() gin.H {
	return wrapResponse(constants.Errcode_Param, "Bad request", struct{}{})
}

func ServerError() gin.H {
	return wrapResponse(constants.Errcode_Server, "Server error", struct{}{})
}

func WrongParam(errmsg string) gin.H {
	return wrapResponse(constants.Errcode_Param, errmsg, struct{}{})
}

func WrapResponse(response interface{}) gin.H {
	return wrapResponse(constants.Errcode_Ok, "", response)
}
