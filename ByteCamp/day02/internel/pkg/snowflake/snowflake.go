package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

//雪花算法

var node *snowflake.Node

func Init(t time.Time, machineID int64) error {
	var err error
	snowflake.Epoch = t.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return err
}

func GetID() int64 {
	return node.Generate().Int64()
}
