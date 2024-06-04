package models

type (
	Seller struct {
		ID    int    `json:"id"`
		Name  string `json:"name,omitempty"`
		Phone string `json:"phone,omitempty"`
	}
)
