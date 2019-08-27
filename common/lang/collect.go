package lang

const (
	CollectTitle = CollectBase + iota
	CollectContent
	CollectPointNoNum
	CollectChooseNoCollecting
	CollectChooseNoSpeaial
	CollectNowCollecting
	CollectNotCollectNPC
	CollectMiZangNotCollect
	CollectMiZangOpenTypeWrong
	CollectMiZangTitle
	CollectMiZangContent
	CollectMiZangDisappear
)

var (
	collectLangMap = map[LangCode]string{
		CollectTitle:               "采集物奖励",
		CollectContent:             "采集物奖励内容",
		CollectPointNoNum:          "采集点已无采集次数",
		CollectChooseNoCollecting:  "不在采集中",
		CollectChooseNoSpeaial:     "不是特殊采集物",
		CollectNowCollecting:       "已经在采集中",
		CollectNotCollectNPC:       "不是采集物",
		CollectMiZangNotCollect:    "没有采集过密藏",
		CollectMiZangOpenTypeWrong: "密藏开启类型错误",
		CollectMiZangTitle:         "秘藏",
		CollectMiZangContent:       "成功开启BOSS秘藏宝箱，但由于背包空间不足，奖励通过邮件发送给您，请查收！",
		CollectMiZangDisappear:     "该秘藏已消失",
	}
)

func init() {
	mergeLang(collectLangMap)
}
