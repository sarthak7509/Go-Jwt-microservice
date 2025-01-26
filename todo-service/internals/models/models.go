package models

import (
	"time"

	"github.com/sarthak7509/go-micro/todo-service/internals/database"
)

// Todo represents the structure of a TODO item
type Todo struct {
	ID          int64     `json:"id"`                    // Unique identifier
	UserID      int64     `json:"user_id"`               // Associated user ID
	Title       string    `json:"title"`                 // Task title
	Description string    `json:"description,omitempty"` // Task description, omits if empty
	Status      string    `json:"status"`                // Status of the task (e.g., pending, in progress, completed)
	Priority    string    `json:"priority"`              // Priority level (e.g., low, medium, high)
	DueDate     time.Time `json:"due_date,omitempty"`    // Task due date, omits if empty
	CreatedAt   time.Time `json:"created_at,omitempty"`  // Timestamp when the task was created
	UpdatedAt   time.Time `json:"updated_at,omitempty"`  // Timestamp when the task was last updated
}

func (todo *Todo) Save() error {
	//save it to database
	query := `
	INSERT INTO todo (user_id, title, description, status, priority, due_date, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	` //safe way of using query
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(todo.UserID, todo.Title, todo.Description, todo.Status, todo.Priority, todo.DueDate, time.Now(), time.Now())

	if err != nil {
		return err
	}
	todo.ID, err = result.LastInsertId()
	return err
}

func GetAllTodo() ([]Todo, error) {
	query := `select * from todo`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil

}
