package db

import (
	"gopkg.in/mgo.v2/bson"

	"consts"
	"fmt"
	"models"
)

//Login
func Login(name string) (models.Role, error) {

	c := Connect("role")
	role := models.Role{}
	err := c.Find(bson.M{"name": name}).One(&role)
	if err != nil {
		fmt.Println(consts.RNE)
	}
	return role, err
}

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
func GetRoleItemList(role_id int) ([]models.RoleItem, error) {
	c := Connect("role_item")
	role_items := []models.RoleItem{}
	err := c.Find(bson.M{"roleId": role_id}).Sort("id").All(&role_items)
	if err != nil {
		fmt.Println(consts.INE)
	}
	return role_items, err
}

//Item
func GetItemById(id int) (models.Item, error) {
	c := Connect("item")
	item := models.Item{}
	err := c.Find(bson.M{"id": id}).One(&item)
	if err != nil {
		fmt.Println(consts.INE)
	}
	return item, err
}

func GetItemByName(name string) (models.Item, error) {
	c := Connect("item")
	item := models.Item{}
	err := c.Find(bson.M{"name": name}).One(&item)
	if err != nil {
		fmt.Println(consts.INE)
	}
	return item, err
}
