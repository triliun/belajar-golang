package repository

import "belajar-golang-unit-test/entity"

// Bikin Kontrak
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
