package game

import (
	/*"bufio"
	"os"*/
	"fmt"

	"consts"
	"db"
)

// 将命令与处理函数对应
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"帮助":    Help,
		"退出":    Quit,
		"查看玩家":  ViewRole,
		"查看NPC": ViewNpc,
		"选择NPC": SelectNpc,
		"查看物品":  ViewRoleItem,
		"赠送物品":  PresentNpcItem,
	}
}

func Login() error {
	for !IsLogin {
		/*fmt.Print(consts.LOGIN)
		r := bufio.NewReader(os.Stdin)
		b, _, _ := r.ReadLine()
		name := string(b)*/
		role, err := db.Login("草未眠")
		roleGo = &role
		if err != nil {
			fmt.Println(consts.LOGIN_FAIL)
		} else {
			IsLogin = true
			fmt.Println(roleGo.Name, consts.LOGIN_SUCCESS)
		}
	}
	return nil
}
