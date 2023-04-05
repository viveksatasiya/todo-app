# ToDo App

This repository contains a simple ToDo application implemented in Go using the Test-Driven Development (TDD) approach and a layered architecture with clean code principles.

## Features

- Create a ToDo item
- List all ToDo items
- Retrieve a ToDo item by ID
- Update a ToDo item by ID
- Delete a ToDo item by ID

## Project Structure

- `cmd/api`: Main application entry point and server setup
- `internal/app/service`: Application layer containing the business logic
- `internal/domain`: Domain layer containing domain entities, repository interfaces, and service interfaces
- `internal/infrastructure/repository`: Infrastructure layer with repository implementations (memory and PostgreSQL)
- `internal/presentation/handler`: Presentation layer with API handlers

## Requirements

- Go 1.16+
- PostgreSQL (optional, for persistent storage)

## Setup

1. Clone the repository:

```
git clone https://github.com/yourusername/todo-app.git
```

2. Change to the project directory:

```
cd todo-app
```

3. Install dependencies:

```
go mod tidy
```

4. Set up the PostgreSQL database (optional):

- Create a new database and user in PostgreSQL with the necessary privileges
- Update the connection string in the `main.go` file with your database credentials

## Running the Application

1. Start the application:

```
go run cmd/api/main.go
```

2. The server will start on port 8080. Use an API client (e.g., Postman, curl) to make requests to the endpoints:

- `POST /todos`: Create a new ToDo item
- `GET /todos`: List all ToDo items
- `GET /todos/{id}`: Retrieve a ToDo item by ID
- `PUT /todos/{id}`: Update a ToDo item by ID
- `DELETE /todos/{id}`: Delete a ToDo item by ID

## Running Tests

To run tests for the entire project, use the following command:

```
go test ./...
```

## Contributing

If you would like to contribute to this project, please feel free to submit a pull request or open an issue on GitHub.

## License

This project is licensed under the [MIT License](LICENSE).
