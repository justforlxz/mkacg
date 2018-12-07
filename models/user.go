package models

type User struct {
	ID       int
	Salt     string
	Password string
	Profile  *Profile `orm:"reverse(one)"`
}
