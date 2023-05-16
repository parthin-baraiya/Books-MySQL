// Hello this is model of books and some basic functionalities.

package models

import (
	"github.com/jinzhu/gorm"
	"github.com/parthin-12/movies-mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("id=?", ID).Find(&getBook)
	return &getBook, db
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("id=?", ID).Delete(book)
	return book
}
