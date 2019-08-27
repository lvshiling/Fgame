package bagua

import (
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	baguatypes "fgame/fgame/game/bagua/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"sync"
)

//八卦秘境接口处理
type BaGuaService interface {
	Heartbeat()
	//夫妻助战邀请
	PairInvite(playerId int64, playerName string, spouseId int64, spouseName string, level int32)
	//配偶决策
	PairInviteDeal(pl player.Player, result bool) (inviteName string, codeResult baguatypes.BaGuaPairCodeType)
	//玩家取消夫妻助战邀请
	CanclePairInvite(pl player.Player) (codeResult baguatypes.BaGuaPairCodeType)
	//夫妻助战挑战失败
	PairChallegeFail(playerId int64)
	//夫妻助战配偶中途退出
	PairSpouseExit(spouseId int64)
	//是否存在夫妻助战
	IsExistPairKill(playerId int64) (spouseId int64, pairFlag bool)
	//获取邀请
	GetBaGuaInvite(playerId int64) *baguatypes.BaGuaInvite
	//移除invite
	RemoveInvite(playerId int64, spouseId int64)
}

type baGuaService struct {
	//夫妻助战读写锁
	rwm sync.RWMutex
	//邀请交互使用
	inviteMap map[int64]*baguatypes.BaGuaInvite
	//夫妻助战 下一关使用
	pairMap map[int64]*baguatypes.BaGuaPair
}

//初始化
func (rs *baGuaService) init() (err error) {
	rs.inviteMap = make(map[int64]*baguatypes.BaGuaInvite)
	rs.pairMap = make(map[int64]*baguatypes.BaGuaPair)

	return
}

//移除invite
func (rs *baGuaService) RemoveInvite(playerId int64, spouseId int64) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	delete(rs.inviteMap, playerId)
	delete(rs.inviteMap, spouseId)
}

//获取邀请
func (rs *baGuaService) GetBaGuaInvite(playerId int64) *baguatypes.BaGuaInvite {
	rs.rwm.RLock()
	defer rs.rwm.RUnlock()

	baGuaInvite, exist := rs.inviteMap[playerId]
	if !exist {
		return nil
	}
	return baGuaInvite
}

//是否存在夫妻助战
func (rs *baGuaService) IsExistPairKill(playerId int64) (spouseId int64, pairFlag bool) {
	rs.rwm.RLock()
	defer rs.rwm.RUnlock()

	baGuaPair, exist := rs.pairMap[playerId]
	if !exist {
		return 0, false
	}
	if baGuaPair.GetPlayerId() != playerId {
		return 0, false
	}
	return baGuaPair.GetSpouseId(), true
}

func (rs *baGuaService) removeBaGuaPair(playerId int64) {
	baGuaPair, exist := rs.pairMap[playerId]
	if !exist {
		return
	}
	delete(rs.pairMap, baGuaPair.GetPlayerId())
	delete(rs.pairMap, baGuaPair.GetSpouseId())
}

//夫妻助战挑战失败
func (rs *baGuaService) PairChallegeFail(playerId int64) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	rs.removeBaGuaPair(playerId)

}

//夫妻助战配偶中途退出
func (rs *baGuaService) PairSpouseExit(spouseId int64) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	rs.removeBaGuaPair(spouseId)
}

//夫妻助战邀请
func (rs *baGuaService) PairInvite(playerId int64, playerName string, spouseId int64, spouseName string, level int32) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	baGuaInvite := baguatypes.NewBaGuaInvite(playerId, playerName, spouseId, spouseName, level, now)

	rs.inviteMap[playerId] = baGuaInvite
	rs.inviteMap[spouseId] = baGuaInvite
}

//玩家取消夫妻助战邀请
func (rs *baGuaService) CanclePairInvite(pl player.Player) (codeResult baguatypes.BaGuaPairCodeType) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	codeResult = baguatypes.BaGuaPairCodeTypeSucess
	baGuaInvite, exist := rs.inviteMap[pl.GetId()]
	if !exist {
		codeResult = baguatypes.BaGuaPairCodeTypeDeal
		return
	}
	delete(rs.inviteMap, baGuaInvite.PlayerId)
	delete(rs.inviteMap, baGuaInvite.SpouseId)
	//发送事件
	gameevent.Emit(baguaeventtypes.EventTypeBaGuaPairInviteCancle, pl, baGuaInvite.SpouseId)

	return
}

func (rs *baGuaService) checkInviteIsOverdue() {
	diffTime := int64(30 * common.SECOND)
	now := global.GetGame().GetTimeService().Now()
	for _, baGuaInvite := range rs.inviteMap {
		createTime := baGuaInvite.CreateTime
		if now-createTime >= diffTime {
			delete(rs.inviteMap, baGuaInvite.PlayerId)
			delete(rs.inviteMap, baGuaInvite.SpouseId)
			gameevent.Emit(baguaeventtypes.EventTypeBaGuaPairNoAnswer, baGuaInvite.PlayerId, baGuaInvite.SpouseName)
		}
	}
}

//心跳
func (rs *baGuaService) Heartbeat() {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	rs.checkInviteIsOverdue()
}

//配偶决策
func (rs *baGuaService) PairInviteDeal(pl player.Player, result bool) (inviteName string, codeResult baguatypes.BaGuaPairCodeType) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	codeResult = baguatypes.BaGuaPairCodeTypeSucess

	baGuaInvite, exist := rs.inviteMap[pl.GetId()]
	if !exist {
		codeResult = baguatypes.BaGuaPairCodeTypeCancle
		return
	}
	inviteName = baGuaInvite.PlayerName
	//对方已下线
	spl := player.GetOnlinePlayerManager().GetPlayerById(baGuaInvite.PlayerId)
	if spl == nil {
		codeResult = baguatypes.BaGuaPairCodeTypeCancle
		return
	}

	playerId := baGuaInvite.PlayerId
	spouseId := baGuaInvite.SpouseId
	delete(rs.inviteMap, playerId)
	delete(rs.inviteMap, spouseId)

	baGuaPair := baguatypes.NewBaGuaPair(playerId, spouseId)
	rs.pairMap[playerId] = baGuaPair
	rs.pairMap[spouseId] = baGuaPair

	//发送事件
	eventData := baguaeventtypes.CreateBaGuaPairInviteDealEventData(playerId, pl, baGuaInvite.Level, result)
	gameevent.Emit(baguaeventtypes.EventTypeBaGuaPairInviteDeal, nil, eventData)

	return
}

var (
	once sync.Once
	cs   *baGuaService
)

func Init() (err error) {
	once.Do(func() {
		cs = &baGuaService{}
		err = cs.init()
	})
	return err
}

func GetBaGuaService() BaGuaService {
	return cs
}
