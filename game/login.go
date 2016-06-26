package game

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"consts"
	"db"
)

func Login(args []string) int {
	if len(args) != 2 {
		fmt.Println(consts.LOGIN_FORMAT)
	} else {
		c := db.Connect("role")
		err := c.Find(bson.M{"name": args[1]}).One(&roleGo)
		if err != nil {
			fmt.Println("登录失败，请重试")
			return 0
		}
		fmt.Println(roleGo.Name, "登录成功")
	}
	return 0
}
