package lang

const (
	RingNotEquip LangCode = RingBase + iota
	RingAdvanceAlreadyTop
	RingStrengthenAlreadyTop
	RingJingLingAlreadyTop
	RingFuseItemNotSuit
	RingTempalteNotExist
	RingBaoKuNotGetRewards
	RingLuckyBoxTitle
	RingLuckyBoxContent
	RingAttendPointsNotEnough
	RingAdvanceNotice
	RingFuseNotice
	RingIsNotRing
)

var (
	ringLangMap = map[LangCode]string{
		RingNotEquip:              "玩家未穿戴该特戒",
		RingAdvanceAlreadyTop:     "特戒进阶已达到最高",
		RingStrengthenAlreadyTop:  "特戒强化等级已达到最高",
		RingJingLingAlreadyTop:    "特戒净灵等级已达到最高",
		RingFuseItemNotSuit:       "特戒融合需要物品与当前物品不符",
		RingTempalteNotExist:      "特戒模板不存在",
		RingBaoKuNotGetRewards:    "宝库没有掉落",
		RingLuckyBoxTitle:         "特戒宝库寻宝额外奖励",
		RingLuckyBoxContent:       "恭喜您在特戒宝库寻宝中额外获得如下奖励",
		RingAttendPointsNotEnough: "宝库积分不足",
		RingAdvanceNotice:         "恭喜%s的%s成功升到%d阶激活特殊属性：%s",
		RingFuseNotice:            "恭喜%s将%s成功融合%d次，技能效果加强，战力飙升%d!",
		RingIsNotRing:             "不是特戒",
	}
)

func init() {
	mergeLang(ringLangMap)
}
