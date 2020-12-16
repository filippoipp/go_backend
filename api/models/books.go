package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	ID      			uint64    	`gorm:"primary_key;auto_increment" json:"id"`
	Title     			string    	`gorm:"size:255;not null" json:"title"`
	Pages   			string    	`gorm:"not null;" json:"pages"`
	LoggedUserID 		uint64		`gorm:"not null" json:"logged_user_id"`
	ToUserID			uint64		`gorm:"not null" json:"to_user_id"`
	CreatedAt 			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Book) Prepare() {
	b.ID = 0
	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
	b.Pages = html.EscapeString(strings.TrimSpace(b.Pages))
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Book) Validate() error {

	if b.Title == "" {
		return errors.New("Required Title")
	}
	return nil
}

func (b *Book) SaveBook(db *gorm.DB) (*Book, error) {
	var err error
	err = db.Debug().Model(&Book{}).Create(&b).Error
	if err != nil {
		return &Book{}, err
	}
	if b.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?").Error
		if err != nil {
			return &Book{}, err
		}
	}
	return b, nil
}

func (b *Book) UpdateABook(db *gorm.DB, touser uint64, bookid uint64) (*Book, error) {
	err := db.Debug().Model(&Book{}).Where("id = ?", bookid).Update("to_user_id", touser).Error
	if err != nil {
		return &Book{}, err
	}

	return b, nil
}