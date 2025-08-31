package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("PokiInfo >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}
