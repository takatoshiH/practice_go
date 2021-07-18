package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct {
	//文字列展開したいときはこんな書き方(JS的な)
	// ID: 1 的なjsonで返すことができる
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	NumberOfPages int     `json:"numberofpages"`
	Author        *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

// GET All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// jsonを返します
	w.Header().Set("Content-Type", "application/json")

	// booksをjsonEncodeして返している
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func getLatestBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books[len(books)-1])
}

// 書籍の登録
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func main() {

	// ルーターのイニシャライズ
	r := mux.NewRouter()

	// モックデータの作成
	books = append(books, Book{ID: "1", Title: "Book one", NumberOfPages: 100, Author: &Author{FirstName: "Philip", LastName: "Williams"}})
	books = append(books, Book{ID: "2", Title: "Book Two", NumberOfPages: 999, Author: &Author{FirstName: "John", LastName: "Johnson"}})

	// エンドポイント
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/latest_book", getLatestBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Method("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
