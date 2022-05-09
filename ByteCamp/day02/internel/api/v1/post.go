package v1

import (
	"github.com/gin-gonic/gin"
	"go_dance/day_2/1_post/internel/global"
	"go_dance/day_2/1_post/internel/logic"
	"go_dance/day_2/1_post/internel/model/request"
	"go_dance/day_2/1_post/internel/pkg/app"
	"go_dance/day_2/1_post/internel/pkg/app/errcode"
)

type post struct{}

func (api *post) PublicPost(ctx *gin.Context) {
	reply := app.NewReply(ctx)
	params := new(request.PublicPost)
	if err := ctx.ShouldBindJSON(params); err != nil {
		global.Log.Println(err)
		reply.ToResponseError(errcode.ErrParamNotValid)
		return
	}
	if err := params.Check(); err != nil {
		reply.ToResponseError(err)
		return
	}
	data, err := logic.Group.Post.PublicPost(ctx, params)
	if err != nil {
		reply.ToResponseError(err)
		return
	}
	reply.ToResponseData(data)
}
