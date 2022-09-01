package session3

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type UserService struct {
	User []*User
}

func NewUserService(db []*User) IUserService {
	return &UserService{
		User: db,
	}
}

func (us *UserService) Register(u *User) string {
	us.User = append(us.User, u)
	fmt.Println("User : ", u.Name)
	return "Yeay berhasil ditambahkan " + u.Name
}

func (us *UserService) GetUser() []*User {
	return us.User
}
