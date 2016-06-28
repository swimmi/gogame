package models

import (
	"fmt"
	"time"

	"consts"
)

type Role struct {
	Id      int
	Name    string
	Gender  int
	Age     int
	Created time.Time
}

func (role *Role) ViewInfo() {

	fmt.Sprintf(consts.ROLE_INFO, role.Name)
}
