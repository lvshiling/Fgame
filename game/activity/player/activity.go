package player

import (
	"fgame/fgame/game/activity/dao"
	activityeventtypes "fgame/fgame/game/activity/event/types"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	player.RegisterPlayerDataManager(types.PlayerActivityDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerActivityDataManager))
}

//玩家活动数据管理器
type PlayerActivityDataManager struct {
	p player.Player
	//玩家活动map
	playerActivityMap map[activitytypes.ActivityType]*PlayerActivityObject
	//玩家活动pk数据
	playerActivityPkMap map[activitytypes.ActivityType]*PlayerActivityPkObject
	//玩家活动排行数据
	playerActivityRankMap map[activitytypes.ActivityType]*PlayerActivityRankObject
	//玩家活动采集数据
	playerActivityCollectMap map[activitytypes.ActivityType]*PlayerActivityCollectObject

	//活动进入时间
	enterTimeMap map[activitytypes.ActivityType]int64
	//上次奖励时间
	lastRewTimeMap map[activitytypes.ActivityType]int64
	//上次奖励积分时间
	lastRewPointTimeMap map[activitytypes.ActivityType]int64
}

//玩家
func (m *PlayerActivityDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerActivityDataManager) Load() (err error) {
	//加载玩家活动数据列表
	activitysEntityArr, err := dao.GetActivityDao().GetActivitys(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range activitysEntityArr {
		newObj := CreatePlayerActivityObject(m.p)
		newObj.FromEntity(entity)
		m.playerActivityMap[newObj.activityType] = newObj
	}

	activitysPkListEntityArr, err := dao.GetActivityDao().GetActivityPkList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range activitysPkListEntityArr {
		newObj := CreatePlayerActivityPkObject(m.p)
		newObj.FromEntity(entity)
		m.playerActivityPkMap[newObj.activityType] = newObj
	}

	activitysRankListEntityArr, err := dao.GetActivityDao().GetActivityRankList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range activitysRankListEntityArr {
		newObj := CreatePlayerActivityRankObject(m.p)
		newObj.FromEntity(entity)
		m.playerActivityRankMap[newObj.activityType] = newObj
	}

	//
	collectEntityList, err := dao.GetActivityDao().GetActivityCollectList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range collectEntityList {
		newObj := CreatePlayerActivityCollectObject(m.p)
		newObj.FromEntity(entity)
		m.playerActivityCollectMap[newObj.activityType] = newObj
	}

	return nil
}

//加载后
func (m *PlayerActivityDataManager) AfterLoad() error {
	err := m.RefreshData()
	if err != nil {
		return err
	}
	return nil
}

func (m *PlayerActivityDataManager) RefreshData() error {
	now := global.GetGame().GetTimeService().Now()

	for _, obj := range m.playerActivityMap {
		isSame, err := timeutils.IsSameFive(obj.updateTime, now)
		if err != nil {
			return err
		}
		if !isSame {
			obj.attendTimes = 0
			obj.updateTime = now
			obj.SetModified()
		}
	}
	return nil
}

//心跳
func (m *PlayerActivityDataManager) Heartbeat() {

}

//生成玩家初始活动数据
func (m *PlayerActivityDataManager) initNewActivityObject(activityType activitytypes.ActivityType) *PlayerActivityObject {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	newObj := CreatePlayerActivityObject(m.p)
	newObj.id = id
	newObj.activityType = activityType
	newObj.createTime = now

	m.playerActivityMap[activityType] = newObj
	newObj.SetModified()

	return newObj
}

//是否有活动参与次数
func (m *PlayerActivityDataManager) IsHaveTimes(activityType activitytypes.ActivityType) bool {
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activityType)
	limitTimes := activityTemplate.JoinCount
	if limitTimes == 0 {
		return true
	}

	activityObj := m.GetActivity(activityType)
	attendTimes := activityObj.attendTimes
	return attendTimes < limitTimes
}

//更新挑战次数
func (m *PlayerActivityDataManager) AttendActivity(activityType activitytypes.ActivityType) {
	now := global.GetGame().GetTimeService().Now()
	activityObj := m.GetActivity(activityType)
	activityObj.attendTimes += 1
	activityObj.updateTime = now
	activityObj.SetModified()

	//发送记录事件
	gameevent.Emit(activityeventtypes.EventTypeActivityJoin, m.p, activityType)
}

//获取玩家活动信息
func (m *PlayerActivityDataManager) GetActivity(activityType activitytypes.ActivityType) *PlayerActivityObject {
	for typ, activityObj := range m.playerActivityMap {
		if typ == activityType {
			return activityObj
		}
	}

	return m.initNewActivityObject(activityType)
}

//获取玩家活动信息
func (m *PlayerActivityDataManager) GetActivityPkDataList() (dataList []*scene.PlayerActvitiyKillData) {
	for _, activityObj := range m.playerActivityPkMap {
		data := scene.CreatePlayerActvitiyKillData(activityObj.activityType, activityObj.killedNum, activityObj.lastKilledTime)
		dataList = append(dataList, data)
	}
	return dataList
}

//获取玩家活动信息
func (m *PlayerActivityDataManager) GetActivityRankDataList() (dataList []*scene.PlayerActvitiyRankData) {
	for _, activityObj := range m.playerActivityRankMap {
		data := scene.CreatePlayerActvitiyRankData(activityObj.activityType, activityObj.rankMap, activityObj.endTime)
		dataList = append(dataList, data)
	}
	return dataList
}

//获取玩家采集信息
func (m *PlayerActivityDataManager) GetActivityCollectDataList() (dataList []*scene.PlayerActvitiyCollectData) {
	for _, collectObj := range m.playerActivityCollectMap {
		data := scene.CreatePlayerActvitiyCollectData(collectObj.activityType, collectObj.countMap, collectObj.endTime)
		dataList = append(dataList, data)
	}
	return dataList
}

//获取玩家活动信息
func (m *PlayerActivityDataManager) UpdateActivityRankData(activityType activitytypes.ActivityType, rankMap map[int32]int64, endTime int64) {
	now := global.GetGame().GetTimeService().Now()
	activityRankObj, ok := m.playerActivityRankMap[activityType]
	if ok {
		activityRankObj.endTime = endTime
		activityRankObj.rankMap = rankMap
		activityRankObj.updateTime = now
		return
	}
	activityRankObj = CreatePlayerActivityRankObject(m.p)
	id, _ := idutil.GetId()
	activityRankObj.id = id
	activityRankObj.endTime = endTime
	activityRankObj.activityType = activityType
	activityRankObj.rankMap = rankMap
	activityRankObj.createTime = now
	m.playerActivityRankMap[activityType] = activityRankObj
}

//获取玩家活动信息
func (m *PlayerActivityDataManager) UpdateActivityPkData(activityType activitytypes.ActivityType, killedNum int32, lastKillTime int64) {
	now := global.GetGame().GetTimeService().Now()
	activityPkObj, ok := m.playerActivityPkMap[activityType]
	if ok {
		activityPkObj.killedNum = killedNum
		activityPkObj.lastKilledTime = lastKillTime
		activityPkObj.updateTime = now
		return
	}
	activityPkObj = CreatePlayerActivityPkObject(m.p)
	id, _ := idutil.GetId()
	activityPkObj.id = id
	activityPkObj.killedNum = killedNum
	activityPkObj.activityType = activityType
	activityPkObj.lastKilledTime = lastKillTime
	activityPkObj.createTime = now
	m.playerActivityPkMap[activityType] = activityPkObj
}

//更新采集数据
func (m *PlayerActivityDataManager) UpdateActivityCollectData(activityType activitytypes.ActivityType, countMap map[int32]int32, endTime int64) {
	now := global.GetGame().GetTimeService().Now()
	collectObj, ok := m.playerActivityCollectMap[activityType]
	if ok {
		collectObj.endTime = endTime
		collectObj.countMap = countMap
		collectObj.updateTime = now
		return
	}
	collectObj = CreatePlayerActivityCollectObject(m.p)
	id, _ := idutil.GetId()
	collectObj.id = id
	collectObj.endTime = endTime
	collectObj.activityType = activityType
	collectObj.countMap = countMap
	collectObj.createTime = now
	m.playerActivityCollectMap[activityType] = collectObj
}

//更新进入场景时间
func (m *PlayerActivityDataManager) UpdateEnterTime(typ activitytypes.ActivityType, endTime int64) {
	now := global.GetGame().GetTimeService().Now()
	m.enterTimeMap[typ] = now
	m.lastRewTimeMap[typ] = 0
	m.lastRewPointTimeMap[typ] = 0
}

//更新获取奖励时间
func (m *PlayerActivityDataManager) UpdateLastRewTime(typ activitytypes.ActivityType) {
	now := global.GetGame().GetTimeService().Now()
	m.lastRewTimeMap[typ] = now
}

//更新获取积分奖励时间
func (m *PlayerActivityDataManager) UpdateLastRewPointTime(typ activitytypes.ActivityType) {
	now := global.GetGame().GetTimeService().Now()
	m.lastRewPointTimeMap[typ] = now
}

func (m *PlayerActivityDataManager) GetEnterTime(typ activitytypes.ActivityType) int64 {
	return m.enterTimeMap[typ]
}

func (m *PlayerActivityDataManager) GetPreRewTime(typ activitytypes.ActivityType) int64 {
	return m.lastRewTimeMap[typ]
}

func (m *PlayerActivityDataManager) GetPreRewPointTime(typ activitytypes.ActivityType) int64 {
	return m.lastRewPointTimeMap[typ]
}

//退出活动
func (m *PlayerActivityDataManager) ExitActivity(typ activitytypes.ActivityType) {
	//修改排行榜
	activityRankObj, ok := m.playerActivityRankMap[typ]
	if ok {
		activityRankObj.SetModified()
	}
	//杀人cd
	activityPkObj, ok := m.playerActivityPkMap[typ]
	if ok {
		activityPkObj.SetModified()
	}
	//采集信息
	activityCollectObj, ok := m.playerActivityCollectMap[typ]
	if ok {
		activityCollectObj.SetModified()
	}
}

func CreatePlayerActivityDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerActivityDataManager{}
	m.p = p
	m.playerActivityMap = make(map[activitytypes.ActivityType]*PlayerActivityObject)
	m.playerActivityPkMap = make(map[activitytypes.ActivityType]*PlayerActivityPkObject)
	m.playerActivityRankMap = make(map[activitytypes.ActivityType]*PlayerActivityRankObject)
	m.playerActivityCollectMap = make(map[activitytypes.ActivityType]*PlayerActivityCollectObject)
	m.enterTimeMap = make(map[activitytypes.ActivityType]int64)
	m.lastRewTimeMap = make(map[activitytypes.ActivityType]int64)
	m.lastRewPointTimeMap = make(map[activitytypes.ActivityType]int64)

	return m
}
