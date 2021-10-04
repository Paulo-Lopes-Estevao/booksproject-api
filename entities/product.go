package entities

import "time"

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Edittion    string `json:"edition"`
	Releasedata string `json:"releasedata"`
	Type        Categorie
	Status      string `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
