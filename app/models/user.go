package models

type User struct {
	Username string
	Password string
	Token    string
}

func GetUser(token string) error {
	// err = errors.New("this is error")
	user := &User{
		Username: "abc",
		Password: "pass",
		Token:    "aaa",
	}
	c := Dbm.DB("go-sample").C("user")
	return c.Insert(user)
}
