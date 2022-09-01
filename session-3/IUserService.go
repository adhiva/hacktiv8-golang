package session3

type IUserService interface {
	Register(user *User) string
	GetUser() []*User
}
