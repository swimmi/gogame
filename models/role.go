package models

import (
	"fmt"

	"consts"
)

type Role struct {
	Id     int
	Name   string
	Gender int
	Age    int
}

func (role *Role) ViewInfo() {

	fmt.Sprintf(consts.ROLE_INFO, role.Name)
}
