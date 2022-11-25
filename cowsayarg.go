package main

import (
	"fmt"
	"os"
)

// ohne scanner, stattdessen mit argument Uebergabe als string
func main() {

	if len(os.Args) == 1 {
		fmt.Println("error: bitte argument eingeben!")
		os.Exit(0)
	}

	text := os.Args[1]

	for k := 0; k < len(text); k++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println(text)

	for k := 0; k < len(text); k++ {
		fmt.Printf("-")
	}

	printCow() // Aufrufen von Methode
}

func printCow() {
	cow := ` 

       \   ^__^
        \  (oo)\_______
           (__)\       )\/\
               ||----w |
               ||     ||  			           
			       
			   `

	fmt.Println(cow) // ausgabe in Terminal
}
