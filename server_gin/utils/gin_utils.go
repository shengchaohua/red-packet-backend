package serverutils

import (
	"github.com/gin-gonic/gin"

	"github.com/shengchaohua/red-packet-backend/internal/constants"
)

func Response(errcode int, errmsg string, response interface{}) gin.H {
	return gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
		"data":    response,
	}
}

func BadRequest() gin.H {
	return Response(int(constants.Errcode_WrongParam), "Bad request", struct{}{})
}

func ServerError() gin.H {
	return Response(int(constants.Errcode_Server), "Server error", struct{}{})
}

func WrongParam(errmsg string) gin.H {
	return Response(int(constants.Errcode_WrongParam), errmsg, struct{}{})
}

func OkResponse(response interface{}) gin.H {
	return Response(int(constants.Errcode_Ok), "", response)
}
