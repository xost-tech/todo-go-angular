package models

type Task struct {
	TaskID            int    `json:"task_id"`
	UserID            int    `json:"user_id"`
	CategoryID        int    `json:"category_id"`
	TaskName          string `json:"task_name"`
	TaskDescription   string `json:"task_description"`
	TaskStatus        int    `json:"task_status"`
	DueDate           string `json:"due_date"`
	CreatedTimestamp  string `json:"created_timestamp"`
	ModifiedTimestamp string `json:"modified_timestamp"`
}

const (
	Pending    = 0
	InProgress = 1
	Completed  = 2
)
