package main

import (
	"log"

	"github.com/pitakill/http-cli-concurrency/cli"
	server "github.com/pitakill/http-cli-concurrency/http-server"
)

func main() {
	chHTTP := make(chan error)
	chCLI := make(chan string)

	go server.Start(chHTTP)
	go cli.Start(chCLI)

	for {
		select {
		case err := <-chHTTP:
			log.Println(err)
		case data := <-chCLI:
			log.Println(data)
		}
	}
}
