package book

import (
	"context"
)

// BookRepository defines CRUD operations for managing books.
type BookRepository interface {
	// Create adds a new book to the repository.
	Create(ctx context.Context, book Book) (Book, error)

	// GetAll retrieves all books from the repository.
	// TODO: Add pagination e.g. limit, offset
	GetAll(ctx context.Context) ([]Book, error)

	// SaveAll saves a slice of books to a json file
	SaveAll(ctx context.Context, books []Book) error

	// GetById retrieves a book by its ID.
	GetById(ctx context.Context, id string) (*Book, error)

	// Update modifies an existing book in the repository.
	Update(ctx context.Context, book Book) error

	// Delete removes a book by its ID.
	Delete(ctx context.Context, id string) error
}
