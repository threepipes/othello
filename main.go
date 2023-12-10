package main

import (
	"log"
	"othello/cli"
)

func main() {
	if err := cli.GamePlay(); err != nil {
		log.Fatal(err)
	}
}
