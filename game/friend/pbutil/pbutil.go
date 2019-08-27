package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	friend "fgame/fgame/game/friend/friend"
	playerfriend "fgame/fgame/game/friend/player"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/player"
	playercommon "fgame/fgame/game/player/common"
	playerpbutil "fgame/fgame/game/player/pbutil"
	playertypes "fgame/fgame/game/player/types"
)

func BuildSCFriendsGet(friendList []*uipb.FriendInfo, blackList []*uipb.FriendInfo, feedbackList []*playerfriend.PlayerFriendFeedbackObject, record []int32, dummyNum int32) *uipb.SCFriendsGet {
	friendsGet := &uipb.SCFriendsGet{}
	friendsGet.FriendList = friendList
	friendsGet.BlackList = blackList
	for _, feedback := range feedbackList {
		friendsGet.FeedbackList = append(friendsGet.FeedbackList, buildFriendFeedback(feedback))
	}
	friendsGet.RewRecord = record
	friendsGet.DummyFrNum = &dummyNum
	return friendsGet
}

func BuildFriend(friendId int64, point int32, info *playercommon.PlayerInfo, isBlacked bool) *uipb.FriendInfo {
	friendInfo := &uipb.FriendInfo{}
	friendInfo.FriendId = &friendId
	friendInfo.Point = &point
	friendInfo.PlayerBasicInfo = playerpbutil.BuildPlayerBasicInfo(info, isBlacked)
	return friendInfo
}

func BuildSCFriendAdd(friendId int64, point int32, info *playercommon.PlayerInfo, agree bool) *uipb.SCFriendAdd {
	scFriendAdd := &uipb.SCFriendAdd{}
	scFriendAdd.Agree = &agree
	if info != nil {
		scFriendAdd.Friend = BuildFriend(friendId, point, info, false)
	}
	return scFriendAdd
}

func BuildSCFriendBlack(friendId int64) *uipb.SCFriendBlack {
	scFriendBlack := &uipb.SCFriendBlack{}
	scFriendBlack.FriendId = &friendId
	return scFriendBlack
}

func BuildSCFriendRemoveBlack(friendId int64, isFriend bool) *uipb.SCFriendRemoveBlack {
	scFriendRemoveBlack := &uipb.SCFriendRemoveBlack{}
	scFriendRemoveBlack.FriendId = &friendId
	scFriendRemoveBlack.IsFriend = &isFriend
	return scFriendRemoveBlack
}

func BuildSCFriendInvite(friendId int64, inviteTime int64) *uipb.SCFriendInvite {
	scFriendInvite := &uipb.SCFriendInvite{}
	scFriendInvite.FriendId = &friendId
	scFriendInvite.InviteTime = &inviteTime
	return scFriendInvite
}

func BuildSCFriendInvitePushPeer(info *playercommon.PlayerInfo, inviteTime int64) *uipb.SCFriendInvitePushPeer {
	scFriendInvitePushPeer := &uipb.SCFriendInvitePushPeer{}
	scFriendInvitePushPeer.PlayerBasicInfo = playerpbutil.BuildPlayerBasicInfo(info, false)
	scFriendInvitePushPeer.InviteTime = &inviteTime
	return scFriendInvitePushPeer
}

func BuildSCFriendInviteRefusePushPeer(friendId int64, name string) *uipb.SCFriendInviteRefusePushPeer {
	scFriendInviteRefusePushPeer := &uipb.SCFriendInviteRefusePushPeer{}
	scFriendInviteRefusePushPeer.FriendId = &friendId
	scFriendInviteRefusePushPeer.Name = &name
	return scFriendInviteRefusePushPeer
}

func BuildSCFriendDelete(friendId int64) *uipb.SCFriendDelete {
	scFriendDelete := &uipb.SCFriendDelete{}
	scFriendDelete.FriendId = &friendId
	return scFriendDelete
}

func BuildSCFriendSearch(info *playercommon.PlayerInfo) *uipb.SCFriendSearch {
	scFriendSearch := &uipb.SCFriendSearch{}
	if info != nil {
		scFriendSearch.PlayerList = append(scFriendSearch.PlayerList, playerpbutil.BuildPlayerBasicInfo(info, false))
	}
	return scFriendSearch
}
func BuildSCRecommentFriendsGet(playerList []player.Player) *uipb.SCRecommentFriendsGet {
	scRecommentFriendsGet := &uipb.SCRecommentFriendsGet{}
	for _, p := range playerList {
		scRecommentFriendsGet.PlayerList = append(scRecommentFriendsGet.PlayerList, playerpbutil.BuildPlayerBasicInfoByPlayer(p, false))
	}
	return scRecommentFriendsGet
}

func BuildSCFriendGift(friendId int64, itemId int32, num int32, msgId int32, auto int32) *uipb.SCFriendGift {
	scFriendGift := &uipb.SCFriendGift{}
	scFriendGift.FriendId = &friendId
	scFriendGift.ItemId = &itemId
	scFriendGift.Num = &num
	scFriendGift.MsgId = &msgId
	scFriendGift.Auto = &auto
	return scFriendGift
}

func BuildSCFriendGiftRecv(sendId int64, itemId int32, num int32, msgId int32, msgContent string) *uipb.SCFriendGiftRecv {
	scFriendGiftRecv := &uipb.SCFriendGiftRecv{}
	scFriendGiftRecv.SendId = &sendId
	scFriendGiftRecv.ItemId = &itemId
	scFriendGiftRecv.Num = &num
	scFriendGiftRecv.MsgId = &msgId
	scFriendGiftRecv.MsgContent = &msgContent
	return scFriendGiftRecv
}

func BuildSCFriendGiftFeedback(friendId int64) *uipb.SCFriendGiftFeedback {
	scFriendGiftFeedback := &uipb.SCFriendGiftFeedback{}
	scFriendGiftFeedback.FriendId = &friendId
	return scFriendGiftFeedback
}

func BuildSCFriendGiftFeedbackRecv(sendId int64) *uipb.SCFriendGiftFeedbackRecv {
	scFriendGiftFeedbackRecv := &uipb.SCFriendGiftFeedbackRecv{}
	scFriendGiftFeedbackRecv.SendId = &sendId

	return scFriendGiftFeedbackRecv
}

func BuildSCFriendPointChange(friendId int64, point int32) *uipb.SCFriendPointChange {
	friendPointChange := &uipb.SCFriendPointChange{}
	friendPointChange.FriendId = &friendId
	friendPointChange.Point = &point
	return friendPointChange
}

func BuildSCFrinedBatch(agree bool) *uipb.SCFriendBatch {
	scFriendBatch := &uipb.SCFriendBatch{}
	scFriendBatch.Agree = &agree
	return scFriendBatch
}

func BuildSCFriendGiveFlowerLight(itemId int32, num int32, playerIdLists []int64, playerNameLists []string) *uipb.SCFriendGiveFlowerLight {
	scFriendGiveFlowerLight := &uipb.SCFriendGiveFlowerLight{}
	scFriendGiveFlowerLight.ItemId = &itemId
	scFriendGiveFlowerLight.Num = &num
	scFriendGiveFlowerLight.PlayerIdList = append(scFriendGiveFlowerLight.PlayerIdList, playerIdLists...)
	scFriendGiveFlowerLight.PlayerNameList = append(scFriendGiveFlowerLight.PlayerNameList, playerNameLists...)
	return scFriendGiveFlowerLight
}

func BuildSCFriendInviteList(inviteMap map[int64]*playerfriend.PlayerFriendInviteObject) *uipb.SCFriendInviteList {
	scFriendInviteList := &uipb.SCFriendInviteList{}
	for _, friendInvite := range inviteMap {
		scFriendInviteList.InviteList = append(scFriendInviteList.InviteList, buildInvite(friendInvite))
	}

	return scFriendInviteList
}

func BuildSCFriendNoticeBroadcase(noticeType friendtypes.FriendNoticeType, frId int64, frName string, frRole playertypes.RoleType, frSex playertypes.SexType, condition int32, args string) *uipb.SCFriendNoticeBroadcase {
	scMsg := &uipb.SCFriendNoticeBroadcase{}
	typ := int32(noticeType)
	role := int32(frRole)
	sex := int32(frSex)
	scMsg.FriendId = &frId
	scMsg.NoticeType = &typ
	scMsg.FriendName = &frName
	scMsg.FriendRole = &role
	scMsg.FriendSex = &sex
	scMsg.Args = &args
	scMsg.Condition = &condition
	return scMsg
}

func BuildSCFriendNoticeFeedback(feedbackType friendtypes.FriendFeedbackType, condition int32) *uipb.SCFriendNoticeFeedback {
	scMsg := &uipb.SCFriendNoticeFeedback{}
	fType := int32(feedbackType)
	scMsg.FeedbackType = &fType
	scMsg.Condition = &condition
	return scMsg
}

func BuildSCFriendFeedbackAskFor() *uipb.SCFriendFeedbackAskFor {
	scMsg := &uipb.SCFriendFeedbackAskFor{}
	return scMsg
}

func BuildSCFriendFeedbackAskForReply(isAgree bool) *uipb.SCFriendFeedbackAskForReply {
	scMsg := &uipb.SCFriendFeedbackAskForReply{}
	scMsg.IsAgree = &isAgree
	return scMsg
}

func BuildSCFriendFeedbackAskForReplyNotice(isAgree bool, friendName string) *uipb.SCFriendFeedbackAskForReplyNotice {
	scMsg := &uipb.SCFriendFeedbackAskForReplyNotice{}
	scMsg.IsAgree = &isAgree
	scMsg.FriendName = &friendName
	return scMsg
}

func BuildSCFriendFeedbackAskForNotice(friendId int64, frName, allianceName string, frRole, frSex, itemId int32) *uipb.SCFriendFeedbackAskForNotice {
	scMsg := &uipb.SCFriendFeedbackAskForNotice{}
	scMsg.FrName = &frName
	scMsg.FriendId = &friendId
	scMsg.AllianceName = &allianceName
	scMsg.ItemId = &itemId
	scMsg.Sex = &frSex
	scMsg.Role = &frRole
	return scMsg
}

func BuildSCFriendNoticeFeedbackNotice(frId int64, frName string, noticeType friendtypes.FriendNoticeType, feedbackType friendtypes.FriendFeedbackType, condition int32) *uipb.SCFriendNoticeFeedbackNotice {
	scMsg := &uipb.SCFriendNoticeFeedbackNotice{}
	nType := int32(noticeType)
	fType := int32(feedbackType)

	scMsg.FriendId = &frId
	scMsg.FriendName = &frName
	scMsg.FeedbackType = &fType
	scMsg.NoticeType = &nType
	scMsg.Condition = &condition
	return scMsg
}

func BuildSCFriendNoticeFeedbackRead(feedbackList []*playerfriend.PlayerFriendFeedbackObject) *uipb.SCFriendNoticeFeedbackRead {
	scMsg := &uipb.SCFriendNoticeFeedbackRead{}
	for _, obj := range feedbackList {
		scMsg.FeedbackList = append(scMsg.FeedbackList, buildFriendFeedback(obj))
	}
	return scMsg
}

func BuildSCFriendAddRew(frNum int32) *uipb.SCFriendAddRew {
	scMsg := &uipb.SCFriendAddRew{}
	scMsg.FrNum = &frNum
	return scMsg
}

func BuildSCFriendDummyFriendNumChanged(frNum int32) *uipb.SCFriendDummyFriendNumChanged {
	scMsg := &uipb.SCFriendDummyFriendNumChanged{}
	scMsg.DummyFrNum = &frNum
	return scMsg
}

func BuildSCFriendAddAll(cdTime int64) *uipb.SCFriendAddAll {
	friendAddAll := &uipb.SCFriendAddAll{}
	friendAddAll.InviteAllTime = &cdTime
	return friendAddAll
}

func buildInvite(friendInvite *playerfriend.PlayerFriendInviteObject) *uipb.FriendInviteInfo {
	friendInviteInfo := &uipb.FriendInviteInfo{}
	inviteTime := friendInvite.UpdateTime
	inviteId := friendInvite.InviteId
	name := friendInvite.Name
	role := friendInvite.Role
	sex := friendInvite.Sex
	level := friendInvite.Level
	force := friendInvite.Force

	friendInviteInfo.InviteTime = &inviteTime
	friendInviteInfo.InviteId = &inviteId
	friendInviteInfo.Name = &name
	friendInviteInfo.Role = &role
	friendInviteInfo.Sex = &sex
	friendInviteInfo.Level = &level
	friendInviteInfo.Force = &force
	return friendInviteInfo
}

func buildFriendFeedback(obj *playerfriend.PlayerFriendFeedbackObject) *uipb.FriendFeedback {
	info := &uipb.FriendFeedback{}
	frId := obj.GetFriendId()
	frName := obj.GetFriendName()
	noticeType := int32(obj.GetNoticeType())
	feedbackType := int32(obj.GetFeedbackType())
	condition := obj.GetCondition()
	info.FriendId = &frId
	info.FriendName = &frName
	info.NoticeType = &noticeType
	info.FeedbackType = &feedbackType
	info.Condition = &condition

	return info
}

func buildMarryDevelopLog(log *friend.FriendMarryDevelopLogObject) *uipb.FriendMarryDevelopLog {
	logInfo := &uipb.FriendMarryDevelopLog{}
	sendName := log.SendName
	recvName := log.RecvName
	itemId := log.ItemId
	itemNum := log.ItemNum
	charmNum := log.CharmNum
	developExp := log.DevelopExp
	contextStr := log.ContextStr
	updateTime := log.UpdateTime

	logInfo.CreateTime = &updateTime
	logInfo.SendName = &sendName
	logInfo.RecvName = &recvName
	logInfo.ItemId = &itemId
	logInfo.ItemNum = &itemNum
	logInfo.CharmNum = &charmNum
	logInfo.DevelopExp = &developExp
	logInfo.ContextStr = &contextStr

	return logInfo
}

func buildMarryDevelopSendLog(log *playerfriend.PlayerFriendMarryDevelopSendLogObject) *uipb.FriendMarryDevelopLog {
	logInfo := &uipb.FriendMarryDevelopLog{}
	sendName := log.GetPlayerName()
	recvName := log.RecvName
	itemId := log.ItemId
	itemNum := log.ItemNum
	charmNum := log.CharmNum
	developExp := log.DevelopExp
	contextStr := log.ContextStr
	updateTime := log.UpdateTime

	logInfo.CreateTime = &updateTime
	logInfo.SendName = &sendName
	logInfo.RecvName = &recvName
	logInfo.ItemId = &itemId
	logInfo.ItemNum = &itemNum
	logInfo.CharmNum = &charmNum
	logInfo.DevelopExp = &developExp
	logInfo.ContextStr = &contextStr

	return logInfo
}

func buildMarryDevelopRecvLog(log *playerfriend.PlayerFriendMarryDevelopRecvLogObject) *uipb.FriendMarryDevelopLog {
	logInfo := &uipb.FriendMarryDevelopLog{}
	sendName := log.SendName
	recvName := log.GetPlayerName()
	itemId := log.ItemId
	itemNum := log.ItemNum
	charmNum := log.CharmNum
	developExp := log.DevelopExp
	contextStr := log.ContextStr
	updateTime := log.UpdateTime

	logInfo.CreateTime = &updateTime
	logInfo.SendName = &sendName
	logInfo.RecvName = &recvName
	logInfo.ItemId = &itemId
	logInfo.ItemNum = &itemNum
	logInfo.CharmNum = &charmNum
	logInfo.DevelopExp = &developExp
	logInfo.ContextStr = &contextStr

	return logInfo
}

func BuildSCFriendMarryDevelopLogIncr(logType friendtypes.MarryDevelopLogType, logList []*friend.FriendMarryDevelopLogObject) *uipb.SCFriendMarryDevelopLogIncr {
	scLogIncr := &uipb.SCFriendMarryDevelopLogIncr{}
	for _, log := range logList {
		scLogIncr.LogList = append(scLogIncr.LogList, buildMarryDevelopLog(log))
	}
	logTypeInt := int32(logType)
	scLogIncr.LogType = &logTypeInt
	return scLogIncr
}

func BuildSCFriendMarryDevelopSendLogIncr(logType friendtypes.MarryDevelopLogType, logSendList []*playerfriend.PlayerFriendMarryDevelopSendLogObject) *uipb.SCFriendMarryDevelopLogIncr {
	scLogIncr := &uipb.SCFriendMarryDevelopLogIncr{}
	for _, log := range logSendList {
		scLogIncr.LogList = append(scLogIncr.LogList, buildMarryDevelopSendLog(log))
	}
	logTypeInt := int32(logType)
	scLogIncr.LogType = &logTypeInt
	return scLogIncr
}

func BuildSCFriendMarryDevelopRecvLogIncr(logType friendtypes.MarryDevelopLogType, logRecvList []*playerfriend.PlayerFriendMarryDevelopRecvLogObject) *uipb.SCFriendMarryDevelopLogIncr {
	scLogIncr := &uipb.SCFriendMarryDevelopLogIncr{}
	for _, log := range logRecvList {
		scLogIncr.LogList = append(scLogIncr.LogList, buildMarryDevelopRecvLog(log))
	}
	logTypeInt := int32(logType)
	scLogIncr.LogType = &logTypeInt
	return scLogIncr
}
