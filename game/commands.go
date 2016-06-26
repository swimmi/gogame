package game

// 将命令与处理函数对应
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"帮助":    Help,
		"退出":    Quit,
		"登录":    Login,
		"查看玩家":  ViewRole,
		"查看npc": ViewNpc,
		"选择npc": SelectNpc,
		"查看物品":  ViewItem,
	}
}
