package consts

const (
	CHT = "命令："     //Command Hint Text
	CNE = "命令不存在"   //Command Not Exist
	NNE = "NPC不存在"  //NPC Not Exist
	EID = "请输入对应ID" //Enter ID

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

	LOGIN_FORMAT      = "格式：登录 <用户名>"
	SELECT_NPC_FORMAT = "格式：选择npc <ID/名字>"

	NPC_LIST = `<NPC>	ID	名字	性别	年龄`
	NPC_HEAD = "\n+++\t\t\t以下是NPC信息\t\t\t+++\n\n╬--%s\t\t--╬\n"
	NPC_ITEM = "╬--< %v >\t%v\t%s\t%s\t%v\t\t--╬\n"
	NPC_FOOT = "\n+++\t\t\tNPC信息展示完毕\t\t\t+++\n\n"
)
