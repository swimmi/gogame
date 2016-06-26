package game

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"consts"
	"db"
)

func SelectNpc(args []string) int {
	if len(args) != 2 {
		fmt.Println(consts.SELECT_NPC_FORMAT)
	} else {
		cond := bson.M{"name": args[1]}
		if id, err := strconv.Atoi(args[1]); err == nil {
			cond = bson.M{"id": id}
		}
		c := db.Connect("npc")
		err := c.Find(cond).One(&npcGo)
		if err != nil {
			fmt.Println(consts.NNE)
		} else {
			npcGo.ViewInfo()
		}
	}
	return 0
}
