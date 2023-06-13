package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/phouse512/go-coda"
	"golang.org/x/exp/slices"
)

const (
	ELEADS_DOC = "gf7mKlD0Hg"
	MAIN_TABLE = "grid-wwu6CDBkoL"
)

var COL_IDS = map[string]string{
	"name":     "c-mcp8xL01YO",
	"status":   "c-LGgu1oc3QD",
	"assignee": "c-wDankSaUyF",
	"priority": "c-NsTx8GD-Pn",
	"type":     "c-y2KEdTNNfM",
}

func getUserName(codaClient *coda.Client) string {
	userInfo, err := codaClient.GetUserInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return userInfo.Name
}

func viewTasks(codaClient *coda.Client, assignee string) {
	var params coda.ListRowsParameters
	if assignee != "" {
		params = coda.ListRowsParameters{
			Query: fmt.Sprintf(`%s:"%s"`, COL_IDS["assignee"], assignee),
		}
	} else {
		params = coda.ListRowsParameters{}
	}

	rows, err := codaClient.ListTableRows(ELEADS_DOC, MAIN_TABLE, params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// use tabwriter to print the table in terminal
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	defer w.Flush()
	fmt.Fprint(w, "Name\tStatus\tAssignee\tPriority\tType\n")
	for _, row := range rows.Rows {
		data := row.Values
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", data[COL_IDS["name"]], data[COL_IDS["status"]], data[COL_IDS["assignee"]], data[COL_IDS["priority"]], data[COL_IDS["type"]])
	}
}

func addTask(codaClient *coda.Client, name string, status string, assignee string, priority string, taskType string) {
	cells := []coda.CellParam{
		{Column: COL_IDS["name"], Value: name},
		{Column: COL_IDS["status"], Value: status},
		{Column: COL_IDS["assignee"], Value: assignee},
		{Column: COL_IDS["priority"], Value: priority},
		{Column: COL_IDS["type"], Value: taskType},
	}
	params := coda.InsertRowsParameters{
		Rows: []coda.RowParam{
			{Cells: cells},
		},
	}

	codaClient.InsertRows(ELEADS_DOC, MAIN_TABLE, false, params)
}

func main() {
	listSet := flag.NewFlagSet("list", flag.ExitOnError)
	addSet := flag.NewFlagSet("add", flag.ExitOnError)

	listAll := listSet.Bool("all", false, "List all tasks")

	var addPriority string
	addSet.StringVar(&addPriority, "priority", "P2", "Priority of the task")
	addSet.StringVar(&addPriority, "p", "P2", "Priority of the task")

	var addStatus string
	addSet.StringVar(&addStatus, "status", "Not started", "Status of the task")
	addSet.StringVar(&addStatus, "s", "Not started", "Status of the task")

	var addType string
	addSet.StringVar(&addType, "type", "", "Type of the task")
	addSet.StringVar(&addType, "t", "Other", "Type of the task")

	if len(os.Args) < 2 {
		fmt.Println("Expected a subcommand. Run 'coda help' for usage.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listSet.Parse(os.Args[2:])
	case "add":
		addSet.Parse(os.Args[2:])
	case "help":
		fmt.Println("Usage: coda <command> [<args>]\n")
		fmt.Println("Commands and arguments:")
		fmt.Println("  list\t\tList all tasks")
		listSet.PrintDefaults()
		fmt.Println()
		fmt.Println("  add\t\tAdd a new task")
		addSet.PrintDefaults()
		fmt.Println()
		fmt.Println("  help\t\tShow this help message")
		os.Exit(0)
	default:
		fmt.Printf("%q is not valid command. Run 'coda help' for usage\n", os.Args[1])
	}

	// read API key from file
	file, err := os.Open("api_key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	apiKey := scanner.Text()

	codaClient := coda.DefaultClient(apiKey)
	if listSet.Parsed() {
		if *listAll {
			fmt.Println("Listing all tasks")
			viewTasks(codaClient, "")
		} else if listSet.NArg() > 0 {
			fmt.Println("Listing tasks for", listSet.Arg(0))
			viewTasks(codaClient, listSet.Arg(0))
		} else {
			fmt.Println("Listing your tasks")
			viewTasks(codaClient, getUserName(codaClient))
		}
	}
	if addSet.Parsed() {
		fmt.Println("Adding a task")

		addAsignee := os.Args[len(os.Args)-1]

		// check that the status is valid
		statusChoices := []string{"Not started", "In progress", "Done", "Blocked"}
		if !slices.Contains(statusChoices, addStatus) {
			fmt.Printf("Invalid status: %s\n", addStatus)
			os.Exit(1)
		}

		// check that the priority is valid
		priorityChoices := []string{"P0", "P1", "P2", "P3"}
		if !slices.Contains(priorityChoices, addPriority) {
			fmt.Printf("Invalid priority: %s\n", addPriority)
			os.Exit(1)
		}

		// check that the type is valid
		typeChoices := []string{"Firmware", "Hardware", "Administrative", "Other"}
		if !slices.Contains(typeChoices, addType) {
			fmt.Printf("Invalid type: %s\n", addType)
			os.Exit(1)
		}

		// get task name from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter task name: ")
		addName, _ := reader.ReadString('\n')

		// pass in all the info to addTask
		addTask(codaClient, addName, addStatus, addAsignee, addPriority, addType)
	}
}
