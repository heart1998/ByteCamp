package database

import "time"

// Post 回帖
type Post struct {
	Id         int64
	ParentId   int64
	Content    string
	CreateTime time.Time
}
