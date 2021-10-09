package usecase

import (
	"github.com/Paulo-Lopes-Estevao/bookproject/app/repository"
	"github.com/Paulo-Lopes-Estevao/bookproject/entities"
)

type CategorieUseCase struct {
	CategorieRepository repository.CategorieRepository
}

func (usecase CategorieUseCase) ListCategory() entities.Categorie {
	categories := usecase.CategorieRepository.List()

	return categories
}

func (usecase CategorieUseCase) CreateCategory(categorie *entities.Categorie) (*entities.Categorie, error) {
	categories, err := usecase.CategorieRepository.Insert(categorie)

	if err != nil {
		return categories, err
	}

	return categories, nil

}

func (usecase CategorieUseCase) FindCategory(id string) *entities.Categorie {
	categories, _ := usecase.CategorieRepository.Find(id)
	return categories
}
