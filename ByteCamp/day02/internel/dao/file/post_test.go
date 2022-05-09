package file

import (
	"context"
	"github.com/stretchr/testify/require"
	"go_dance/day_2/1_post/internel/model/database"
	"go_dance/day_2/1_post/internel/utils"
	"os"
	"testing"
)

func Test_PostDB(t *testing.T) {
	t.Parallel()
	f, err := os.CreateTemp("", "postDB")
	defer os.Remove(f.Name())
	require.NoError(t, err)
	postDB, err := initPostDB(f)
	require.NoError(t, err)
	require.NotNil(t, postDB)
}

func TestPostDB(t *testing.T) {
	f := testCreateFile(t)
	defer os.Remove(f.Name())
	postDB, err := initPostDB(f)
	require.NoError(t, err)
	cnts := make(map[int64]int)
	posts := make(map[int64]database.Post)
	testList := make([]database.Post, utils.RandomInt(1, 100))
	for i := range testList {
		post := database.Post{
			ParentId: utils.RandomInt(1, 20),
			Content:  utils.RandomString(100),
		}
		testList[i] = post
		cnts[post.ParentId]++
		postID, err := postDB.SavePost(context.Background(), post.ParentId, post.Content)
		require.NoError(t, err)
		require.NotZero(t, postID)
		posts[postID] = post
	}
	for topicID, cnt := range cnts {
		res, err := postDB.QueryPostsByTopicID(context.Background(), topicID)
		require.NoError(t, err)
		require.Len(t, res, cnt)
		for i := range res {
			require.Equal(t, posts[res[i].Id].ParentId, res[i].ParentId)
		}
	}
}
