package main

import (
	"fmt"

	"github.com/DLuCJ/gonant/gonant"
)

func main() {
	fmt.Println(gonant.Column{})
	fmt.Println(gonant.Song{})
	fmt.Println(gonant.Instrument{})

	fmt.Printf("Hello, world!  This be Gonant.\n")
	gonant.Gonant_Init()
}
