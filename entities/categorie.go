package entities

import "time"

type Categorie struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Status  bool      `json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func NewCategorie(name string) *Categorie {
	categorie := &Categorie{
		Name:    name,
		Status:  true,
		Created: time.Now(),
		Updated: time.Now(),
	}

	return categorie
}
