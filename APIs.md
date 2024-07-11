# API Documentation

## Create User

### POST /users

Creates a new user.

#### Request Body

```json
{
  "email": "test.user@example.com",
  "full_name": "test user"
}
```

#### Response

1. 201 Created

```json
{
  "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
  "email": "test.user@example.com",
  "full_name": "test user"
}

```

1. 400 Bad Request

```json
{
  "error": "Invalid input data"
}

```


## Create ToDo Item

### POST /todo

Creates a new ToDo item.

#### Request Body

```json
{
  "title": "Complete Project Report",
  "description": "Write a detailed report on the project progress.",
  "status": "pending",
  "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7"
}

```

#### Response

1. 201 Created

```json
{
  "id": "a1b2c3d4-5678-90ab-cdef-1234567890ab",
  "title": "Complete Project Report",
  "description": "Write a detailed report on the project progress.",
  "status": "pending",
  "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
  "created": "2024-07-10T12:00:00Z",
  "updated": "2024-07-10T12:00:00Z"
}

```

1. 400 Bad Request

```json
{
  "error": "Invalid input data"
}

```

## Get ToDo Item

### GET /todo

Retrieves a list of ToDo items for a specific user.

#### Request Body

```json
{
  "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
}


```

#### Response

1. 200 OK

```json
[
  {
    "id": "a1b2c3d4-5678-90ab-cdef-1234567890ab",
    "title": "Complete Project Report",
    "description": "Write a detailed report on the project progress.",
    "status": "pending",
    "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
    "created": "2024-07-10T12:00:00Z",
    "updated": "2024-07-10T12:00:00Z"
  },
]


```

1. 400 Bad Request

```json
{
  "error": "user_id is required"
}

```

## Update ToDo Item

### PUT /todo/{id}

Updates an existing ToDo item.

#### Request Body

```json
{
  "title": "Complete Final Report",
  "description": "Write a final detailed report on the project.",
  "status": "completed"
}


```

#### Response

1. 200 OK

```json
{
  "id": "a1b2c3d4-5678-90ab-cdef-1234567890ab",
  "title": "Complete Final Report",
  "description": "Write a final detailed report on the project.",
  "status": "completed",
  "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
  "created": "2024-07-10T12:00:00Z",
  "updated": "2024-07-10T14:00:00Z"
}


```

1. 400 Bad Request

```json
{
  "error": "Invalid input data"
}

```

## Delete ToDo Item

### DELETE /todo/{id}

Deletes an existing ToDo item.

#### Request Body

```json
{
  "title": "Complete Final Report",
  "description": "Write a final detailed report on the project.",
  "status": "completed"
}


```

#### Response

1. 204 No Content

```
{}

```

2. 400 Bad Request

```json
{
  "error": "Invalid ToDo item ID"
}

```

## Filter ToDo Items by Status

### GET /todo/{user_id}/{status}

Retrieves TODO items for a specified user filtered by status (e.g., pending, completed).

#### Request Body

```json
{}

```

#### Response

1. 200 OK

```
[
  {
    "id": "a1b2c3d4-5678-90ab-cdef-1234567890ab",
    "title": "Complete Project Report",
    "description": "Write a detailed report on the project progress.",
    "status": "pending",
    "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
    "created": "2024-07-10T12:00:00Z",
    "updated": "2024-07-10T12:00:00Z"
  },
  {
    "id": "b1c2d3e4-5678-90ab-cdef-2345678901bc",
    "title": "Prepare Presentation",
    "description": "Create slides for the project presentation.",
    "status": "pending",
    "user_id": "96d9a69b-3e7f-11ef-8968-e454e835abe7",
    "created": "2024-07-10T12:30:00Z",
    "updated": "2024-07-10T12:30:00Z"
  }
]


```

2. 400 Bad Request

```json
{
  "error": "Invalid userID or status"
}

```

3. 404 Not Found

```json
{
  "error": "Invalid userID or status"
}

```
## Sort ToDo Items

### POST /todo/sort

Retrieves TODO items for a specified user sorted by creation date.

#### Request Body

```json
{
  "sort": "desc",
  "user_id": "d5b661dd-3f45-11ef-8b30-e454e835abe7"
}

```

#### Response

1. 200 OK

```
[
  {
    "id": "a1b2c3d4-5678-90ab-cdef-1234567890ab",
    "title": "Complete Project Report",
    "description": "Write a detailed report on the project progress.",
    "status": "pending",
    "user_id": "d5b661dd-3f45-11ef-8b30-e454e835abe7",
    "created": "2024-07-10T12:00:00Z",
    "updated": "2024-07-10T12:00:00Z"
  },
  {
    "id": "b1c2d3e4-5678-90ab-cdef-2345678901bc",
    "title": "Prepare Presentation",
    "description": "Create slides for the project presentation.",
    "status": "pending",
    "user_id": "d5b661dd-3f45-11ef-8b30-e454e835abe7",
    "created": "2024-07-10T12:30:00Z",
    "updated": "2024-07-10T12:30:00Z"
  }
]



```

2. 400 Bad Request

```json
{
  "error": "Invalid request body"
}

```

This documentation outlines how the API works, what it expects in terms of input, and what users can expect in terms of output. Adjust the endpoint path (/identify) and any other specific details as per your project setup.
