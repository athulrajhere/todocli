# todocli

A fast, minimal todo manager for the terminal built with Go.

## Install
```bash
go install github.com/athulrajhere/todocli@latest
```

## Usage
```bash
# Add a todo
todocli add "Buy groceries"

# List all todos
todocli list

# Mark as complete (interactive)
todocli complete

# Delete a todo (interactive)
todocli delete
```

## Built With

- [Go](https://golang.org/)
- [promptui](https://github.com/manifoldco/promptui) — interactive terminal prompts
