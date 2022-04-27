package dao

import (
	"github.com/dev-sota/go-handson/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

// データ操作 Data access object dao, SQL
func (u *User) FindByEmail(email string) (*model.User, error) {
	var m model.User
	err := u.db.
		Where("email = ?", email).
		Find(&m).
		Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (u *User) Find(id int64) (*model.User, error) {
	var m model.User
	err := u.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (u *User) Create(name, email string, age int) (*model.User, error) {
	m := model.NewUser(name, email, age)
	err := u.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
