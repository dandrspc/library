# 📚 Biblioteca CLI

A simple command-line application written in Go to manage a personal library. It allows users to add and list books using JSON file-based storage.

---

## ✨ Features

- Book management via CLI
- JSON file persistence
- Clean Architecture separation of concerns
- Console input validation
- Go modules support
- Modular and extensible codebase

---

## 🧱 Project Structure

```bash
biblioteca/
├── cmd/
│   └── main.go               # Application entry point
├── books/
│   ├── book.go               # Book struct and validations
│   ├── repository.go         # Repository interface
│   └── service.go            # Business logic (service layer)
├── storage/
│   └── json_storage.go       # JSON-based repository implementation
├── internal/
│   └── input.go              # Console input helpers
├── data/
│   └── libros.json           # JSON file to store books (initial content: `[]`)
├── go.mod                    # Go module definition
└── README.md                 # This file
```

---

## 🛠️ Requirements

- Go 1.18 or later
- A `libros.json` file inside the `/data` folder with an initial empty array: `[]`

---

## 🚀 Usage

```bash
# Clone the repository
git clone https://github.com/dandrspc/library.git
cd biblioteca

# Install dependencies
go mod tidy

# Run the application
go run ./cmd
```

---

## 📦 Sample CLI Interaction

```
====== LIBRARY ======

    1. Add book
    2. List all books
    3. Get book by ID
    4. Update book
    5. Delete book
    6. Exit

	Select an option:
```

---

## 🧪 Testing

> Unit tests will be included in a future version.

---

## 🪪 License

This project is licensed under the MIT License. Free for personal and commercial use.

---

## 💡 Author

Developed by Daniel Andrés Palacios Carabalí.
