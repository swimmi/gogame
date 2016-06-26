package db

import (
	"gopkg.in/mgo.v2/bson"

	"consts"
	"fmt"
	"models"
)

//Npc
func GetNpcList() ([]models.Npc, error) {
	c := Connect("npc")
	npcs := []models.Npc{}
	err := c.Find(nil).Sort("id").All(&npcs)
	if err != nil {
		fmt.Println(consts.NNE)
	}
	return npcs, err
}

func GetNpcById(id int) (models.Npc, error) {
	c := Connect("npc")
	npc := models.Npc{}
	err := c.Find(bson.M{"id": id}).One(&npc)
	if err != nil {
		fmt.Println(consts.NNE)
	}
	return npc, err
}

func GetNpcByName(name string) (models.Npc, error) {
	c := Connect("npc")
	npc := models.Npc{}
	err := c.Find(bson.M{"name": name}).One(&npc)
	if err != nil {
		fmt.Println(consts.NNE)
	}
	return npc, err
}

//RoleItem
func GetRoleItemList() ([]models.RoleItem, error) {
	c := Connect("role_item"))
	items := []models.RoleItem{}
	err := c.Find(nil).Sort("id").All(&items)
	if err != nil {
		fmt.Println(consts.INE)
	}
	return items, err
}
