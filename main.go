package main

import (
	"github.com/tanya-lyubimaya/translate/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
