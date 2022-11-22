
package main

import
(
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin) // scanner
	fmt.Print("Enter Text:" )
	scanner.Scan()
	text := scanner.Text()
	
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