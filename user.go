package gitlab

type User struct {
	Id    int
	Name  string
	Token string
	Url   string
}

func NewUser(id int, name, token, url string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Token: token,
		Url:   url,
	}
}
