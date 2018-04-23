package main

import (
	"fmt"
	"webdemo/dao"
	"webdemo/model"
)

func main() {
	var u model.User
	u.Name = "yusheng"
	u.Password = "123456"
	u.Age = 20
	u.Sex = false

	fmt.Println(dao.Save(u))
	fmt.Println(dao.Query(1).Name)
	fmt.Println(dao.Delete(2))

	user := dao.Query(3)
	user.Name = "ccnu"
	fmt.Println(dao.Update(user))
}
