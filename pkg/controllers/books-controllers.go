package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/parthin-12/movies-mysql/pkg/models"
	"github.com/parthin-12/movies-mysql/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ID := vars["id"]
	bookId, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing")
	}
	bookDetails, _ := models.GetBookByID(bookId)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	utils.ParseBody(r, &book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("Error in parsing")
	}
	book := models.DeleteBook(bookID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateBook models.Book
	utils.ParseBody(r, &updateBook)

	vars := mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("Error in parsing")
	}

	book, db := models.GetBookByID(bookID)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
