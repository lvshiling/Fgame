package lang

const (
	RepeatCreateJob LangCode = PlayerBase + iota
	NameInvalid
	JobInvalid
	SexInvalid
	NameAlreadyExist
	PlayerNoExist
	PlayerNoOnline
	PlayerSilverNoEnough
	PlayerGoldNoEnough
	PlayerLevelTooLow
	PlayerLevelTooHigh
	PlayerRoleWrong
	PlayerSexWrong
	PlayerZhuanShengTooLow
	PlayerZhuanShengTooHigh
	PlayerNoInScene
	PlayerNotInTeam
	PlayerOnlineTimeNoEnough
	PlayerInPVP
	PlayerInCross
	PlayerIn3v3Match
	PlayerInMarryCruise
	PlayerInFuBen
	PlayerInTeamCopyBattle
	PlayerInForbidChat
	PlayerGoldYuanLevelMax
	PlayerOffline
	PlayerNotRedState
	PlayerInLianYuLineUp
	PlayerInGodSiegeLineUp
	PlayerTrackNoOnline
	PlayerTrackInCross
	NetworkNotStable
	PlayerGongXunNoEnough
	PlayerArenaPointNoEnough
	PlayerKilledRewardCd
	PlayerNoInAlliance
	PlayerArenapvpPointNoEnough
	PlayerLineUpNoExist
)

var playerLangMap = map[LangCode]string{
	RepeatCreateJob:             "重复创建角色",
	NameInvalid:                 "名字无效",
	JobInvalid:                  "角色无效",
	SexInvalid:                  "性别无效",
	NameAlreadyExist:            "名字已经存在",
	PlayerNoExist:               "角色不存在",
	PlayerNoOnline:              "用户不在线",
	PlayerSilverNoEnough:        "银两不足",
	PlayerGoldNoEnough:          "元宝不足",
	PlayerLevelTooLow:           "等级太低",
	PlayerLevelTooHigh:          "等级太高",
	PlayerRoleWrong:             "职业不符",
	PlayerSexWrong:              "性别不符",
	PlayerZhuanShengTooLow:      "转数太低",
	PlayerZhuanShengTooHigh:     "转数太高",
	PlayerNoInScene:             "玩家不在场景中",
	PlayerNotInTeam:             "玩家不在组队中",
	PlayerOnlineTimeNoEnough:    "在线时间不足",
	PlayerInPVP:                 "玩家当前处于PK状态，无法传送",
	PlayerInCross:               "玩家当前处于跨服中，无法传送",
	PlayerIn3v3Match:            "玩家当前处于3v3匹配,无法传送",
	PlayerInMarryCruise:         "玩家当前处于婚礼游街状态,无法传送",
	PlayerInFuBen:               "玩家当前处于副本场景，无法传送",
	PlayerInTeamCopyBattle:      "玩家当前处于组队副本战斗中,无法传送",
	PlayerInForbidChat:          "玩家当前处于禁言中",
	PlayerOffline:               "对方不在线",
	PlayerGoldYuanLevelMax:      "玩家当前元神等级已达上限,无法继续吞噬装备",
	PlayerNotRedState:           "玩家不是红名状态",
	PlayerInLianYuLineUp:        "您当前正在无间炼狱排队中,无法传送",
	PlayerInGodSiegeLineUp:      "您当前正在神兽攻城排队中,无法传送",
	PlayerTrackNoOnline:         "该玩家不在线,无法进行追踪",
	PlayerTrackInCross:          "跨服地图无法直接传送,请自行前往",
	NetworkNotStable:            "网络波动,请稍后重试",
	PlayerGongXunNoEnough:       "玩家功勋不足",
	PlayerArenaPointNoEnough:    "玩家3V3积分不足",
	PlayerKilledRewardCd:        "该玩家处于%s秒保护cd中，击杀后无法获得击杀奖励",
	PlayerNoInAlliance:          "用户不在仙盟",
	PlayerArenapvpPointNoEnough: "玩家比武大会积分不足",
	PlayerLineUpNoExist:         "您当前未在排队中",
}

func init() {
	mergeLang(playerLangMap)
}
