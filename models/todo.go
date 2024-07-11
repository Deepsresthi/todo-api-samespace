package models

import (
	"fmt"
	"time"
	"todo-api/config"

	"github.com/gocql/gocql"
)

type TodoItem struct {
	ID          gocql.UUID `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created     time.Time  `json:"created"`
	Updated     time.Time  `json:"updated"`
	CreatedUnix int64      `json:"created_unix"`
}

func (item *TodoItem) Save() error {
	item.ID = gocql.TimeUUID()
	item.Created = time.Now()
	item.CreatedUnix = item.Created.Unix()
	item.Updated = item.Created

	return config.Session.Query(`INSERT INTO todo (id, user_id, title, description, status, created, updated, created_unix) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		item.ID, item.UserID, item.Title, item.Description, item.Status, item.Created, item.Updated, item.CreatedUnix).Exec()
}

func GetTodoItems(userID string, status string, limit int) ([]TodoItem, error) {

	var query string
	var args []interface{}

	query = `SELECT id, user_id, title, description, status, created, updated FROM todo WHERE user_id = ?`
	args = append(args, userID)

	if status != "" {
		query += ` AND status = ?`
		args = append(args, status)
	}

	query += ` LIMIT ? ALLOW FILTERING`
	args = append(args, limit)

	iter := config.Session.Query(query, args...).Iter()

	var items []TodoItem
	var item TodoItem
	for iter.Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Status, &item.Created, &item.Updated) {
		items = append(items, item)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return items, nil
}

func (item *TodoItem) Update() error {
	item.Updated = time.Now()
	return config.Session.Query(`UPDATE todo SET title = ?, description = ?, status = ?, updated = ? WHERE id = ?`,
		item.Title, item.Description, item.Status, item.Updated, item.ID).Exec()
}

func DeleteTodoItem(id gocql.UUID) error {
	return config.Session.Query(`DELETE FROM todo WHERE id = ?`, id).Exec()
}

func GetTodoItemsByStatus(userID string, status string) ([]TodoItem, error) {
	var query string
	var args []interface{}

	query = `SELECT id, user_id, title, description, status, created, updated FROM todo WHERE user_id = ? AND status = ? ALLOW FILTERING`
	args = append(args, userID, status)

	iter := config.Session.Query(query, args...).Iter()

	var items []TodoItem
	var item TodoItem
	for iter.Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Status, &item.Created, &item.Updated) {
		items = append(items, item)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return items, nil
}

func GetSortTodoItems(userID string, sort string, limit int) ([]TodoItem, error) {
	var query string
	var items []TodoItem

	// Validate sort order (optional)
	if sort != "ASC" && sort != "DESC" {
		return nil, fmt.Errorf("invalid sort order. Allowed values: ASC, DESC")
	}

	// Construct query based on sort order
	query = `SELECT id, user_id, title, description, status, created, updated, created_unix 
            FROM todo 
            WHERE user_id = ? 
            ORDER BY created_unix %s 
            LIMIT ?`

	// Format the query with the appropriate sort order
	query = fmt.Sprintf(query, sort)

	fmt.Println("Query:", query)
	// Execute query
	iter := config.Session.Query(query, userID, limit).Iter()

	// Iterate through results and populate items
	var item TodoItem
	for iter.Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Status, &item.Created, &item.Updated, &item.CreatedUnix) {
		items = append(items, item)
	}

	// Check for errors during iteration
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return items, nil
}
