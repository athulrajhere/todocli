package cli

import (
	"fmt"
	"strings"

	"github.com/athulrajhere/todocli/todo"
	"github.com/manifoldco/promptui"
)

type Handler struct {
	service todo.TodoService
}

func (h *Handler) Run(args []string) {
	if len(args) == 0 {
		fmt.Println("no command provided")
		return
	}

	switch args[0] {
	case "add":
		h.handleAdd(args[1:])
	case "list":
		h.handleList()
	case "complete":
		h.handleComplete(args[1:])
	case "delete":
		h.handleDelete(args[1:])
	default:
		fmt.Printf("unknown command: %s\n", args[0])
	}
}

func (h *Handler) handleAdd(args []string) {
	if len(args) == 0 {
		fmt.Println("title is required")
		return
	}
	title := strings.Join(args, " ")
	t, err := h.service.Add(title)
	if err != nil {
		fmt.Printf("failed to add todo: %v\n", err)
		return
	}
	fmt.Println("─────────────────────────────────────────")
	fmt.Println("  ✔ Todo created successfully")
	fmt.Println("─────────────────────────────────────────")
	printTodo(t)
	fmt.Println("─────────────────────────────────────────")
}

func (h *Handler) handleList() {
	todos, err := h.service.List()
	if err != nil {
		fmt.Printf("failed to list todos: %v\n", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("no todos yet. add one with: todo add <title>")
		return
	}
	fmt.Println("─────────────────────────────────────────────────────────")
	fmt.Printf("  %-4s  %-8s  %-28s  %-10s\n", "Done", "ID", "Title", "Created")
	fmt.Println("─────────────────────────────────────────────────────────")
	for _, t := range todos {
		status := " "
		if t.Completed {
			status = "✓"
		}
		fmt.Printf("  [%s]   %-8s  %-28s  %s\n",
			status,
			t.ID[:8],
			t.Title,
			t.CreatedAt.Format("2006-01-02"),
		)
	}
	fmt.Println("─────────────────────────────────────────────────────────")
	fmt.Printf("  %d todo(s) total\n", len(todos))
}

func (h *Handler) handleComplete(_ []string) {
	todos, err := h.service.List()
	if err != nil {
		fmt.Printf("failed to fetch todos: %v\n", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("no todos found")
		return
	}

	labels := make([]string, len(todos))
	for i, t := range todos {
		status := " "
		if t.Completed {
			status = "✓"
		}
		labels[i] = fmt.Sprintf("[%s] %s", status, t.Title)
	}

	prompt := promptui.Select{
		Label: "Select todo to complete",
		Items: labels,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("prompt failed: %v\n", err)
		return
	}

	selected := todos[index]
	err = h.service.Complete(selected.ID)
	if err != nil {
		fmt.Printf("failed to complete todo: %v\n", err)
		return
	}
	fmt.Printf("✅ marked as completed: %s\n", selected.Title)
}

func (h *Handler) handleDelete(_ []string) {
	todos, err := h.service.List()
	if err != nil {
		fmt.Printf("failed to fetch todos: %v\n", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("no todos found")
		return
	}

	labels := make([]string, len(todos))
	for i, t := range todos {
		status := " "
		if t.Completed {
			status = "✓"
		}
		labels[i] = fmt.Sprintf("[%s] %s", status, t.Title)
	}

	prompt := promptui.Select{
		Label: "Select todo to delete",
		Items: labels,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("prompt failed: %v\n", err)
		return
	}

	selected := todos[index]
	err = h.service.Delete(selected.ID)
	if err != nil {
		fmt.Printf("failed to delete todo: %v\n", err)
		return
	}
	fmt.Printf("🗑️  deleted: %s\n", selected.Title)
}

func NewHandler(service todo.TodoService) *Handler {
	return &Handler{service: service}
}

// printTodo is a private helper to display a single todo nicely
func printTodo(t todo.Todo) {
	status := "Pending"
	if t.Completed {
		status = "Completed"
	}
	fmt.Printf("  ID      : %s\n", t.ID)
	fmt.Printf("  Title   : %s\n", t.Title)
	fmt.Printf("  Status  : %s\n", status)
	fmt.Printf("  Created : %s\n", t.CreatedAt.Format("2006-01-02 15:04:05"))
}