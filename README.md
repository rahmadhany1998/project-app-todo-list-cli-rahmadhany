# 📝 To-Do List CLI App (Golang)

A simple Command Line Interface (CLI) application for managing a To-Do List using the Go programming language. All data is stored locally in JSON format.

---

## 📝 Explanatory Video

[Explanatory Video](https://drive.google.com/file/d/1kjqaM3fwRxtQBr3hBZLDaO5WZC7_v1Eg/view?usp=sharing)

---

## 📁 Folder Structure

```
todo_cli/
├── data/              # JSON storage file
│   └── todo.json
├── dto/               # Data Transfer Objects
│   └── request.go
├── handler/           # CLI handler
│   └── cli_handler.go
├── model/             # Task model structure
│   └── task.go
├── service/           # Task business logic
│   └── task_service.go
├── utils/             # JSON file handling
│   └── file_utils.go
└── main.go            # Application entry point
```

---

## 🚀 How to Run

1. Clone this repository or extract the zip.
2. Initialize Go module:
   ```bash
   go mod init todo-list-cli
   ```
3. Use the following commands:

### Add Task
```bash
go run main.go --add --title="Learn Go" --desc="Practicing CLI and JSON"
```

### List All Tasks
```bash
go run main.go --list
```

### Mark Task as Done
```bash
go run main.go --done=1
```

### Delete a Task
```bash
go run main.go --delete=1
```

### Search Task
```bash
go run main.go --search="json"
```

---

## ✨ Features

- ✅ Add new task
- 📋 View tasks in table format
- 🔍 Search tasks by title or description
- ✅ Mark tasks as completed
- ❌ Delete tasks by ID
- 🧠 Input validation (no empty titles, no duplicate names, no empty search keyword)
- 💾 Data is stored in `data/todo.json` in JSON format

---

## 🧱 Built With

- Language: [Go](https://golang.org/)
- Data Format: JSON
- CLI: Go’s built-in `flag` package

---

## 📌 Notes

- This app is purely terminal-based and does not require a database

