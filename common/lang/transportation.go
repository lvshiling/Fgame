package lang

const (
	TransportationOnDoing LangCode = TransportationBase + iota
	TransportationAcceptNumNoEnough
	TransportationAcceptSystemBroadcast
	TransportationAllianceAcceptSystemBroadcast
	TransportationAllianceRobSystemBroadcast
	TransportationPersonalRobSystemBroadcast
	TransportationNotAllianceTransportation
	TransportationNotExist
	TransportationDistressCD
	TransportationRobFull
	TransportationAttackedFailed
	TransportationNotEnoughSlot
)

var (
	TransportationLangMap = map[LangCode]string{
		TransportationOnDoing:                       "当前正在押镖!",
		TransportationAcceptNumNoEnough:             "今日押镖次数已达上限，请明日再来",
		TransportationAllianceAcceptSystemBroadcast: "%s仙盟开始押送仙盟镖车，夺下仙盟镖车，本仙盟所有成员都将获得珍贵劫镖宝箱奖励!",
		TransportationAllianceRobSystemBroadcast:    "恭喜%s玩家成功劫走镖车，该玩家仙盟所有成员都将获得劫镖宝箱奖励",
		TransportationAcceptSystemBroadcast:         "玩家%s开始押送%s镖车%s",
		TransportationPersonalRobSystemBroadcast:    "恭喜%s玩家成功劫走镖车，获得%d%s",
		TransportationNotAllianceTransportation:     "不是仙盟镖车",
		TransportationNotExist:                      "镖车不存在",
		TransportationDistressCD:                    "穿云箭CD中",
		TransportationRobFull:                       "您今天已经劫镖%d次，无法再进行劫镖",
		TransportationAttackedFailed:                "您当前pk模式为和平模式，无法攻击镖车，可以切换PK模式进行攻击",
		TransportationNotEnoughSlot:                 "押镖获得物品，背包空间不足",
	}
)

func init() {
	mergeLang(TransportationLangMap)
}
