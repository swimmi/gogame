package game

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"consts"
	"models"
	"utils"
)

func ViewRole(args []string) int {
	fmt.Println(consts.ROLE_INFO)
	return 0
}

func ViewNpc(args []string) int {
	switch len(args) {
	case 1:
		if npcs, err := db.GetNpcList(); err == nil {
			fmt.Printf(consts.NPC_HEAD, consts.NPC_LIST)
			for num, npc := range npcs {
				fmt.Printf(consts.NPC_ITEM, num+1, npc.Id, npc.Name, utils.PsGender(npc.Gender), npc.Age)
			}
			fmt.Printf(consts.NPC_FOOT)
		}
	case 2:
		id, err := strconv.Atoi(args[1])
		if err != nil {
			if npc, err := db.GetNpcByName(args[1]); err == nil {
				npc.ViewInfo()
			}
		} else {
			if npc, err := db.GetNpcById(id); err == nil {
				npc.ViewInfo()
			}
		}
	default:

	}
	return 0
}

func ViewItem(args []string) int {

	c := db.Connect("role_item")
	switch len(args) {
	case 1:
		npcs := []models.Npc{}
		err := c.Find(nil).Sort("id").All(&npcs)
		if err != nil {
		} else {
			fmt.Printf(consts.NPC_HEAD, consts.NPC_LIST)
			for num, npc := range npcs {
				fmt.Printf(consts.NPC_ITEM, num+1, npc.Id, npc.Name, utils.PsGender(npc.Gender), npc.Age)
			}
			fmt.Printf(consts.NPC_FOOT)
		}
	case 2:

		npc := new(models.Npc)
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(consts.EID)
		} else {
			err = c.Find(bson.M{"id": id}).One(&npc)
			if err != nil {
				fmt.Println(consts.NNE)
			} else {
				npc.ViewInfo()
			}
		}
	default:

	}
	return 0
}
