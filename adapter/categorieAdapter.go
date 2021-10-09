package adapter

import (
	"github.com/Paulo-Lopes-Estevao/bookproject/app/usecase"
)

type CategorieAdapter struct {
	CategorieUseCase usecase.CategorieUseCase
}

func NewCategoryAdapter() *CategorieAdapter {
	return &CategorieAdapter{}
}
