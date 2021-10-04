package repository

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/bookproject/entities"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	List() (*entities.Product, error)
	Insert(product *entities.Product) (*entities.Product, error)
	Find(id string) (*entities.Product, error)
}

type ProductRepository struct {
	Db *gorm.DB
}

var products entities.Product

func (repository ProductRepository) List() (*entities.Product, error) {
	repository.Db.Raw("select *from products").Scan(&products)
	return &products, nil
}

func (repository ProductRepository) Insert(product *entities.Product) (*entities.Product, error) {
	err := repository.Db.Create(product).Error

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repository ProductRepository) Find(id string) (*entities.Product, error) {

	repository.Db.Find(&products, "id=?", id)

	if products.Id == "" {
		return nil, fmt.Errorf("Products does not exist")
	}

	return &products, nil
}
