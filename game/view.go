package game

import (
	"fmt"
	"strconv"

	"consts"
	"db"
	"utils"
)

//查看角色
func ViewRole(args []string) int {
	fmt.Println(consts.ROLE_INFO)
	return 0
}

//查看NPC
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

//查看物品
func ViewRoleItem(args []string) int {

	switch len(args) {
	case 1:
		if role_items, err := db.GetRoleItemList(roleGo.Id); err == nil {
			fmt.Printf(consts.ROLE_ITEM_HEAD, consts.ROLE_ITEM_LIST)
			for num, role_item := range role_items {
				fmt.Printf(consts.ROLE_ITEM, num+1, role_item.Id, role_item.ItemId, role_item.Count)
			}
			fmt.Printf(consts.ROLE_ITEM_FOOT)
		}
	case 2:
		id, err := strconv.Atoi(args[1])
		if err != nil {
			if item, err := db.GetItemByName(args[1]); err == nil {
				item.ViewInfo()
			}
		} else {
			if item, err := db.GetItemById(id); err == nil {
				item.ViewInfo()
			}
		}
	default:

	}
	return 0
}
