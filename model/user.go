package model

import (
	"fmt"

	"github.com/dev-sota/go-handson/database"
)

// ドメイン、業務の登場人物
type User struct {
	ID    int64
	Name  string
	Email string
	Age   int
}

func NewUser(name, email string, age int) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

// データ操作 Data access object dao, SQL
func (u *User) FindByEmail(email string) error {
	return database.Get().
		Where("email = ?", email).
		Find(u).
		Error
}

func (u *User) Find(id int64) error {
	return database.Get().First(u, id).Error
}

func (u *User) Create() error {
	return database.Get().Create(u).Error
}

func (u *User) Update() error {
	return database.Get().Save(u).Error
}

func (u *User) Delete() error {
	return database.Get().Delete(u).Error
}

// ビジネスロジック, 業務ルール
func (u *User) Signup(name, email string, age int) error {
	if err := u.FindByEmail(email); err != nil {
		return err
	}
	if u.ID != 0 {
		return fmt.Errorf("already registerd")
	}

	if name == "ito" {
		fmt.Println("クーポン発行")
	}
	if age < 20 {
		fmt.Println("クーポン発行")
	}
	// if birthDayMonth == time.Now().Month() {
	// 	fmt.Println("クーポン発行")
	// }

	u.Name = name
	u.Email = email
	u.Age = age
	return u.Create()
}
