package lang

const (
	AllianceSceneGuardNoExist LangCode = AllianceSceneBase + iota
	AllianceSceneGuardCalled
	AllianceSceneNotBelongFirstDefend
	AllianceSceneNotBelongDefend
	AllianceSceneNotInAllianceScene
	AllianceSceneDoorNotBroke
	AllianceSceneBelongFirstDefend
	AllianceSceneDoorRewardAlreadyGet
	AllianceSceneDoorRewardNoExist
	AllianceSceneWinTitle
	AllianceSceneWinContent
	AllianceSceneContinueWinTitle
	AllianceSceneContinueWinContent
	AllianceScenePointNotEnough
	AllianceSceneDoorRewTitle
	AllianceSceneDoorRewContent
	AllianceScenePlayerIsDefence
)

var allianceSceneLangMap = map[LangCode]string{
	AllianceSceneGuardNoExist:         "城战守卫不存在",
	AllianceSceneGuardCalled:          "城战守卫已经召唤过",
	AllianceSceneNotBelongFirstDefend: "不属于城战最初守方",
	AllianceSceneNotBelongDefend:      "不属于城战守方",
	AllianceSceneNotInAllianceScene:   "不在城战",
	AllianceSceneDoorNotBroke:         "城门还没攻破",
	AllianceSceneBelongFirstDefend:    "属于城战最初守方",
	AllianceSceneDoorRewardAlreadyGet: "城门奖励已经领取",
	AllianceSceneDoorRewardNoExist:    "城门没有奖励",
	AllianceSceneWinTitle:             "城战获胜奖励",
	AllianceSceneWinContent:           "城战获胜奖励内容",
	AllianceSceneContinueWinTitle:     "城战连胜奖励",
	AllianceSceneContinueWinContent:   "城战连胜奖励内容",
	AllianceScenePointNotEnough:       "积分不足",
	AllianceSceneDoorRewTitle:         "城战攻破城门奖励",
	AllianceSceneDoorRewContent:       "城战攻破城门奖励内容",
	AllianceScenePlayerIsDefence:      "玉玺仅攻方可采集，请守方保护好玉玺",
}

func init() {
	mergeLang(allianceSceneLangMap)
}
