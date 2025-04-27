package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/dandrspc/library/book"
	"github.com/dandrspc/library/storage"
)

var jsonStorage book.BookRepository

func main() {
	// ctx := context.Background()
	storagePath := "data/books.json"
	jsonStorage = storage.NewJsonBookRepository(storagePath)
}

func InitBookDataFile(filepath string) error {
	_, err := os.Stat(filepath)
	if errors.Is(err, fs.ErrNotExist) {
		os.WriteFile(filepath, []byte("[]"), 0644)
	} else if err != nil {
		return fmt.Errorf("error chechig file: %w", err)
	}
	return nil
}
