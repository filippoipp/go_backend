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
	Title     			string    	`gorm:"size:255;not" json:"title"`
	Pages   			string    	`gorm:"not null;" json:"pages"`
	Owner  				User		`json:"owner"`
	OwnerID 			uint		`sql:"type:int REFERENCES users(id)" json:"owner_id"`
	CreatedAt 			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 			time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Book) Prepare() {
	b.ID = 0
	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
	b.Pages = html.EscapeString(strings.TrimSpace(b.Pages))
	b.Owner = User{}
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
		err = db.Debug().Model(&User{}).Where("id = ?", b.OwnerID).Take(&b.Pages).Error
		if err != nil {
			return &Book{}, err
		}
	}
	return b, nil
}