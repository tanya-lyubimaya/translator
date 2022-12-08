package main

import (
	"github.com/tanya-lyubimaya/translator/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
