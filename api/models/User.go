package models

import (
	"errors"
	"time"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        		uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	Name  			string    	`gorm:"size:255;not null;unique" json:"name"`
	Email     		string    	`gorm:"size:100;not null;unique" json:"email"`
	CreatedAt 		time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Collection  	[]string  	`json: "collection"`
	LentBooks  		[]string  	`json: "lent_books"`
	BorrowedBooks  	[]string  	`json: "borrowed_books"`
}

func (u *User) Prepare() {
	u.ID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	if u.Name == "" {
		return errors.New("Nome obrigatorio")
	}
	if u.Email == "" {
		return errors.New("Email obrigatório")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Email inválido")
	}
	return nil
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("Usuário não encontrado")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":  			u.Name,
			"email":     		u.Email,
			"collection": 		u.Collection,
			"lent_books": 		u.LentBooks,
			"borrowed_books": 	u.BorrowedBooks,
			"update_at": 		time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	return u, nil
}