package file

import (
	"context"
	"go_dance/day_2/1_post/internel/model/database"
	"go_dance/day_2/1_post/internel/pkg/snowflake"
	"io"
	"sync"
	"time"
)

type postDB struct {
	list map[int64][]database.Post
	rw   sync.RWMutex
	db   io.ReadWriteCloser
}

func initPostDB(db io.ReadWriteCloser) (*postDB, error) {
	postDB := &postDB{list: map[int64][]database.Post{}, rw: sync.RWMutex{}, db: db}
	if err := fileToData[database.Post](db, func(post database.Post) {
		postDB.list[post.ParentId] = append(postDB.list[post.ParentId], post)
	}); err != nil {
		return nil, err
	}
	return postDB, nil
}

func (D *postDB) SavePost(ctx context.Context, ParentId int64, Content string) (int64, error) {
	D.rw.Lock()
	defer D.rw.Unlock()
	post := database.Post{
		Id:         snowflake.GetID(),
		ParentId:   ParentId,
		Content:    Content,
		CreateTime: time.Now(),
	}
	D.list[ParentId] = append(D.list[ParentId], post)
	if err := dataToFile(post, D.db); err != nil {
		D.list[ParentId] = D.list[ParentId][:len(D.list)-1]
		return 0, err
	}
	return post.Id, nil
}

func (D *postDB) QueryPostsByTopicID(ctx context.Context, topicID int64) ([]database.Post, error) {
	D.rw.RLock()
	defer D.rw.RUnlock()
	if posts, ok := D.list[topicID]; ok {
		return posts, nil
	}
	return []database.Post{}, nil
}

func (D *postDB) Close() {
	D.db.Close()
}
