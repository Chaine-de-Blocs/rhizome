package main

import (
	"educoin/src"
	"flag"
	"log"
)

func main() {
	cmdEntry := flag.String("entry", "0.0.0.0", "Boostrap node IP address")

	flag.Parse()

	if err := src.OpenConnection(*cmdEntry + src.DefaultPort); err != nil {
		log.Println(err) // TODO handle errors
	}

	src.Listen()
}
