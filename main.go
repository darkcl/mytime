package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/darkcl/mytime/config"
	"github.com/darkcl/mytime/db"
	"github.com/darkcl/mytime/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
