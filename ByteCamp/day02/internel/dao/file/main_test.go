package file

import (
	"github.com/stretchr/testify/require"
	"go_dance/day_2/1_post/internel/pkg/snowflake"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := snowflake.Init(time.Now(), 1); err != nil {
		panic(err)
	}
	m.Run()
}

func testCreateFile(t *testing.T) *os.File {
	f, err := os.CreateTemp("", "postDB")
	require.NoError(t, err)
	return f
}
