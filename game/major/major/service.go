package major

import (
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	"sync"
)

//双休服务接口处理
type MajorService interface {
	Heartbeat()
	//双修邀请
	MajorInvite(playerId int64, playerName string, spouseId int64, spouseName string, majorType majortypes.MajorType, fubenId int32)
	//双休配偶决策
	MajorInviteDeal(pl player.Player, result bool) (inviteName string, codeResult majortypes.MajorPairCodeType)
	//玩家取消双休邀请
	CancleMajorInvite(pl player.Player) (codeResult majortypes.MajorPairCodeType)
	//获取邀请
	GetMajorInvite(playerId int64) *majortypes.MajorInvite
	//移除invite
	RemoveInvite(playerId int64, spouseId int64)
}

type majorService struct {
	//双休读写锁
	rwm sync.RWMutex
	//邀请交互使用
	inviteMap map[int64]*majortypes.MajorInvite
}

//初始化
func (ms *majorService) init() (err error) {
	ms.inviteMap = make(map[int64]*majortypes.MajorInvite)

	return
}

//移除invite
func (ms *majorService) RemoveInvite(playerId int64, spouseId int64) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	delete(ms.inviteMap, playerId)
	delete(ms.inviteMap, spouseId)
}

//获取邀请
func (ms *majorService) GetMajorInvite(playerId int64) *majortypes.MajorInvite {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	majorInvite, exist := ms.inviteMap[playerId]
	if !exist {
		return nil
	}
	return majorInvite
}

//双修邀请
func (ms *majorService) MajorInvite(playerId int64, playerName string, spouseId int64, spouseName string, majorType majortypes.MajorType, fubenId int32) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	majorInvite := majortypes.NewMajorInvite(playerId, playerName, spouseId, spouseName, now, majorType, fubenId)

	ms.inviteMap[playerId] = majorInvite
	ms.inviteMap[spouseId] = majorInvite
}

//玩家取消双修邀请
func (ms *majorService) CancleMajorInvite(pl player.Player) (codeResult majortypes.MajorPairCodeType) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	codeResult = majortypes.MajorPairCodeTypeSucess
	majorInvite, exist := ms.inviteMap[pl.GetId()]
	if !exist {
		codeResult = majortypes.MajorPairCodeTypeDeal
		return
	}
	delete(ms.inviteMap, majorInvite.PlayerId)
	delete(ms.inviteMap, majorInvite.SpouseId)
	//发送事件
	gameevent.Emit(majoreventtypes.EventTypeMajorInviteCancle, pl, majorInvite)
	return
}

func (ms *majorService) checkInviteIsOverdue() {
	diffTime := int64(30 * common.SECOND)
	now := global.GetGame().GetTimeService().Now()
	for _, majorInvite := range ms.inviteMap {
		createTime := majorInvite.CreateTime
		if now-createTime >= diffTime {
			delete(ms.inviteMap, majorInvite.PlayerId)
			delete(ms.inviteMap, majorInvite.SpouseId)
			gameevent.Emit(majoreventtypes.EventTypeMajorInviteNoAnswer, majorInvite.PlayerId, majorInvite)
		}
	}
}

//心跳
func (ms *majorService) Heartbeat() {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()
	ms.checkInviteIsOverdue()
}

//配偶决策
func (ms *majorService) MajorInviteDeal(pl player.Player, agree bool) (inviteName string, codeResult majortypes.MajorPairCodeType) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	codeResult = majortypes.MajorPairCodeTypeSucess

	majorInvite, exist := ms.inviteMap[pl.GetId()]
	if !exist {
		codeResult = majortypes.MajorPairCodeTypeCancle
		return
	}
	inviteName = majorInvite.PlayerName
	//对方已下线
	spl := player.GetOnlinePlayerManager().GetPlayerById(majorInvite.PlayerId)
	if spl == nil {
		codeResult = majortypes.MajorPairCodeTypeCancle
		return
	}

	playerId := majorInvite.PlayerId
	spouseId := majorInvite.SpouseId
	delete(ms.inviteMap, playerId)
	delete(ms.inviteMap, spouseId)

	//发送事件
	eventData := majoreventtypes.CreateMajorInviteDealEventData(playerId, pl, agree, majorInvite.FuBenType, majorInvite.FuBenId)
	gameevent.Emit(majoreventtypes.EventTypeMajorInviteDeal, nil, eventData)
	return
}

var (
	once sync.Once
	cs   *majorService
)

func Init() (err error) {
	once.Do(func() {
		cs = &majorService{}
		err = cs.init()
	})
	return err
}

func GetMajorService() MajorService {
	return cs
}
