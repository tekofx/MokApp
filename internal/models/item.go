package models

type Item struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Stock       int64   `json:"stock,omitempty"`
	Price       float32 `json:"price"`
}
