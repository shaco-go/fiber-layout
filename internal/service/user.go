package service

import "githut.com/shaco-go/fiber-kit/internal/repo"

func NewUser(ur *repo.User) *User {
	return &User{ur: ur}
}

type User struct {
	ur *repo.User
}

func (u *User) GetUser() string {
	return u.ur.GetUser()
}
