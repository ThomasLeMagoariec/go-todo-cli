package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <command>\nfor a list of available commands use:\n%s help\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	args := os.Args[1:]
	entries := loadEntries("./tasks.csv")

	switch args[0] {
	case "list":
		listEntries(entries)
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

func listEntries(entries [][]string) {

	w := tabwriter.NewWriter(os.Stdout, 5, 4, 1, ' ', tabwriter.DiscardEmptyColumns)

	fmt.Fprintln(w, "ID\t", entries[0][0], "\t", entries[0][1])

	for i := 1; i < len(entries); i++ {
		if entries[i][0] == "" || entries[i][1] == "" {
			w.Flush()
			fmt.Println("'tasks.csv' seems to be missing data. (id:", i, ")")
			os.Exit(1)
		}

		fmt.Fprintln(w, i, "\t", entries[i][0], "\t", entries[i][1])
	}

	w.Flush()
}

func displayHelpMessage() {
	fmt.Println("TODO APP")
	fmt.Println("\nAvailable commands:")
	fmt.Println("\thelp\tdisplays this message")
	fmt.Println("\tlist\tdisplays tasks")

	//? this messes up alignment
	//? fmt.Println("\tdosomething\tsmth")
}
