
# Task Manager

A simple CLI-based task manager, developed in Go. This tool allows you to efficiently manage tasks by adding, listing, completing, and deleting them.

## Features

- **Add Tasks:** Add new tasks with a description and priority.
- **List Tasks:** Display all tasks or filter by priority.
- **Complete Tasks:** Mark tasks as completed.
- **Delete Tasks:** Remove unnecessary tasks from the list.

## Prerequisites

Ensure you have Go installed on your computer. The project was developed and tested with Go 1.15 and later. You can download Go from the official [Go website](https://golang.org/dl/).

## Installation

To use this project, first clone the repository to your local computer:

```bash
git clone https://github.com/robertwade/my-task-manager.git
cd my-task-manager
```

## Usage

To start the Task Manager, navigate to the project directory and run the program:

```bash
go run cmd/main.go
```

Follow the instructions in the CLI to manage tasks. You can choose between different options such as adding or listing tasks.

## Project Structure

- `cmd/main.go`: The main entry point of the application.
- `pkg/task/task.go`: Defines the Task structure and constructor function.
- `pkg/storage/storage.go`: Responsible for managing the storage of tasks in a CSV file.

## License

This project is released under the MIT License. Further details can be found in the [LICENSE](LICENSE.md) file.

## Contact

If you have any questions or need support, do not hesitate to contact me via [GitHub Issues](https://github.com/robertwade/my-task-manager/issues).
