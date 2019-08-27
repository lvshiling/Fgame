package lang

const (
	TeamPlayerInTeam LangCode = TeamBase + iota
	TeamPlayerOff
	TeamPlayerFull
	TeamDissolve
	TeamNoExist
	TeamCaptainIsOther
	TeamPlayerNoMember
	TeamFull
	TeamApplyNoExist
	TeamPlayerNotInTeam
	TeamInInit
	TeamInMatch
	TeamNotInMatch
	TeamInGame
	TeamInMatchLevelLow
	TeamInMatchInviteJion
	TeamInMatchInCross
	TeamInMatchInFuBen
	TeamInMatchInPvp
	TeamInMatchNoActivityTime
	TeamInMatchJionFail
	TeamInMatchDealAgreeFail
	TeamInMatchInLianYuLineUp
	TeamInMatchInGodSiegeLineUp
	TeamCreateHouseInOther
	TeamCreateHouseNoCaptain
	TeamInTeamCopyBattle
	TeamInTeamCopyJionFail
	TeamMemberPlayerOffline
	TeamNotInTeamCopyBattle
	TeamBattlePuroseIsNormal
	TeamSelfPuroseIsHouseInfo
	TeamCreateHouseLevelLow
	TeamCreateHouseSelfLevelLow
	TeamApplyJoinFuncNoOpen
	TeamHouseIsBatting
	TeamInMatchSelfInCross
	TeamInMatchSelfInLianYuLineUp
	TeamInMatchInSelfGodSiegeLineUp
	TeamInMatchSelfInFuBen
	TeamInMatchSelfInPvp
	TeamRushMatchIsExist
	TeamRushMatchCaptainOffline
	TeamInMatchInSelfShenMoLineUp
	TeamJionByLeavedInCd
)

var (
	teamLangMap = map[LangCode]string{
		TeamPlayerInTeam:                "当前已有队伍,无法加入他人队伍",
		TeamPlayerOff:                   "对方已经离线",
		TeamPlayerFull:                  "当前队伍人数已达上限",
		TeamDissolve:                    "队伍已解散",
		TeamNoExist:                     "队伍不存在",
		TeamCaptainIsOther:              "您不是队长",
		TeamPlayerNoMember:              "被操作者不是队员",
		TeamApplyNoExist:                "操作不存在的申请",
		TeamPlayerNotInTeam:             "玩家不在队伍中",
		TeamInInit:                      "队伍还没开始匹配",
		TeamInMatch:                     "队伍正在匹配中",
		TeamNotInMatch:                  "队伍还没在匹配",
		TeamInMatchLevelLow:             "您当前匹配战队存在玩家角色等级不足,无法进入匹配!",
		TeamInMatchInviteJion:           "当前队伍正在匹配,无法邀请加入",
		TeamInMatchInCross:              "您当前匹配战队内有玩家正在跨服中,无法进入匹配!",
		TeamInMatchInFuBen:              "您当前匹配战队内有玩家正在副本中,无法进入匹配!",
		TeamInMatchInPvp:                "您当前匹配战队内有玩家正在战斗中,无法进入匹配!",
		TeamInMatchNoActivityTime:       "当前不在活动时间内,无法进行匹配,活动时间为每天8:00-24:00",
		TeamInMatchJionFail:             "加入失败!对方战队已开始进行匹配",
		TeamInMatchDealAgreeFail:        "您加入的队伍正在3V3匹配中,暂时无法加入",
		TeamInMatchInLianYuLineUp:       "您当前匹配战队内有玩家正在无间炼狱排队中,无法进入匹配!",
		TeamInMatchInGodSiegeLineUp:     "您当前匹配战队内有玩家正在神兽攻城排队中,无法进入匹配!",
		TeamCreateHouseInOther:          "当前已在别的房间中",
		TeamInTeamCopyBattle:            "队伍正在组队副本战斗中",
		TeamInTeamCopyJionFail:          "加入失败!对方战队已开始进行组队副本战斗",
		TeamMemberPlayerOffline:         "您当前匹配战斗内有玩家处于离线中,无法进入匹配",
		TeamNotInTeamCopyBattle:         "队伍还没开始战斗",
		TeamBattlePuroseIsNormal:        "您当前的队伍是通常组队,无法开始战斗",
		TeamSelfPuroseIsHouseInfo:       "您当前的队伍已是房间信息",
		TeamCreateHouseLevelLow:         "您当前战队存在玩家还未开启该房间功能,无法创建房间",
		TeamCreateHouseSelfLevelLow:     "您当前还未开启创建该房间功能,无法创建房间",
		TeamApplyJoinFuncNoOpen:         "您当前功能还未开启相应的组队副本,无法加入组队副本的队伍",
		TeamHouseIsBatting:              "当前房间已开始战斗,无法加入",
		TeamInMatchSelfInCross:          "您当前正在跨服中,无法进入匹配!",
		TeamInMatchSelfInLianYuLineUp:   "您当前正在无间炼狱排队中,无法进入匹配!",
		TeamInMatchInSelfGodSiegeLineUp: "您当前正在神兽攻城排队中,无法进入匹配!",
		TeamInMatchSelfInFuBen:          "您当前正在副本中,无法进入匹配!",
		TeamInMatchSelfInPvp:            "您当前正在战斗中,无法进入匹配!",
		TeamRushMatchIsExist:            "已经有人在催了,稍安勿躁",
		TeamRushMatchCaptainOffline:     "队长当前不在线",
		TeamInMatchInSelfShenMoLineUp:   "您当前正在神魔战场中,无法进入匹配!",
		TeamJionByLeavedInCd:            "您方才被请出了队伍,短时间内无法加入这个队伍哦",
	}
)

func init() {
	mergeLang(teamLangMap)
}
