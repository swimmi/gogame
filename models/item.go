package models

import (
	"fmt"
)

type Item struct {
	Id          int
	Name        string
	Type        string
	Desc        string
	MaxCount    int
	FavourValue int
}

func (item *Item) ViewInfo() {
	fmt.Println(item.Name, item.Desc)
}
