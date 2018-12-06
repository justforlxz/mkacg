package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	ID         int
	User       *User `orm:"rel(fk)"`
	Title      string
	Body       string
	Categories []*Categories `orm:"rel(m2m)"`
	Created    time.Time     `orm:"auto_new_add:type(datetime)"`
	Updated    time.Time     `orm:"auto_now;type(datetime)"`
}

func AddPost(title, content string) error {
	o := orm.NewOrm()

	post := &Post{
		Title:   title,
		Body:    content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(post)

	return err
}
