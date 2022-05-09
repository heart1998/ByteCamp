package file

import (
	"context"
	"github.com/stretchr/testify/require"
	"go_dance/day_2/1_post/internel/model/database"
	"go_dance/day_2/1_post/internel/utils"
	"os"
	"testing"
)

func Test_TopicDB(t *testing.T) {
	t.Parallel()
	f, err := os.CreateTemp("", "topicDB")
	defer os.Remove(f.Name())
	require.NoError(t, err)
	topicDB, err := initTopicDB(f)
	require.NoError(t, err)
	require.NotNil(t, topicDB)
}

func TestTopicDB(t *testing.T) {
	f := testCreateFile(t)
	defer os.Remove(f.Name())
	topicDB, err := initTopicDB(f)
	require.NoError(t, err)
	require.NotNil(t, topicDB)
	topics := make([]database.Topic, utils.RandomInt(1, 100))
	for i := range topics {
		topic := database.Topic{
			Title:   utils.RandomString(100),
			Content: utils.RandomString(100),
		}
		id, err := topicDB.SaveTopic(context.Background(), topic.Title, topic.Content)
		require.NoError(t, err)
		require.NotZero(t, id)
		topic.Id = id
		topics[i] = topic
	}
	for i := range topics {
		topic, err := topicDB.QueryTopic(context.Background(), topics[i].Id)
		require.NoError(t, err)
		require.Equal(t, topic.Title, topics[i].Title)
		require.Equal(t, topic.Content, topics[i].Content)
	}
}
