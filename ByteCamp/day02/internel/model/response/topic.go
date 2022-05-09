package response

import "go_dance/day_2/1_post/internel/model/database"

type PublicTopic struct {
	ID int64 `json:"id"`
}

type QueryTopic struct {
	Topic database.Topic  `json:"topic"`
	Post  []database.Post `json:"post"`
}
