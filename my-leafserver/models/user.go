package models

import "gopkg.in/mgo.v2"

const (
	userCollectionName = "user"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// 获取 user 的集合
func getUserCollection() *mgo.Collection {
	return mgSessionDB.C(userCollectionName)
}

// 添加一个用户
func (u *User) AddUser() error {
	c := getUserCollection()
	return c.Insert(u)
}
