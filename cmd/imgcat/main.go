package main

import (
	"log"

	"github.com/nugget/iterm2"
)

func main() {
	err := iterm2.PrintImage("sample.png")
	if err != nil {
		log.Fatal(err)
	}
}
