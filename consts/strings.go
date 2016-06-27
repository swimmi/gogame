package consts

const (
	LOGIN = "登录："
	CHT   = "命令："     //Command Hint Text
	CNE   = "命令不存在"   //Command Not Exist
	RNE   = "角色不存在"   //Role Not Exist
	NNE   = "NPC不存在"  //NPC Not Exist
	INE   = "物品不存在"   //Item Not Exist
	EID   = "请输入对应ID" //Enter ID

	HELP = `
		以下命令可使用：
	`
	WELCOME = `
		╬刻╬╬骨╬╬╬━━╬╬一╬╬笑╬╬
		╬╬╬╬╬╬╬┇ 累 ┇╬╬╬╬╬╬╬
		╬╬╬╬╬╬╬┇ 了 ┇╬╬╬╬╬╬╬
		╬铭╬╬心╬╬╬━━╬╬而╬╬过╬╬
	`
	ROLE_INFO = `
		姓名: %s,
	`

	LOGIN_FAIL    = "登录失败，请重试"
	LOGIN_SUCCESS = "登录成功，欢迎你"

	LOGIN_FORMAT      = "格式：登录 <用户名>"
	SELECT_NPC_FORMAT = "格式：选择npc <ID/名字>"

	NPC_LIST = `<NPC>	ID	名字	性别	年龄`
	NPC_HEAD = "\n+++\t\t\t以下是NPC信息\t\t\t+++\n\n╬--%s\t\t--╬\n"
	NPC_ITEM = "╬--< %v >\t%v\t%s\t%s\t%v\t\t--╬\n"
	NPC_FOOT = "\n+++\t\t\tNPC信息展示完毕\t\t\t+++\n\n"

	ROLE_ITEM_LIST = `<物品>	ID	物品ID	数量`
	ROLE_ITEM_HEAD = "\n+++\t\t\t以下是物品信息\t\t\t+++\n\n╬--%s\t\t--╬\n"
	ROLE_ITEM      = "╬--< %v >\t%v\t%v\t%v\t\t--╬\n"
	ROLE_ITEM_FOOT = "\n+++\t\t\t物品信息展示完毕\t\t\t+++\n\n"
)
