package main

import (
	"fmt"
	"os"

	"github.com/redite/kleng/config"
	"github.com/redite/kleng/server"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Got error: %s", err.Error())
		os.Exit(-1)
	}

	server.RunServer(conf)
}
