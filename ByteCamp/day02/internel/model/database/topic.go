package database

import "time"

// Topic 主题
type Topic struct {
	Id         int64
	Title      string
	Content    string
	CreateTime time.Time
}
