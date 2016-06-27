package models

import (
	"fmt"

	"consts"
)

type Npc struct {
	Id       int
	Name     string
	Nickname string
	Gender   int
	Age      int
	Desc     string
	Title    string
	Favour   int
}

func (npc *Npc) ViewInfo() {
	fmt.Println(npc.Name, npc.Desc)
}

func (npc *Npc) ThankPresent(favour int) {
	fmt.Println(npc.Name, consts.THANKS)
}
