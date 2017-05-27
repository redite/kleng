package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/redite/kleng/config"
	"github.com/redite/kleng/server"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Got error: %s", err.Error())
		os.Exit(-1)
	}

	pp.Println(conf)
	err = server.RunServer(conf)
	if err != nil {
		pp.Println(err)
	}
}
