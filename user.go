package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Usernmae string `json:"username"`
	Passwrod string `json:"password"`
}
