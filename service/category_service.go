package service

import (
	"errors"
	"unit-test-golang/entity"
	"unit-test-golang/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(Id string) (*entity.Category, error) {
	category := service.Repository.FindById(Id)
	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}
