package usecase

import (
	"github.com/Paulo-Lopes-Estevao/bookproject/app/repository"
	"github.com/Paulo-Lopes-Estevao/bookproject/entities"
)

type ProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func (usecase ProductUseCase) ListProdct() *entities.Product {
	products, _ := usecase.productRepository.List()
	return products
}

func (usecase ProductUseCase) CreateProduct(product *entities.Product) (*entities.Product, error) {
	products, err := usecase.productRepository.Insert(product)

	if err != nil {
		return products, err
	}

	return products, nil

}

func (usecase ProductUseCase) FindProdct(id string) *entities.Product {
	products, _ := usecase.productRepository.Find(id)
	return products
}
