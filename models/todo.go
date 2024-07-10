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
}

func (item *TodoItem) Save() error {
	item.ID = gocql.TimeUUID()
	item.Created = time.Now()
	item.Updated = item.Created

	return config.Session.Query(`INSERT INTO todo (id, user_id, title, description, status, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		item.ID, item.UserID, item.Title, item.Description, item.Status, item.Created, item.Updated).Exec()
}

func GetTodoItems(userID string, status string, limit int, offset int) ([]TodoItem, error) {
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

	fmt.Println("Query:", query) // Print the query string for debugging purposes
	fmt.Println("Args:", args)   // Print the arguments for debugging purposes

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
