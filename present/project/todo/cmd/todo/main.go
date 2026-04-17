package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/example/foss-go-101/present/project/todo"
)

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  add <title>     add a new task")
	fmt.Println("  done <id>       mark a task as done")
	fmt.Println("  remove <id>     remove a task")
	fmt.Println("  list            show all tasks")
	fmt.Println("  pending         show pending task count")
	fmt.Println("  help            show commands")
	fmt.Println("  exit            quit")
}

func printList(list *todo.List) {
	items := list.Items()
	if len(items) == 0 {
		fmt.Println("No tasks yet.")
		return
	}

	for _, item := range items {
		status := " "
		if item.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d - %s\n", status, item.ID, item.Title)
	}
}

func parseID(raw string) (int, error) {
	id, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid id: %q", raw)
	}
	return id, nil
}

func main() {
	fmt.Println("Todo CLI (teaching demo)")
	fmt.Println("Type 'help' to see commands.")

	list := todo.New()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		cmd := strings.ToLower(parts[0])
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		switch cmd {
		case "add":
			item, err := list.Add(arg)
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			fmt.Printf("added task %d\n", item.ID)
		case "done":
			id, err := parseID(arg)
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			if err := list.Done(id); err != nil {
				fmt.Println("error:", err)
				continue
			}
			fmt.Printf("task %d marked as done\n", id)
		case "remove":
			id, err := parseID(arg)
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			if err := list.Remove(id); err != nil {
				fmt.Println("error:", err)
				continue
			}
			fmt.Printf("task %d removed\n", id)
		case "list":
			printList(list)
		case "pending":
			fmt.Printf("pending tasks: %d\n", list.PendingCount())
		case "help":
			printHelp()
		case "exit", "quit":
			fmt.Println("bye")
			return
		default:
			fmt.Println("unknown command:", cmd)
			printHelp()
		}
	}
}
