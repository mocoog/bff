package main

import (
	"flag"
	"log"

	"github.com/mocoog/bff/interface/provider"
	"github.com/mocoog/bff/utils"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")

func main() {
	server, err := provider.InitializeAPIServer()
	if err != nil {
		log.Fatalln(utils.CustomErrorLog(err))
	}
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag

	if err := server.Serve(); err != nil {
		log.Fatalln(utils.CustomErrorLog(err))
	}
}
