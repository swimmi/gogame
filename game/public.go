package game

import (
	"fmt"
	"models"
	"time"

	"utils"
)

var IsLogin = false
var roleGo *models.Role
var npcGo *models.Npc

var TimeNow = ""
var WeatherNow = ""

var WEATHER = []string{"天晴", "阴天", "雨天"}

func Tick() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			SetEnv()
		}
	}()
}

func SetEnv() {
	now := time.Now()
	from := roleGo.Created
	year := now.Year() - from.Year() + 1
	month := int(now.Month())
	day := now.Day()
	ap := "上午"
	hour := now.Hour()
	if hour > 12 {
		hour -= 12
		ap = "下午"
	}
	minute := now.Minute()
	TimeNow = fmt.Sprintf("第%s年 %s月%s日 %s%v时%v分", utils.PsSmNum(year), utils.PsSmNum(month), utils.PsSmNum(day), ap, hour, minute)
	WeatherNow = WEATHER[0]
}
