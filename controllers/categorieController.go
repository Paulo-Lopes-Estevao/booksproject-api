package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/bookproject/adapter"
	"github.com/Paulo-Lopes-Estevao/bookproject/app/repository"
	"github.com/Paulo-Lopes-Estevao/bookproject/app/usecase"
	"github.com/Paulo-Lopes-Estevao/bookproject/entities"
	"github.com/Paulo-Lopes-Estevao/bookproject/utils/db"
)

func ListCategoryAdapter(w http.ResponseWriter, r *http.Request) {
	categorieAdapter := configCategorie()

	categorie := categorieAdapter.CategorieUseCase.ListCategory()

	value, _ := json.Marshal(categorie)
	w.Write(value)
}

var categorie entities.Categorie

func CreateCategoryAdapter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&categorie)

	categorieAdapter := configCategorie()

	registreCategorie := entities.NewCategorie(categorie.Name)

	categorieAdapter.CategorieUseCase.CreateCategory(registreCategorie)

	fmt.Fprintln(w, "Cadastrado!")
}

func configCategorie() *adapter.CategorieAdapter {
	db := db.ConnectionDB()
	categorieAdapter := adapter.NewCategoryAdapter()
	categorieRepository := repository.CategorieRepository{Db: db}
	categorieAdapter.CategorieUseCase = usecase.CategorieUseCase{CategorieRepository: categorieRepository}
	return categorieAdapter
}
