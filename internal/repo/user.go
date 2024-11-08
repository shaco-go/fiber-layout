package repo

func NewUser() *User {
	return &User{}
}

type User struct {
}

func (u *User) GetUser() string {
	return "hello word"
}
