package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go_dance/day_2/1_post/internel/api/v1"
)

type topic struct{}

func (router *topic) Init(group *gin.RouterGroup) {
	topicGroup := group.Group("topic")
	{
		topicGroup.POST("public", v1.Group.Topic.PublicTopic)
		topicGroup.GET("get", v1.Group.Topic.QueryTopic)
	}
}
