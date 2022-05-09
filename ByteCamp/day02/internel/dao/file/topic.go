package file

import (
	"context"
	"go_dance/day_2/1_post/internel/model/database"
	"go_dance/day_2/1_post/internel/pkg/snowflake"
	"io"
	"sync"
	"time"
)

type topicDB struct {
	list map[int64]database.Topic
	rw   sync.RWMutex
	db   io.ReadWriteCloser
}

func initTopicDB(db io.ReadWriteCloser) (*topicDB, error) {
	topicDB := &topicDB{rw: sync.RWMutex{}, list: map[int64]database.Topic{}, db: db}
	if err := fileToData[database.Topic](db, func(topic database.Topic) {
		topicDB.list[topic.Id] = topic
	}); err != nil {
		return nil, err
	}
	return topicDB, nil
}

func (D *topicDB) SaveTopic(ctx context.Context, Title, Content string) (int64, error) {
	D.rw.Lock()
	defer D.rw.Unlock()
	topic := database.Topic{
		Id:         snowflake.GetID(),
		Title:      Title,
		Content:    Content,
		CreateTime: time.Now(),
	}
	D.list[topic.Id] = topic
	if err := dataToFile(topic, D.db); err != nil {
		delete(D.list, topic.Id)
		return 0, err
	}
	return topic.Id, nil
}

func (D *topicDB) QueryTopic(ctx context.Context, topicID int64) (*database.Topic, error) {
	D.rw.RLock()
	defer D.rw.RUnlock()
	if topic, ok := D.list[topicID]; ok {
		return &topic, nil
	}
	return nil, nil
}

func (D *topicDB) Close() {
	D.db.Close()
}
