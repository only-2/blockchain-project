package apiControl

import (
	"github.com/astaxie/beego"
)

type Hash_control struct {
	beego.Controller
}

func (this *Hash_control) Hashget() {
	this.TplName = "hash.html"
}
