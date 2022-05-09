package global

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "[log]", log.Llongfile|log.Ltime)
