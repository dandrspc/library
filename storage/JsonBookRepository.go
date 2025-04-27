package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dandrspc/library/book"
)

type JsonBookRepository struct {
	filepath string
}

func NewJsonBookRepository(filepath string) *JsonBookRepository {
	repo := &JsonBookRepository{
		filepath: filepath,
	}
	repo.ensureFileExists()
	return repo
}

func (r *JsonBookRepository) ensureFileExists() {
	if _, err := os.Stat(r.filepath); os.IsNotExist(err) {
		file, err := os.Create(r.filepath)
		if err != nil {
			fmt.Printf("error creating file: %v", err.Error())
			return
		}
		defer file.Close()

		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Printf("error writing initial JSON to file: %v", err.Error())
		}
	}
}

func (r *JsonBookRepository) Create(cxt context.Context, b book.Book) (book.Book, error) {
	books, err := r.loadBooks()
	if err != nil {
		return book.Book{}, err 
	}

	books = append(books, b)

	if err:= r.saveBooks(books); err != nil {
		return book.Book{}, nil
	}
	return b, nil
}

func (r *JsonBookRepository) GetById(ctx context.Context, id string) (*book.Book, error) {
	books, err := r.loadBooks()
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func (r *JsonBookRepository) Update(ctx context.Context, updatedBook book.Book) error {
	books, err := r.loadBooks()
	if err != nil {
		return err
	}

	updated := false
	for i, book := range books {
		if book.ID == updatedBook.ID {
			books[i] = updatedBook
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("book not found")
	}
	
	return r.saveBooks(books)
}

func (r *JsonBookRepository) Delete(ctx context.Context, id string) error {
	books, err := r.loadBooks()
	if err != nil {
		return err
	}

	var updatedBooks []book.Book
	deleted := false
	for _, book := range books {
		if book.ID != id {
			updatedBooks = append(updatedBooks, book)
		} else {
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("failed to delete the book with id: %s", id)
	}

	return r.saveBooks(updatedBooks)
}

func (r *JsonBookRepository) GetAll(ctx context.Context) ([]book.Book, error) {
	return r.loadBooks()
}

func (r *JsonBookRepository) SaveAll(ctx context.Context, books []book.Book) error {
	return r.saveBooks(books)
}

func (r *JsonBookRepository) loadBooks() ([]book.Book, error) {
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return []book.Book{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var books []book.Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, fmt.Errorf("failed to unmarshall file: %w", err)
	}
	return books, nil
}

func (r *JsonBookRepository) saveBooks(books []book.Book) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal books: %w", err)
	}
	return os.WriteFile(r.filepath, data, 0644)
}

