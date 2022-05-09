package request

import "go_dance/day_2/1_post/internel/pkg/app/errcode"

type PublicTopic struct {
	Title   string `binding:"required" json:"title"`
	Content string `binding:"required" json:"content"`
}

func (topic *PublicTopic) Check() errcode.Err {
	switch {
	case len(topic.Title) > 100, len(topic.Content) > 1000:
		return errcode.ErrLengthOver
	case len(topic.Title) == 0, len(topic.Content) == 0:
		return errcode.ErrLengthZero
	}
	return nil
}

type QueryTopic struct {
	ID int64 `binding:"required" json:"id"`
}
