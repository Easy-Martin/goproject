package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

const USER_MODEL_TABLE_NAME string = "user"

type User struct {
	Id       string `orm:"pk"`
	Name     string
	Password string
}

func GetUsers() ([]*User) {
	o := orm.NewOrm()
	o.Using("default")
	var users []*User
	num, err := o.QueryTable(USER_MODEL_TABLE_NAME).All(&users)
	CheckError(err)
	fmt.Printf("Returned Rows Num: %s", num)
	return users
}
