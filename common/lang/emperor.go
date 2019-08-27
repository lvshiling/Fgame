package lang

const (
	EmperorWorshipReachLimit LangCode = EmperorBase + iota
	EmperorNoWorshipMyself
	EmperorStorageGetNotReach
	EmperorWasMyself
	EmperorNoExist
	EmperorRobbedByOther
	EmperorGetStorageNotice
	EmperorRobNotice
	EmperorMergeServerTitle
	EmperorMergeServerContent
	EmperorOpenBoxNoOwner
	EmperorOpenBoxNoStorage
	EmperorOpenBoxRobbed
	EmperorOpenBoxNoticeGood
	EmperorOpenBoxNotice
	EmperorRobRewardTitle
	EmperorRobRewardContent
	EmperorOpenBoxTitle
	EmperorOpenBoxContent
	EmperorRobTitle
	EmperorRobGiveBlackContent
)

var (
	emperorLangMap = map[LangCode]string{
		EmperorWorshipReachLimit:   "膜拜次数已到上限",
		EmperorNoWorshipMyself:     "不可以对自己进行膜拜",
		EmperorStorageGetNotReach:  "条件不满足,无法领取",
		EmperorWasMyself:           "已经是帝王了",
		EmperorNoExist:             "帝王无主",
		EmperorRobbedByOther:       "帝王已被其它人抢先抢走",
		EmperorGetStorageNotice:    "%s从龙椅旁边的宝库中乐滋滋的领走了%s银两，再也不缺钱了",
		EmperorRobNotice:           "%s一脚将%s踹下龙椅，战力上升%s",
		EmperorMergeServerTitle:    "合服龙椅重置",
		EmperorMergeServerContent:  "由于当前服务器进行合服,龙椅进行重置,您失去了原先的龙椅位置,系统补偿您上次抢夺龙椅花费的%d%%元宝,请进行查收",
		EmperorOpenBoxNoOwner:      "当前帝王不是您",
		EmperorOpenBoxNoStorage:    "当前宝箱库存为0",
		EmperorOpenBoxRobbed:       "您的帝王已被抢走,无法开启帝王宝箱",
		EmperorOpenBoxNoticeGood:   "鸿运当头,%s开启帝王宝箱，获得极品道具%s",
		EmperorOpenBoxNotice:       "时运一般,%s开启帝王宝箱，获得道具%s",
		EmperorRobRewardTitle:      "帝王奖励",
		EmperorRobRewardContent:    "抢夺帝王奖励内容",
		EmperorOpenBoxTitle:        "帝王宝箱奖励",
		EmperorOpenBoxContent:      "帝王宝箱奖励内容",
		EmperorRobTitle:            "龙椅抢夺",
		EmperorRobGiveBlackContent: "您的龙椅被%s抢走,系统返还给您100%%的绑元,请查收!",
	}
)

func init() {
	mergeLang(emperorLangMap)
}
