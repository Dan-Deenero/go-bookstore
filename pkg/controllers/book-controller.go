package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dan-deenero/go-bookstore/pkg/models"
	"github.com/dan-deenero/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newbooks := models.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
        http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
