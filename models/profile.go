package models

type Profile struct {
	Id   int
	User *User `orm:"reverse(one)"`
}
