package models

type Categories struct {
	ID   int
	Post []*Post `orm:"reverse(many)"`
}
