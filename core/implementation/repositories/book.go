package imp_repo

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	domain_entities "github.com/abdelrhman-basyoni/golang-assignment/core/domain/entities"
)

type BookRepoSql struct {
	db *sql.DB
}

func NewBookSqlRepo(db *sql.DB) *BookRepoSql {
	if db == nil {
		panic("Missing Database for BookRepo")
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY,
		name VARCHAR(255) UNIQUE,
		genre VARCHAR(255),
		price  REAL
	)`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic("Error creating table 'books'")
	}

	return &BookRepoSql{db: db}
}

func (b *BookRepoSql) Create(name, genre string, Price float32) error {
	insertSQL := `
	INSERT INTO books (name, genre, price)
	VALUES ($1, $2, $3)
`
	_, err := b.db.Exec(insertSQL, name, genre, Price)

	return err
}

func (b *BookRepoSql) GetAll() []domain_entities.Book {
	selectSQL := `
	SELECT id, name, genre, price
	FROM books
`

	var books []domain_entities.Book
	rows, err := b.db.Query(selectSQL)
	if err != nil {
		return books
	}
	defer rows.Close()

	var idInt int
	for rows.Next() {
		var book domain_entities.Book

		if err := rows.Scan(&idInt, &book.Name, &book.Genre, &book.Price); err != nil {
			fmt.Println(err)
			return books
		}
		fmt.Println(book)
		book.ID = strconv.Itoa(idInt)
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return books
	}

	return books
}

func (b *BookRepoSql) GetByID(id string) (*domain_entities.Book, error) {

	selectQuery := `SELECT * FROM books WHERE id = $1`

	row := b.db.QueryRow(selectQuery, id)

	var book domain_entities.Book
	if err := row.Scan(&book.ID, &book.Name, &book.Genre, &book.Price); err != nil {
		return nil, err
	}

	return &book, nil

}

func (b *BookRepoSql) Update(id string, update map[string]interface{}) error {

	var placeholders []string
	var values []interface{}
	index := 1
	for key, value := range update {
		placeholders = append(placeholders, fmt.Sprintf("%s = $%d", key, index))
		values = append(values, value)
		index++
	}
	// Combine placeholders into a comma-separated string
	setClause := strings.Join(placeholders, ", ")

	// Define the SQL update statement
	updateSQL := fmt.Sprintf(`
		UPDATE books
		SET %s
		WHERE id = $%d
	`, setClause, index)

	values = append(values, id)

	_, err := b.db.Exec(updateSQL, values...)

	return err
}

func (b *BookRepoSql) Delete(id string) error {

	deleteStatement := `DELETE books where id = %1`

	_, err := b.db.Exec(deleteStatement, id)

	return err
}
