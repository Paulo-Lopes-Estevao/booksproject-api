package entities

import "time"

type User struct {
	Name      string `json:"name"`
	Bi        string `json:"idade"`
	Birthdate string `json:"Birthdate"`
	Adress    string `json:"adress"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
