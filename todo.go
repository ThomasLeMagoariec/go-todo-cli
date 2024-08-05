package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <command>\nfor a list of available commands use:\n%s help\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	args := os.Args[1:]

	switch args[0] {
	case "list":
		fmt.Println("list!")
	case "help":
		displayHelpMessage()
	default:
		fmt.Printf("unknown command: %s\n", args[0])
	}
}

func displayHelpMessage() {
	fmt.Println("TODO APP")
	fmt.Println("\nAvailable commands:")
	fmt.Println("\thelp\tdisplays this message")
	fmt.Println("\tlsit\tdisplays tasks")
}
