package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/jieyi/jieyi"
)

func BuildSCJieYiPlayerInfo(isJieYi int32, name string) *uipb.SCJieYiPlayerInfo {
	scMsg := &uipb.SCJieYiPlayerInfo{}
	scMsg.IsJieYi = &isJieYi
	scMsg.Name = &name
	return scMsg
}

func BuildSCJieYiInfoOnChange(obj *jieyi.JieYiObject) *uipb.SCJieYiInfoOnChange {
	scMsg := &uipb.SCJieYiInfoOnChange{}
	id := obj.GetDBId()
	name := obj.GetName()
	num := jieyi.GetJieYiService().GetJieYiMemberNum(id)
	scMsg.JieYiId = &id
	scMsg.Name = &name
	scMsg.PlayerNum = &num
	return scMsg
}

func BuildSCJieYiTokenActivite(token int32) *uipb.SCJieYiTokenActivite {
	scMsg := &uipb.SCJieYiTokenActivite{}
	scMsg.Token = &token
	return scMsg
}

func BuildSCJieYiTokenChange(token int32, level int32) *uipb.SCJieYiTokenChange {
	scMsg := &uipb.SCJieYiTokenChange{}
	scMsg.Token = &token
	scMsg.Level = &level
	return scMsg
}

func BuildSCJieYiDaoJuChange(daoJu int32) *uipb.SCJieYiDaoJuChange {
	scMsg := &uipb.SCJieYiDaoJuChange{}
	scMsg.DaoJu = &daoJu
	return scMsg
}

func BuildSCJieYiPost(postTime int64, leaveWord string) *uipb.SCJieYiPost {
	scMsg := &uipb.SCJieYiPost{}
	scMsg.PostTime = &postTime
	scMsg.LeaveWord = &leaveWord
	return scMsg
}

func BuildSCJieYiTokenGiveNotice(playerId int64, playerName string, leaveWord string, token int32, lastToken int32) *uipb.SCJieYiTokenGiveNotice {
	scMsg := &uipb.SCJieYiTokenGiveNotice{}
	scMsg.PlayerId = &playerId
	scMsg.PlayerName = &playerName
	scMsg.LeaveWord = &leaveWord
	scMsg.Token = &token
	scMsg.LastToken = &lastToken
	return scMsg
}

func BuildSCJieYiLaoDaTiRen(playerId int64) *uipb.SCJieYiLaoDaTiRen {
	scMsg := &uipb.SCJieYiLaoDaTiRen{}
	scMsg.PlayerId = &playerId
	return scMsg
}

func BuildSCJieYiLaoDaTiRenNotice(playerId int64, playerName string, laoDaId int64, laoDaName string) *uipb.SCJieYiLaoDaTiRenNotice {
	scMsg := &uipb.SCJieYiLaoDaTiRenNotice{}
	scMsg.PlayerId = &playerId
	scMsg.PlayerName = &playerName
	scMsg.LaoDaId = &laoDaId
	scMsg.LaoDaName = &laoDaName
	return scMsg
}

func BuildSCJieYiTokenSuoYao(playerId int64, token int32, leaveWord string) *uipb.SCJieYiTokenSuoYao {
	scMsg := &uipb.SCJieYiTokenSuoYao{}
	scMsg.PlayerId = &playerId
	scMsg.Token = &token
	scMsg.LeaveWord = &leaveWord
	return scMsg
}

func BuildSCJieYiTokenSuoYaoNotice(token int32, playerId int64, playerName string, leaveWord string) *uipb.SCJieYiTokenSuoYaoNotice {
	scMsg := &uipb.SCJieYiTokenSuoYaoNotice{}
	scMsg.Token = &token
	scMsg.PlayerId = &playerId
	scMsg.PlayerName = &playerName
	scMsg.LeaveWord = &leaveWord
	return scMsg
}

func BuildSCJieYiLaoDaTiRenOtherNotice(playerId int64, playerName string) *uipb.SCJieYiLaoDaTiRenOtherNotice {
	scMsg := &uipb.SCJieYiLaoDaTiRenOtherNotice{}
	scMsg.LaoDaId = &playerId
	scMsg.LaoDaName = &playerName
	return scMsg
}

func BuildSCJieYiHandleTokenSuoYao(playerId int64, token int32) *uipb.SCJieYiHandleTokenSuoYao {
	scMsg := &uipb.SCJieYiHandleTokenSuoYao{}
	scMsg.PlayerId = &playerId
	scMsg.Token = &token
	return scMsg
}

func BuildSCJieYiQiuYuan(lastQiuYuanTime int64) *uipb.SCJieYiQiuYuan {
	scMsg := &uipb.SCJieYiQiuYuan{}
	scMsg.LastQiuYuanTime = &lastQiuYuanTime
	return scMsg
}

func BuildSCJieYiQiuYuanNotice(playerName string, mapId int32, pos coretypes.Position) *uipb.SCJieYiQiuYuanNotice {
	scMsg := &uipb.SCJieYiQiuYuanNotice{}
	scMsg.PlayerName = &playerName
	scMsg.MapId = &mapId
	scMsg.Pos = buildJieYiPosition(pos)
	return scMsg
}

func BuildSCJieYiJiuYuan() *uipb.SCJieYiJiuYuan {
	scMsg := &uipb.SCJieYiJiuYuan{}
	return scMsg
}

func BuildSCJieYiDaoJuHelpChangeNotice(daoJu int32) *uipb.SCJieYiDaoJuHelpChangeNotice {
	scMsg := &uipb.SCJieYiDaoJuHelpChangeNotice{}
	scMsg.DaoJuType = &daoJu
	return scMsg
}

func BuildSCJieYiTokenGive(token int32, playerId int64, level int32) *uipb.SCJieYiTokenGive {
	scMsg := &uipb.SCJieYiTokenGive{}
	scMsg.Token = &token
	scMsg.PlayerId = &playerId
	scMsg.Level = &level
	return scMsg
}

func BuildSCJieYiDaoJuHelpChange(daoJu int32, playerId int64) *uipb.SCJieYiDaoJuHelpChange {
	scMsg := &uipb.SCJieYiDaoJuHelpChange{}
	scMsg.DaoJuType = &daoJu
	scMsg.PlayerId = &playerId
	return scMsg
}

func BuildSCJieYiLeaveWordInfo(objList []*jieyi.JieYiLeaveWordObject) *uipb.SCJieYiLeaveWordInfo {
	scMsg := &uipb.SCJieYiLeaveWordInfo{}
	for _, obj := range objList {
		scMsg.LeaveWordInfo = append(scMsg.LeaveWordInfo, buildLeaveWordInfo(obj))
	}
	return scMsg
}

func BuildSCJieYiTokenUpLev(token int32, level int32, randBless int32, pro int32, success bool) *uipb.SCJieYiTokenUpLev {
	scMsg := &uipb.SCJieYiTokenUpLev{}
	scMsg.Typ = &token
	scMsg.Level = &level
	scMsg.RandBless = &randBless
	scMsg.Pro = &pro
	scMsg.Success = &success
	return scMsg
}

func BuildSCJieBrotherInfoOnChange(memberObj *jieyi.JieYiMemberObject) *uipb.SCJieYiBrotherInfoOnChange {
	scMsg := &uipb.SCJieYiBrotherInfoOnChange{}
	playerId := memberObj.GetPlayerId()
	name := memberObj.GetPlayerName()
	role := int32(memberObj.GetRole())
	sex := int32(memberObj.GetSex())
	level := memberObj.GetLevel()
	force := memberObj.GetForce()
	zhuanSheng := memberObj.GetZhuanSheng()
	jieYiTime := memberObj.GetJieYiTime()
	tokenType := int32(memberObj.GetTokenType())
	tokenLev := memberObj.GetTokenLev()
	daoJu := int32(memberObj.GetDaoJuType())
	isOnline := int32(memberObj.GetOnLineState())
	shengWeiZhi := memberObj.GetShengWeiZhi()
	nameLev := memberObj.GetNameLev()
	scMsg.PlayerId = &playerId
	scMsg.Name = &name
	scMsg.Role = &role
	scMsg.Sex = &sex
	scMsg.Level = &level
	scMsg.Force = &force
	scMsg.ZhuanSheng = &zhuanSheng
	scMsg.JieYiTime = &jieYiTime
	scMsg.TokenType = &tokenType
	scMsg.TokenLev = &tokenLev
	scMsg.JieYiDaoJu = &daoJu
	scMsg.IsOnline = &isOnline
	scMsg.ShengWeiZhi = &shengWeiZhi
	scMsg.NameLev = &nameLev
	return scMsg
}

func BuildSCJieYiMemberInfo(jieYiObj *jieyi.JieYiObject, memberObj []*jieyi.JieYiMemberObject, playerId int64, tokenPro int32, shengWeiZhi int32) *uipb.SCJieYiMemberInfo {
	scMsg := &uipb.SCJieYiMemberInfo{}
	scMsg.JieYiInfo = buildJieYiInfo(jieYiObj)
	scMsg.JieYiMember = buildJieYiMemberInfo(memberObj, playerId, tokenPro, shengWeiZhi)
	return scMsg
}

func BuildSCJieYiInviteNotice(playerId int64, playerName string, sex int32, role int32, force int64, name string, memberNum int32, daoJu int32) *uipb.SCJieYiInviteNotice {
	scMsg := &uipb.SCJieYiInviteNotice{}
	scMsg.JieYiName = &name
	scMsg.MemberNum = &memberNum
	scMsg.JieYiDaoJu = &daoJu
	scMsg.PlayerId = &playerId
	scMsg.PlayerName = &playerName
	scMsg.Sex = &sex
	scMsg.Role = &role
	scMsg.Force = &force

	return scMsg
}

func BuildSCJieYiInvite(daoJu int32, name string) *uipb.SCJieYiInvite {
	scMsg := &uipb.SCJieYiInvite{}
	scMsg.JieYiDaoJu = &daoJu
	scMsg.Name = &name

	return scMsg
}

func BuildSCJieYiHandleInvite(inviteId int64, agree bool, sex int32) *uipb.SCJieYiHandleInvite {
	scMsg := &uipb.SCJieYiHandleInvite{}
	scMsg.PlayerId = &inviteId
	scMsg.Agree = &agree
	scMsg.Sex = &sex
	return scMsg
}

func BuildSCJieYiHandleInviteNotice(inviteeId int64, inviteeName string, agree bool, laoDaId int64, sex int32) *uipb.SCJieYiHandleInviteNotice {
	scMsg := &uipb.SCJieYiHandleInviteNotice{}
	scMsg.PlayerId = &inviteeId
	scMsg.PlayerName = &inviteeName
	scMsg.Agree = &agree
	scMsg.LaoDaId = &laoDaId
	scMsg.Sex = &sex
	return scMsg
}

func BuildSCJieYiNameUpLev(success bool, level int32, pro int32, randBless int32) *uipb.SCJieYiNameUpLev {
	scMsg := &uipb.SCJieYiNameUpLev{}
	scMsg.Level = &level
	scMsg.Pro = &pro
	scMsg.Success = &success
	scMsg.RandBless = &randBless
	return scMsg
}

func BuildSCJieYiShengWeiZhiDrop(dropNum int32, dropLev int32, name string, mapId int32, pos coretypes.Position, playerId int64, playerName string) *uipb.SCJieYiShengWeiZhiDrop {
	scMsg := &uipb.SCJieYiShengWeiZhiDrop{}
	scMsg.DropNum = &dropNum
	scMsg.DropLev = &dropLev
	scMsg.KillerName = &name
	scMsg.MapId = &mapId
	scMsg.PlayerId = &playerId
	scMsg.Name = &playerName
	scMsg.Pos = buildJieYiPosition(pos)
	return scMsg
}

func BuildSCJieYiShengWeiZhiTuiSong(shengWeiZhi int32) *uipb.SCJieYiShengWeiZhiTuiSong {
	scMsg := &uipb.SCJieYiShengWeiZhiTuiSong{}
	scMsg.ShengWeiZhi = &shengWeiZhi
	return scMsg
}

func BuildSCJieYiJieChu() *uipb.SCJieYiJieChu {
	scMsg := &uipb.SCJieYiJieChu{}
	return scMsg
}

func BuildSCJieYiJieChuNotice(playerId int64, playerName string) *uipb.SCJieYiJieChuNotice {
	scMsg := &uipb.SCJieYiJieChuNotice{}
	scMsg.PlayerId = &playerId
	scMsg.PlayerName = &playerName
	return scMsg
}

func buildJieYiInfo(jieYiObj *jieyi.JieYiObject) *uipb.JieYiInfo {
	scMsg := &uipb.JieYiInfo{}
	jieyiId := jieYiObj.GetDBId()
	name := jieYiObj.GetName()
	memberNum := jieyi.GetJieYiService().GetJieYiMemberNum(jieyiId)
	scMsg.JieYiId = &jieyiId
	scMsg.Name = &name
	scMsg.PlayerNum = &memberNum
	return scMsg
}

func buildJieYiMemberInfo(memberObjList []*jieyi.JieYiMemberObject, memberId int64, tokenPro int32, shengWeiZhi int32) []*uipb.JieYiMemberInfo {
	var scMsgList []*uipb.JieYiMemberInfo
	for _, memberObj := range memberObjList {
		scMsg := &uipb.JieYiMemberInfo{}
		playerId := memberObj.GetPlayerId()
		name := memberObj.GetPlayerName()
		role := int32(memberObj.GetRole())
		sex := int32(memberObj.GetSex())
		level := memberObj.GetLevel()
		force := memberObj.GetForce()
		zhuanSheng := memberObj.GetZhuanSheng()
		jieYiTime := memberObj.GetJieYiTime()
		tokenType := int32(memberObj.GetTokenType())
		tokenLev := memberObj.GetTokenLev()
		daoJu := int32(memberObj.GetDaoJuType())
		isOnline := int32(memberObj.GetOnLineState())
		nameLev := memberObj.GetNameLev()
		scMsg.PlayerId = &playerId
		scMsg.Name = &name
		scMsg.Role = &role
		scMsg.Sex = &sex
		scMsg.Level = &level
		scMsg.Force = &force
		scMsg.ZhuanSheng = &zhuanSheng
		scMsg.JieYiTime = &jieYiTime
		scMsg.TokenType = &tokenType
		scMsg.TokenLev = &tokenLev
		scMsg.JieYiDaoJu = &daoJu
		scMsg.IsOnline = &isOnline
		scMsg.NameLev = &nameLev

		if playerId == memberId {
			scMsg.TokenPro = &tokenPro
			scMsg.ShengWeiZhi = &shengWeiZhi
		}

		scMsgList = append(scMsgList, scMsg)
	}
	return scMsgList
}

func buildLeaveWordInfo(obj *jieyi.JieYiLeaveWordObject) *uipb.LeaveWordInfo {
	scMsg := &uipb.LeaveWordInfo{}
	playerId := obj.GetPlayerId()
	force := obj.GetForce()
	name := obj.GetPlayerName()
	role := int32(obj.GetRole())
	sex := int32(obj.GetSex())
	level := obj.GetLevel()
	leaveWord := obj.GetLeaveWord()
	lastPostTime := obj.GetLastPostTime()
	isOnline := int32(obj.GetOnLineState())

	scMsg.PlayerId = &playerId
	scMsg.Force = &force
	scMsg.PlayerName = &name
	scMsg.LeaveWord = &leaveWord
	scMsg.Level = &level
	scMsg.LastPostTime = &lastPostTime
	scMsg.Sex = &sex
	scMsg.Role = &role
	scMsg.IsOnline = &isOnline
	return scMsg
}

func buildJieYiPosition(pos coretypes.Position) *uipb.JieYiPosition {
	scMsg := &uipb.JieYiPosition{}
	x := float32(pos.X)
	y := float32(pos.Y)
	z := float32(pos.Z)
	scMsg.PosX = &x
	scMsg.PosY = &y
	scMsg.PosZ = &z
	return scMsg
}
