# ToDo Application
 
The ToDo application is a robust backend system designed to manage tasks efficiently. It provides a set of APIs that enable users to create, update, retrieve, and delete tasks (ToDo items) with associated details such as title, description, status, and timestamps. Built using Golang and ScyllaDB, the application leverages these technologies for high performance and scalability.

## Running the code

### Prerequisites

- Install `Docker` to manage containerized applications.

### Setting Up ScyllaDB with Docker:

- Pull ScyllaDB Docker Image:
  
```
docker pull scylladb/scylla
```
- Run the ScyllaDB container:

```
docker run -d --name todo-scylla -p 9042:9042 scylladb/scylla
```
- Check if the container is running:

```
docker ps -a
```
- Connect to ScyllaDB

```
docker exec -it todo-scylla cqlsh
```

-  Create Keyspace and Table

```
CREATE KEYSPACE todo WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE todo;

CREATE TABLE todo (
    id UUID PRIMARY KEY,
    user_id UUID,
    title TEXT,
    description TEXT,
    status TEXT,
    created TIMESTAMP,
    updated TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    user_id TEXT PRIMARY KEY,
    email TEXT,
    full_name TEXT,
    created BIGINT
);

```

## Code Structure

```
todo-api/
├── main.go
├── controllers/
│   ├── todo_controller.go
│   ├── user_controller.go
├── models/
│   ├── todo.go
│   ├── user.go
├── views/
│   ├── response.go
├── config/
│   ├── config.go
├── routes/
│   ├── routes.go
├── go.mod
└── go.sum
```

## Data Model Descriptions

### TodoItem

A TodoItem represents a task or item in a todo list. Here are the properties of a TodoItem:

- `ID`: Unique identifier for the todo item.
- `UserID`: The user ID associated with the todo item.
- `Title`: The title or name of the todo item.
- `Description`: Additional details or description of the todo item.
- `Status`: Current status of the todo item, such as 'pending', 'in progress', or 'completed'.
- `Created`: The date and time when the todo item was created.
- `Updated`: The date and time when the todo item was last updated.

### User

A User represents a user entity with associated information. Here are the properties of a User:

- `UserID`: Unique identifier for the user.
- `Email`: The email address associated with the user.
- `FullName`: The full name of the user.



## API Documentation

[API Documentation](APIs.md)
