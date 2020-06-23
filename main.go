package main

import (
	"fmt"
)

func main() {

	var name string
	var alphabet_count int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &alphabet_count)

	// Printing the given texts
	fmt.Printf("The word %s containg %d number of alphabets.",
		name, alphabet_count)

}
