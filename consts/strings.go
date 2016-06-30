package consts

const (
	LOGIN = "登录："
	CNE   = "命令不存在"   //Command Not Exist
	RNE   = "角色不存在"   //Role Not Exist
	NNE   = "NPC不存在"  //NPC Not Exist
	INE   = "物品不存在"   //Item Not Exist
	PNE   = "地点不存在"   //Place Not Exist
	GNE   = "商品不存在"   //Goods Not Exist
	EID   = "请输入对应ID" //Enter ID

	HELP = `
		以下命令可使用：
	`

	COMMAND = `
	┍--------------------------------------------------------- ┑
	╬	  %s	%s	    ╬
	┕--------------------------------------------------------- ┙
命令：` //Command Hint Text
	WELCOME = `
	┍--------------------------------------------------------- ┑
	╬           多情自古伤离别  更那堪  冷落清秋节             ╬
	╬                                                %s    ╬
	┕--------------------------------------------------------- ┙
	`
	ROLE_INFO = `
	姓名: %s,
	`

	SITE_INFO = `
	欢迎您的到来，这里是%s
	`

	LOGIN_FAIL    = "登录失败，请重试"
	LOGIN_SUCCESS = "登录成功，欢迎你"

	LOGIN_FORMAT      = "格式：登录 <用户名>"
	SELECT_NPC_FORMAT = "格式：选择NPC <ID/名字>"
	PRESENT_FORMAT    = "格式：赠送物品 <NPC> <物品> [数量]"
	GOTO_SITE_FORMAT  = "格式：前往地点 <ID/名称>"

	NPC_LIST = `	╬--0	ID	名字	性别	年龄	好感	--╬`
	NPC_ITEM = `	╬--%v	%v	%s	%s	%v	%v	--╬
`

	ROLE_ITEM_LIST = `	╬--0	ID	物品	数量	--╬`
	ROLE_ITEM      = `	╬--%v	%v	%v	%v	--╬
`

	FARM_UNIT_LIST = `	╬--0	ID	作物	时期	--╬`
	FARM_UNIT      = `	╬--%v	%v	%s	%v	--╬
`

	SITE_LIST = `	╬--0	ID	名称	--╬`
	SITE_ITEM = `	╬--%v	%v	%s	--╬
`

	GOODS_LIST = `	╬--0	ID	商品		价格	数量	--╬`
	GOODS_ITEM = `	╬--%v	%v	%s		%v	%v	--╬
`
	THANKS = "	谢谢你的礼物~"

	GOTO_SITE = "	请先选择地点前往，不要停留在%s大街上~\n"
	NO_GOODS  = "	这里没有东西可以卖给你哦~"
	LOG_FAIL  = "	日志写入失败"
)
