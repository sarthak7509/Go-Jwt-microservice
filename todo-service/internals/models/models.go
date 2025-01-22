package models

import (
	"time"

	"github.com/sarthak7509/go-micro/todo-service/internals/database"
)

// Todo represents the structure of a TODO item
type Todo struct {
	ID          int       `json:"id"`                    // Unique identifier
	UserID      int       `json:"user_id"`               // Associated user ID
	Title       string    `json:"title"`                 // Task title
	Description string    `json:"description,omitempty"` // Task description, omits if empty
	Status      string    `json:"status"`                // Status of the task (e.g., pending, in progress, completed)
	Priority    string    `json:"priority"`              // Priority level (e.g., low, medium, high)
	DueDate     time.Time `json:"due_date,omitempty"`    // Task due date, omits if empty
	CreatedAt   time.Time `json:"created_at"`            // Timestamp when the task was created
	UpdatedAt   time.Time `json:"updated_at"`            // Timestamp when the task was last updated
}

func (todo *Todo) Save() {
	//save it to database
	query := `
	INSERT INTO todo (user_id, title, description, status, priority, due_date, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	` //safe way of using query
	stmt, err := database.DB.Prepare(query)
}
