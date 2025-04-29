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
		t.Errorf("`GetAll` returned incorrect books. Expeted: %v. Got: %v", initialBooks, allBooks)
	}

	// Test an empty file
	cleanupTestFile(t)
	emptyRepo := NewJsonBookRepository(testFile)
	emptyBooks, err := emptyRepo.GetAll(ctx)
	if err != nil {
		t.Fatalf("`GetAll` on empty file returned an error: %v", err)
	}

	if len(emptyBooks) > 0 {
		t.Errorf("`GetAll` on empty file should return an empty file. Got: %+v", emptyBooks)
	}
}

func TestGetById(t *testing.T) {
	initialBooks := []book.Book{
		{ID: "1", Title: "Title One", Author: "Author One", Year: 2000},
		{ID: "2", Title: "Title Two", Author: "Author Two", Year: 2000},
	}
	setupTestFile(t, initialBooks)
	defer cleanupTestFile(t)

	repo := NewJsonBookRepository(testFile)
	ctx := context.Background()

	t.Run("Book Exists", func(t *testing.T) {
		targetID := "2"
		retrievedBook, err := repo.GetById(ctx, targetID)
		if err != nil {
			t.Fatalf("GetByid returned an error for existing ID: %v", err)
		}

		var expectedBook *book.Book
		for i := range initialBooks {
			if initialBooks[i].ID == targetID {
				expectedBook = &initialBooks[i]
				break
			}
		}

		if !reflect.DeepEqual(retrievedBook, expectedBook) {
			t.Errorf("GetbyId returned incorrect book for ID. Expected: %v. Got: %v", expectedBook, retrievedBook)
		}
	})

	t.Run("Book Does Not Exist", func(t *testing.T) {
		nonExistentID := "3"
		retrievedBook, err := repo.GetById(ctx, nonExistentID)
		if err != nil {
			t.Fatalf("GetById returned an unexpected error for non-existent ID: %v", err)
		}
		if retrievedBook != nil {
			t.Errorf("GetById should return nil for non-existent ID '%s'. Got: %+v", nonExistentID, retrievedBook)
		}
	})
}

func TestSaveAll(t *testing.T) {
	booksToSave := []book.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Year: 2000},
		{ID: "2", Title: "Book 2", Author: "Author 2", Year: 2000},
	}

	setupTestFile(t, booksToSave)
	defer cleanupTestFile(t)

	repo := NewJsonBookRepository(testFile)
	ctx := context.Background()

	err := repo.SaveAll(ctx, booksToSave)

	if err != nil {
		t.Fatalf("Save all returned an error: %v", err)
	}

	savedBooks, err := repo.loadBooks()
	if err != nil {
		t.Fatalf("loadbooks returned ")
	}

	if !reflect.DeepEqual(booksToSave, savedBooks) {
		t.Errorf("SaveAll returned wrong books. Expected: %v. Got: %v", booksToSave, savedBooks)
	}

	//Test saving an empty slice
	emptyRepo := NewJsonBookRepository(testFile)
	emptyCollection := []book.Book{}

	err = emptyRepo.SaveAll(context.Background(), emptyCollection)
	if err != nil {
		t.Fatalf("SaveAll on empty slice returned an error: %v", err)
	}
	emptyBooks, err := emptyRepo.loadBooks()
	if len(emptyBooks) > 0 {
		t.Errorf("should return an empty file. Expected: %v, Got: %v", emptyCollection, emptyBooks)
	}
}

func TestUpdate(t *testing.T) {
	initialData := []book.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Year: 2000},
	}
	setupTestFile(t, initialData)
	defer cleanupTestFile(t)

	updatedBook := book.Book{ID: "1", Title: "Book Updated", Author: "Author 1", Year: 2000}

	repo := NewJsonBookRepository(testFile)
	ctx := context.Background()

	err := repo.Update(ctx, updatedBook)
	if err != nil {
		t.Fatalf("Update returned an error: %v", err)
	}

	retrievedBook, err := repo.GetById(ctx, "1")
	if err != nil {
		t.Errorf("Error retrieving book after update: %v", err)
	}

	if !reflect.DeepEqual(*retrievedBook, updatedBook) {
		t.Errorf("Update incorrect book information. Expected: %+v. Got: %+v", updatedBook, *retrievedBook)
	}

	// Update empty book.
	err = repo.Update(ctx, book.Book{})
	if err == nil {
		t.Errorf("Update should have returned an error.")
	}
}

func TestDelete(t *testing.T) {
	initialBooks := []book.Book{
		{ID: "1", Title: "Title One", Author: "Author One", Year: 2000},
		{ID: "2", Title: "Title Two", Author: "Author Two", Year: 2000},
	}
	setupTestFile(t, initialBooks)
	defer cleanupTestFile(t)

	repo := NewJsonBookRepository(testFile)
	ctx := context.Background()

	err := repo.Delete(ctx, "1")
	if err != nil {
		t.Fatalf("Delete returned an error: %v", err)
	}

	deletedBook, err := repo.GetById(ctx, "1")
	if err != nil {
		t.Fatalf("Error retrieving deleted book: %v", err)
	}
	if deletedBook != nil {
		t.Errorf("Book with ID '1' should have been deleted: %+v", deletedBook)
	}

	// Test deleting a non-existent book
	err = repo.Delete(ctx, "3")
	if err == nil {
		t.Errorf("Delete should have returned an error for a non-existent book.")
	}
}
