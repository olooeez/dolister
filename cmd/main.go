package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/olooeez/dolister/internal/task"
)

var tl *task.TaskList = task.NewTaskList()
var tasksFilePath string

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: Can't get user's home directory: %v", err)
	}

	tasksFilePath = filepath.Join(homeDir, ".dolister.json")

	tl.LoadFromFile(tasksFilePath)

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	addTitle := addCmd.String("title", "", "title of the todo item")
	doneID := doneCmd.Int("id", 0, "id of the todo item to mark as done")
	deleteID := deleteCmd.Int("id", 0, "id of the todo item to delete")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand(addCmd, addTitle)
	case "list":
		listCommand(listCmd)
	case "done":
		doneCommand(doneCmd, doneID)
	case "delete":
		deleteCommand(deleteCmd, deleteID)
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}

func printUsage() {
	fmt.Println("Usage of todo:")
	fmt.Println("  -add cmd")
	fmt.Println("     Add a new todo item")
	fmt.Println("  -list cmd")
	fmt.Println("     List all todo items")
	fmt.Println("  -done cmd")
	fmt.Println("     Mark a todo item as done")
	fmt.Println("  -delete cmd")
	fmt.Println("     Delete a todo item")
}

func addCommand(addCmd *flag.FlagSet, addTitle *string) {
	addCmd.Parse(os.Args[2:])

	if *addTitle == "" {
		log.Fatalln("Error: title of the task cannot be empty.")
	}

	tl.AddTask(*addTitle)
	tl.SaveToFile(tasksFilePath)
}

func listCommand(listCmd *flag.FlagSet) {
	listCmd.Parse(os.Args[2:])

	tl.LoadFromFile(tasksFilePath)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.StripEscape)

	fmt.Fprintln(tw, "ID\tTitle\tCompleted")

	for _, task := range tl.Tasks {
		fmt.Fprintf(tw, "%d\t%s\t%v\n", task.ID, task.Title, task.Completed)
	}

	tw.Flush()
}

func doneCommand(doneCmd *flag.FlagSet, doneID *int) {
	doneCmd.Parse(os.Args[2:])

	if *doneID <= 0 {
		log.Fatalln("Error: You need to pass a valid id number.")
	}

	tl.CompleteTask(*doneID)
	tl.SaveToFile(tasksFilePath)
}

func deleteCommand(deleteCmd *flag.FlagSet, deleteID *int) {
	deleteCmd.Parse(os.Args[2:])

	if *deleteID <= 0 {
		log.Fatalln("Error: You need to pass a valid id number.")
	}

	tl.RemoveTask(*deleteID)
	tl.SaveToFile(tasksFilePath)
}
