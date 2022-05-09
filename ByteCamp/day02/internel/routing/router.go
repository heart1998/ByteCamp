package routing

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	root := r.Group("/")
	{
		Group.Topic.Init(root)
		Group.Post.Init(root)
	}
	return r
}
