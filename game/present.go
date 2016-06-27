package game

import (
	"fmt"
	"strconv"

	"consts"
	"db"
	"models"
)

func PresentNpcItem(args []string) int {

	count := 1
	switch len(args) {
	case 3:
		break
	case 4:
		count, _ = strconv.Atoi(args[3])
		break
	default:
		fmt.Println(consts.PRESENT_FORMAT)
		return 0
	}
	npc := models.Npc{}
	role_item := models.RoleItem{}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		npc, err = db.GetNpcByName(args[1])
	} else {
		npc, err = db.GetNpcById(id)
	}
	id, err = strconv.Atoi(args[2])
	if err != nil {
		role_item, err = db.GetRoleItemByName(args[2])
	} else {
		role_item, err = db.GetRoleItemById(id)
	}
	db.PresentNpcItem(npc, role_item, count)
	return 0
}
