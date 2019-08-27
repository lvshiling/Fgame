package lang

const (
	AllianceNameIllegal LangCode = AllianceBase + iota
	AllianceNameDirty
	AllianceNoticeIllegal
	AllianceNoticeDirty
	AllianceNameExist
	AllianceCreateFailed
	AllianceNoExist
	AllianceUserAlreadyInAlliance
	AllianceUserInAllianceToApply
	AllianceAlreadyFullApply
	AllianceAlreadyFullInvitation //10
	AllianceAlreadyApply
	AllianceApplyNoExist
	AllianceUserNotInAlliance
	AllianceUserPrivilegeNoEnough
	AllianceMengZhuNoExit
	AllianceNotInSameAlliance
	AllianceUserNotMengZhu
	AllianceUserCanNotImpeachSelf
	AllianceSceneNoOpen
	AllianceRefuseApply //20
	AllianceUserCanNotKickSelf
	AllianceUserCanNotCommitSelf
	AllianceAlreadyMember
	AlliancePositionAlreadyFull
	AllianceImpeachConditionNotEnough
	AllianceDonateMaxTimes
	AllianceSkillMaxLevel
	AllianceSkillNotOpen
	AllianceSkillNotEnoughGongXian
	AllianceInviterAlreadyInAlliance //30
	AllianceCreateLog
	AllianceJoinLog
	AllianceExitLog
	AllianceKickLog
	AllianceDonateItemLog
	AllianceImpeachLog
	AllianceCommitLog
	AllianceReleaseLog
	AllianceTransferLog
	AllianceCharmSystemNotice //40
	AllianceKickNotice
	AllianceCommitNotice
	AllianceExitNotice
	AllianceImpeachNotice
	AllianceTransferNotice
	AllianceConvertTimesNotEnough
	AllianceDismissCanNotInSceneOpened
	AllianceDismissCanNotInTransportation
	AllianceTransportationNotInAlliance
	AllianceTransportationPositionNoEnough //50
	AllianceUserNotInAllianceForHegemon
	AllianceInvitationPositionNotEnough
	AllianceInviteMergeRefuse
	AllianceTransportationAcceptNumNoEnough
	AllianceDonateHuFuNotice
	AllianceCreateNotice
	AllianceHuangGongCloseCannotEnter
	AllianceSceneGuardInactive
	AllianceTickInCrossTuLong
	AllianceExitInCrossTuLong //60
	AllianceDismissInCrossTuLong
	AllianceImpeachInCrossTuLong
	AllianceTransferInCrossTuLong
	AllianceDepotHasNotEnoughSlot
	AllianceDepotHasNotEnoughPoint
	AllianceJoinNotice
	AllianceJoinPersonalNotice
	AllianceDepotSaveNotice
	AllianceMemberDeadNotice
	AllianceMemberCallNotice //70
	AllianceMemberCallNoticeCD
	AllianceMemberNotRescue
	AllianceDepotTakeOutNotice
	AllianceBatchJoinCD
	AllianceDouShenJoinMailTitle
	AllianceDouShenJoinMailContentSenior
	AllianceDouShenJoinMailContentJunior
	AllianceDouShenExitMailTitle
	AllianceDouShenExitMailContent
	AllianceHuangGongCloseCannotExit //80
	AllianceDonateResouceLog
	AllianceBossSummonNoMengZhu
	AllianceBossSummonedBoss
	AllianceBossTitle
	AllianceBossContent
	AllianceBossEnterNoStart
	AllianceBossTodayFinish
	AllianceDismissCanNotInBoss
	AllianceBossSummonSucessChat
	AllianceSendInvitation
	AllianceDismissInYuXiWar
	AllianceTickInYuXiWar
	AllianceExitInYuXiWar
	AllianceImpeachInYuXiWar
	AllianceTransferInYuXiWar
	AllianceDepotClose
	AllianceInviteAllianceMergeMemberToMuch
	AllianceInviteAllianceMergeMemberNotOnline
	AllianceInviteAllianceMergeNotice
	AllianceInviteAllianceMergeAllianceNotice
	AllianceInviteAllianceMergeTitle
	AllianceInviteAllianceMergeContent
	AllianceInviteAllianceMergeLog
	AllianceRenameMailTitle
	AllianceRenameMailContent
	AllianceRenameNotice
	AllianceOnAllianceActivity
	AllianceRenameSame
	AllianceMergeApplyExpire
	AllianceMergeApplyCd
	AllianceMergeApplying
	AllianceCampNotSame
)

var allianceLangMap = map[LangCode]string{
	AllianceNameIllegal:                        "仙盟名称需要在1-6个汉字之间",
	AllianceNameDirty:                          "仙盟名称中包含敏感非法字符",
	AllianceNoticeIllegal:                      "仙盟公告需要在0-50个汉字之间",
	AllianceNoticeDirty:                        "仙盟公告中包含敏感非法字符",
	AllianceNameExist:                          "仙盟名字已经存在",
	AllianceCreateFailed:                       "仙盟创建失败",
	AllianceNoExist:                            "仙盟不存在",
	AllianceUserAlreadyInAlliance:              "该玩家当前已有仙盟",
	AllianceAlreadyFullApply:                   "仙盟成员已达上限，无法申请加入",
	AllianceUserInAllianceToApply:              "您当前已有仙盟，无法加入",
	AllianceAlreadyFullInvitation:              "仙盟人数已达上限，无法邀请新成员",
	AllianceAlreadyApply:                       "已向该仙盟发送入盟申请",
	AllianceApplyNoExist:                       "仙盟加入申请不存在",
	AllianceUserNotInAlliance:                  "用户不在仙盟",
	AllianceUserPrivilegeNoEnough:              "你的权限不足",
	AllianceMengZhuNoExit:                      "盟主不能退出",
	AllianceNotInSameAlliance:                  "不在同一个仙盟",
	AllianceUserNotMengZhu:                     "不是盟主",
	AllianceUserCanNotImpeachSelf:              "不能弹劾自己",
	AllianceSceneNoOpen:                        "九霄城战还未开始",
	AllianceRefuseApply:                        "拒绝了入盟申请",
	AllianceUserCanNotKickSelf:                 "不能将自己逐出仙盟",
	AllianceUserCanNotCommitSelf:               "不能任命自己",
	AllianceAlreadyMember:                      "已经是普通成员",
	AlliancePositionAlreadyFull:                "仙盟职位满员",
	AllianceImpeachConditionNotEnough:          "盟主离线时间不足1周，无法弹劾",
	AllianceDonateMaxTimes:                     "捐献已达最大次数",
	AllianceSkillMaxLevel:                      "仙术等级已达上限",
	AllianceSkillNotOpen:                       "仙术未开启",
	AllianceSkillNotEnoughGongXian:             "仙盟贡献不足",
	AllianceInviterAlreadyInAlliance:           "被邀请人已经在仙盟",
	AllianceCreateLog:                          "%s创建了仙盟！",
	AllianceJoinLog:                            "%s加入仙盟，仙盟阵容变得更加强大！",
	AllianceExitLog:                            "%s离开了仙盟！",
	AllianceKickLog:                            "%s%s将%s逐出仙盟！",
	AllianceDonateResouceLog:                   "%s捐献了%s%s",
	AllianceImpeachLog:                         "%s通过弹劾成为了盟主！",
	AllianceCommitLog:                          "盟主%s将%s任命为仙盟%s",
	AllianceTransferLog:                        "玩家%s觉得自己能力不足，将盟主职位转让给了%s!",
	AllianceReleaseLog:                         "盟主%s撤销%s的仙盟%s职位",
	AllianceCharmSystemNotice:                  "%s昭告天下,欲广纳天下志同道合之士共同成长变强,有意者需尽快加入%s",
	AllianceKickNotice:                         "%s实力不佳！严重影响仙盟实力，已被%s逐出仙盟！",
	AllianceCommitNotice:                       "盟主%s将%s任命为仙盟%s",
	AllianceExitNotice:                         "玩家%s离开了仙盟",
	AllianceImpeachNotice:                      "玩家%s经过弹劾成为了盟主！",
	AllianceTransferNotice:                     "玩家%s将盟主转让给了玩家%s",
	AllianceConvertTimesNotEnough:              "今日兑换次数已满，请明日再来",
	AllianceDismissCanNotInSceneOpened:         "当前处于九霄城战关键时刻，不允许玩家更换仙盟",
	AllianceUserNotInAllianceForHegemon:        "九霄城战需要拥有仙盟的玩家才可参与，赶紧去加入仙盟吧",
	AllianceDismissCanNotInTransportation:      "仙盟押镖期间，无法解散仙盟",
	AllianceTransportationNotInAlliance:        "您当前还没有仙盟，无法参与仙盟押镖",
	AllianceTransportationPositionNoEnough:     "仙盟镖车仅有盟主可以发起押送，赶紧催盟主押镖吧！",
	AllianceInvitationPositionNotEnough:        "只有仙盟的管理层可以邀请他人入帮！",
	AllianceInviteMergeRefuse:                  "%s拒绝了您的仙盟合并邀请",
	AllianceTransportationAcceptNumNoEnough:    "今日仙盟押镖次数已达上限，请明日再来",
	AllianceDonateHuFuNotice:                   "%s为%s仙盟捐献虎符%s个，九霄城战又多一分胜算！",
	AllianceCreateNotice:                       "恭喜%s创建%s仙盟，望广大修真侠士加入！",
	AllianceHuangGongCloseCannotEnter:          "皇宫已经关闭,无法进入",
	AllianceSceneGuardInactive:                 "该守卫未被激活，无法参与战斗",
	AllianceTickInCrossTuLong:                  "您的仙盟正在跨服屠龙,期间无法踢人",
	AllianceExitInCrossTuLong:                  "您所在的仙盟正在跨服屠龙,期间无法退盟",
	AllianceDismissInCrossTuLong:               "您所在的仙盟正在跨服屠龙,期间无法解散仙盟",
	AllianceImpeachInCrossTuLong:               "您所在的仙盟正在跨服屠龙,期间无法弹劾",
	AllianceTransferInCrossTuLong:              "您所在的仙盟正在跨服屠龙,期间无法转让盟主",
	AllianceDepotHasNotEnoughSlot:              "仓库空间不足，无法存入物品，请通知盟主及时清理",
	AllianceDepotHasNotEnoughPoint:             "仓库积分不足，无法取出物品",
	AllianceJoinNotice:                         "欢迎%s加入仙盟！仙盟实力大增，必能一统三界！",
	AllianceJoinPersonalNotice:                 "帮会仓库里的装备有用的可以随便拿，打到有用的装备不要忘记放到帮会仓库！",
	AllianceDepotSaveNotice:                    "%s将%s放入了仙盟仓库，获取了%d点仓库积分！%s",
	AllianceMemberDeadNotice:                   "仙盟成员%s在%s%s被%s残忍的杀害了！%s如此轻视本盟，不可饶恕！%s",
	AllianceMemberCallNotice:                   "%s在%s%s发出了召集信号，想必遇到危急之事，请仙盟各位道友前往支援！%s",
	AllianceMemberCallNoticeCD:                 "仙盟召集CD中，请稍后再试",
	AllianceMemberNotRescue:                    "该地图不支持传送", // "该地图不支持仙盟求救",
	AllianceDepotTakeOutNotice:                 "%s取出了仙盟仓库中的%s，实力大增！%s",
	AllianceBatchJoinCD:                        "仙盟一键加入CD中，请稍后再试",
	AllianceDouShenJoinMailTitle:               "仙盟斗神殿成员更新",
	AllianceDouShenJoinMailContentSenior:       "本盟成员%s因实力强大而加入了斗神殿，您享受到了他带来的领域加成，战斗力大幅上升！",
	AllianceDouShenJoinMailContentJunior:       "本盟成员%s因实力强大而加入了斗神殿，但由于他的领域等级过低，%s，着实遗憾！",
	AllianceDouShenExitMailTitle:               "仙盟斗神殿成员退出",
	AllianceDouShenExitMailContent:             "本盟斗神殿成员%s离开了仙盟，由于他的离开，您失去了他的领域加成，%s！",
	AllianceHuangGongCloseCannotExit:           "皇宫已经关闭,无法退出",
	AllianceDonateItemLog:                      "%s捐献了%s个%s",
	AllianceBossSummonNoMengZhu:                "只有盟主和副盟主可以召唤,您无法召唤仙盟boss",
	AllianceBossSummonedBoss:                   "您所在的仙盟,今日已经召唤过仙盟boss了",
	AllianceBossTitle:                          "仙盟boss",
	AllianceBossContent:                        "在仙盟众位仙友的努力下,历经一场大战,成功击杀了%s,获得以下奖励,敬请查收!",
	AllianceBossEnterNoStart:                   "盟主还没召唤仙盟boss,无法进入仙盟boss",
	AllianceBossTodayFinish:                    "您所在的仙盟今日已经结束了仙盟boss",
	AllianceDismissCanNotInBoss:                "正在仙盟boss",
	AllianceBossSummonSucessChat:               "%s%s成功召唤%s,希望盟内仙友能鼎力相助%s",
	AllianceSendInvitation:                     "已向玩家发起入盟邀请",
	AllianceDismissInYuXiWar:                   "您所在的仙盟正在玉玺之战,期间无法解散仙盟",
	AllianceTickInYuXiWar:                      "您的仙盟正在玉玺之战,期间无法踢人",
	AllianceExitInYuXiWar:                      "您所在的仙盟正在玉玺之战,期间无法退盟",
	AllianceImpeachInYuXiWar:                   "您所在的仙盟正在玉玺之战,期间无法弹劾",
	AllianceTransferInYuXiWar:                  "您所在的仙盟正在玉玺之战,期间无法转让盟主",
	AllianceDepotClose:                         "仙盟仓库维护中",
	AllianceInviteAllianceMergeMemberToMuch:    "仙盟人数过多，无法合并",
	AllianceInviteAllianceMergeMemberNotOnline: "该仙盟盟主和副盟主不在线，无法邀请！",
	AllianceInviteAllianceMergeNotice:          "%s与%s强强联手进行了合并，仙盟阵容越发强大！",
	AllianceInviteAllianceMergeAllianceNotice:  "欢迎%s的兄弟加入我们的大家庭，以后大家有福同享，有难同当！一起共进退！",
	AllianceInviteAllianceMergeTitle:           "仙盟合并",
	AllianceInviteAllianceMergeContent:         "您的仙盟与%s进行了合并，您现在是%s这个大家庭的一份子了！",
	AllianceInviteAllianceMergeLog:             "%s的所有成员合并加入大家庭，仙盟阵容越发强大！",
	AllianceRenameMailTitle:                    "仙盟改名",
	AllianceRenameMailContent:                  "玩家%s成功将%s改名为%s",
	AllianceRenameNotice:                       "玩家%s成功将%s改名为%s",
	AllianceOnAllianceActivity:                 "仙盟当前处于仙盟活动中，无法操作",
	AllianceRenameSame:                         "仙盟名字是一样的",
	AllianceMergeApplyExpire:                   "合并申请已经过期",
	AllianceMergeApplyCd:                       "合并申请cd中",
	AllianceMergeApplying:                      "仙盟合并邀请中",
	AllianceCampNotSame:                        "阵营不同",
}

func init() {
	mergeLang(allianceLangMap)
}
