package main

import (
	"log"

	"github.com/ArthurQR98/challenge_fiber/cmd/api/boostrap"
)

func main() {
	if err := boostrap.Run(); err != nil {
		log.Fatal(err)
	}
}
