package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/alliance/alliance"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/center/center"
	chargetemplate "fgame/fgame/game/charge/template"
	"fgame/fgame/game/global"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorypbutil "fgame/fgame/game/inventory/pbutil"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
)

const (
	bossStatusLive = int32(0)
	bossStatusDead = int32(1)
)

func BuildAllianceMemberInfo(mem *alliance.AllianceMemberObject) *uipb.AllianceMemberInfo {
	memInfo := &uipb.AllianceMemberInfo{}
	memberId := mem.GetMemberId()
	memInfo.MemberId = &memberId
	joinTime := mem.GetJoinTime()
	memInfo.JoinTime = &joinTime
	position := int32(mem.GetPosition())
	memInfo.Position = &position
	level := mem.GetLevel()
	memInfo.Level = &level
	name := mem.GetName()
	memInfo.Name = &name
	gongXian := mem.GetGongXian()
	memInfo.GongXian = &gongXian
	force := mem.GetForce()
	memInfo.Force = &force
	zhuanSheng := mem.GetZhuanSheng()
	memInfo.ZhuanSheng = &zhuanSheng
	sex := int32(mem.GetSex())
	memInfo.Sex = &sex
	role := int32(mem.GetRole())
	memInfo.Role = &role
	onlineStatus := int32(mem.GetOnlineStatus())
	memInfo.OnlineStatus = &onlineStatus
	memVip := mem.GetVip()
	memInfo.MemberVip = &memVip

	var time int64
	now := global.GetGame().GetTimeService().Now()
	if mem.GetOnlineStatus() == playertypes.PlayerOnlineStateOffline {
		time = now - mem.GetLastLogoutTime()
	}
	memInfo.Time = &time

	return memInfo
}

func BuildAllianceBriefInfo(al *alliance.Alliance) *uipb.AllianceBriefInfo {
	briefInfo := &uipb.AllianceBriefInfo{}
	allianceId := al.GetAllianceId()
	briefInfo.AllianceId = &allianceId
	level := al.GetAllianceLevel()
	// 判断是否超过新版本最大仙盟等级
	version := getAllianceVersion()
	if version == alliancetypes.AllianceVersionTypeNew {
		maxLevel := alliancetemplate.GetAllianceTemplateService().GetAllianceMaxLevel(version)
		if level > maxLevel {
			level = maxLevel
		}
	}
	briefInfo.Level = &level
	jianShe := al.GetAllianceObject().GetJianShe()
	briefInfo.JianShe = &jianShe
	numOfMembers := al.NumOfMembers()
	briefInfo.Num = &numOfMembers
	huFu := al.GetAllianceObject().GetHuFu()
	briefInfo.HuFu = &huFu
	totalForce := al.GetAllianceObject().GetTotalForce()
	briefInfo.TotalForce = &totalForce
	name := al.GetAllianceObject().GetName()
	briefInfo.Name = &name
	createId := al.GetAllianceObject().GetMengzhuId()
	briefInfo.CreateId = &createId
	createName := al.GetMengzhuName()
	briefInfo.CreateName = &createName
	notice := al.GetAllianceObject().GetNotice()
	briefInfo.Notice = &notice
	createTime := al.GetAllianceObject().GetCreateTime()
	briefInfo.CreateTime = &createTime
	isAuto := al.GetIsAutoAgree()
	briefInfo.IsAuto = &isAuto
	isAutoRemoveDepot := al.GetAllianceObject().GetIsAutoRemoveDepot()
	briefInfo.IsAutoRemoveDepot = &isAutoRemoveDepot
	maxZhuanSheng := al.GetAllianceObject().GetMaxRemoveZhuanSheng()
	briefInfo.MaxZhuanSheng = &maxZhuanSheng
	maxQuality := int32(al.GetAllianceObject().GetMaxRemoveQuality())
	briefInfo.MaxQuality = &maxQuality
	mengzhuVip := al.GetMengzhuVip()
	briefInfo.MengzhuVip = &mengzhuVip
	briefInfo.MengzhuInfo = buildMengZhuInfo(al)
	mergeTime := al.GetAllianceObject().GetMergeAllianceTime()
	briefInfo.MergeAllianceTime = &mergeTime

	return briefInfo
}

func BuildSCAllianceMengZhuInfoNotice(al *alliance.Alliance) *uipb.SCAllianceMengZhuInfoNotice {
	scMsg := &uipb.SCAllianceMengZhuInfoNotice{}
	scMsg.MengzhuInfo = buildMengZhuInfo(al)

	return scMsg
}

func BuildAllianceInfo(al *alliance.Alliance) *uipb.AllianceInfo {
	info := &uipb.AllianceInfo{}
	info.BriefInfo = BuildAllianceBriefInfo(al)
	for _, member := range al.GetMemberList() {
		info.Members = append(info.Members, BuildAllianceMemberInfo(member))
	}
	info.DepotItemList = buildItemList(al.GetDepotItemList())

	return info
}

func BuildSCAllianceCreate(al *alliance.Alliance) *uipb.SCAllianceCreate {
	scAllianceCreate := &uipb.SCAllianceCreate{}
	scAllianceCreate.Info = BuildAllianceInfo(al)
	return scAllianceCreate
}

func BuildSCAllianceList(alList []*alliance.Alliance) *uipb.SCAllianceList {
	scAllianceList := &uipb.SCAllianceList{}
	for _, al := range alList {
		info := BuildAllianceBriefInfo(al)
		scAllianceList.AllianceList = append(scAllianceList.AllianceList, info)
	}
	return scAllianceList
}

func BuildSCAlliance(al *alliance.Alliance) *uipb.SCAlliance {
	scAlliance := &uipb.SCAlliance{}
	info := BuildAllianceBriefInfo(al)
	scAlliance.Allianc = info

	return scAlliance
}

func BuildSCAllianceMemberList(alList []*alliance.AllianceMemberObject) *uipb.SCAllianceMemberList {
	scAllianceMemberList := &uipb.SCAllianceMemberList{}
	for _, al := range alList {
		info := BuildAllianceMemberInfo(al)
		scAllianceMemberList.Members = append(scAllianceMemberList.Members, info)
	}
	return scAllianceMemberList
}

func BuildSCAllianceMemberChanged(memList []*alliance.AllianceMemberObject) *uipb.SCAllianceMemberChanged {
	scAllianceMemberChanged := &uipb.SCAllianceMemberChanged{}
	for _, mem := range memList {
		info := BuildAllianceMemberInfo(mem)
		scAllianceMemberChanged.Members = append(scAllianceMemberChanged.Members, info)
	}
	return scAllianceMemberChanged
}

func BuildSCAllianceDouShenMemberList(alList []*alliance.AllianceMemberObject) *uipb.SCAllianceDouShenList {
	scllianceDouShenList := &uipb.SCAllianceDouShenList{}
	for _, al := range alList {
		info := buildSCAllianceDouShen(al)
		scllianceDouShenList.Doushen = append(scllianceDouShenList.Doushen, info)
	}
	return scllianceDouShenList
}

func buildSCAllianceDouShen(mem *alliance.AllianceMemberObject) *uipb.DouShenInfo {
	douShenInfo := &uipb.DouShenInfo{}
	playerId := mem.GetMemberId()
	douShenInfo.PlayerId = &playerId
	name := mem.GetName()
	douShenInfo.Name = &name
	force := mem.GetForce()
	douShenInfo.Force = &force
	lingyuId := mem.GetLingyuId()
	douShenInfo.LingyuId = &lingyuId
	sex := int32(mem.GetSex())
	douShenInfo.Sex = &sex
	role := int32(mem.GetRole())
	douShenInfo.Role = &role
	level := mem.GetLevel()
	douShenInfo.Level = &level

	return douShenInfo
}

func BuildSCAllianceJoinApply(allianceId int64) *uipb.SCAllianceJoinApply {
	scAllianceJoinApply := &uipb.SCAllianceJoinApply{}
	scAllianceJoinApply.AllianceId = &allianceId
	return scAllianceJoinApply
}

func BuildSCAllianceJoinApplyBatch() *uipb.SCAllianceJoinApplyBatch {
	scMsg := &uipb.SCAllianceJoinApplyBatch{}
	return scMsg
}

func BuildJoinApplyInfo(applyObj *alliance.AllianceJoinApplyObject) *uipb.AllianceJoinApplyInfo {
	joinApplyInfo := &uipb.AllianceJoinApplyInfo{}
	name := applyObj.GetName()
	joinId := applyObj.GetJoinId()
	force := applyObj.GetForce()
	level := applyObj.GetLevel()
	sex := int32(applyObj.GetSex())
	role := int32(applyObj.GetRole())

	joinApplyInfo.Name = &name
	joinApplyInfo.Sex = &sex
	joinApplyInfo.Role = &role
	joinApplyInfo.MemberId = &joinId
	joinApplyInfo.Force = &force
	joinApplyInfo.Level = &level
	return joinApplyInfo
}

func BuildSCAllianceJoinApplyBroadcast(applyObj *alliance.AllianceJoinApplyObject) *uipb.SCAllianceJoinApplyBroadcast {
	scAllianceJoinApplyBroadcast := &uipb.SCAllianceJoinApplyBroadcast{}
	scAllianceJoinApplyBroadcast.ApplyInfo = BuildJoinApplyInfo(applyObj)
	return scAllianceJoinApplyBroadcast
}

func BuildSCAllianceDismiss(allianceId int64) *uipb.SCAllianceDismiss {
	scAllianceDismiss := &uipb.SCAllianceDismiss{}
	scAllianceDismiss.AllianceId = &allianceId
	return scAllianceDismiss
}

func BuildSCAllianceAgreeJoinApply(joinId int64, agree bool) *uipb.SCAllianceAgreeJoinApply {
	scAllianceAgreeJoinApply := &uipb.SCAllianceAgreeJoinApply{}
	scAllianceAgreeJoinApply.JoinId = &joinId
	scAllianceAgreeJoinApply.Agree = &agree
	return scAllianceAgreeJoinApply
}

func BuildSCAllianceAgreeJoinApplyToApply(allianceId int64, name string, agree bool) *uipb.SCAllianceAgreeJoinApplyToApply {
	scAllianceAgreeJoinApplyToApply := &uipb.SCAllianceAgreeJoinApplyToApply{}
	scAllianceAgreeJoinApplyToApply.AllianceId = &allianceId
	scAllianceAgreeJoinApplyToApply.Name = &name
	scAllianceAgreeJoinApplyToApply.Agree = &agree
	return scAllianceAgreeJoinApplyToApply
}

func BuildSCAllianceAgreeJoinApplyToManager(joinId int64, agree bool) *uipb.SCAllianceAgreeJoinApplyToManager {
	scAllianceAgreeJoinApplyToManager := &uipb.SCAllianceAgreeJoinApplyToManager{}
	scAllianceAgreeJoinApplyToManager.JoinId = &joinId
	scAllianceAgreeJoinApplyToManager.Agree = &agree
	return scAllianceAgreeJoinApplyToManager
}

func BuildSCAllianceJoinApplyList(applyObjList []*alliance.AllianceJoinApplyObject) *uipb.SCAllianceJoinApplyList {
	scAllianceJoinApplyList := &uipb.SCAllianceJoinApplyList{}
	for _, applyObj := range applyObjList {
		scAllianceJoinApplyList.ApplyList = append(scAllianceJoinApplyList.ApplyList, BuildJoinApplyInfo(applyObj))
	}
	return scAllianceJoinApplyList
}

var (
	scAllianceExit = &uipb.SCAllianceExit{}
)

func BuildSCAllianceExit() *uipb.SCAllianceExit {

	return scAllianceExit
}

func BuildSCAllianceKick(kickMemberId int64, kickMemberName string) *uipb.SCAllianceKick {
	scAllianceKick := &uipb.SCAllianceKick{}
	scAllianceKick.KickMemberId = &kickMemberId
	scAllianceKick.KickMemberName = &kickMemberName
	return scAllianceKick
}

func BuildSCAllianceInvitation(invitationId int64) *uipb.SCAllianceInvitation {
	scAllianceInvitation := &uipb.SCAllianceInvitation{}
	scAllianceInvitation.InvitationId = &invitationId
	return scAllianceInvitation
}

func BuildSCAllianceCommit(commitMemberId int64, pos alliancetypes.AlliancePosition) *uipb.SCAllianceCommit {
	scAllianceCommit := &uipb.SCAllianceCommit{}
	posInt := int32(pos)
	scAllianceCommit.CommitMemberId = &commitMemberId
	scAllianceCommit.Position = &posInt
	return scAllianceCommit
}
func BuildSCAllianceCommitNotice(playerId int64, name string, commitPosition alliancetypes.AlliancePosition) *uipb.SCAllianceCommitNotice {
	scAllianceCommitNotice := &uipb.SCAllianceCommitNotice{}
	commitPosInt := int32(commitPosition)

	scAllianceCommitNotice.MemberId = &playerId
	scAllianceCommitNotice.MemberName = &name
	scAllianceCommitNotice.Positio = &commitPosInt

	return scAllianceCommitNotice
}

func BuildSCAllianceTransfer(transferMemberId int64) *uipb.SCAllianceTransfer {
	scAllianceTransfer := &uipb.SCAllianceTransfer{}
	scAllianceTransfer.TransferMemberId = &transferMemberId

	return scAllianceTransfer
}

func BuildSCAllianceTransferBroadcast(playerId, transferId int64, memberName, transferName string) *uipb.SCAllianceTransferBroadcast {
	scAllianceTransferBroadcast := &uipb.SCAllianceTransferBroadcast{}

	scAllianceTransferBroadcast.MemberId = &playerId
	scAllianceTransferBroadcast.MemberName = &memberName
	scAllianceTransferBroadcast.TransferId = &transferId
	scAllianceTransferBroadcast.TransferName = &transferName

	return scAllianceTransferBroadcast
}

func BuildSCAllianceDismissBroadcast(allianceId int64) *uipb.SCAllianceDismissBroadcast {
	scAllianceDismissBroadcast := &uipb.SCAllianceDismissBroadcast{}
	scAllianceDismissBroadcast.AllianceId = &allianceId

	return scAllianceDismissBroadcast
}

func BuildSCAllianceImpeacheBroadcast(memberId int64, memberName string) *uipb.SCAllianceImpeachBroadcast {
	scAllianceImpeachBroadcast := &uipb.SCAllianceImpeachBroadcast{}

	scAllianceImpeachBroadcast.MemberId = &memberId
	scAllianceImpeachBroadcast.MemberName = &memberName

	return scAllianceImpeachBroadcast
}

func BuildSCAllianceDepotSettingNotice(isAuto, zhuansheng, quality int32) *uipb.SCAllianceDepotSettingNotice {
	scMsg := &uipb.SCAllianceDepotSettingNotice{}
	scMsg.IsAuto = &isAuto
	scMsg.MaxZhuanSheng = &zhuansheng
	scMsg.MaxQuality = &quality

	return scMsg
}

var (
	scAllianceImpeach = &uipb.SCAllianceImpeach{}
)

func BuildSCAllianceImpeach() *uipb.SCAllianceImpeach {
	scAllianceImpeach := &uipb.SCAllianceImpeach{}

	return scAllianceImpeach
}

func BuildSCAllianceInfo(al *alliance.Alliance, mem *alliance.AllianceMemberObject) *uipb.SCAllianceInfo {
	scAllianceInfo := &uipb.SCAllianceInfo{}
	scAllianceInfo.Info = BuildAllianceInfo(al)

	applyNum := int32(0)
	if mem.GetPosition() != alliancetypes.AlliancePositionMember {
		applyNum = int32(len(al.GetApplyList()))
	}
	scAllianceInfo.ApplyNum = &applyNum
	return scAllianceInfo
}

var (
	scAllianceCharm = &uipb.SCAllianceCharm{}
)

func BuildSCAllianceCharm() *uipb.SCAllianceCharm {
	scAllianceCharm := &uipb.SCAllianceCharm{}
	return scAllianceCharm
}

func BuildSCAllianceDonate(typ alliancetypes.AllianceJuanXianType) *uipb.SCAllianceDonate {
	scAllianceDonate := &uipb.SCAllianceDonate{}
	typInt := int32(typ)
	scAllianceDonate.Typ = &typInt
	return scAllianceDonate
}

var (
	scAllianceDonateHuFu = &uipb.SCAllianceDonateHuFu{}
)

func BuildSCAllianceDonateHuFu() *uipb.SCAllianceDonateHuFu {
	return scAllianceDonateHuFu
}

func BuildSCAllianceInvitationNotice(allianceId int32, allianceName string, memId int64, memName string) *uipb.SCAllianceInvitationNotice {
	now := global.GetGame().GetTimeService().Now()

	scAllianceAgreeInvitation := &uipb.SCAllianceInvitationNotice{}
	scAllianceAgreeInvitation.AllianceId = &allianceId
	scAllianceAgreeInvitation.AllianceName = &allianceName
	scAllianceAgreeInvitation.MemberId = &memId
	scAllianceAgreeInvitation.MemberName = &memName
	scAllianceAgreeInvitation.InvitationTime = &now
	return scAllianceAgreeInvitation
}

func BuildSCAllianceAgreeInvitationNotice(invitationId int64, invitationName string, agree bool) *uipb.SCAllianceAgreeInvitationNotice {
	scAllianceAgreeInvitationNotice := &uipb.SCAllianceAgreeInvitationNotice{}
	scAllianceAgreeInvitationNotice.InvitationId = &invitationId
	scAllianceAgreeInvitationNotice.InvitationName = &invitationName
	scAllianceAgreeInvitationNotice.Agree = &agree

	return scAllianceAgreeInvitationNotice
}

func BuildSCAllianceAgreeInvitation(memberId int64, agree bool) *uipb.SCAllianceAgreeInvitation {
	scAllianceAgreeInvitation := &uipb.SCAllianceAgreeInvitation{}
	scAllianceAgreeInvitation.MemberId = &memberId
	scAllianceAgreeInvitation.Agree = &agree

	return scAllianceAgreeInvitation
}

func BuildAllianceDonate(typ alliancetypes.AllianceJuanXianType, num int32) *uipb.AllianceDonate {
	allianceDonate := &uipb.AllianceDonate{}
	typInt := int32(typ)
	allianceDonate.Typ = &typInt
	allianceDonate.Num = &num
	return allianceDonate
}

func BuildSCAlliancePlayerInfo(obj *playeralliance.PlayerAllianceObject, skillObjArr map[alliancetypes.AllianceSkillType]*playeralliance.PlayerAllianceSkillObject) *uipb.SCAlliancePlayerInfo {
	scAlliancePlayerInfo := &uipb.SCAlliancePlayerInfo{}
	currentGongXian := obj.GetCurrentGongXian()
	scAlliancePlayerInfo.CurrentGongXian = &currentGongXian
	for typ, num := range obj.GetDonateMap() {
		donate := BuildAllianceDonate(typ, num)
		scAlliancePlayerInfo.DonateList = append(scAlliancePlayerInfo.DonateList, donate)
	}
	for _, skillObj := range skillObjArr {
		scAlliancePlayerInfo.SkillInfoList = append(scAlliancePlayerInfo.SkillInfoList, buildSkillInfo(skillObj))
	}
	yaoPai := obj.GetYaoPai()
	scAlliancePlayerInfo.CurrentYaoPai = &yaoPai
	hadConvertTimes := obj.GetConvertTimes()
	scAlliancePlayerInfo.HadConvertTimes = &hadConvertTimes
	depotPoint := obj.GetDepotPoint()
	scAlliancePlayerInfo.CurDepotPoint = &depotPoint
	callTime := obj.GetLastMemberCallTime()
	scAlliancePlayerInfo.LastMemberCallTime = &callTime
	return scAlliancePlayerInfo
}

func BuildSCAllianceSkillUpgrade(id int32) *uipb.SCAllianceSkillUpgrade {
	scAllianceSkillUpgrade := &uipb.SCAllianceSkillUpgrade{}
	scAllianceSkillUpgrade.TemId = &id

	return scAllianceSkillUpgrade
}

func buildSkillInfo(obj *playeralliance.PlayerAllianceSkillObject) *uipb.AllianceSkillInfo {
	allianceSkillInfo := &uipb.AllianceSkillInfo{}
	level := obj.GetLevel()
	allianceSkillInfo.Level = &level
	typ := int32(obj.GetSkillType())
	allianceSkillInfo.SkillType = &typ

	return allianceSkillInfo
}

func BuildSCAllianceNoticeBroadcast(content string) *uipb.SCAllianceNoticeBroadcast {
	scAllianceNoticeBroadcast := &uipb.SCAllianceNoticeBroadcast{}
	scAllianceNoticeBroadcast.Content = &content

	return scAllianceNoticeBroadcast
}

func BuildSCAllianceNoticeChange(content string) *uipb.SCAllianceNoticeChange {
	scAllianceNoticeChange := &uipb.SCAllianceNoticeChange{}
	scAllianceNoticeChange.Content = &content

	return scAllianceNoticeChange
}

func BuildSCAllianceLogList(alLogList []*alliance.AllianceLogObject) *uipb.SCAllianceLog {
	scAllianceLog := &uipb.SCAllianceLog{}
	for _, logObj := range alLogList {
		log := buildAllianceLog(logObj)
		scAllianceLog.LogList = append(scAllianceLog.LogList, log)
	}
	return scAllianceLog
}

func buildAllianceLog(logObj *alliance.AllianceLogObject) *uipb.AllianceLog {
	allianceLog := &uipb.AllianceLog{}
	createTiem := logObj.GetCreateTime()
	content := logObj.GetContent()
	allianceLog.Content = &content
	allianceLog.CreateTime = &createTiem

	return allianceLog
}

func BuildSCAlliancePlayerYaoPaiChanged(yaoPai int32) *uipb.SCAlliancePlayerYaoPaiChanged {
	scAlliancePlayerYaoPaiChanged := &uipb.SCAlliancePlayerYaoPaiChanged{}
	scAlliancePlayerYaoPaiChanged.CurrentYaoPai = &yaoPai

	return scAlliancePlayerYaoPaiChanged
}

func BuildSCAllianceSceneHuFuChanged(hufu, donorId int64, donorName string) *uipb.SCAllianceDonateHuFuBroadcast {
	scAllianceDonateHuFuBroadcast := &uipb.SCAllianceDonateHuFuBroadcast{}
	scAllianceDonateHuFuBroadcast.HuFu = &hufu
	scAllianceDonateHuFuBroadcast.DonorId = &donorId
	scAllianceDonateHuFuBroadcast.DonorName = &donorName

	return scAllianceDonateHuFuBroadcast
}

func BuildSCAllianceKickNotice(memId int64, memName string) *uipb.SCAllianceKickNotice {
	scAllianceKickNotice := &uipb.SCAllianceKickNotice{}
	scAllianceKickNotice.MemberId = &memId
	scAllianceKickNotice.MemberName = &memName

	return scAllianceKickNotice
}

func BuildSCYaoPaiConvert(itemId, num int32) *uipb.SCYaoPaiConvert {
	scYaoPaiConvert := &uipb.SCYaoPaiConvert{}

	scYaoPaiConvert.ItemId = &itemId
	scYaoPaiConvert.Num = &num

	return scYaoPaiConvert
}

func BuildSCAllianceAgreeJoinApplyBatch() *uipb.SCAllianceAgreeJoinApplyBatch {
	scAllianceAgreeJoinApplyBatch := &uipb.SCAllianceAgreeJoinApplyBatch{}
	return scAllianceAgreeJoinApplyBatch
}

func BuildSCAllianceDouShenLingyuChangedBroadcast(lingyuId int32, pl player.Player) *uipb.SCAllianceDouShenLingyuChangedBroadcast {
	scAllianceDouShenLingyuChangedBroadcast := &uipb.SCAllianceDouShenLingyuChangedBroadcast{}
	scAllianceDouShenLingyuChangedBroadcast.LingyuId = &lingyuId
	name := pl.GetName()
	scAllianceDouShenLingyuChangedBroadcast.Name = &name
	playerId := pl.GetId()
	scAllianceDouShenLingyuChangedBroadcast.PlayerId = &playerId
	sex := int32(pl.GetSex())
	scAllianceDouShenLingyuChangedBroadcast.Sex = &sex
	role := int32(pl.GetRole())
	scAllianceDouShenLingyuChangedBroadcast.Role = &role

	return scAllianceDouShenLingyuChangedBroadcast
}

func BuildSCSaveInAllianceDepot(curPoint int32) *uipb.SCSaveInAllianceDepot {
	scMsg := &uipb.SCSaveInAllianceDepot{}
	scMsg.TotalPoint = &curPoint
	return scMsg
}

func BuildSCTakeOutAllianceDepot(curPoint int32) *uipb.SCTakeOutAllianceDepot {
	scMsg := &uipb.SCTakeOutAllianceDepot{}
	scMsg.TotalPoint = &curPoint
	return scMsg
}

func BuildSCAllianceDepotChangedNotice(itemList []*alliance.AllianceDepotItemObject) *uipb.SCAllianceDepotChangedNotice {
	scMsg := &uipb.SCAllianceDepotChangedNotice{}
	scMsg.ChangedItemList = buildItemList(itemList)
	return scMsg
}

func BuildSCAllianceDepotMergeNotice(itemList []*alliance.AllianceDepotItemObject) *uipb.SCAllianceDepotMergeNotice {
	scMsg := &uipb.SCAllianceDepotMergeNotice{}
	scMsg.ChangedItemList = buildItemList(itemList)
	return scMsg
}

func BuildSCAllianceDepotMerge() *uipb.SCAllianceDepotMerge {
	scMsg := &uipb.SCAllianceDepotMerge{}
	return scMsg
}

func BuildSCAllianceMemberCall(callTime int64, callType int32) *uipb.SCAllianceMemberCall {
	scMsg := &uipb.SCAllianceMemberCall{}
	scMsg.LastMemberCallTime = &callTime
	scMsg.CallType = &callType
	return scMsg
}

func BuildSCAllianceMemberCallBroadcast(playerName string, mapId int32, pos coretypes.Position, callType int32) *uipb.SCAllianceMemberCallBroadcast {
	scMsg := &uipb.SCAllianceMemberCallBroadcast{}
	scMsg.CallPlayerName = &playerName
	scMsg.Pos = buildAlliancePos(pos)
	scMsg.MapId = &mapId
	scMsg.CallType = &callType
	return scMsg
}

func BuildSCAllianceMemberRescue() *uipb.SCAllianceMemberRescue {
	scMsg := &uipb.SCAllianceMemberRescue{}
	return scMsg
}

func BuildSCAllianceMemberPos(pl scene.Player) *uipb.SCAllianceMemberPos {
	allianceMemberPos := &uipb.SCAllianceMemberPos{}
	s := pl.GetScene()
	allPlayers := s.GetAllPlayers()
	for _, spl := range allPlayers {
		if spl.GetId() == pl.GetId() {
			continue
		}
		if spl.GetAllianceId() != pl.GetAllianceId() {
			continue
		}
		allianceMemberPos.MemberPosList = append(allianceMemberPos.MemberPosList, buildMemberPos(spl))
	}
	return allianceMemberPos
}

func BuildSCAllianceAutoAgreeJoin(isAuto int32) *uipb.SCAllianceAutoAgreeJoin {
	scMsg := &uipb.SCAllianceAutoAgreeJoin{}
	scMsg.IsAuto = &isAuto
	return scMsg
}

func BuildSCAllianceDepotAutoRemove(isAuto, zhuansheng, quality int32) *uipb.SCAllianceDepotAutoRemove {
	scMsg := &uipb.SCAllianceDepotAutoRemove{}
	scMsg.IsAuto = &isAuto
	scMsg.MaxZhuanSheng = &zhuansheng
	scMsg.MaxQuality = &quality
	return scMsg
}

func buildMemberPos(pl scene.Player) *uipb.AllianceMemberPos {
	allianceMemberPos := &uipb.AllianceMemberPos{}
	playerId := pl.GetId()
	pos := pl.GetPos()
	allianceMemberPos.PlayerId = &playerId
	allianceMemberPos.Pos = buildAlliancePos(pos)
	return allianceMemberPos
}

func buildItemList(itemList []*alliance.AllianceDepotItemObject) []*uipb.AllianceSlotItem {
	var allianceSlotList []*uipb.AllianceSlotItem
	for _, item := range itemList {
		allianceSlotList = append(allianceSlotList, buildItem(item))
	}
	return allianceSlotList
}

func buildItem(item *alliance.AllianceDepotItemObject) *uipb.AllianceSlotItem {
	slotItem := &uipb.AllianceSlotItem{}
	itemId := item.GetItemId()
	num := item.GetNum()
	index := item.GetIndex()
	level := item.GetLevel()
	bindInt := int32(item.GetBindType())

	slotItem.ItemId = &itemId
	slotItem.Num = &num
	slotItem.Index = &index
	slotItem.Level = &level
	slotItem.BindType = &bindInt
	slotItem.PropertyData = inventorypbutil.BuildItemPropertyData(item.GetPropertyData())
	return slotItem
}

func buildGoldEquipProperty(data *goldequiptypes.GoldEquipPropertyData) *uipb.GoldEquipProperty {
	openLightLevel := data.OpenLightLevel
	upstarLevel := data.UpstarLevel

	info := &uipb.GoldEquipProperty{}
	info.OpenLightLevel = &openLightLevel
	info.UpstarLevel = &upstarLevel
	info.AttrList = data.AttrList
	return info
}

func buildBaseProperty(expireType int32, expireTime, itemGetTime int64) *uipb.BaseProperty {
	info := &uipb.BaseProperty{}
	info.ExpireType = &expireType
	info.ExpireTime = &expireTime
	info.ItemGetTime = &itemGetTime
	return info
}

func BuildSCAllianceBossSummon() *uipb.SCAllianceBossSummon {
	scAllianceBossSummon := &uipb.SCAllianceBossSummon{}
	return scAllianceBossSummon
}

func BuildSCAllianceBossEnter(sd alliancebossscene.AllianceBossSceneData) *uipb.SCAllianceBossEnter {
	scAllianceBossEnter := &uipb.SCAllianceBossEnter{}
	scAllianceBossEnter.SceneInfo = buildAllianceBossSceneInfo(sd)
	return scAllianceBossEnter
}

func BuildSCAllianceBossRank(sd alliancebossscene.AllianceBossSceneData) *uipb.SCAllianceBossRank {
	scAllianceBossRank := &uipb.SCAllianceBossRank{}
	rankList := sd.GetRankTop()
	for _, rankInfo := range rankList {
		playerId := rankInfo.GetPlayerId()
		playerName := rankInfo.GetPlayerName()
		damage := rankInfo.GetDamage()
		scAllianceBossRank.RankList = append(scAllianceBossRank.RankList, buildAllianceBossDamage(playerId, playerName, damage))
	}
	return scAllianceBossRank
}

func BuildSCAllianceBossChanged(level int32, npc scene.NPC) *uipb.SCAllianceBossChanged {
	scAllianceBossChanged := &uipb.SCAllianceBossChanged{}
	scAllianceBossChanged.BossInfo = buildAllianceBossInfo(level, npc)
	return scAllianceBossChanged
}

func BuildSCAllianceBossEnd(sucess bool) *uipb.SCAllianceBossEnd {
	scAllianceBossEnd := &uipb.SCAllianceBossEnd{}
	scAllianceBossEnd.Sucess = &sucess
	return scAllianceBossEnd
}

func BuildSCAllianceBoss(status int32, level int32, exp int32, summonTime int64) *uipb.SCAllianceBoss {
	scAllianceBoss := &uipb.SCAllianceBoss{}
	scAllianceBoss.BossStatus = &status
	scAllianceBoss.Level = &level
	scAllianceBoss.Exp = &exp
	scAllianceBoss.SummonTime = &summonTime
	return scAllianceBoss
}

func BuildSCAllianceBossSummonSucess() *uipb.SCAllianceBossSummonSucess {
	scAllianceBossSummonSucess := &uipb.SCAllianceBossSummonSucess{}
	return scAllianceBossSummonSucess
}

func BuildSCAllianceInviteMerge(inviteAllianceId int64, inviteAllianceName string) *uipb.SCAllianceInviteMerge {
	scMsg := &uipb.SCAllianceInviteMerge{}
	scMsg.InviteAllianceId = &inviteAllianceId
	scMsg.InviteAllianceName = &inviteAllianceName
	return scMsg
}

func BuildSCAllianceInviteMergeNotice(al *alliance.Alliance) *uipb.SCAllianceInviteMergeNotice {
	scMsg := &uipb.SCAllianceInviteMergeNotice{}
	mengzhuName := al.GetMengzhuName()
	allianceId := al.GetAllianceId()
	allianceName := al.GetAllianceName()
	scMsg.PlayerName = &mengzhuName
	scMsg.AllianceId = &allianceId
	scMsg.AllianceName = &allianceName
	return scMsg
}

func BuildSCAllianceInviteMergeFeedback() *uipb.SCAllianceInviteMergeFeedback {
	scMsg := &uipb.SCAllianceInviteMergeFeedback{}
	return scMsg
}

func buildAllianceBossSceneInfo(sd alliancebossscene.AllianceBossSceneData) *uipb.AllianceBossSceneInfo {
	allianceBossSceneInfo := &uipb.AllianceBossSceneInfo{}
	summonTime := sd.GetSummonTime()
	allianceBossSceneInfo.SummonTime = &summonTime
	level := sd.GetLevel()
	allianceBossSceneInfo.BossInfo = buildAllianceBossInfo(level, sd.GetBossNpc())
	rankList := sd.GetRankTop()
	for _, rankInfo := range rankList {
		playerId := rankInfo.GetPlayerId()
		playerName := rankInfo.GetPlayerName()
		damage := rankInfo.GetDamage()
		allianceBossSceneInfo.RankList = append(allianceBossSceneInfo.RankList, buildAllianceBossDamage(playerId, playerName, damage))
	}
	return allianceBossSceneInfo

}

func buildAllianceBossInfo(level int32, npc scene.NPC) *uipb.AllianceBossInfo {
	allianceBossInfo := &uipb.AllianceBossInfo{}
	allianceBossInfo.Level = &level

	hp := npc.GetHP()
	maxHP := npc.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	status := bossStatusLive
	if hp == 0 {
		status = bossStatusDead
	}
	allianceBossInfo.Status = &status
	allianceBossInfo.Hp = &hp
	allianceBossInfo.MaxHp = &maxHP
	return allianceBossInfo
}

func buildAllianceBossDamage(playerId int64, playerName string, damage int64) *uipb.AllianceBossRank {
	allianceBossRank := &uipb.AllianceBossRank{}
	allianceBossRank.PlayerId = &playerId
	allianceBossRank.PlayerName = &playerName
	allianceBossRank.Damage = &damage
	return allianceBossRank
}

func buildMengZhuInfo(al *alliance.Alliance) *uipb.MengZhuInfo {
	info := &uipb.MengZhuInfo{}
	weaponId := al.GetMengzhuWeaponId()
	fashionId := al.GetMengzhuFashionId()
	wingId := al.GetMengzhuWingId()

	info.WingId = &wingId
	info.FashionId = &fashionId
	info.WeaponId = &weaponId

	return info
}

func buildAlliancePos(pos coretypes.Position) *uipb.AlliancePosition {
	posion := &uipb.AlliancePosition{}
	x := float32(pos.X)
	y := float32(pos.Y)
	z := float32(pos.Z)
	posion.PosX = &x
	posion.PosY = &y
	posion.PosZ = &z
	return posion
}

// 获取仙盟版本
func getAllianceVersion() alliancetypes.AllianceVersionType {
	versionType := alliancetypes.AllianceVersionTypeOld
	sdkList := center.GetCenterService().GetSdkList()
	for _, sdk := range sdkList {
		temp := chargetemplate.GetChargeTemplateService().GetQuDaoTemplateByType(sdk)
		if temp == nil {
			continue
		}

		if temp.GetAllianceVersion() == alliancetypes.AllianceVersionTypeNew {
			versionType = alliancetypes.AllianceVersionTypeNew
			break
		}
	}
	return versionType
}
