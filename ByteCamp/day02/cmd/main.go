package main

import (
	"go_dance/day_2/1_post/internel/dao"
	"go_dance/day_2/1_post/internel/dao/file"
	"go_dance/day_2/1_post/internel/pkg/snowflake"
	"go_dance/day_2/1_post/internel/routing"
	"time"
)

const DATAPATH = "/home/raja/workspace/go/src/go_dance/day_2/1_post/storge"

func main() {
	if err := snowflake.Init(time.Now(), 1); err != nil {
		panic(err)
	}
	db, err := file.InitByPath(DATAPATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dao.InitDB(db)
	router := routing.NewRouter()
	if err := router.Run(":5001"); err != nil {
		panic(err)
	}
}
