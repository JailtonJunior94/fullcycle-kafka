package models

type Customer struct {
	ID       int64  `json:"id"`
	Document string `json:"document"`
	Name     string `json:"name"`
}
