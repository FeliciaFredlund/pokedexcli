package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replLoop(config *config) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		if err := reader.Err(); err != nil {
			fmt.Printf("Scanning error occured: %v\n", err)
			os.Exit(1)
		}

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		if command, exists := config.cmds[input[0]]; exists {
			var arg string
			if len(input) > 1 {
				arg = input[1]
			}
			if err := command.callback(config, arg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)
	return words
}
