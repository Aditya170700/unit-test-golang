package repository

import "unit-test-golang/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
