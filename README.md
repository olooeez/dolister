[![Test](https://github.com/olooeez/dolister/actions/workflows/test.yml/badge.svg)](https://github.com/olooeez/dolister/actions/workflows/test.yml)

# dolister

O dolister is a simple command-line tool for managing your to-do list. It allows you to add, list, mark tasks as done, and delete tasks. It's designed to help you stay organized and keep track of your tasks effortlessly.

## Prerequisites

Before using the dolister, make sure your system meets the following requirements:

1. **Go Installation**: You need to have Go installed on your system. You can download and install Go from the [official website](https://golang.org/dl/). To check if Go is installed, run the following command in your terminal:

   ```
   go version
   ```

## Getting Started

Follow these steps to install and use the dolister:

### Step 1: Install dolister

Install using the go command to your local machine:

```
go install github.com/olooeez/dolister:latest
```

### Step 2: Run the dolister

You can now run the dolister with the following commands:

- Add a new task:

  ```
  dolister add --title "Your task title"
  ```

- List all tasks:

  ```
  dolister list
  ```

- Mark a task as done:

  ```
  dolister done --id 1
  ```

- Delete a task:

  ```
  dolister delete --id 1
  ```

## Contributing

If you're interested in contributing to this project, feel free to open a pull request. We welcome all forms of collaboration!

## License

This dolister is available under the [Apache License 2.0](https://github.com/olooeez/dolister/blob/main/LICENSE). For more information, please see the LICENSE file.