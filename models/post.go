package models

type Post struct {
	ID         int
	User       *User `orm:"rel(fk)"`
	Title      string
	Body       string
	Categories []*Categories `orm:"rel(m2m)"`
}
