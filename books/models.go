package books

import (
	"errors"
	"github.com/go-mongo-server/config"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type Book struct {
	Isbn string
	Title string
	Author string
	Price float32
}

func AllBooks() ([]Book, error) {
	books := []Book{}
	err := config.Books.Find(bson.M{}).All(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func OneBook(r *http.Request) (Book, error) {
	book := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return book, errors.New("400: Bad Request")
	}

	err := config.Books.Find(bson.M{ "isbn": isbn }).One(&book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400: Bad Request")
	}

	err := config.Books.Remove(bson.M{ "isbn": isbn })
	if err != nil {
		return errors.New("500: Internal Server Error")
	}
	return nil
}

func CreateBook(r *http.Request) (Book, error) {
	// Get Form Values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// Validate Form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		return book, errors.New("400: All fields are required")
	}

	// Convert Form Values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return book, errors.New("400: Not acceptable. Price should be float")
	}
	book.Price = float32(f64)

	// Insert to DB
	err = config.Books.Insert(book)
	if err != nil {
		return book, errors.New("500 Internal Server error" + err.Error())
	}
	return book, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	// Get Form Values
	book := Book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// Validate Form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		return book, errors.New("400: All fields are required")
	}

	// Convert Form Values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return book, errors.New("400: Not acceptable. Price should be float")
	}
	book.Price = float32(f64)

	// Insert to DB
	err = config.Books.Update(bson.M{ "isbn": book.Isbn }, &book)
	if err != nil {
		return book, errors.New("500 Internal Server error" + err.Error())
	}
	return book, nil
}