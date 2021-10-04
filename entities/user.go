package entities

type User struct {
	Nome      string `json:"nome"`
	Bi        string `json:"idade"`
	Birthdate string `json:"Birthdate"`
	Morada    string `json:"morada"`
}
