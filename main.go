package main

import (
	"github.com/lancelrq/go-iris-starter-kit/src"
	"github.com/lancelrq/go-iris-starter-kit/src/utils"
	"log"
)

func main() {
	utils.InitGlobal()
	server := src.NewServer()
	if err := server.RunServer(); err != nil {
		log.Fatal(err)
	}
}


