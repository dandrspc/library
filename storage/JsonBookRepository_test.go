package storage

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/dandrspc/library/book"
)

const testFile = "test_books.json"

func setupTestFile(t *testing.T, initialData []book.Book) {
	data, err := json.MarshalIndent(initialData, "", "  ")
	if err != nil {
		t.Fatalf("error marshalling initial data: %v", err)
	}
	err = os.WriteFile(testFile, data, 0644)
	if err != nil {
		t.Fatalf("error writing test file: %v", err)
	}
}

func cleanupTestFile(t *testing.T) {
	err := os.Remove(testFile)
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("error cleaning the test file: %v", err)
	}
}

func TestCreate(t *testing.T) {
	cleanupTestFile(t)
	defer cleanupTestFile(t)

	repo := NewJsonBookRepository(testFile)
	ctx := context.Background()
	newBook := book.Book{ID: "4", Title: "New Book", Author: "New Author", Year: 2000}

	createdBook, err := repo.Create(ctx, newBook)
	if err != nil {
		t.Fatalf("create returned an error: %v", err)
	}

	if !reflect.DeepEqual(createdBook, newBook) {
		t.Errorf("create returned the wrong book. Expected: %v, Got: %v", newBook, createdBook)
	}
}

func TestGetAll(t *testing.T) {
	initialBooks := []book.Book{
		{ID: "1", Title: "Title One", Author: "Author One", Year: 2000},
		{ID: "2", Title: "Title Two", Author: "Author Two", Year: 2000},
	}
	setupTestFile(t, initialBooks)
	defer cleanupTestFile(t)

	repo := NewJsonBookRepository(testFile)
	ctx := context.Background()

	allBooks, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatalf("`GetAll` returned an error: %v", err)
	}

	if !reflect.DeepEqual(allBooks, initialBooks) {
		t.Fatalf("`GetAll` returned incorrect books. Expeted: %v. Got: %v", initialBooks, allBooks)
	}

	// Test an empty file
	cleanupTestFile(t)
	emptyRepo := NewJsonBookRepository(testFile)
	emptyBooks, err := emptyRepo.GetAll(ctx)
	if err != nil {
		t.Fatalf("`GetAll` on empty file returned an error: %v", err)
	}

	if len(emptyBooks) > 0 {
		t.Fatalf("`GetAll` on empty file should return an empty file. Got: %+v", emptyBooks)
	}

}