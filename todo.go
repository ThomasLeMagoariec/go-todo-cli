package main

import (
	"encoding/csv"
	"fmt"
	"log"
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

// ? i didn't write this
// ? found it on stackoverflow
// ? may rewrite later
func loadEntries(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func displayHelpMessage() {
	fmt.Println("TODO APP")
	fmt.Println("\nAvailable commands:")
	fmt.Println("\thelp\tdisplays this message")
	fmt.Println("\tlist\tdisplays tasks")

	//? this messes up alignment
	//? fmt.Println("\tdosomething\tsmth")
}
