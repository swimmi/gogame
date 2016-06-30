package models

import (
	"fmt"

	"consts"
)

type Site struct {
	Id    int
	Name  string
	Goods []SiteGoods
}

func (site *Site) ViewInfo() {
	fmt.Printf(consts.SITE_INFO, site.Name)
}
