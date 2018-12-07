package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Profile struct {
	ID          int
	User        *User   `orm:"rel(one)"`
	Post        []*Post `orm:"reverse(many)"`
	Control     int
	DisplayName string
	Email       string
}

func (this *Profile) findUserByEmail(email string) (Profile, error) {
	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable("profile")
	var profile Profile
	err := qs.Filter("Email", email).RelatedSel().One(&profile)

	if err == orm.ErrMultiRows {
		fmt.Println("Returned Multi Rows Not One")
		return profile, err
	}

	if err == orm.ErrNoRows {
		fmt.Println("Not row one!")
		return profile, err
	}

	return profile, err
}

func (this *Profile) addUser() error {
	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable("profile")

	var profile Profile
	err := qs.Filter("Email", this.Email).One(&profile)

	if err == orm.ErrNoRows {
		// insert new user
		id, idErr := o.Insert(&profile)
		if idErr != nil {
			fmt.Println("insert row failed, please check database!")
			return idErr
		}

		fmt.Println("Insert success!", id)

		return nil
	}

	fmt.Println("email exist!")
	return err
}
