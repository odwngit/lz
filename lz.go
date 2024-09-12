package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
 
	for _, entry := range entries {
        	if entry.IsDir() {
			fmt.Printf("%v%v%v  %v\n", Ansi("Blue"), "<./>", Ansi("Reset"), entry.Name())
		} else {
			var id FileIdentification = Identify(entry.Name()) 
			fmt.Printf("%v%v%v  %v %v(%v)%v\n", id.ansi, id.symbol, Ansi("Reset"), entry.Name(), Ansi("Gray"), id.name, Ansi("Reset"))
		}
	}
}
