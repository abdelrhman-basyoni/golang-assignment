package domain

import (
	"database/sql"

	domain_entities "github.com/abdelrhman-basyoni/golang-assignment/core/domain/entities"
	domain_Repos "github.com/abdelrhman-basyoni/golang-assignment/core/domain/repositories"
	imp_repo "github.com/abdelrhman-basyoni/golang-assignment/core/implementation/repositories"
)

type BookUseCases struct {
	bookRepo domain_Repos.BookRepository
}

func NewBookUseCases(db *sql.DB) *BookUseCases {
	repo := imp_repo.NewBookSqlRepo(db)

	return &BookUseCases{bookRepo: repo}
}

func (buc *BookUseCases) Create(name, genre string, price float32) error {

	return buc.bookRepo.Create(name, genre, price)
}

func (buc *BookUseCases) GetAll() []domain_entities.Book {
	return buc.bookRepo.GetAll()
}

func (buc *BookUseCases) GetByID(id string) (*domain_entities.Book, error) {

	return buc.bookRepo.GetByID(id)
}

func (buc *BookUseCases) Update(id string, update domain_entities.Book) error {
	return buc.bookRepo.Update(id, update)
}

func (buc *BookUseCases) Delete(id string) error {
	return buc.bookRepo.Delete(id)
}
