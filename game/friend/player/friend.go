package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/dao"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/friend"
	friendtemplate "fgame/fgame/game/friend/template"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playercommon "fgame/fgame/game/player/common"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家好友管理器
type PlayerFriendDataManager struct {
	p player.Player
	//好友map
	friendListMap map[int64]*friend.FriendObject
	//黑名单
	blackListMap map[int64]*PlayerFriendBlackObject
	//逆向黑名单
	revBlackListMap map[int64]*PlayerFriendBlackObject
	//添加好友邀请
	friendInviteMap map[int64]*PlayerFriendInviteObject
	//邀请时间
	inviteCdMap map[int64]int64
	//赞赏数据
	feedbackList []*PlayerFriendFeedbackObject
	//添加好友奖励记录
	rewRecordObj *PlayerFriendAddRewObject
	//邀请一键添加cd
	inviteAddAllCd int64
	//赞赏记录数据
	admireRecordList []*PlayerFriendAdmireObject
	//我的表白记录数据
	marryDevelopSendLogList []*PlayerFriendMarryDevelopSendLogObject
	//对我的表白记录数据对象
	marryDevelopRecvLogList []*PlayerFriendMarryDevelopRecvLogObject
	//好友异步日志
	friendLogList []*friend.PlayerFriendLogObject
	//
	hbRunner heartbeat.HeartbeatTaskRunner
}

const (
	admireLimit = 50
)

func (m *PlayerFriendDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFriendDataManager) Load() (err error) {
	m.friendListMap = make(map[int64]*friend.FriendObject)
	m.blackListMap = make(map[int64]*PlayerFriendBlackObject)
	m.revBlackListMap = make(map[int64]*PlayerFriendBlackObject)
	m.friendInviteMap = make(map[int64]*PlayerFriendInviteObject)
	m.inviteCdMap = make(map[int64]int64)

	m.friendListMap = friend.GetFriendService().GetFriendList(m.p.GetId())
	friendBlackEntityList, err := dao.GetFriendDao().GetFriendBlackList(m.p.GetId())
	if err != nil {
		return
	}
	for _, friendBlackEntity := range friendBlackEntityList {
		blackFriend := newPlayerFriendBlackObject(m.p)
		err = blackFriend.FromEntity(friendBlackEntity)
		if err != nil {
			return err
		}
		if blackFriend.Black == 1 {
			m.addBlackFriend(blackFriend)
		}
		if blackFriend.RevBlack == 1 {
			m.addRevBlackFriend(blackFriend)
		}
	}

	friendInviteEntityList, err := dao.GetFriendDao().GetFriendInviteList(m.p.GetId())
	if err != nil {
		return
	}
	for _, friendInviteEntity := range friendInviteEntityList {
		friendInvite := newPlayerFriendInviteObject(m.p)
		friendInvite.FromEntity(friendInviteEntity)
		m.friendInviteMap[friendInvite.InviteId] = friendInvite
	}

	//玩家赞赏
	feedbackEntityList, err := dao.GetFriendDao().GetFriendFeedbackList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range feedbackEntityList {
		feedback := newPlayerFriendFeedbackObject(m.p)
		feedback.FromEntity(entity)
		m.feedbackList = append(m.feedbackList, feedback)
	}

	//玩家领奖记录
	rewRecordEntity, err := dao.GetFriendDao().GetFriendAddRew(m.p.GetId())
	if err != nil {
		return
	}

	if rewRecordEntity != nil {
		obj := newPlayerFriendAddRewObject(m.p)
		obj.FromEntity(rewRecordEntity)
		m.rewRecordObj = obj
	} else {
		m.initFriendAddRewRecord()
	}

	//玩家赞赏记录
	admireEntityList, err := dao.GetFriendDao().GetFriendAdmireList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range admireEntityList {
		obj := newPlayerFriendAdmireObject(m.p)
		obj.FromEntity(entity)
		m.admireRecordList = append(m.admireRecordList, obj)
	}

	//我的表白记录数据
	sendLogEntityList, err := dao.GetFriendDao().GetMarryDevelopSendLogList(m.p.GetId())
	if err != nil {
		return
	}
	for _, sendLogEntity := range sendLogEntityList {
		obj := newPlayerFriendMarryDevelopSendLogObject(m.p)
		obj.FromEntity(sendLogEntity)
		m.marryDevelopSendLogList = append(m.marryDevelopSendLogList, obj)
	}

	//对我的表白记录数据
	recvLogEntityList, err := dao.GetFriendDao().GetMarryDevelopRecvLogList(m.p.GetId())
	if err != nil {
		return
	}
	for _, recvLogEntity := range recvLogEntityList {
		obj := newPlayerFriendMarryDevelopRecvLogObject(m.p)
		obj.FromEntity(recvLogEntity)
		m.marryDevelopRecvLogList = append(m.marryDevelopRecvLogList, obj)
	}

	friendLogEntityList, err := dao.GetFriendDao().GetFriendLogList(m.p.GetId())
	if err != nil {
		return
	}
	for _, friendLogEntity := range friendLogEntityList {
		friendLogObj := friend.NewPlayerFriendLogObject()
		friendLogObj.FromEntity(friendLogEntity)
		m.friendLogList = append(m.friendLogList, friendLogObj)
	}

	return nil
}

//加载后
func (m *PlayerFriendDataManager) AfterLoad() (err error) {
	m.hbRunner.AddTask(CreateAddDummyFriendTask(m.p))

	err = m.refreshTimes()
	if err != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	for _, friendLogObj := range m.friendLogList {
		m.OffonlineChange(friendLogObj.FrinedId, friendLogObj.Type)
		friendLogObj.DeleteTime = now
		friendLogObj.SetModified()
	}
	m.friendLogList = nil
	return nil
}

//心跳
func (m *PlayerFriendDataManager) Heartbeat() {
	m.hbRunner.Heartbeat()
}

func (m *PlayerFriendDataManager) initFriendAddRewRecord() {
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj := newPlayerFriendAddRewObject(m.p)
	obj.id = id
	obj.rewRecord = map[int32]int32{}
	obj.createTime = now
	obj.SetModified()
	m.rewRecordObj = obj
}

func (m *PlayerFriendDataManager) initBlackFriend(friendId int64) (bf *PlayerFriendBlackObject) {
	now := global.GetGame().GetTimeService().Now()
	bf = newPlayerFriendBlackObject(m.p)
	id, _ := idutil.GetId()
	bf.Id = id
	bf.FriendId = friendId
	bf.Black = 1
	bf.RevBlack = 0
	bf.CreateTime = now
	bf.SetModified()
	return bf
}

func (m *PlayerFriendDataManager) initRevBlackFriend(friendId int64) (rbf *PlayerFriendBlackObject) {
	now := global.GetGame().GetTimeService().Now()
	rbf = newPlayerFriendBlackObject(m.p)
	id, _ := idutil.GetId()
	rbf.Id = id
	rbf.FriendId = friendId
	rbf.Black = 0
	rbf.RevBlack = 1
	rbf.CreateTime = now
	rbf.SetModified()
	return rbf
}

func (m *PlayerFriendDataManager) addBlackFriend(bf *PlayerFriendBlackObject) {
	if bf.Black == 0 {
		return
	}
	m.blackListMap[bf.FriendId] = bf
}

func (m *PlayerFriendDataManager) addRevBlackFriend(bf *PlayerFriendBlackObject) {
	if bf.RevBlack == 0 {
		return
	}
	m.revBlackListMap[bf.FriendId] = bf
}

func (m *PlayerFriendDataManager) GetFriends() map[int64]*friend.FriendObject {
	return m.friendListMap
}

func (m *PlayerFriendDataManager) GetBlacks() map[int64]*PlayerFriendBlackObject {
	return m.blackListMap
}

func (m *PlayerFriendDataManager) GetRevBlacks() map[int64]*PlayerFriendBlackObject {
	return m.revBlackListMap
}

//是否应该加好友
func (m *PlayerFriendDataManager) ShouldAddFriend(friendId int64) (flag bool) {
	f := m.GetFriend(friendId)
	if f != nil {
		return
	}
	return true
}

func (m *PlayerFriendDataManager) IsFriend(friendId int64) (flag bool) {
	f := m.GetFriend(friendId)
	if f == nil {
		return false
	}
	return true
}

//好友数量
func (m *PlayerFriendDataManager) NumOfFriend() int {
	return len(m.friendListMap)
}

//黑名单
func (m *PlayerFriendDataManager) NumOfBlack() int {
	return len(m.blackListMap)
}

//添加好友
func (m *PlayerFriendDataManager) AddFrined(fo *friend.FriendObject) {
	friendId := fo.FriendId
	if friendId == m.p.GetId() {
		friendId = fo.PlayerId
	}
	bf := m.getBlack(friendId)
	if bf != nil {
		m.removeFromBlack(friendId)

	}
	m.addFriend(fo, friendId)
}

func (m *PlayerFriendDataManager) ShouldAddBlack(friendId int64) (flag bool) {
	f := m.getBlack(friendId)
	if f == nil {
		return true
	}
	return false
}

func (m *PlayerFriendDataManager) AddBlack(friendId int64) {
	flag := m.ShouldAddBlack(friendId)
	if !flag {
		return
	}
	rbf := m.revBlackListMap[friendId]
	now := global.GetGame().GetTimeService().Now()
	if rbf == nil {
		rbf = m.initBlackFriend(friendId)
	} else {
		rbf.Black = 1
		rbf.UpdateTime = now
		rbf.SetModified()
	}
	m.addBlackFriend(rbf)
	gameevent.Emit(friendeventtypes.EventTypeFriendBlack, m.p, friendId)
}

func (m *PlayerFriendDataManager) ShouldRemoveBlack(friendId int64) (flag bool) {
	f := m.getBlack(friendId)
	if f != nil {
		return true
	}
	return false
}

func (m *PlayerFriendDataManager) RemoveBlack(friendId int64) {
	flag := m.ShouldRemoveBlack(friendId)
	if !flag {
		return
	}
	m.removeFromBlack(friendId)
	gameevent.Emit(friendeventtypes.EventTypeFriendRemoveBlack, m.p, friendId)
}

func (m *PlayerFriendDataManager) addFriend(fo *friend.FriendObject, friendId int64) {
	m.friendListMap[friendId] = fo
}

func (m *PlayerFriendDataManager) GetFriend(friendId int64) (f *friend.FriendObject) {
	f, exist := m.friendListMap[friendId]
	if !exist {
		return nil
	}
	return
}

//对方是否添加过我
func (m *PlayerFriendDataManager) IsAddedFriend(friendId int64) bool {
	f := m.GetFriend(friendId)
	if f != nil {
		return true
	}

	bf := m.getBlack(friendId)
	if bf != nil {
		return true
	}
	return false
}

func (m *PlayerFriendDataManager) getBlack(friendId int64) (f *PlayerFriendBlackObject) {
	f, exist := m.blackListMap[friendId]
	if !exist {
		return nil
	}
	return
}

func (m *PlayerFriendDataManager) getReverseBlackFriend(friendId int64) (f *PlayerFriendBlackObject) {
	f, exist := m.revBlackListMap[friendId]
	if !exist {
		return nil
	}
	return
}

func (m *PlayerFriendDataManager) IsBlack(friendId int64) (flag bool) {
	_, ok := m.blackListMap[friendId]
	if ok {
		flag = true
	}
	return
}

//是否被对方拉黑
func (m *PlayerFriendDataManager) IsBlacked(friendId int64) (flag bool) {
	_, ok := m.revBlackListMap[friendId]
	if ok {
		flag = true
	}
	return
}

//移除从黑名单
func (m *PlayerFriendDataManager) removeFromBlack(friendId int64) {
	now := global.GetGame().GetTimeService().Now()
	bf := m.blackListMap[friendId]
	if bf == nil {
		return
	}
	bf.Black = 0
	bf.UpdateTime = now
	if bf.RevBlack == 0 {
		bf.DeleteTime = now
	}
	bf.SetModified()
	delete(m.blackListMap, friendId)
}

func (m *PlayerFriendDataManager) removeFromFriend(friendId int64) {
	delete(m.friendListMap, friendId)
}

func (m *PlayerFriendDataManager) removeFromReverseBlack(friendId int64) {
	now := global.GetGame().GetTimeService().Now()
	rbf := m.revBlackListMap[friendId]
	if rbf == nil {
		return
	}
	rbf.RevBlack = 0
	rbf.UpdateTime = now
	if rbf.Black == 0 {
		rbf.DeleteTime = now
	}
	rbf.SetModified()
	delete(m.revBlackListMap, friendId)
}

func (m *PlayerFriendDataManager) DeleteFriend(friendId int64) {
	m.removeFromFriend(friendId)
}

func (m *PlayerFriendDataManager) ReverseBlackFriend(friendId int64) {
	rbf := m.getReverseBlackFriend(friendId)
	if rbf != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	bf := m.getBlack(friendId)
	if bf == nil {
		bf = m.initRevBlackFriend(friendId)
	} else {
		bf.RevBlack = 1
		bf.UpdateTime = now
		bf.SetModified()
	}
	m.addRevBlackFriend(bf)
}

func (m *PlayerFriendDataManager) ReverseRemoveBlackFriend(friendId int64) {
	rbf := m.getReverseBlackFriend(friendId)
	if rbf == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	rbf.RevBlack = 0
	rbf.UpdateTime = now
	rbf.SetModified()
	delete(m.revBlackListMap, friendId)
}

func (m *PlayerFriendDataManager) OffonlineChange(friendId int64, logType friendtypes.FriendLogType) {
	now := global.GetGame().GetTimeService().Now()
	switch logType {
	case friendtypes.FriendLogTypeBlack:
		{
			rbf := m.getReverseBlackFriend(friendId)
			if rbf != nil {
				return
			}
			bf := m.getBlack(friendId)
			if bf == nil {
				bf = m.initRevBlackFriend(friendId)
			} else {
				bf.RevBlack = 1
				bf.UpdateTime = now
				bf.SetModified()
			}
			m.addRevBlackFriend(bf)
			break
		}
	case friendtypes.FriendLogTypeRemoveBlack:
		{
			rbf := m.getReverseBlackFriend(friendId)
			if rbf == nil {
				return
			}
			m.removeFromReverseBlack(friendId)
			break
		}
	}
}

func (m *PlayerFriendDataManager) initFriendInvite(inviteId int64, info *playercommon.PlayerInfo) (fi *PlayerFriendInviteObject) {
	now := global.GetGame().GetTimeService().Now()
	fi = newPlayerFriendInviteObject(m.p)

	id, _ := idutil.GetId()
	fi.Id = id
	fi.InviteId = inviteId
	fi.Name = info.Name
	fi.Role = int32(info.Role)
	fi.Sex = int32(info.Sex)
	fi.Force = info.Force
	fi.Level = info.Level
	fi.CreateTime = now
	fi.UpdateTime = now
	fi.SetModified()
	return
}

func (m *PlayerFriendDataManager) FriendInvite(inviteId int64) {
	now := global.GetGame().GetTimeService().Now()
	playerInfo, err := player.GetPlayerService().GetPlayerInfo(inviteId)
	if err != nil {
		return
	}
	friendInvite, exist := m.friendInviteMap[inviteId]
	if !exist {
		friendInvite = m.initFriendInvite(inviteId, playerInfo)
		m.friendInviteMap[inviteId] = friendInvite
	} else {
		friendInvite.Level = playerInfo.Level
		friendInvite.Force = playerInfo.Force
		friendInvite.UpdateTime = now
		friendInvite.SetModified()
	}
}

func (m *PlayerFriendDataManager) RemoveFriendInvite(friendId int64) {
	fi := m.friendInviteMap[friendId]
	if fi == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	fi.DeleteTime = now
	fi.SetModified()
	delete(m.friendInviteMap, friendId)
}

func (m *PlayerFriendDataManager) HasedInvite(friendId int64) (flag bool) {
	_, flag = m.friendInviteMap[friendId]
	return
}

//是否邀请过于频繁
func (m *PlayerFriendDataManager) InviteFrequent(friendId int64) bool {
	inviteTime, exist := m.inviteCdMap[friendId]
	if !exist {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	cdTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendInviteCdTime) * 1000)
	if now-inviteTime < cdTime {
		return true
	}
	return false
}

func (m *PlayerFriendDataManager) InviteTime(friendId int64) int64 {
	now := global.GetGame().GetTimeService().Now()
	m.inviteCdMap[friendId] = now
	return now
}

func (m *PlayerFriendDataManager) GetFriendInviteMap() map[int64]*PlayerFriendInviteObject {
	return m.friendInviteMap
}

// 添加好友赞赏
func (m *PlayerFriendDataManager) AddFriendFeedback(frId int64, frName string, condition int32, nType friendtypes.FriendNoticeType, fType friendtypes.FriendFeedbackType) {
	obj := newPlayerFriendFeedbackObject(m.p)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.friendId = frId
	obj.friendName = frName
	obj.noticeType = nType
	obj.feedbackType = fType
	obj.condition = condition
	obj.createTime = now
	obj.SetModified()

	m.feedbackList = append(m.feedbackList, obj)
}

// 获取赞赏列表
func (m *PlayerFriendDataManager) GetFriendFeedbackList() []*PlayerFriendFeedbackObject {
	return m.feedbackList
}

//阅读赞赏
func (m *PlayerFriendDataManager) ReadFeedback() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.feedbackList {
		obj.updateTime = now
		obj.deleteTime = now
		obj.SetModified()
	}

	m.feedbackList = []*PlayerFriendFeedbackObject{}
}

// 领取添加好友奖励
func (m *PlayerFriendDataManager) AddRewRecord(frNum int32) {
	_, ok := m.rewRecordObj.rewRecord[frNum]
	if ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.rewRecordObj.rewRecord[frNum] = 1
	m.rewRecordObj.updateTime = now
	m.rewRecordObj.SetModified()
}

// 是否领取添加好友奖励
func (m *PlayerFriendDataManager) IsCanReceiveRew(frNum int32) bool {
	_, ok := m.rewRecordObj.rewRecord[frNum]
	if !ok {
		return true
	}
	return false
}

// 获取领奖记录
func (m *PlayerFriendDataManager) GetReceiveRewRecord() (record []int32) {
	for frNum, _ := range m.rewRecordObj.rewRecord {
		record = append(record, frNum)
	}

	return
}

// 获取虚拟好友数量
func (m *PlayerFriendDataManager) GetDummyFriendNum() int32 {
	return m.rewRecordObj.frDummyNum
}

// 添加虚拟好友
func (m *PlayerFriendDataManager) addDummyFriend() {
	now := global.GetGame().GetTimeService().Now()
	noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
	interval := int64(noticeConstantTemp.TianjiaJiarenTime)
	diff := now - m.rewRecordObj.lastAddDummyTime
	if diff < interval {
		return
	}

	curFriNum := int32(len(m.friendListMap))
	maxRewNum := friendtemplate.GetFriendNoticeTemplateService().GetFriendRewMaxAddNum()
	curTotalNum := curFriNum + m.rewRecordObj.frDummyNum
	if curTotalNum >= maxRewNum {
		return
	}

	m.rewRecordObj.frDummyNum += 1
	m.rewRecordObj.lastAddDummyTime = now
	m.rewRecordObj.updateTime = now
	m.rewRecordObj.SetModified()

	gameevent.Emit(friendeventtypes.EventTypeFriendDummyNumChanged, m.p, m.rewRecordObj.frDummyNum)
	return
}

func (m *PlayerFriendDataManager) GetExcludeForAddAll() (excludeIdMap map[int64]struct{}) {
	excludeIdMap = make(map[int64]struct{})
	for _, fri := range m.GetFriends() {
		friendId := fri.PlayerId
		if friendId == m.p.GetId() {
			friendId = fri.FriendId
		}
		excludeIdMap[friendId] = struct{}{}
	}
	for playerId, _ := range m.GetBlacks() {
		excludeIdMap[playerId] = struct{}{}
	}
	for playerId, _ := range m.GetRevBlacks() {
		excludeIdMap[playerId] = struct{}{}
	}
	excludeIdMap[m.p.GetId()] = struct{}{}
	return
}

func (m *PlayerFriendDataManager) IfCanAddAll() bool {
	now := global.GetGame().GetTimeService().Now()
	cdTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAddFriendInviteBatchCD))
	return (now - m.inviteAddAllCd) >= cdTime
}

func (m *PlayerFriendDataManager) AddAllLeftTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	cdTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAddFriendInviteBatchCD))
	elapse := now - m.inviteAddAllCd
	return cdTime - elapse
}

func (m *PlayerFriendDataManager) InviteAddAllCdTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	m.inviteAddAllCd = now
	return m.inviteAddAllCd
}

func (m *PlayerFriendDataManager) AddCongratulateTimes() {
	m.refreshTimes()
	now := global.GetGame().GetTimeService().Now()
	m.rewRecordObj.congratulateTimes += 1
	m.rewRecordObj.lastCongratulateTime = now
	m.rewRecordObj.updateTime = now
	m.rewRecordObj.SetModified()
}

func (m *PlayerFriendDataManager) IsLimitCongraluteRew() bool {
	m.refreshTimes()
	noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
	if m.rewRecordObj.congratulateTimes >= noticeConstantTemp.ShouliLimitCount {
		return true
	}
	return false
}

func (m *PlayerFriendDataManager) refreshTimes() (err error) {
	now := global.GetGame().GetTimeService().Now()
	lastTime := m.rewRecordObj.lastCongratulateTime
	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return err
	}
	if !flag {
		m.rewRecordObj.congratulateTimes = 0
		m.rewRecordObj.lastCongratulateTime = now
		m.rewRecordObj.SetModified()
	}
	return
}

func (m *PlayerFriendDataManager) AddAdmireTimes(friId int64) {
	m.refreshAdmireTimes()

	now := global.GetGame().GetTimeService().Now()
	obj := m.getAdmireObj(friId)
	if obj == nil {
		id, _ := idutil.GetId()
		now := global.GetGame().GetTimeService().Now()
		obj = newPlayerFriendAdmireObject(m.p)
		obj.id = id
		obj.friId = friId
		obj.admireTimes = 0
		obj.createTime = now
		obj.SetModified()
		m.admireRecordList = append(m.admireRecordList, obj)
	}

	obj.admireTimes += 1
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerFriendDataManager) IsLimitAdmire(friId int64) bool {
	m.refreshAdmireTimes()
	obj := m.getAdmireObj(friId)
	if obj == nil {
		return false
	}

	if obj.admireTimes >= admireLimit {
		return true
	}

	return false
}

func (m *PlayerFriendDataManager) refreshAdmireTimes() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.admireRecordList {
		lastTime := obj.updateTime
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			obj.admireTimes = 0
			obj.updateTime = now
			obj.SetModified()
		}
	}
	return
}

func (m *PlayerFriendDataManager) getAdmireObj(friId int64) *PlayerFriendAdmireObject {
	for _, obj := range m.admireRecordList {
		if obj.friId != friId {
			continue
		}

		return obj

	}
	return nil
}

func (m *PlayerFriendDataManager) GetMarryDevelopSendLogByTime(time int64) []*PlayerFriendMarryDevelopSendLogObject {
	for index, log := range m.marryDevelopSendLogList {
		if time < log.UpdateTime {
			return m.marryDevelopSendLogList[index:]
		}
	}

	return nil
}

func (m *PlayerFriendDataManager) AddMarryDevelopSendLog(logData *friendtypes.MarryDevelopLogData) {
	obj := m.createMarryDevelopSendLogObj(logData)
	m.marryDevelopSendLogList = append(m.marryDevelopSendLogList, obj)
}

func (m *PlayerFriendDataManager) createMarryDevelopSendLogObj(logData *friendtypes.MarryDevelopLogData) *PlayerFriendMarryDevelopSendLogObject {
	now := global.GetGame().GetTimeService().Now()
	maxLogLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMarryDevelopLogMaxNum)
	var obj *PlayerFriendMarryDevelopSendLogObject
	if len(m.marryDevelopSendLogList) >= int(maxLogLen) {
		obj = m.marryDevelopSendLogList[0]
		m.marryDevelopSendLogList = m.marryDevelopSendLogList[1:]
	} else {
		obj = newPlayerFriendMarryDevelopSendLogObject(m.p)
		id, _ := idutil.GetId()
		obj.Id = id
		obj.CreateTime = now
	}

	obj.RecvId = logData.RecvId
	obj.RecvName = logData.RecvName
	obj.ItemId = logData.ItemId
	obj.ItemNum = logData.ItemNum
	obj.CharmNum = logData.CharmNum
	obj.DevelopExp = logData.DevelopExp
	obj.ContextStr = logData.ContextStr
	obj.UpdateTime = now
	obj.SetModified()

	return obj
}

func (m *PlayerFriendDataManager) GetMarryDevelopRecvLogByTime(time int64) []*PlayerFriendMarryDevelopRecvLogObject {
	for index, log := range m.marryDevelopRecvLogList {
		if time < log.UpdateTime {
			return m.marryDevelopRecvLogList[index:]
		}
	}

	return nil
}

func (m *PlayerFriendDataManager) AddMarryDevelopRecvLog(logData *friendtypes.MarryDevelopLogData) {
	obj := m.createMarryDevelopRecvLogObj(logData)
	m.marryDevelopRecvLogList = append(m.marryDevelopRecvLogList, obj)
}

func (m *PlayerFriendDataManager) createMarryDevelopRecvLogObj(logData *friendtypes.MarryDevelopLogData) *PlayerFriendMarryDevelopRecvLogObject {
	now := global.GetGame().GetTimeService().Now()
	maxLogLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMarryDevelopLogMaxNum)
	var obj *PlayerFriendMarryDevelopRecvLogObject
	if len(m.marryDevelopRecvLogList) >= int(maxLogLen) {
		obj = m.marryDevelopRecvLogList[0]
		m.marryDevelopRecvLogList = m.marryDevelopRecvLogList[1:]
	} else {
		obj = newPlayerFriendMarryDevelopRecvLogObject(m.p)
		id, _ := idutil.GetId()
		obj.Id = id
		obj.CreateTime = now
	}

	obj.SendId = logData.SendId
	obj.SendName = logData.SendName
	obj.ItemId = logData.ItemId
	obj.ItemNum = logData.ItemNum
	obj.CharmNum = logData.CharmNum
	obj.DevelopExp = logData.DevelopExp
	obj.ContextStr = logData.ContextStr
	obj.UpdateTime = now
	obj.SetModified()

	return obj
}

func CreatePlayerFriendManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFriendDataManager{}
	m.p = p
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFriendDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFriendManager))
}
