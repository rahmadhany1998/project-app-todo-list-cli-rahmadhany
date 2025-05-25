package model

//Task Model
type Task struct {
	ID          int    `json:"id"`          // Unique ID
	Title       string `json:"title"`       // Task Title
	Description string `json:"description"` // Task Description
	Done        bool   `json:"done"`        // Status whether the task is completed
}
