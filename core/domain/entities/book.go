package domain_entities

type Book struct {
	ID    string  `json:"_id"` // defining id as string so it can work with any database not just sql types
	Name  string  `jason:"name" validate:"required"`
	Genre string  `jason:"genre" validate:"required"`
	Price float32 `jason:"price" validate:"required"`
}
