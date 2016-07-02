package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"consts"
	"game"
)

func main() {
	//登录
	game.Login()
	if game.IsLogin {
		r := bufio.NewReader(os.Stdin)
		handlers := game.GetCommandHandlers()
		for { // 循环读取用户输入
			fmt.Printf(consts.COMMAND, game.TimeNow, game.WeatherNow, game.PlaceNow, game.SiteNow)
			b, _, _ := r.ReadLine()
			line := string(b)

			tokens := strings.Split(line, " ")

			if handler, ok := handlers[tokens[0]]; ok {
				fmt.Println()
				if ret := handler(tokens); ret != 0 {
					break
				}
			} else {
				fmt.Println(tokens[0], consts.CNE)
			}
		}
	}
}
