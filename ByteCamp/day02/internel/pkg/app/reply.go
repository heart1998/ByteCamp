package app

import (
	"github.com/gin-gonic/gin"
	"go_dance/day_2/1_post/internel/pkg/app/errcode"
	"net/http"
)

type Reply struct {
	ctx *gin.Context
}

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type list struct {
	List   interface{} `json:"list"`
	Length int         `json:"length"`
}

func NewReply(ctx *gin.Context) *Reply {
	return &Reply{ctx: ctx}
}

func (r *Reply) response(code errcode.Err, data interface{}) {
	r.ctx.JSON(http.StatusOK, response{
		Code: code.Code(),
		Data: data,
		Msg:  code.Msg(),
	})
}

func (r *Reply) ToResponseError(code errcode.Err) {
	r.response(code, nil)
}

func (r *Reply) ToResponseData(data interface{}) {
	r.response(errcode.StatusOK, data)
}

func (r *Reply) ToResponseList(data interface{}, length int) {
	r.response(errcode.StatusOK, list{
		List:   data,
		Length: length,
	})
}
