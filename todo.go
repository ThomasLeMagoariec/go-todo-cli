package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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
	case "add":
		addTask(args[1:])
	case "help":
		displayHelpMessage()
	case "test":
		res := mergeIntoOne(args)
		fmt.Println(res)
	case "status":
		if len(args) >= 4 {
			updateTaskStatus(entries, args[1], args[2], mergeIntoOne(args[3:]))
		} else {
			fmt.Println("wrong number of arguments passed")
			os.Exit(1)
		}
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

		//! check for ghost entries
		if entries[i][0] == "" || entries[i][1] == "" {
			w.Flush()
			fmt.Println("'tasks.csv' seems to be missing data. (id:", i, ")")
			os.Exit(1)
		}

		fmt.Fprintln(w, i, "\t", entries[i][0], "\t", entries[i][1])
	}

	w.Flush()
}

// ! technically useless if you surround the name in "
func mergeIntoOne(strs []string) string {
	var message string
	for i, str := range strs {
		if i == len(strs)-1 {
			message += str
		} else {
			message += str + " "
		}
	}

	return message
}

func addTask(name []string) {
	fmt.Println("add task:", name)

	f, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data [][]string
	data = append(data, []string{mergeIntoOne(name), "incomplete"})

	w := csv.NewWriter(f)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task Added!")
}

func updateTaskStatus(entries [][]string, id string, field string, value string) {
	int_id, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if int_id == 0 || int_id > len(entries) {
		fmt.Println("invalid task ID")
		os.Exit(1)
	}

	if !(field == "status" || field == "name") {
		fmt.Println("invalid field '", field, "'. Use 'status' or 'name'")
		os.Exit(1)
	}

	f, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i, entry := range entries {
		if int_id == i {
			fmt.Println("found task |", entry)
			if field == "status" {
				entries[i][1] = value
			} else {
				entries[i][0] = value
			}
			break
		}
	}

	f.Truncate(0)
	f.Seek(0, 0)

	w := csv.NewWriter(f)
	w.WriteAll(entries)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}

func displayHelpMessage() {
	fmt.Println("TODO APP")
	fmt.Println("\nAvailable commands:")
	fmt.Println("\thelp\tdisplays this message")
	fmt.Println("\tlist\tdisplays tasks")

	//? this messes up alignment
	//? fmt.Println("\tdosomething\tsmth")
}
