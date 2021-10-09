package repository

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/bookproject/entities"
	"github.com/jinzhu/gorm"
)

type CategorieRepositoryInterface interface {
	List() *entities.Categorie
	Insert(categorie *entities.Categorie) (*entities.Categorie, error)
	Find(id string) (*entities.Categorie, error)
}

type CategorieRepository struct {
	Db *gorm.DB
}

var categories entities.Categorie

func (repository CategorieRepository) List() entities.Categorie {
	repository.Db.Table("tb_categorie").Scan(&categories)

	return categories
}

func (repository CategorieRepository) Insert(categorie *entities.Categorie) (*entities.Categorie, error) {
	err := repository.Db.Table("tb_categorie").Create(categorie).Error

	if err != nil {
		return nil, err
	}
	return categorie, nil
}

func (repository CategorieRepository) Find(id string) (*entities.Categorie, error) {

	repository.Db.Find(&categories, "id=?", id)

	if categories.Id == "" {
		return nil, fmt.Errorf("Categories does not exist")
	}

	return &categories, nil
}
