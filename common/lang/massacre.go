package lang

const (
	MassacreAdanvacedReachedLimit = MassacreBase + iota
	MassacreAdvanceToLow
	MassacreAdvanceToHigh
	MassacreAdvanceNotGen
	MassacreAdvancedOneNotice
	MassacreAdvancedTwoNotice
	MassacreAdvancedWeaponNotice
)

var (
	massacreLangMap = map[LangCode]string{
		MassacreAdanvacedReachedLimit: "戮仙刃已达最高阶",
		MassacreAdvanceToLow:          "您戮仙刃系统的阶别不足，无法使用物品",
		MassacreAdvanceToHigh:         "您戮仙刃系统的阶别过高，无法使用物品",
		MassacreAdvanceNotGen:         "您的杀气不足，无法进行操作",
		MassacreAdvancedOneNotice:     "大杀四方，恭喜%s成功将戮仙刃提升至%s，战力飙升%s，达到%s阶即可获取酷炫神兵外形",
		MassacreAdvancedTwoNotice:     "大杀四方，恭喜%s成功将戮仙刃提升至%s，战力飙升%s",
		MassacreAdvancedWeaponNotice:  "神兵浴血，恭喜%s将戮仙刃提升至%s，战力飙升%s，成功激活浴血神兵外形%s",
	}
)

func init() {
	mergeLang(massacreLangMap)
}
