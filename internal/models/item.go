package models

type Item struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"em,omitempty"`
	Description string `json:"pp,omitempty"`
	Stock       int64  `json:"ps,omitempty"`
}
