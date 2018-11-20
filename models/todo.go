package models

// Todo .
type Todo struct {
	ID     string `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	Done   bool   `json:"done,omitempty"`
	UserID string `json:"user,omitempty"`
}
