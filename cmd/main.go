package main

import (
	"flag"
	"log"

	"github.com/Trother555/portctl/app"
	server "github.com/Trother555/portctl/web"
)

func main() {
	inPorts := flag.Int64("inPorts", 1, "number of input ports")
	outPorts := flag.Int64("outPorts", 1, "number of output ports")
	flag.Parse()

	log.Printf("starting server with params: inPorts: %d, outPorts: %d", *inPorts, *outPorts)
	app := app.New(&app.Config{InPorts: *inPorts, OutPorts: *outPorts})
	s := server.New(app)
	s.ListenAndServe()
}
