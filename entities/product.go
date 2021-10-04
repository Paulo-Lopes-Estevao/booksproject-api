package entities

type Product struct {
	Name     string `json:"name"`
	Edittion string `json:"edition"`
	Data     string `json:"data"`
	Type     Categorie
}
