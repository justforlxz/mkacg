package models

type Profile struct {
	ID   int
	User *User `orm:"reverse(one)"`
}
