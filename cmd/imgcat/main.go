package main

import (
	"log"

	"github.com/nugget/imgcat"
)

func main() {
	err := imgcat.PrintImage("sample.png")
	if err != nil {
		log.Fatal(err)
	}
}
