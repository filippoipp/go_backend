package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        		uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	Name  			string    	`gorm:"size:255;not null" json:"name"`
	Email     		string    	`gorm:"size:100;not null;unique" json:"email"`
	CreatedAt 		time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Collection      []Book		`gorm:"foreignkey:LoggedUserID" json:"collection"`
	LentBooks      	[]Book		`gorm:"foreignkey:LoggedUserID" json:"lent_books"`
	BorrowedBooks   []Book		`gorm:"foreignkey:ToUserID" json:"borrowed_books"`
}

func (u *User) Prepare() {
	u.Name = 			html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = 			html.EscapeString(strings.TrimSpace(u.Email))
	u.Collection = 		[]Book{}
	u.LentBooks = 		[]Book{}
	u.BorrowedBooks = 	[]Book{}
	u.CreatedAt = 		time.Now()
	u.UpdatedAt = 		time.Now()
}

func (u *User) Validate(action string) error {
	if u.Name == "" {
		return errors.New("Required Name")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Preload("Collection").Preload("BorrowedBooks").Preload("LentBooks", "to_user_id != 0").Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":  		u.Name,
			"email":     	u.Email,
			"update_at": 	time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}