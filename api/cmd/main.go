package main

import (
	"github.com/kose-yusuke/gocrud/api/cmd/server"
    "github.com/kose-yusuke/gocrud/api/cmd/db"
)

func main() {
    db.Init()
    server.Init()
	
    db.Close()
}

