package lang

const (
	MountEatCulDanReachedLimit LangCode = MountBase + iota
	MountEatCulDanReachedFull
	MountEatUnDanReachedLimit
	MountEatUnDanReachedFull
	MountUnrealCondNotReached
	MountAdanvacedReachedLimit
	MountUnrealNoExist
	MountOtherIsNoActived
	MountAdvanceToLow
	MountAdvanceToHigh
	MountAdvanceNotEqual
	MountSkinUpstarNoActive
	MountSkinReacheFullStar
	MountAdvancedNotice
	MountUnrealActivateNotice
	MountDragonActivateNotice
)

var (
	mountLangMap = map[LangCode]string{
		MountEatCulDanReachedLimit: "食用培养丹等级已达最大,请进阶后再试",
		MountEatCulDanReachedFull:  "培养丹食丹等级满级",
		MountEatUnDanReachedLimit:  "幻化丹食丹等级已达最大,请进阶后再试",
		MountEatUnDanReachedFull:   "幻化丹食丹等级满级",
		MountUnrealCondNotReached:  "还有条件未达成，无法解锁幻化",
		MountAdanvacedReachedLimit: "坐骑已达最高阶",
		MountUnrealNoExist:         "当前没有幻化",
		MountOtherIsNoActived:      "该坐骑还未激活",
		MountAdvanceToLow:          "您坐骑系统的阶别不足，无法使用物品",
		MountAdvanceToHigh:         "您坐骑系统的阶别过高，无法使用物品",
		MountAdvanceNotEqual:       "您坐骑系统的阶别不符，无法使用物品",
		MountSkinUpstarNoActive:    "未激活的坐骑皮肤,无法升星",
		MountSkinReacheFullStar:    "坐骑皮肤已满星",
		MountAdvancedNotice:        "恭喜%s成功将坐骑提升至%s，战力飙升%s",
		MountUnrealActivateNotice:  "恭喜%s成功将坐骑幻化为%s，战力飙升%s",
		MountDragonActivateNotice:  "神龙现世，%s将龙蛋沐浴神血，提升至第%s阶段，不知将孵化出哪个属性的神龙",
	}
)

func init() {
	mergeLang(mountLangMap)
}
