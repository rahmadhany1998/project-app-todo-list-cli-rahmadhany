package dto

// CreateTaskRequest is used for input when creating a new task
type CreateTaskRequest struct {
	Title       string // Task Title
	Description string // Task Description
}

// SearchTaskRequest is used to input task searches based on keywords
type SearchTaskRequest struct {
	Keyword string // Search keyword (can be from title or description)
}
