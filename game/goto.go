package game

import (
	"fmt"
	"strconv"

	"consts"
	"db"
)

//前往地点
func GotoSite(args []string) int {
	if len(args) == 2 {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			if site, err := db.GetSiteByName(*placeGo, args[1]); err == nil {
				siteGo = &site
				site.ViewInfo()
			}
		} else {
			if site, err := db.GetSiteById(*placeGo, id); err == nil {
				siteGo = &site
				SiteNow = siteGo.Name
				site.ViewInfo()
			}
		}
	} else {
		fmt.Println(consts.GOTO_SITE_FORMAT)
	}
	return 0
}
