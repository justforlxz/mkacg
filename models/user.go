package models

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`
}
