package main

import (
	"github.com/kose-yusuke/tripblog/api/cmd/server"
    "github.com/kose-yusuke/tripblog/api/cmd/db"
)

func main() {
    db.Init()
    server.Init()
	
    db.Close()
}

