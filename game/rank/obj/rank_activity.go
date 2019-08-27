package obj

import (
	"fgame/fgame/game/global"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type RankActivityMap struct {
	rankDataMap map[int32]RankTypeData
	//榜单配置
	rankConfigMap map[int32]*ranktypes.RankConfig
	//类型map
	rankTypeMap map[int32]ranktypes.RankType
	//刷新周期
	refreshTime int64
}

func NewRankActivityMap(refreshTime int64) *RankActivityMap {
	rankMap := &RankActivityMap{
		rankDataMap:   make(map[int32]RankTypeData),
		rankConfigMap: make(map[int32]*ranktypes.RankConfig),
		rankTypeMap:   make(map[int32]ranktypes.RankType),
		refreshTime:   refreshTime,
	}

	return rankMap
}

//初始化排行榜
func (r *RankActivityMap) Init() (err error) {
	for groupId, config := range r.rankConfigMap {
		rankType := r.rankTypeMap[groupId]
		handler, exist := handlerMap[rankType]
		if !exist {
			return fmt.Errorf("rank:排行榜初始化，排行榜数据初始化处理器不存在")
		}

		_, exist = r.rankDataMap[groupId]
		if !exist {
			rankTypeData := handler.CreateRankObj(config)
			r.rankDataMap[groupId] = rankTypeData

			now := global.GetGame().GetTimeService().Now()
			timestamp, flag := timeutils.GetIntervalTimeStampMs(now, r.refreshTime)
			if !flag {
				return fmt.Errorf("rank:刷新时间应该是能被24小时整除的")
			}
			err = rankTypeData.init(timestamp)
			if err != nil {
				return
			}
		}
	}
	return
}

//启动
func (r *RankActivityMap) Star() (err error) {
	return
}

//更新本服活动排行榜
func (r *RankActivityMap) UpdateRank() (err error) {
	now := global.GetGame().GetTimeService().Now()
	timestamp, _ := timeutils.GetIntervalTimeStampMs(now, r.refreshTime)
	for groupId, rankData := range r.rankDataMap {
		// 超过最大展示时间移除
		config := r.rankConfigMap[groupId]
		maxExpire := config.EndTime + config.MaxExpireTime
		if timestamp > maxExpire && config.IsCanExpire {
			delete(r.rankDataMap, groupId)
			delete(r.rankConfigMap, groupId)
			delete(r.rankTypeMap, groupId)
			continue
		}

		err = rankData.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	return
}

//注册榜单
func (r *RankActivityMap) RegisterRank(rankType ranktypes.RankType, config *ranktypes.RankConfig) bool {
	// _, exist := r.rankTypeMap[config.GroupId]
	// if exist {
	// 	return false
	// }
	r.rankTypeMap[config.GroupId] = rankType

	orignConfig, exist := r.rankConfigMap[config.GroupId]
	if exist {
		//结束时间一样
		if config.EndTime == orignConfig.EndTime {
			return false
		}
	}

	log.Infof("注册%s活动排行榜,groupId:%d", rankType.String(), config.GroupId)
	config.RankType = rankType
	r.rankConfigMap[config.GroupId] = config
	err := r.initActivityRank(config.GroupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"rankType": rankType.String(),
				"groupId":  config.GroupId,
			}).Error("rank:注册活动排行榜失败")
		return false
	}
	return true
}

func (r *RankActivityMap) initActivityRank(groupId int32) (err error) {
	config, ok := r.rankConfigMap[groupId]
	if !ok {
		return
	}

	rankType := r.rankTypeMap[groupId]
	handler, exist := handlerMap[rankType]
	if !exist {
		return
	}
	_, exist = r.rankDataMap[groupId]
	//移除排行榜数据
	delete(r.rankDataMap, groupId)

	rankTypeData := handler.CreateRankObj(config)
	r.rankDataMap[groupId] = rankTypeData

	now := global.GetGame().GetTimeService().Now()
	timestamp, flag := timeutils.GetIntervalTimeStampMs(now, r.refreshTime)
	if !flag {
		return fmt.Errorf("rank:刷新时间应该是能被24小时整除的")
	}
	err = rankTypeData.init(timestamp)
	if err != nil {
		return
	}
	return
}

func (r *RankActivityMap) GetRankDataMap() map[int32]RankTypeData {
	return r.rankDataMap
}

func (r *RankActivityMap) GetRankData(groupId int32) RankTypeData {
	return r.rankDataMap[groupId]
}

func (r *RankActivityMap) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.GetRankingInfoList()
}
