package models

type User struct {
	ID      int
	Name    string
	Profile *Profile `orm:"rel(one)"`
	Post    []*Post  `orm:"reverse(many)"`
}
