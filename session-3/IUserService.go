package session3

type IUserService interface {
	Register(*User) string
	GetUser() []*User
}
