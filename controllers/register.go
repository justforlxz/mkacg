package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {

}

func (this *RegisterController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")

	fmt.Println(uname, pwd)
}
