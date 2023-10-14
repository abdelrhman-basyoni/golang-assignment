package domain_Repos

import domain_entities "github.com/abdelrhman-basyoni/golang-assignment/core/domain/entities"

type BookRepository interface {
	Create(name, genre string, Price float32) error
	GetAll() []domain_entities.Book
	GetByID(id string) (*domain_entities.Book, error)
	Update(id string, update map[string]interface{}) error
	Delete(id string) error
}
