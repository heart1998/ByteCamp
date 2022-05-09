package logic

import (
	"github.com/gin-gonic/gin"
	"go_dance/day_2/1_post/internel/dao"
	"go_dance/day_2/1_post/internel/global"
	"go_dance/day_2/1_post/internel/model/database"
	"go_dance/day_2/1_post/internel/model/request"
	"go_dance/day_2/1_post/internel/model/response"
	"go_dance/day_2/1_post/internel/pkg/app/errcode"
	"sync"
)

type topic struct {
}

func (t *topic) PublicTopic(ctx *gin.Context, params *request.PublicTopic) (int64, errcode.Err) {
	id, err := dao.Store.SaveTopic(ctx, params.Title, params.Content)
	if err != nil {
		global.Log.Println(err)
		return 0, errcode.ErrSave
	}
	return id, nil
}

func (t *topic) QueryTopic(ctx *gin.Context, params *request.QueryTopic) (*response.QueryTopic, errcode.Err) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var (
		topic              *database.Topic
		posts              []database.Post
		topicErr, postsErr error
	)
	go func() {
		defer wg.Done()
		topic, topicErr = dao.Store.QueryTopic(ctx, params.ID)
	}()
	go func() {
		defer wg.Done()
		posts, postsErr = dao.Store.QueryPostsByTopicID(ctx, params.ID)
	}()
	wg.Wait()
	switch {
	case topic == nil:
		return nil, errcode.NotFound
	case topicErr != nil, postsErr != nil:
		global.Log.Println(topicErr, postsErr)
		return nil, errcode.ErrQuery
	default:
		return &response.QueryTopic{Topic: *topic, Post: posts}, nil
	}
}
