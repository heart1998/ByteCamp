package file

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestInitByPath(t *testing.T) {
	root, err := os.Getwd()
	require.NoError(t, err)
	db, err := InitByPath(root)
	require.NoError(t, err)
	require.NotNil(t, db)
	require.NotNil(t, db.postDB)
	require.NotNil(t, db.topicDB)
	file1 := path.Join(root, "postDB")
	file2 := path.Join(root, "topicDB")
	require.FileExists(t, file1)
	require.FileExists(t, file2)
	db.Close()
	err = os.Remove(file1)
	require.NoError(t, err)
	err = os.Remove(file2)
	require.NoError(t, err)
}
