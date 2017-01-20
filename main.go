package main

import (
	"fmt"
	"os"

	"github.com/redite/kleng/server"
	"github.com/redite/kleng/utils"
)

func main() {
	conf, err := utils.LoadConfig()
	if err != nil {
		fmt.Printf("Got error: %s", err.Error())
		os.Exit(-1)
	}

	server.RunServer(conf)
}
