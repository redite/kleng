package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/redite/kleng/config"
	"github.com/redite/kleng/core"
	"github.com/redite/kleng/server"
)

var (
	mode = flag.String("M", "", "run with `mode`. available mode: 'populate', 'http'")
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Got error: %s", err.Error())
		os.Exit(-1)
	}

	switch *mode {
	case "http":
		log.Fatal(server.RunServer(conf))
		break
	case "populate":
		err = core.Populate()
		if err != nil {
			log.Printf("Error Populating: %s\n", err.Error())
		}
		break
	default:
		flag.Usage()
	}

}
