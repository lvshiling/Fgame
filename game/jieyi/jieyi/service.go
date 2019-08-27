package jieyi

import (
	"fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/jieyi/dao"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sort"
	"sync"
)

// 结义信息服务
type JieYiService interface {
	// 添加邀请数据
	AddJieYiInviteData(inviteId int64, inviteeId int64, name string, daoJu jieyitypes.JieYiDaoJuType, inviteDaoJu jieyitypes.JieYiDaoJuType, inviteToken jieyitypes.JieYiTokenType, inviteTokenLev, nameLev int32) bool
	// 结义成功
	JieYiSuccess(inviteId int64, inviteeId int64, inviteeDaoJu jieyitypes.JieYiDaoJuType, inviteeToken jieyitypes.JieYiTokenType, inviteeTokenLev, inviteeNameLev int32) (jieYi *JieYi, mem *JieYiMemberObject, flag bool)
	// 玩家解除结义
	JieChuJieYiSuccess(playerId int64) (jieYi *JieYi, tiMemList []*JieYiMemberObject, flag bool)
	// 踢人
	JieYiTiRen(playerId int64, receiverId int64) (jieYi *JieYi, tiMemList []*JieYiMemberObject, flag bool)
	// 邀请失败
	JieYiInviteFail(inviteId int64, inviteeId int64) bool
	// 信物改变成功
	TokenChangeSucess(playerId int64, token jieyitypes.JieYiTokenType) bool
	TokenChangeLevel(playerId int64, level int32) bool
	// 道具改变成功
	DaoJuChangeSucess(playerId int64, daoJu jieyitypes.JieYiDaoJuType) bool

	// 玩家登入
	PlayerLogin(playerId int64)
	// 玩家登出
	PlayerLogout(playerId int64)
	// 玩家名字发生改变
	UpdatePlayerName(playerId int64, name string)
	// 玩家战斗力发生改变
	UpdatePlayerForce(playerId int64, force int64)
	// 玩家等级发生改变
	UpdatePlayerLevel(playerId int64, level int32)
	// 玩家转生发生改变
	UpdatePlayerZhuanSheng(playerId int64, zhuanSheng int32)
	// 声威值改变
	UpdateShengWeiZhi(playerId int64, lev int32, num int32)
	// 信物等级改变
	UpdateTokenLevel(playerId int64, level int32)

	// 修改结义威名
	SetJieYiName(jieYiId int64, name string) bool
	// 添加结义留言数据
	AddJieYiLeaveWord(playerId int64, leaveWord string) bool

	// 获取成员排名
	GetJieYiMemberRank(playerId int64) int32
	// 获取结义留言墙信息
	GetJieYiLeaveWord() []*JieYiLeaveWordObject
	// 获取结义信息
	GetJieYi(jieYiId int64) *JieYi
	GetJieYiInfo(jieYiId int64) *JieYiObject
	// 获取结义成员信息
	GetJieYiMemberInfo(playerId int64) *JieYiMemberObject
	// 获取所有结义成员信息
	GetJieYiMemberList(jieYiId int64) []*JieYiMemberObject
	// 获取成员人数
	GetJieYiMemberNum(jieYiId int64) int32
	// 获取邀请数据
	GetInviteData(inviteId int64) *JieYiInviteObject

	// 玩家是否已经结义
	IsAlreadyJieYi(playerId int64) bool
	// 玩家是否为结义创始人
	IsJieYiLaoDa(playerId int64) bool
	// 结义是否满人
	IsFullMember(playerId int64) bool
	// 威名是否重复
	IsNameRepetitive(name string) bool
	Heartbeat()
}

type jieYiService struct {
	rwm sync.RWMutex
	// 结义信息
	jieYiInfoMap map[int64]*JieYi
	// 结义成员信息
	jieYiMemberMap map[int64]*JieYiMemberObject
	// 结义留言数据
	jieYiLeaveWordMap map[int64]*JieYiLeaveWordObject
	// 邀请数据
	jieYiInviteMap map[int64]*JieYiInviteObject
}

func (s *jieYiService) init() (err error) {
	serverId := global.GetGame().GetServerIndex()
	jieYiEntityList, err := dao.GetJieYiDao().GetJieYiListEntity(serverId)
	if err != nil {
		return
	}

	s.jieYiInfoMap = make(map[int64]*JieYi)
	s.jieYiMemberMap = make(map[int64]*JieYiMemberObject)

	//创建结义数据
	for _, jieYiEntity := range jieYiEntityList {
		jieYiObj := NewJieYiObject()
		err = jieYiObj.FromEntity(jieYiEntity)
		if err != nil {
			return
		}
		s.addJieYi(jieYiObj)
	}

	memberEntityList, err := dao.GetJieYiDao().GetJieYiMemberListEntity(serverId)
	if err != nil {
		return err
	}
	// 添加结义成员对象
	for _, memberEntity := range memberEntityList {
		jieYi := s.getJieYiInfo(memberEntity.JieYiId)
		if jieYi == nil {
			continue
		}
		memberObj := newJieYiMemberObject(jieYi)
		err = memberObj.FromEntity(memberEntity)
		if err != nil {
			return
		}

		s.addJieYiMember(memberObj)
	}

	// 加载结义留言数据
	leaveWordEntityList, err := dao.GetJieYiDao().GetJieYiLeaveWordListEntity(serverId)
	if err != nil {
		return
	}

	s.jieYiLeaveWordMap = make(map[int64]*JieYiLeaveWordObject)

	for _, leaveWordEntity := range leaveWordEntityList {
		obj := NewJieYiLeaveWordObject()
		err = obj.FromEntity(leaveWordEntity)
		if err != nil {
			return
		}

		s.addJieYiLeaveWord(obj)
	}

	for _, info := range s.jieYiInfoMap {
		info.updateJieYiRank()
	}

	// 加载邀请数据
	inviteEntityList, err := dao.GetJieYiDao().GetJieYiInviteListEntity(serverId)
	if err != nil {
		return
	}

	s.jieYiInviteMap = make(map[int64]*JieYiInviteObject)

	for _, entity := range inviteEntityList {
		obj := NewJieYiInviteObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}

		s.addJieYiInivteObj(obj)
	}

	return
}

func (s *jieYiService) getLeaveWordObj(playerId int64) *JieYiLeaveWordObject {
	obj, ok := s.jieYiLeaveWordMap[playerId]
	if !ok {
		return nil
	}
	return obj
}

func (s *jieYiService) addJieYiLeaveWord(jieYiObj *JieYiLeaveWordObject) {
	obj := s.getLeaveWordObj(jieYiObj.GetPlayerId())
	if obj != nil {
		return
	}

	s.jieYiLeaveWordMap[jieYiObj.GetPlayerId()] = jieYiObj
}

// 玩家结义，删除该玩家结义墙数据
func (s *jieYiService) deletePlayerLevelWord(playerId int64) {
	obj := s.getLeaveWordObj(playerId)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.deleteTime = now
	obj.SetModified()
	delete(s.jieYiLeaveWordMap, playerId)
}

func (s *jieYiService) getJieYiInivteObj(inviteId int64) *JieYiInviteObject {
	obj, ok := s.jieYiInviteMap[inviteId]
	if !ok {
		return nil
	}
	return obj
}

func (s *jieYiService) addJieYiInivteObj(inviteObj *JieYiInviteObject) {
	obj := s.getJieYiInivteObj(inviteObj.inviteId)
	if obj != nil {
		return
	}
	s.jieYiInviteMap[inviteObj.inviteId] = inviteObj
}

func (s *jieYiService) addJieYi(jieYiObj *JieYiObject) {
	obj := s.getJieYiObj(jieYiObj.GetDBId())
	if obj != nil {
		return
	}
	jieYiInfo := &JieYi{}
	jieYiInfo.jieYiObject = jieYiObj
	s.jieYiInfoMap[jieYiObj.GetDBId()] = jieYiInfo
}

func (s *jieYiService) getJieYiObj(id int64) *JieYiObject {
	info, ok := s.jieYiInfoMap[id]
	if !ok {
		return nil
	}
	return info.jieYiObject
}

func (s *jieYiService) getJieYiInfo(jieYiId int64) *JieYi {
	info, ok := s.jieYiInfoMap[jieYiId]
	if !ok {
		return nil
	}
	return info
}

func (s *jieYiService) addJieYiMember(memberObj *JieYiMemberObject) {
	obj := s.getJieYiMember(memberObj.GetPlayerId())
	if obj != nil {
		return
	}
	s.jieYiMemberMap[memberObj.GetPlayerId()] = memberObj

	info := s.getJieYiInfo(memberObj.GetJieYiId())
	if info == nil {
		return
	}
	info.addJieYiMember(memberObj)
}

func (s *jieYiService) removeJieYiMember(memberObj *JieYiMemberObject) {
	now := global.GetGame().GetTimeService().Now()
	memberObj.deleteTime = now
	memberObj.SetModified()
	memberObj.GetJieYi().removeJieYiMember(memberObj.GetPlayerId())
	delete(s.jieYiMemberMap, memberObj.GetPlayerId())
}

func (s *jieYiService) removeJieYi(jieYi *JieYi) (memList []*JieYiMemberObject) {
	memList = make([]*JieYiMemberObject, jieYi.getMemberNum())
	copy(memList, jieYi.jieYiMemberObjectList)
	for _, mem := range memList {
		s.removeJieYiMember(mem)
	}
	now := global.GetGame().GetTimeService().Now()
	jieYi.jieYiObject.deleteTime = now
	jieYi.jieYiObject.SetModified()
	delete(s.jieYiInfoMap, jieYi.getJieYiId())
	return
}

func (s *jieYiService) getJieYiMember(playerId int64) *JieYiMemberObject {
	obj, ok := s.jieYiMemberMap[playerId]
	if !ok {
		return nil
	}
	return obj
}

func (s *jieYiService) AddJieYiInviteData(
	inviteId int64,
	inviteeId int64,
	name string,
	daoJu jieyitypes.JieYiDaoJuType,
	inviteDaoJu jieyitypes.JieYiDaoJuType,
	inviteToken jieyitypes.JieYiTokenType,
	inviteTokenLev int32,
	nameLev int32) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.getJieYiInivteObj(inviteId)
	//已经邀请过
	if obj != nil {
		return false
	}

	//判断邀请人
	inviteObj := s.getJieYiMember(inviteId)
	//判断是否结义
	if inviteObj != nil {
		jieYi := inviteObj.GetJieYi()
		//不是老大
		if jieYi.getLaoDa() != inviteId {
			return false
		}
		//满人
		if jieYi.isFull() {
			return false
		}
	}

	// 判断被邀请人
	inviteeObj := s.getJieYiMember(inviteeId)
	if inviteeObj != nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	obj = NewJieYiInviteObject()
	id, _ := idutil.GetId()
	obj.id = id
	obj.serverId = global.GetGame().GetServerIndex()
	obj.state = jieyitypes.InviteStateInit
	obj.daoJu = daoJu
	obj.inviteDaoJu = inviteDaoJu
	obj.inviteToken = inviteToken
	obj.inviteTokenLev = inviteTokenLev
	obj.nameLev = nameLev
	obj.name = name
	obj.inviteId = inviteId
	obj.inviteeId = inviteeId
	obj.createTime = now
	obj.updateTime = now
	obj.SetModified()
	s.jieYiInviteMap[inviteId] = obj
	return true
}

func (s *jieYiService) JieYiInviteFail(inviteId int64, inviteeId int64) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	return s.jieYiInviteFail(inviteId, inviteeId)
}

func (s *jieYiService) jieYiInviteFail(inviteId int64, inviteeId int64) bool {
	obj := s.getJieYiInivteObj(inviteId)
	if obj == nil {
		return false
	}
	if obj.inviteeId != inviteeId {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	if !obj.Fail(now) {
		return false
	}

	obj.deleteTime = now
	obj.SetModified()
	delete(s.jieYiInviteMap, inviteId)
	event.Emit(jieyieventtypes.JieYiEventTypeJieYiInviteFail, obj, nil)
	return true
}

func (s *jieYiService) JieYiSuccess(inviteId int64, inviteeId int64, inviteeDaoJu jieyitypes.JieYiDaoJuType, inviteeToken jieyitypes.JieYiTokenType, inviteeTokenLev, inviteeNameLev int32) (jieYi *JieYi, inviteMemObj *JieYiMemberObject, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//验证有没有这个邀请
	inviteObj := s.getJieYiInivteObj(inviteId)
	if inviteObj == nil {
		return nil, nil, false
	}

	//验证是否同一个人
	if inviteObj.inviteeId != inviteeId {
		return nil, nil, false
	}

	//判断被邀请人
	inviteeMemObj := s.getJieYiMember(inviteeId)
	if inviteeMemObj != nil {
		return nil, nil, false
	}
	inviteePlayer := player.GetOnlinePlayerManager().GetPlayerById(inviteeId)
	invitePlayer := player.GetOnlinePlayerManager().GetPlayerById(inviteId)
	//有一个不在线
	if inviteePlayer == nil || invitePlayer == nil {
		return nil, nil, false
	}
	//邀请状态不对
	if !inviteObj.IfCanSuccess() {
		return nil, nil, false
	}
	now := global.GetGame().GetTimeService().Now()
	name := inviteObj.GetJieYiName()
	daoJu := inviteObj.GetJieYiDaoJu()
	//判断邀请人
	inviteMemObj = s.getJieYiMember(inviteId)

	//判断是否结义
	if inviteMemObj != nil {
		jieYi = inviteMemObj.GetJieYi()
		//不是老大
		if jieYi.getLaoDa() != inviteId {
			return nil, nil, false
		}
		//满人
		if jieYi.isFull() {
			return nil, nil, false
		}

		// 判断邀请人的道具是否能被替换
		if inviteObj.inviteDaoJu < daoJu {
			inviteMemObj.jieYiDaoJu = daoJu
			inviteMemObj.updateTime = now
			inviteMemObj.SetModified()
		}

		// 判断被邀请人的道具是否能被替换
		if inviteeDaoJu < daoJu {
			inviteeDaoJu = daoJu
		}

		inviteeMemObj = newJieYiMemberObjectWithPlayer(jieYi, inviteePlayer, inviteeDaoJu, inviteeToken, inviteeTokenLev, inviteeNameLev, now)
		inviteeMemObj.SetModified()
		s.addJieYiMember(inviteeMemObj)
		// 删除结义留言数据
		s.deletePlayerLevelWord(inviteeId)
		jieYi.updateJieYiRank()

		event.Emit(jieyieventtypes.JieYiEventTypeJionJieYiLog, jieYi, inviteeId)
	} else {
		inviteDaoJu := inviteObj.GetInviteDaoJu()
		if inviteDaoJu < daoJu {
			inviteDaoJu = daoJu
		}
		//创建
		jieYi, inviteMemObj = s.createJieYi(invitePlayer, name, inviteDaoJu, inviteObj.inviteToken, inviteObj.inviteTokenLev, inviteObj.nameLev, now)
		if inviteeDaoJu < daoJu {
			inviteeDaoJu = daoJu
		}
		inviteeMemObj = newJieYiMemberObjectWithPlayer(jieYi, inviteePlayer, inviteeDaoJu, inviteeToken, inviteeTokenLev, inviteeNameLev, now)
		inviteeMemObj.SetModified()
		s.addJieYiMember(inviteeMemObj)
		// 删除结义留言数据
		s.deletePlayerLevelWord(inviteeId)
		s.deletePlayerLevelWord(inviteId)
		jieYi.updateJieYiRank()

		event.Emit(jieyieventtypes.JieYiEventTypeJionJieYiLog, jieYi, inviteId)
		event.Emit(jieyieventtypes.JieYiEventTypeJionJieYiLog, jieYi, inviteeId)
	}

	flag = inviteObj.Success(now)
	if !flag {
		panic(fmt.Errorf("jieyi:应该成功"))
	}
	inviteObj.deleteTime = now
	inviteObj.SetModified()
	delete(s.jieYiInviteMap, inviteId)

	return jieYi, inviteMemObj, true
}

func (s *jieYiService) JieChuJieYiSuccess(playerId int64) (jieYi *JieYi, tiMemList []*JieYiMemberObject, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	memberObj := s.getJieYiMember(playerId)
	if memberObj == nil {
		return nil, nil, false
	}
	jieYi = memberObj.GetJieYi()
	memNum := memberObj.GetJieYi().getMemberNum()
	if memNum <= 2 {
		tiMemList = s.removeJieYi(memberObj.GetJieYi())
	} else {
		s.removeJieYiMember(memberObj)
		tiMemList = append(tiMemList, memberObj)
	}

	// 离开结义日志
	for _, tiMem := range tiMemList {
		event.Emit(jieyieventtypes.JieYiEventTypeLeaveJieYiLog, jieYi, tiMem.GetPlayerId())
	}

	return jieYi, tiMemList, true
}

func (s *jieYiService) JieYiTiRen(laoDaId int64, memId int64) (jieYi *JieYi, tiMemList []*JieYiMemberObject, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if laoDaId == memId {
		return nil, nil, false
	}
	//判断踢人
	laoDaMemObj := s.getJieYiMember(laoDaId)
	if laoDaMemObj == nil {
		return nil, nil, false
	}
	jieYi = laoDaMemObj.GetJieYi()
	//判断是不是老大
	if jieYi.getLaoDa() != laoDaId {
		return nil, nil, false
	}
	//不是成员
	memObj := s.getJieYiMember(memId)
	if memObj == nil {
		return nil, nil, false
	}
	//不是同一个结义
	if memObj.GetJieYi().getJieYiId() != jieYi.getJieYiId() {
		return nil, nil, false
	}
	memNum := jieYi.getMemberNum()
	if memNum <= 2 {
		tiMemList = s.removeJieYi(jieYi)
	} else {
		s.removeJieYiMember(memObj)
		tiMemList = append(tiMemList, memObj)
	}

	// 离开结义日志
	for _, tiMem := range tiMemList {
		event.Emit(jieyieventtypes.JieYiEventTypeLeaveJieYiLog, jieYi, tiMem.GetPlayerId())
	}

	return jieYi, tiMemList, true
}

func (s *jieYiService) GetJieYiMemberRank(playerId int64) int32 {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	memberObj := s.getJieYiMember(playerId)
	if memberObj == nil {
		return 0
	}
	info := s.getJieYiInfo(memberObj.jieYiId)
	if info == nil {
		return 0
	}
	rank, obj := info.getJieYiMemberIndexAndObj(playerId)
	if obj == nil {
		return 0
	}
	return int32(rank) + 1
}

func (s *jieYiService) createJieYi(pl player.Player, name string, daoJu jieyitypes.JieYiDaoJuType, token jieyitypes.JieYiTokenType, tokenLev, nameLev int32, now int64) (jieYi *JieYi, memObj *JieYiMemberObject) {
	playerId := pl.GetId()
	jieYi = s.initJieYi(playerId, name)
	//now := global.GetGame().GetTimeService().Now()
	memObj = newJieYiMemberObjectWithPlayer(jieYi, pl, daoJu, token, tokenLev, nameLev, now-1)
	memObj.SetModified()
	s.addJieYiMember(memObj)
	return
}

func (s *jieYiService) initJieYi(playerId int64, name string) *JieYi {
	now := global.GetGame().GetTimeService().Now()
	obj := NewJieYiObject()
	id, _ := idutil.GetId()
	obj.id = id
	obj.serverId = global.GetGame().GetServerIndex()
	obj.originServerId = global.GetGame().GetServerIndex()
	obj.name = name
	obj.createTime = now
	obj.SetModified()

	info := &JieYi{}
	info.jieYiObject = obj
	s.jieYiInfoMap[id] = info

	return info
}

const (
	maxLeaveWorld = 100
)

func (s *jieYiService) GetJieYiLeaveWord() (objList []*JieYiLeaveWordObject) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	for _, obj := range s.jieYiLeaveWordMap {
		objList = append(objList, obj)
	}
	sort.Sort(sort.Reverse(jieYiLeaveWorldList(objList)))
	if len(objList) >= maxLeaveWorld {
		return objList[:maxLeaveWorld]
	}
	return objList
}

func (s *jieYiService) GetJieYi(jieYiId int64) *JieYi {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	obj := s.getJieYiInfo(jieYiId)
	return obj
}

func (s *jieYiService) GetJieYiInfo(jieYiId int64) *JieYiObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	obj := s.getJieYiInfo(jieYiId)
	return obj.GetJieYiObject()
}

func (s *jieYiService) GetJieYiMemberNum(jieYiId int64) int32 {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	info := s.getJieYiInfo(jieYiId)
	if info == nil {
		return 0
	}
	return info.getMemberNum()
}

func (s *jieYiService) GetJieYiMemberList(jieYiId int64) []*JieYiMemberObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	info := s.getJieYiInfo(jieYiId)
	if info == nil {
		return nil
	}
	return info.jieYiMemberObjectList
}

func (s *jieYiService) IsAlreadyJieYi(playerId int64) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	obj := s.getJieYiMember(playerId)
	if obj == nil {
		return false
	}
	return true
}

func (s *jieYiService) IsFullMember(playerId int64) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	memberObj := s.getJieYiMember(playerId)
	if memberObj == nil {
		return false
	}

	flag := s.isJieyiMemberFull(memberObj.jieYiId)
	return flag
}

func (s *jieYiService) UpdateShengWeiZhi(playerId int64, lev int32, num int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	obj := s.getJieYiMember(playerId)
	if obj == nil {
		return
	}
	obj.nameLev = lev
	obj.shengWeiZhi = num
	obj.updateTime = now
	obj.SetModified()
}

func (s *jieYiService) UpdateTokenLevel(playerId int64, level int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	obj := s.getJieYiMember(playerId)
	if obj == nil {
		return
	}
	obj.tokenLev = level
	obj.updateTime = now
	obj.SetModified()
}

func (s *jieYiService) IsNameRepetitive(name string) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.isNameRepetitive(name)
}

func (s *jieYiService) isNameRepetitive(name string) bool {

	// 结义数据
	for _, info := range s.jieYiInfoMap {
		if info.getMemberNum() == 0 {
			continue
		}

		if info.jieYiObject.GetName() == name {
			return false
		}
	}

	// 结义邀请数据
	for _, obj := range s.jieYiInviteMap {
		if obj.state == jieyitypes.InviteStateFail {
			continue
		}

		if obj.name == name {
			return false
		}
	}
	return true
}

// 结义是否满人
func (s *jieYiService) isJieyiMemberFull(jieYiId int64) bool {
	info := s.getJieYiInfo(jieYiId)
	if info == nil {
		return false
	}
	return info.isFull()
}

func (s *jieYiService) IsJieYiLaoDa(playerId int64) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	memberObj := s.getJieYiMember(playerId)
	if memberObj == nil {
		return true
	}
	info := s.getJieYiInfo(memberObj.jieYiId)
	if info.getLaoDa() != playerId {
		return false
	}
	return true
}

func (s *jieYiService) UpdatePlayerForce(playerId int64, force int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	memberObj := s.getJieYiMember(playerId)
	if memberObj != nil {
		memberObj.force = force
		memberObj.updateTime = now
		memberObj.SetModified()
	}

	leaveObj := s.getLeaveWordObj(playerId)
	if leaveObj != nil {
		leaveObj.force = force
		leaveObj.updateTime = now
		leaveObj.SetModified()
	}
}

func (s *jieYiService) UpdatePlayerLevel(playerId int64, level int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	memberObj := s.getJieYiMember(playerId)
	if memberObj != nil {
		memberObj.level = level
		memberObj.updateTime = now
		memberObj.SetModified()
	}

	leaveObj := s.getLeaveWordObj(playerId)
	if leaveObj != nil {
		leaveObj.level = level
		leaveObj.updateTime = now
		leaveObj.SetModified()
	}
}

func (s *jieYiService) UpdatePlayerZhuanSheng(playerId int64, zhuanSheng int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	memberObj := s.getJieYiMember(playerId)
	if memberObj != nil {
		memberObj.zhuanSheng = zhuanSheng
		memberObj.updateTime = now
		memberObj.SetModified()
	}
}

func (s *jieYiService) UpdatePlayerName(playerId int64, name string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	memberObj := s.getJieYiMember(playerId)
	if memberObj != nil {
		memberObj.name = name
		memberObj.updateTime = now
		memberObj.SetModified()
	}
}

func (s *jieYiService) PlayerLogin(playerId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memberObj := s.getJieYiMember(playerId)
	if memberObj != nil {
		memberObj.onlineStatus = playertypes.PlayerOnlineStateOnline
		memberObj.SetModified()
	}

	leaveObj := s.getLeaveWordObj(playerId)
	if leaveObj != nil {
		leaveObj.onlineStatus = playertypes.PlayerOnlineStateOnline
		leaveObj.SetModified()
	}
}

func (s *jieYiService) SetJieYiName(jieYiId int64, name string) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	flag := s.isNameRepetitive(name)
	if !flag {
		return false
	}
	obj := s.getJieYiObj(jieYiId)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	obj.name = name
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (s *jieYiService) GetInviteData(inviteId int64) *JieYiInviteObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	obj, ok := s.jieYiInviteMap[inviteId]
	if !ok {
		return nil
	}
	return obj
}

func (s *jieYiService) PlayerLogout(playerId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memberObj := s.getJieYiMember(playerId)
	if memberObj != nil {
		memberObj.onlineStatus = playertypes.PlayerOnlineStateOffline
		memberObj.SetModified()
	}

	leaveObj := s.getLeaveWordObj(playerId)
	if leaveObj != nil {
		leaveObj.onlineStatus = playertypes.PlayerOnlineStateOffline
		leaveObj.SetModified()
	}

}

func (s *jieYiService) GetJieYiMemberInfo(playerId int64) *JieYiMemberObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	obj, ok := s.jieYiMemberMap[playerId]
	if !ok {
		return nil
	}

	return obj
}

func (s *jieYiService) TokenChangeSucess(playerId int64, token jieyitypes.JieYiTokenType) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getJieYiMember(playerId)
	if obj == nil {
		return false
	}
	if obj.tokenType >= token {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	obj.tokenType = token
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (s *jieYiService) TokenChangeLevel(playerId int64, level int32) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getJieYiMember(playerId)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	obj.tokenLev = level
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (s *jieYiService) DaoJuChangeSucess(playerId int64, daoJu jieyitypes.JieYiDaoJuType) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.getJieYiMember(playerId)
	if obj == nil {
		return false
	}
	if obj.jieYiDaoJu >= daoJu {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	obj.jieYiDaoJu = daoJu
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (s *jieYiService) initLeaveWordObj(pl player.Player, playerId int64, leaveWord string) *JieYiLeaveWordObject {
	obj := NewJieYiLeaveWordObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.serverId = global.GetGame().GetServerIndex()
	obj.leaveWord = leaveWord
	obj.onlineStatus = playertypes.PlayerOnlineStateOnline
	obj.lastPostTime = now
	obj.name = pl.GetName()
	obj.level = pl.GetLevel()
	obj.role = pl.GetRole()
	obj.sex = pl.GetSex()
	obj.force = pl.GetForce()
	obj.updateTime = now
	obj.playerId = playerId
	obj.createTime = now
	obj.SetModified()
	return obj
}

func (s *jieYiService) AddJieYiLeaveWord(playerId int64, leaveWord string) bool {
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return false
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.getLeaveWordObj(playerId)
	if obj == nil {
		obj = s.initLeaveWordObj(pl, playerId, leaveWord)
		s.addJieYiLeaveWord(obj)
		return true
	}
	now := global.GetGame().GetTimeService().Now()
	obj.leaveWord = leaveWord
	obj.updateTime = now
	obj.SetModified()

	return true
}

func (s *jieYiService) CheckJieYiInviteState() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	existTime := constantTemp.YaoQingExistTime
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range s.jieYiInviteMap {
		if obj.state == jieyitypes.InviteStateInit {
			if now-obj.createTime < existTime {
				continue
			}
			s.jieYiInviteFail(obj.inviteId, obj.inviteeId)
		}
	}

	faBuJieYiMaxTime := constantTemp.FaBuJieYiMaxTime
	for _, obj := range s.jieYiLeaveWordMap {
		if now-obj.createTime < faBuJieYiMaxTime {
			continue
		}
		s.deletePlayerLevelWord(obj.playerId)
	}
}

func (s *jieYiService) Heartbeat() {
	s.CheckJieYiInviteState()
}

var (
	once  sync.Once
	jieyi *jieYiService
)

func Init() (err error) {
	once.Do(func() {
		jieyi = &jieYiService{}
		err = jieyi.init()
	})
	return
}

func GetJieYiService() JieYiService {
	return jieyi
}
