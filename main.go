package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range getCommands() {
		fmt.Printf("%s: %s", v.name, v.description)
		fmt.Println()
	}
	fmt.Println()
	return nil	
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	commands := getCommands()
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if sc.Scan() {
			cmd := sc.Text()
			cmds, exists := commands[cmd]
			if !exists {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
				continue
			}
			err := cmds.callback()
			if err != nil {
				fmt.Printf("Error executing command: %s\n", err)
			}
		}
	}
}
