package models

type Category struct {
	CategoryID        int    `json:"category_id"`
	UserID            int    `json:"user_id"`
	CategoryName      string `json:"category_name"`
	CreatedTimestamp  string `json:"created_timestamp"`
	ModifiedTimestamp string `json:"modified_timestamp"`
}
