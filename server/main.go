package main

import (
	"context"
	"demo1/frontend"
	"flag"
	"github.com/ServiceWeaver/weaver"
	"log"
)



func main() {
	flag.Parse()
	if err := weaver.Run(context.Background(), frontend.Serve); err != nil {
		log.Fatal(err)
	}
}
