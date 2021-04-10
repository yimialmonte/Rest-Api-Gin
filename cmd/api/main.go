package main

import (
	"log"

	"github.com/yimialmonte/rest/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
