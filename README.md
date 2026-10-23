# CLI Task Manager

A simple command-line task manager written in Go. The application allows users to create, view, complete, and delete tasks. Tasks are stored locally in a JSON file.

## Features

- **Add tasks** — create tasks with a custom title.
- **List tasks** — display all saved tasks.
- **Complete tasks** — mark tasks as completed.
- **Delete tasks** — remove tasks by ID.
- **JSON persistence** — store task data locally in `tasks.json`.

## Technologies

- **Go** — application logic.
- **JSON** — local data storage.
- **Go standard library** — input/output and file operations.

## Getting Started

### Prerequisites

Install Go and make sure the `go` command is available in your terminal.

Check your installation:

```bash
go version
```

### Installation

Clone the repository:

```bash
git clone https://github.com/FaLnA1203/CLI-Task-Manager.git
```

Navigate to the project directory:

```bash
cd CLI-Task-Manager
```

### Run the application

Start the application with:

```bash
go run .
```

## Usage

After launching the application, use the interactive menu:

```text
Task Manager

1. Add task
2. List tasks
3. Complete task
4. Delete task
5. Exit

Enter number:
```

Select an option by entering its number.

| Option | Description |
|---|---|
| `1` | Add a new task |
| `2` | Display all tasks |
| `3` | Mark a task as completed |
| `4` | Delete a task |
| `5` | Exit the application |

Tasks have the following properties:

- `Id` — task identifier.
- `TaskName` — task title.
- `Completed` — completion status.

## Project Structure

```text
CLI-Task-Manager/
├── main.go           # Application entry point and menu
├── task/
│   └── task.go       # Task model and related operations
├── storage/
│   └── storage.go    # JSON storage operations
├── tasks.json        # Persistent task data
├── go.mod            # Go module definition
└── README.md         # Project documentation
```

## Data Storage

The application stores tasks in `tasks.json` in JSON format.

Example:

```json
[
  {
    "Id": 1,
    "TaskName": "Learn Go",
    "Completed": false
  },
  {
    "Id": 2,
    "TaskName": "Learn SQL",
    "Completed": true
  }
]
```

The file allows task data to be preserved between application runs.

## Running Tests

Run all project tests with:

```bash
go test ./...
```

To display detailed test output:

```bash
go test -v ./...
```

## Learning Goals

This project is intended as a practical exercise in Go development. It provides experience with:

- Structs and methods.
- Slices and pointers.
- Packages and modules.
- Console input and output.
- File handling and JSON serialization.
- Error handling and unit testing.

## Future Improvements

Potential improvements include:

- Better input validation and error handling.
- More comprehensive unit tests.
- Improved task ID management.
- Filtering tasks by completion status.
- Editing existing tasks.
- Adding priorities and deadlines.
- Migrating from JSON storage to a database.
- Developing a REST API.

## License

No license has been specified yet.
