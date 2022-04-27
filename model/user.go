package model

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
