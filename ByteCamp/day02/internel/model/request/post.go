package request

import "go_dance/day_2/1_post/internel/pkg/app/errcode"

type PublicPost struct {
	ParentId int64  `binding:"required" json:"parent_id"`
	Content  string `binding:"required" json:"content"`
}

func (post *PublicPost) Check() errcode.Err {
	switch {
	case len(post.Content) > 1000:
		return errcode.ErrLengthOver
	case len(post.Content) == 0:
		return errcode.ErrLengthZero
	}
	return nil
}
