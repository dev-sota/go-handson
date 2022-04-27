package usecase

import (
	"fmt"

	"github.com/dev-sota/go-handson/dao"
	"github.com/dev-sota/go-handson/model"
	"github.com/dev-sota/go-handson/repository"
	"gorm.io/gorm"
)

type User struct {
	du repository.User
}

func NewUser(db *gorm.DB) *User {
	return &User{
		du: dao.NewUser(db),
	}
}

func (u *User) Signup(name, email string, age int) (*model.User, error) {
	m, err := u.du.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if m.ID != 0 {
		return nil, fmt.Errorf("already registerd")
	}

	// ビジネスロジック, 業務ルール
	if age < 20 {
		fmt.Println("クーポン発行")
	}

	m, err = u.du.Create(name, email, age)
	if err != nil {
		return nil, err
	}

	return m, nil
}
