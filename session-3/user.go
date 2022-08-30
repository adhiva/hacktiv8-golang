package session3

type User struct {
	Name string
	Age  string
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
	return "Yeay berhasil ditambahkan " + u.Name
}

func (us *UserService) GetUser() []*User {
	return us.User
}
