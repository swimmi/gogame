package db

import (
	"errors"
	"time"

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
		fmt.Println("GetNpcByName", consts.NNE)
	}
	return npc, err
}

func AlterNpcFavour(id int, favour int) {
	c := Connect("npc")
	err := c.Update(bson.M{"id": id}, bson.M{"$inc": bson.M{"favour": favour}})
	if err != nil {
		fmt.Println("AlterNpcFavour", consts.NNE)
	}
}

//RoleItem
func GetRoleItemList() ([]models.RoleItem, error) {
	c := Connect("role_item")
	role_items := []models.RoleItem{}
	err := c.Find(nil).Sort("id").All(&role_items)
	if err != nil {
		fmt.Println("GetRoleItemList", consts.INE)
	}
	return role_items, err
}

//Item
func GetItemById(id int) (models.Item, error) {
	c := Connect("item")
	item := models.Item{}
	err := c.Find(bson.M{"id": id}).One(&item)
	if err != nil {
		fmt.Println("GetItemById", consts.INE)
	}
	return item, err
}

func GetItemByName(name string) (models.Item, error) {
	c := Connect("item")
	item := models.Item{}
	err := c.Find(bson.M{"name": name}).One(&item)
	if err != nil {
		fmt.Println("GetItemByName", consts.INE)
	}
	return item, err
}

func RemoveRoleItem(id int, count int) {
	c := Connect("role_item")
	err := c.Update(bson.M{"id": id}, bson.M{"$inc": bson.M{"count": -count}})
	if err != nil {
		fmt.Println("RemoveRoleItem", err.Error())
	}
}

//RoleItem
func GetRoleItemById(id int) (models.RoleItem, error) {
	c := Connect("role_item")
	role_item := models.RoleItem{}
	err := c.Find(bson.M{"id": id}).One(&role_item)
	if err != nil {
		fmt.Println("GetRoleItemById", consts.INE)
	}
	return role_item, err
}

func GetRoleItemByName(name string) (models.RoleItem, error) {
	role_item := models.RoleItem{}
	item, err := GetItemByName(name)
	if err == nil {
		c := Connect("role_item")
		err = c.Find(bson.M{"item": item.Id}).One(&role_item)
		if err != nil {
			fmt.Println("GetRoleItemByName", consts.INE)
		}
	}
	return role_item, err
}

//Present
func PresentNpcItem(npc models.Npc, role_item models.RoleItem, count int) {
	item, err := GetItemById(role_item.Item)
	if err == nil {
		AlterNpcFavour(npc.Id, item.FavourValue)
		RemoveRoleItem(role_item.Id, count)
		npc.ThankPresent(item.FavourValue)
		log := models.Log{Type: "present", Npc: npc.Id, Item: role_item.Item, Count: count, Timestamp: time.Now()}
		Log(log)
	}
}

func Log(log models.Log) error {
	c := Connect("log")
	err := c.Insert(&log)
	if err != nil {
		fmt.Println("Log", consts.LOG_FAIL)
	}
	return err
}

//Place
func GetPlaceById(id int) (models.Place, error) {
	c := Connect("place")
	place := models.Place{}
	err := c.Find(bson.M{"id": id}).One(&place)
	if err != nil {
		fmt.Println("GetPlaceById", consts.PNE)
	}
	return place, err
}
func GetPlaceByName(name string) (models.Place, error) {
	c := Connect("place")
	place := models.Place{}
	err := c.Find(bson.M{"name": name}).One(&place)
	if err != nil {
		fmt.Println("GetPlaceById", consts.PNE)
	}
	return place, err
}

//Site
func GetSiteById(place models.Place, id int) (models.Site, error) {
	for _, site := range place.Sites {
		if site.Id == id {
			return site, nil
		}
	}
	return models.Site{}, errors.New(consts.PNE)
}
func GetSiteByName(place models.Place, name string) (models.Site, error) {
	for _, site := range place.Sites {
		if site.Name == name {
			return site, nil
		}
	}
	return models.Site{}, errors.New(consts.PNE)
}

//Goods
func GetGoodsById(id int) (models.Goods, error) {
	c := Connect("goods")
	goods := models.Goods{}
	err := c.Find(bson.M{"id": id}).One(&goods)
	if err != nil {
		fmt.Println("GetGoodsById", consts.GNE)
	}
	return goods, err
}
func GetGoodsByName(name string) (models.Goods, error) {
	c := Connect("goods")
	goods := models.Goods{}
	err := c.Find(bson.M{"name": name}).One(&goods)
	if err != nil {
		fmt.Println("GetGoodsByName", consts.GNE)
	}
	return goods, err
}

//Farm
func GetFarmUnitList() ([]models.FarmUnit, error) {
	c := Connect("farm")
	farm_units := []models.FarmUnit{}
	err := c.Find(nil).Sort("id").All(&farm_units)
	if err != nil {
		fmt.Println("GetFarmUnitList", consts.INE)
	}
	return farm_units, err
}

func GetCropById(id int) (models.Crop, error) {
	c := Connect("crop")
	crop := models.Crop{}
	err := c.Find(bson.M{"id": id}).One(&crop)
	if err != nil {
		fmt.Println("GetCropById", consts.INE)
	}
	return crop, err
}
