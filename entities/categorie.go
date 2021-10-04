package entities

import "time"

type Categorie struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
