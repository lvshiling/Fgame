package obj

import (
	"fgame/fgame/game/global"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

type RankTypeData interface {
	init(timestamp int64) (err error)
	updateRankList(timestamp int64) (err error)
	UpdateRankDataList(timestamp int64) (err error)
	GetFirstId() (fristId int64)
	GetPos(id int64) (pos int32)
	ResetRankTime()
	GetRankingInfoList() []*ranktypes.RankingInfo
}

type Handler interface {
	CreateRankObj(config *ranktypes.RankConfig) RankTypeData
}

type HandlerFunc func(config *ranktypes.RankConfig) RankTypeData

func (hf HandlerFunc) CreateRankObj(config *ranktypes.RankConfig) RankTypeData {
	return hf(config)
}

type RankMap struct {
	rankDataMap map[ranktypes.RankType]RankTypeData
	//榜单配置
	rankConfigMap map[ranktypes.RankType]*ranktypes.RankConfig
	//刷新周期
	refreshTime int64
}

func NewRankMap(refreshTime int64) *RankMap {
	rankMap := &RankMap{
		rankDataMap:   make(map[ranktypes.RankType]RankTypeData),
		rankConfigMap: make(map[ranktypes.RankType]*ranktypes.RankConfig),
		refreshTime:   refreshTime,
	}

	return rankMap
}

//初始化排行榜
func (r *RankMap) Init() (err error) {
	for rankType, config := range r.rankConfigMap {
		handler, exist := handlerMap[rankType]
		if !exist {
			continue
		}
		_, exist = r.rankDataMap[rankType]
		if !exist {
			rankTypeData := handler.CreateRankObj(config)
			r.rankDataMap[rankType] = rankTypeData
		}
	}
	return
}

//启动
func (r *RankMap) Star() (err error) {
	now := global.GetGame().GetTimeService().Now()
	timestamp, flag := timeutils.GetIntervalTimeStampMs(now, r.refreshTime)
	if !flag {
		return fmt.Errorf("rank:刷新时间应该是能被24小时整除的")
	}

	for _, rankData := range r.rankDataMap {
		err = rankData.init(timestamp)
		if err != nil {
			return
		}
	}
	return
}

//更新本服排行榜
func (r *RankMap) UpdateRank() (err error) {
	now := global.GetGame().GetTimeService().Now()
	timestamp, _ := timeutils.GetIntervalTimeStampMs(now, r.refreshTime)
	for _, rankData := range r.rankDataMap {
		err = rankData.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	return
}

//注册榜单
func (r *RankMap) RegisterRank(rankType ranktypes.RankType, config *ranktypes.RankConfig) bool {
	_, exist := r.rankConfigMap[rankType]
	if !exist {
		config.RankType = rankType
		r.rankConfigMap[rankType] = config
		return true
	}

	return false
}

func (r *RankMap) GetRankTypeData(rankType ranktypes.RankType) RankTypeData {
	return r.rankDataMap[rankType]
}

func (r *RankMap) GetRankTypeDataMap() map[ranktypes.RankType]RankTypeData {
	return r.rankDataMap
}

func (r *RankMap) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.GetRankingInfoList()
}

var (
	handlerMap = make(map[ranktypes.RankType]Handler)
)

func init() {
	handlerMap[ranktypes.RankTypeForce] = HandlerFunc(newForceRank)
	handlerMap[ranktypes.RankTypeWeapon] = HandlerFunc(newWeaponRank)
	handlerMap[ranktypes.RankTypeWing] = HandlerFunc(newWingRank)
	handlerMap[ranktypes.RankTypeBodyShield] = HandlerFunc(newBodyShieldRank)
	handlerMap[ranktypes.RankTypeMount] = HandlerFunc(newMountRank)
	handlerMap[ranktypes.RankTypeGang] = HandlerFunc(newGangRank)
	handlerMap[ranktypes.RankTypeShenFa] = HandlerFunc(newShenFaRank)
	handlerMap[ranktypes.RankTypeLingYu] = HandlerFunc(newLingYuRank)
	handlerMap[ranktypes.RankTypeFeather] = HandlerFunc(newFeatherRank)
	handlerMap[ranktypes.RankTypeShield] = HandlerFunc(newShieldRank)
	handlerMap[ranktypes.RankTypeCharge] = HandlerFunc(newChargeRank)
	handlerMap[ranktypes.RankTypeCost] = HandlerFunc(newCostRank)
	handlerMap[ranktypes.RankTypeAnQi] = HandlerFunc(newAnQiRank)
	handlerMap[ranktypes.RankTypeCharm] = HandlerFunc(newCharmRank)
	handlerMap[ranktypes.RankTypeCount] = HandlerFunc(newCountRank)
	handlerMap[ranktypes.RankTypeFaBao] = HandlerFunc(newFaBaoRank)
	handlerMap[ranktypes.RankTypeXianTi] = HandlerFunc(newXianTiRank)
	handlerMap[ranktypes.RankTypeShiHunFan] = HandlerFunc(newShiHunFanRank)
	handlerMap[ranktypes.RankTypeTianMoTi] = HandlerFunc(newTianMoTiRank)
	handlerMap[ranktypes.RankTypeLevel] = HandlerFunc(newLevelRank)
	handlerMap[ranktypes.RankTypeFeiSheng] = HandlerFunc(newFeiShengRank)
	handlerMap[ranktypes.RankTypeMarryDevelop] = HandlerFunc(newMarryDevelopRank)

	handlerMap[ranktypes.RankTypeLingQi] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingBing] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingYi] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingShen] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingTongYu] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingBao] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingTi] = HandlerFunc(newLingTongDevRank)
	handlerMap[ranktypes.RankTypeLingTongLevel] = HandlerFunc(newLingTongLevelRank)

	handlerMap[ranktypes.RankTypeLingTongForce] = HandlerFunc(newLingTongForceRank)
	handlerMap[ranktypes.RankTypeGoldEquipForce] = HandlerFunc(newGoldEquipForceRank)
	handlerMap[ranktypes.RankTypeDianXingForce] = HandlerFunc(newDianXingForceRank)
	handlerMap[ranktypes.RankTypeShenQiForce] = HandlerFunc(newShenQiForceRank)
	handlerMap[ranktypes.RankTypeMingGeForce] = HandlerFunc(newMingGeForceRank)
	handlerMap[ranktypes.RankTypeShengHenForce] = HandlerFunc(newShengHenForceRank)
	handlerMap[ranktypes.RankTypeZhenFaForce] = HandlerFunc(newZhenFaForceRank)
	handlerMap[ranktypes.RankTypeTuLongEquipForce] = HandlerFunc(newTuLongEquipForceRank)
	handlerMap[ranktypes.RankTypeBabyForce] = HandlerFunc(newBabyForceRank)
	handlerMap[ranktypes.RankTypeZhuanSheng] = HandlerFunc(newZhuanShengRank)
}
