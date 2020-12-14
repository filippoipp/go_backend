package models

import (
	"backend-go/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"time"
)

var db *gorm.DB

type UserModel struct {
	ID          		uint    `gorm:"primary_key"`
	Name				string  `gorm:"column:username"`
	Email       		string  `gorm:"column:email;unique_index"`
	CreatedAt 			time.Time
	Collection      	pq.StringArray `gorm:"column:collection;type:text[]"`
	LentBooks      		pq.StringArray `gorm:"column:lent_books;type:text[]"`
	BorrowedBooks      	pq.StringArray `gorm:"column:borrowed_books;type:text[]"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&UserModel{})
}

func (b *UserModel) CreateBook() *UserModel {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllBooks() []UserModel {
	var Books []UserModel
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*UserModel , *gorm.DB){
	var getBook UserModel
	db:=db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) UserModel {
	var book UserModel
	db.Where("ID = ?", ID).Delete(book)
	return book
}