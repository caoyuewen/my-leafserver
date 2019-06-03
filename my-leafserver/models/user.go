package models

const (
	userCollectionName = "user"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// 获取 user 的集合

// 添加一个用户
func (u *User) AddUser() error {
	s, c := connection(DBName, userCollectionName)
	defer s.Close()
	return c.Insert(u)
}
