package books

import (
	"github.com/go-mongo-server/config"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	books, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500) + err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", books)
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500) + err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "show.gohtml", book)
}

func Create(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := CreateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500) + err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", book)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500) + err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", book)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500) + err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "updated.gohtml", book)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500) + err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
