package lang

const (
	FourGodKeyNoEnough LangCode = FourGodBase + iota
	FourGodBoxIsOpened
	FourGodNPCNoBox
	FourGodBlackItemNoExist
	FourGodFixedTime
	FourGodPickUpKeyReachLimit
	FourGodBossBorn
	FourGodBossRefresh
	FourGodOpenBox
	FourGodOpenBoxIsExist
	FourGodOpenBoxRepeat
	FourGodBossBlood
	FourGodBossBloodChat
)

var (
	fourGodLangMap = map[LangCode]string{
		FourGodKeyNoEnough:         "开启该宝箱需要%s把钥匙,击杀小怪与BOSS都可以获得钥匙!",
		FourGodBoxIsOpened:         "宝箱已被开启",
		FourGodNPCNoBox:            "npcId不是宝箱",
		FourGodBlackItemNoExist:    "当前未拥有蒙面衣",
		FourGodFixedTime:           "四神遗迹规定时间开始",
		FourGodPickUpKeyReachLimit: "钥匙数量已达上限,无法拾取更多钥匙,开启宝箱将获得丰厚奖励!",
		FourGodBossBorn:            "%s已经刷新,赶紧前往击杀吧!",
		FourGodBossRefresh:         "%s已经刷新,赶紧前往四神遗迹击杀吧!%s",
		FourGodOpenBox:             "玩家%s消耗%d个钥匙开启遗迹宝箱,获得了道具%s",
		FourGodOpenBoxIsExist:      "宝箱有人正在开启",
		FourGodOpenBoxRepeat:       "宝箱正在开启中",
		FourGodBossBlood:           "%s血量剩余%d%%,前往击杀吧!",
		FourGodBossBloodChat:       "%s血量剩余%d%%,赶紧前往击杀吧!",
	}
)

func init() {
	mergeLang(fourGodLangMap)
}
