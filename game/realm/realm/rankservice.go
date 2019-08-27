package realm

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	rankentity "fgame/fgame/game/rank/entity"
	"fgame/fgame/game/realm/dao"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	realmtypes "fgame/fgame/game/realm/types"
	"sort"
	"sync"
)

//境界排行接口处理
type RealmRankService interface {
	Heartbeat()
	//天劫塔第一
	GetTianJieTaFirstId() (playerId int64)
	//玩家天劫塔排行
	GetTianJieTaTopThreeAndMyPos(playerId int64) (list []*rankentity.RankCommonData, pos int32)
	//刷新天劫塔排名
	RefreshTianJieTaRank(playerId int64, playerName string, level int32, usedTime int64)
	//夫妻助战邀请
	PairInvite(playerId int64, playerName string, spouseId int64, spouseName string, level int32)
	//配偶决策
	PairInviteDeal(pl player.Player, result realmtypes.RealmResultType) (inviteName string, codeResult realmtypes.RealmPairCodeType)
	//玩家取消夫妻助战邀请
	CanclePairInvite(pl player.Player) (codeResult realmtypes.RealmPairCodeType)
	//夫妻助战挑战失败
	PairChallegeFail(playerId int64)
	//夫妻助战配偶中途退出
	PairSpouseExit(spouseId int64)
	//是否存在夫妻助战
	IsExistPairKill(playerId int64) (spouseId int64, pairFlag bool)
	//获取邀请
	GetRealmInvite(playerId int64) *realmtypes.RealmInvite
	//移除invite
	RemoveInvite(playerId int64, spouseId int64)
	//重置玩家名字
	PlayerNameChanged(pl player.Player)
}

type realmRankService struct {
	//读写锁
	rwm sync.RWMutex
	//天劫塔topRankMax
	tianJieTaList []*rankentity.RankCommonData
	//排名
	rankMap map[int64]int32
	//夫妻助战读写锁
	pairRwm sync.RWMutex
	//邀请交互使用
	inviteMap map[int64]*realmtypes.RealmInvite
	//夫妻助战 下一关使用
	pairMap map[int64]*realmtypes.RealmPair
}

//初始化
func (rs *realmRankService) init() (err error) {
	rs.tianJieTaList = make([]*rankentity.RankCommonData, 0, 8)
	rs.rankMap = make(map[int64]int32)
	rs.inviteMap = make(map[int64]*realmtypes.RealmInvite)
	rs.pairMap = make(map[int64]*realmtypes.RealmPair)
	rankNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRealTopRank)
	tianJieTaList, err := dao.GetRealmDao().GetRankRealmList(rankNum)
	if err != nil {
		return
	}
	if tianJieTaList != nil {
		for index, tianJieTa := range tianJieTaList {
			usedTime := tianJieTa.UsedTime
			value := realmtypes.Combine(tianJieTa.Level, usedTime)
			data := rankentity.NewRankCommonData(tianJieTa.PlayerId, value, tianJieTa.PlayerName)
			rs.tianJieTaList = append(rs.tianJieTaList, data)
			rs.rankMap[tianJieTa.PlayerId] = int32(index) + 1
		}
	}
	return
}

//移除invite
func (rs *realmRankService) RemoveInvite(playerId int64, spouseId int64) {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()

	delete(rs.inviteMap, playerId)
	delete(rs.inviteMap, spouseId)
}

//获取邀请
func (rs *realmRankService) GetRealmInvite(playerId int64) *realmtypes.RealmInvite {
	rs.pairRwm.RLock()
	defer rs.pairRwm.RUnlock()

	realmInvite, exist := rs.inviteMap[playerId]
	if !exist {
		return nil
	}
	return realmInvite
}

//是否存在夫妻助战
func (rs *realmRankService) IsExistPairKill(playerId int64) (spouseId int64, pairFlag bool) {
	rs.pairRwm.RLock()
	defer rs.pairRwm.RUnlock()

	realmPair, exist := rs.pairMap[playerId]
	if !exist {
		return 0, false
	}
	if realmPair.GetPlayerId() != playerId {
		return 0, false
	}
	return realmPair.GetSpouseId(), true
}

func (rs *realmRankService) removeRealmPair(playerId int64) {
	realmPair, exist := rs.pairMap[playerId]
	if !exist {
		return
	}
	delete(rs.pairMap, realmPair.GetPlayerId())
	delete(rs.pairMap, realmPair.GetSpouseId())
}

//夫妻助战挑战失败
func (rs *realmRankService) PairChallegeFail(playerId int64) {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()
	rs.removeRealmPair(playerId)

}

//夫妻助战配偶中途退出
func (rs *realmRankService) PairSpouseExit(spouseId int64) {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()
	rs.removeRealmPair(spouseId)
}

//夫妻助战邀请
func (rs *realmRankService) PairInvite(playerId int64, playerName string, spouseId int64, spouseName string, level int32) {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	realmInvite := realmtypes.NewRealmInvite(playerId, playerName, spouseId, spouseName, level, now)

	rs.inviteMap[playerId] = realmInvite
	rs.inviteMap[spouseId] = realmInvite
}

//玩家取消夫妻助战邀请
func (rs *realmRankService) CanclePairInvite(pl player.Player) (codeResult realmtypes.RealmPairCodeType) {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()

	codeResult = realmtypes.RealmPairCodeTypeSucess
	realmInvite, exist := rs.inviteMap[pl.GetId()]
	if !exist {
		codeResult = realmtypes.RealmPairCodeTypeDeal
		return
	}
	delete(rs.inviteMap, realmInvite.PlayerId)
	delete(rs.inviteMap, realmInvite.SpouseId)
	//发送事件
	gameevent.Emit(realmeventtypes.EventTypeRealmPairInviteCancle, pl, realmInvite.SpouseId)

	return
}

func (rs *realmRankService) checkInviteIsOverdue() {
	diffTime := int64(30 * common.SECOND)
	now := global.GetGame().GetTimeService().Now()
	for _, realmInvite := range rs.inviteMap {
		createTime := realmInvite.CreateTime
		if now-createTime >= diffTime {
			delete(rs.inviteMap, realmInvite.PlayerId)
			delete(rs.inviteMap, realmInvite.SpouseId)
			gameevent.Emit(realmeventtypes.EventTypeRealmPairNoAnswer, realmInvite.PlayerId, realmInvite.SpouseName)
		}
	}
}

//心跳
func (rs *realmRankService) Heartbeat() {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()
	rs.checkInviteIsOverdue()
}

//配偶决策
func (rs *realmRankService) PairInviteDeal(pl player.Player, result realmtypes.RealmResultType) (inviteName string, codeResult realmtypes.RealmPairCodeType) {
	rs.pairRwm.Lock()
	defer rs.pairRwm.Unlock()

	codeResult = realmtypes.RealmPairCodeTypeSucess

	realmInvite, exist := rs.inviteMap[pl.GetId()]
	if !exist {
		codeResult = realmtypes.RealmPairCodeTypeCancle
		return
	}
	inviteName = realmInvite.PlayerName
	//对方已下线
	spl := player.GetOnlinePlayerManager().GetPlayerById(realmInvite.PlayerId)
	if spl == nil {
		codeResult = realmtypes.RealmPairCodeTypeCancle
		return
	}

	playerId := realmInvite.PlayerId
	spouseId := realmInvite.SpouseId
	delete(rs.inviteMap, playerId)
	delete(rs.inviteMap, spouseId)

	agree := true
	if result == realmtypes.RealmResultTypeNo {
		agree = false
	}

	realmPair := realmtypes.NewRealmPair(playerId, spouseId)
	rs.pairMap[playerId] = realmPair
	rs.pairMap[spouseId] = realmPair

	//发送事件
	eventData := realmeventtypes.CreateRealmPairInviteDealEventData(playerId, pl, realmInvite.Level, agree)
	gameevent.Emit(realmeventtypes.EventTypeRealmPairInviteDeal, nil, eventData)

	return
}

//天劫塔第一
func (rs *realmRankService) GetTianJieTaFirstId() (playerId int64) {
	playerId = 0
	rankLen := len(rs.tianJieTaList)
	if rankLen == 0 {
		return
	}
	return rs.tianJieTaList[0].Id
}

//玩家天劫塔排行
func (rs *realmRankService) GetTianJieTaTopThreeAndMyPos(playerId int64) (list []*rankentity.RankCommonData, pos int32) {
	rs.rwm.RLock()
	rs.rwm.RUnlock()
	pos = 0
	len := len(rs.tianJieTaList)
	if len == 0 {
		return nil, pos
	}
	addLen := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRealTopRank))
	if addLen >= len {
		addLen = len
	}
	pos = rs.rankMap[playerId]
	return rs.tianJieTaList[:addLen], pos
}

func (rs *realmRankService) getFirstId() (playerId int64) {
	playerId = 0
	if len(rs.tianJieTaList) != 0 {
		playerId = rs.tianJieTaList[0].Id
	}
	return
}

//刷新天劫塔排名
func (rs *realmRankService) RefreshTianJieTaRank(playerId int64, playerName string, level int32, usedTime int64) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	maxLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRealTopRank)
	value := realmtypes.Combine(level, usedTime)

	//oldFirstId := rs.getFirstId()
	pos, exist := rs.rankMap[playerId]
	if !exist {
		data := rankentity.NewRankCommonData(playerId, value, playerName)
		rs.tianJieTaList = append(rs.tianJieTaList, data)
		pos = 1
	} else {
		rs.tianJieTaList[pos-1].Value = value
	}

	sort.Sort(sort.Reverse(rankentity.RankCommonDataList(rs.tianJieTaList)))

	if len(rs.tianJieTaList) > int(maxLen) {
		rs.tianJieTaList = rs.tianJieTaList[:maxLen]
	}

	// newFirstId := rs.getFirstId()
	// if oldFirstId != newFirstId {
	// 	gameevent.Emit(realmeventtypes.EventTypeRealmFirstChange, newFirstId, oldFirstId)
	// }

	if exist && rs.tianJieTaList[pos-1].Id == playerId {
		return
	}

	rs.rankMap = make(map[int64]int32)
	for index, tianJieTa := range rs.tianJieTaList {
		rs.rankMap[tianJieTa.Id] = int32(index + 1)
	}
	return

}

func (rs *realmRankService) PlayerNameChanged(pl player.Player) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	playerId := pl.GetId()
	pos, exist := rs.rankMap[playerId]
	if !exist {
		return
	}
	rs.tianJieTaList[pos-1].Name = pl.GetName()
}

var (
	once sync.Once
	cs   *realmRankService
)

func Init() (err error) {
	once.Do(func() {
		cs = &realmRankService{}
		err = cs.init()
	})
	return err
}

func GetRealmRankService() RealmRankService {
	return cs
}
