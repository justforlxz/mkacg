package models

type Post struct {
	Id    int
	User  *User
	Title string
	Body  string
}
