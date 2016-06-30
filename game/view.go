package game

import (
	"fmt"
	"strconv"

	"consts"
	"db"
	"utils"
)

//查看
func View(args []string) int {
	return 0
}

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
			fmt.Println(consts.NPC_LIST)
			for num, npc := range npcs {
				fmt.Printf(consts.NPC_ITEM, num+1, npc.Id, npc.Name, utils.PsGender(npc.Gender), npc.Age, npc.Favour)
			}
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
		if role_items, err := db.GetRoleItemList(); err == nil {
			fmt.Println(consts.ROLE_ITEM_LIST)
			for num, role_item := range role_items {
				if item, err := db.GetItemById(role_item.Item); err == nil {
					fmt.Printf(consts.ROLE_ITEM, num+1, role_item.Id, item.Name, role_item.Count)
				}
			}
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

//查看农场
func ViewFarm(args []string) int {
	if len(args) == 1 {
		if farm_units, err := db.GetFarmUnitList(); err == nil {
			fmt.Println(consts.FARM_UNIT_LIST)
			for num, farm_unit := range farm_units {
				if crop, err := db.GetCropById(farm_unit.Crop); err == nil {
					fmt.Printf(consts.FARM_UNIT, num+1, farm_unit.Id, crop.Name, farm_unit.Day)
				}
			}
		}
	}
	return 0
}

//查看地点
func ViewSite(args []string) int {
	if len(args) == 1 {
		fmt.Println(consts.SITE_LIST)
		for num, site := range placeGo.Sites {
			fmt.Printf(consts.SITE_ITEM, num+1, site.Id, site.Name)
		}
	}
	return 0
}

//查看商品
func ViewGoods(args []string) int {
	if len(args) == 1 {
		if siteGo == nil {
			fmt.Printf(consts.GOTO_SITE, placeGo.Name)
			return 0
		} else if len(siteGo.Goods) == 0 {
			fmt.Println(consts.NO_GOODS)
			return 0
		}

		fmt.Println(consts.GOODS_LIST)
		for num, site_goods := range siteGo.Goods {
			if goods, err := db.GetGoodsById(site_goods.Goods); err == nil {
				fmt.Printf(consts.GOODS_ITEM, num+1, site_goods.Goods, goods.Name, site_goods.Price, site_goods.Count)
			}
		}
	}
	return 0
}
