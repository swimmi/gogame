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

	fmt.Println(consts.WELCOME)

	r := bufio.NewReader(os.Stdin)

	handlers := game.GetCommandHandlers()

	for { // 循环读取用户输入
		fmt.Print(consts.CHT)
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		if handler, ok := handlers[tokens[0]]; ok {
			if ret := handler(tokens); ret != 0 {
				break
			}
		} else {
			fmt.Println(tokens[0], consts.CNE)
		}
	}
}
