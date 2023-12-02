package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"unit-test-golang/entity"
	"unit-test-golang/repository"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestTestCategoryService_GetNotFound(t *testing.T) {
	// setup mock
	Id := "1"
	categoryRepository.Mock.On("FindById", Id).Return(nil)

	category, err := categoryService.Get(Id)

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	data := entity.Category{
		Id:   "1",
		Name: "Makanan",
	}
	categoryRepository.Mock.On("FindById", data.Id).Return(data)

	category, err := categoryService.Get(data.Id)

	assert.NotNil(t, category)
	assert.Nil(t, err)
	assert.Equal(t, data.Id, category.Id)
	assert.Equal(t, data.Name, category.Name)
}
