package alliance

import (
	"fgame/fgame/common/lang"
	gamecommon "fgame/fgame/game/common/common"
)

var (
	//仙盟不存在
	errorAllianceNoExist = gamecommon.CodeError(lang.AllianceNoExist)
	//仙盟名字存在了
	errorAllianceNameExist = gamecommon.CodeError(lang.AllianceNameExist)
	//用户已经在仙盟了
	errorAllianceUserAlreadyInAlliance = gamecommon.CodeError(lang.AllianceUserAlreadyInAlliance)
	//用户已经在仙盟了申请加入
	errorAllianceUserInAllianceToApply = gamecommon.CodeError(lang.AllianceUserInAllianceToApply)
	//仙盟已经满人
	errorAllianceAlreadyFullApply = gamecommon.CodeError(lang.AllianceAlreadyFullApply)
	//已经申请过了
	errorAllianceAlreadyApply = gamecommon.CodeError(lang.AllianceAlreadyApply)
	//申请不存在
	errorAllianceApplyNoExist = gamecommon.CodeError(lang.AllianceApplyNoExist)
	//邀请不存在
	errorAllianceInvitationNoExist = gamecommon.CodeError(lang.AllianceApplyNoExist)
	//权限不足
	errorAlliancePrivilegeNotEnough = gamecommon.CodeError(lang.AllianceUserPrivilegeNoEnough)
	//用户不在仙盟
	errorAllianceUserNotInAlliance = gamecommon.CodeError(lang.AllianceUserNotInAlliance)
	//盟主不能退出
	errorAllianceMengZhuNoExit = gamecommon.CodeError(lang.AllianceMengZhuNoExit)
	//不是在同一个仙盟
	errorAllianceNotInSameAlliance = gamecommon.CodeError(lang.AllianceMengZhuNoExit)
	//不是盟主
	errorAllianceUserNotMengZhu = gamecommon.CodeError(lang.AllianceUserNotMengZhu)
	//不能弹劾自己
	errorAllianceUserCanNotImpeachSelf = gamecommon.CodeError(lang.AllianceUserCanNotImpeachSelf)
	//不能踢自己
	errorAllianceUserCanNotKickSelf = gamecommon.CodeError(lang.AllianceUserCanNotImpeachSelf)
	//不能任命自己
	errorAllianceUserCanNotCommitSelf = gamecommon.CodeError(lang.AllianceUserCanNotCommitSelf)
	//已经是普通成员
	errorAllianceAlreadyMember = gamecommon.CodeError(lang.AllianceAlreadyMember)
	//职位满员
	errorAlliancePositionAlreadyFull = gamecommon.CodeError(lang.AlliancePositionAlreadyFull)
	//弹劾条件不足
	errorAllianceImpeachConditionNotEnough = gamecommon.CodeError(lang.AllianceImpeachConditionNotEnough)
	//被邀请人已经在仙盟了
	errorAllianceInviterAlreadyInAlliance = gamecommon.CodeError(lang.AllianceInviterAlreadyInAlliance)
	//正在开启城战
	errorAllianceDismissCanNotInSceneOpened = gamecommon.CodeError(lang.AllianceDismissCanNotInSceneOpened)
	//正在仙盟押镖
	errorAllianceDismissCanNotInTransportation = gamecommon.CodeError(lang.AllianceDismissCanNotInTransportation)
	//领取仙盟镖
	errorAllianceTransportationNotInAlliance = gamecommon.CodeError(lang.AllianceTransportationNotInAlliance)
	//领取仙盟镖权限不足
	errorAllianceTransportationPositionNoEnough = gamecommon.CodeError(lang.AllianceTransportationPositionNoEnough)
	//邀请加入
	errorAllianceAlreadyFullInvitation = gamecommon.CodeError(lang.AllianceAlreadyFullInvitation)
	//仙盟押镖次数不足
	errorAllianceTransportationAcceptNumNoEnough = gamecommon.CodeError(lang.AllianceTransportationAcceptNumNoEnough)
	// 仙盟仓库空间不足
	errorAllianceDepotSlotNoEnough = gamecommon.CodeError(lang.AllianceDepotHasNotEnoughSlot)
	// 仙盟仓库物品不存在
	errorAllianceDepotItemNotExist = gamecommon.CodeError(lang.InventoryItemNoExist)
	// 仙盟仓库数量不足
	errorAllianceDepotItemNotEnough = gamecommon.CodeError(lang.InventoryItemNoEnough)
	// 仙盟boss还没有召唤
	errorAllianceBossEnterNoStart = gamecommon.CodeError(lang.AllianceBossEnterNoStart)
	// 您所在的仙盟今日已经结束了仙盟boss
	errorAllianceBossTodayFinish = gamecommon.CodeError(lang.AllianceBossTodayFinish)
	//正在仙盟boss
	errorAllianceDismissCanNotInBoss = gamecommon.CodeError(lang.AllianceDismissCanNotInBoss)
	//您不是盟主,无法召唤仙盟boss
	errorAllianceBossSummonNoMengZhu = gamecommon.CodeError(lang.AllianceBossSummonNoMengZhu)
	//您所在的仙盟,今日已经召唤过仙盟boss了
	ErrorAllianceBossSummonedBoss = gamecommon.CodeError(lang.AllianceBossSummonedBoss)
	//合并邀请不存在
	errorAllianceMergeApplyExpired = gamecommon.CodeError(lang.AllianceMergeApplyExpire)
	//合并邀请中
	errorAllianceMergeApplying = gamecommon.CodeError(lang.AllianceMergeApplying)
	//合并邀请cd
	errorAllianceMergeApplyCd = gamecommon.CodeError(lang.AllianceMergeApplyCd)
	// 阵营不同
	errAllianceCampTypeNotSame = gamecommon.CodeError(lang.AllianceCampNotSame)
)
