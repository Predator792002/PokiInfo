package main

import (
	"bufio"
	"fmt"
	"os"
)

// go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PokiInfo >")
		if !scanner.Scan() {
			break
		}
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		cmdName := cleaned[0]
		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println(err)
		}
	}
}
