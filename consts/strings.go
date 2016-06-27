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
	SELECT_NPC_FORMAT = "格式：选择NPC <ID/名字>"
	PRESENT_FORMAT    = "格式：赠送物品 <NPC> <物品> [数量]"

	NPC_LIST = `╬--<NPC>	ID	名字	性别	年龄	好感		--╬`
	NPC_ITEM = `╬--< %v >	%v	%s	%s	%v	%v		--╬
`

	ROLE_ITEM_LIST = `╬--<物品>	ID	物品	数量		--╬`
	ROLE_ITEM      = `╬--< %v >	%v	%v	%v		--╬
`
	THANKS = "谢谢你的礼物~"

	LOG_FAIL = "日志写入失败"
)
