package v1

import (
	"github.com/gin-gonic/gin"
	"go_dance/day_2/1_post/internel/global"
	"go_dance/day_2/1_post/internel/logic"
	"go_dance/day_2/1_post/internel/model/request"
	"go_dance/day_2/1_post/internel/pkg/app"
	"go_dance/day_2/1_post/internel/pkg/app/errcode"
)

type topic struct{}

func (api *topic) PublicTopic(ctx *gin.Context) {
	reply := app.NewReply(ctx)
	params := new(request.PublicTopic)
	if err := ctx.ShouldBindJSON(params); err != nil {
		global.Log.Println(err)
		reply.ToResponseError(errcode.ErrParamNotValid)
		return
	}
	if err := params.Check(); err != nil {
		reply.ToResponseError(err)
		return
	}
	data, err := logic.Group.Topic.PublicTopic(ctx, params)
	if err != nil {
		reply.ToResponseError(err)
		return
	}
	reply.ToResponseData(data)
}

func (api *topic) QueryTopic(ctx *gin.Context) {
	reply := app.NewReply(ctx)
	params := new(request.QueryTopic)
	if err := ctx.ShouldBindJSON(params); err != nil {
		global.Log.Println(err)
		reply.ToResponseError(errcode.ErrParamNotValid)
		return
	}
	data, err := logic.Group.Topic.QueryTopic(ctx, params)
	if err != nil {
		reply.ToResponseError(err)
		return
	}
	reply.ToResponseData(data)
}
