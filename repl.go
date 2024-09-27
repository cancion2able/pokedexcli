package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	commands := getCommands()
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()

		words := cleanInput(sc.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]	
		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
		
		err := cmd.callback()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	words := strings.Fields(lowered)
	return words
}
