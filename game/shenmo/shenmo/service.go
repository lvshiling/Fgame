package shenmo

import (
	"context"
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/core/heartbeat"
	shenmoclient "fgame/fgame/cross/shenmo/client"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/center/center"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shenmo/dao"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	shenmoscene "fgame/fgame/game/shenmo/scene"
	shenmotemplate "fgame/fgame/game/shenmo/template"
	"fgame/fgame/game/shenmo/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sort"
	"sync"
)

type ShenMoService interface {
	Heartbeat()
	Star() (err error)
	//获取我的排名
	GetMyRank(rankType types.RankTimeType, servrId int32, allianceId int64) (pos int32, rankTime int64)
	//获取排名
	GetRankList(rankType types.RankTimeType, page int32) (dataList []*ShenMoRankData, rankTime int64)
	//获取排行时间
	GetRankTime(rankType types.RankTimeType) (rankTime int64)

	//--------------场景----------------

	//创建神魔战场数据
	CreateShenMoScene(mapId int32, endTime int64) scene.Scene
	//获取神魔战场数据
	GetShenMoScene() scene.Scene
	//神魔战场活动结束
	ShenMoSceneFinish()
	//参加神魔战场
	Attend(playerId int64) (lineUpPos int32, flag bool)
	//玩家是否有排队
	GetHasLineUp(playerId int64) (lineUpPos int32, flag bool)
	//玩家取消排队
	CancleLineUp(playerId int64) (flag bool)
	//同步场景人数
	SyncSceneNum(scenePlayerNum int32)
	//获取第一个排队玩家
	RemoveFirstLineUpPlayer(scenePlayerNum int32)
	//获取当前还在排队的
	GetAllLineUpList() (lineUpList []int64)
	//活动是否开始
	IsShenMoActivityTime() (flag bool)

	//--------------排行榜----------------
	//获取排行榜积分
	GetJiFenNum(allianceId int64) int32
	//仙盟增加积分
	AddJiFenNum(serverId int32, allianceId int64, allianceName string, addNum int32)
	//添加本地服务器积分
	AddLocalJiFenNum(allianceId int64, allianceName string, addNum int32)
}

type shenMoData struct {
	//参加记录
	attendMap map[int64]struct{}
	//排队记录
	lineList []int64
	//场景人数
	num int32
	//总人数
	totalNum int32
}

func NewShenMoData() *shenMoData {
	return &shenMoData{}
}

func (s *shenMoData) init() {
	s.attendMap = make(map[int64]struct{})
	s.lineList = make([]int64, 0, 8)
	s.num = 0
	s.totalNum = 0
}

type shenMoService struct {
	rwm          sync.RWMutex
	shenMoClient shenmoclient.ShenMoClient
	hbRunner     heartbeat.HeartbeatTaskRunner

	//  ---------------跨服本服通用-----------------
	//上周排行榜
	lastRankList []*ShenMoRankData
	//本周排行榜
	thisRankList []*ShenMoRankData
	//上周排行榜刷新时间
	lastTime int64
	//本周排行榜刷新时间
	thisTime int64

	// ---------------跨服本服场景通用-----------------
	shenMoDataObj *shenMoData
	//神魔战场场景数据
	sceneData shenmoscene.ShenMoSceneData
	//活动类型
	activityType activitytypes.ActivityType

	// ---------------本服排行榜数据-----------------
	//本周神魔战场排行榜数据
	thisShenMoRankList []*ShenMoRankObject
	//上周神魔战场排行榜数据
	lastShenMoRankList []*ShenMoRankObject
	//排行榜刷新时间对象
	shenMoRankTimeObj *ShenMoRankTimeObject
}

func (s *shenMoService) init(activityType activitytypes.ActivityType) (err error) {
	s.lastRankList = make([]*ShenMoRankData, 0, 8)
	s.thisRankList = make([]*ShenMoRankData, 0, 8)
	s.shenMoDataObj = NewShenMoData()
	s.shenMoDataObj.init()
	s.activityType = activityType
	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()

	if s.activityType == activitytypes.ActivityTypeLocalShenMoWar {
		s.hbRunner.AddTask(CreateShenMoRankDataTask(s))
		s.hbRunner.AddTask(CreateShenMoRankTimeTask(s))

		// conn := center.GetCenterService().GetCross(centertypes.GameServerTypePlatform)
		// if conn == nil {
		// 	return fmt.Errorf("shenmo:跨服连接不存在")
		// }
		// //TODO 修改可能连接变化了
		// s.shenMoClient = shenmoclient.NewShenMoClient(conn)
		err = s.resetClient()
		if err != nil {
			return
		}

		//加载本服排行榜
		err = s.loadLocalShenMoRank()
		if err != nil {
			return
		}
	}
	return nil
}

func (s *shenMoService) loadLocalShenMoRank() (err error) {
	serverId := global.GetGame().GetServerIndex()
	//排行榜列表
	shenMoRankList, err := dao.GetShenMoDao().GetShenMoRankList(serverId)
	if err != nil {
		return
	}

	for _, shenMoRank := range shenMoRankList {
		obj := NewShenMoRankObject()
		obj.FromEntity(shenMoRank)
		s.thisShenMoRankList = append(s.thisShenMoRankList, obj)
		s.lastShenMoRankList = append(s.lastShenMoRankList, obj)
	}

	sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
	sort.Sort(sort.Reverse(LastShenMoRankObjectList(s.lastShenMoRankList)))

	//排行刷新时间
	shenMoRankTimeEntity, err := dao.GetShenMoDao().GetShenMoRankTimeEntity(serverId)
	if err != nil {
		return
	}
	if shenMoRankTimeEntity == nil {
		s.initShenMoRankTimeObject(serverId)
	} else {
		s.shenMoRankTimeObj = NewShenMoRankTimeObject()
		s.shenMoRankTimeObj.FromEntity(shenMoRankTimeEntity)
	}
	return
}

//第一次初始化
func (s *shenMoService) initShenMoRankTimeObject(serverId int32) {
	po := NewShenMoRankTimeObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	po.id = id
	po.serverId = serverId
	po.thisTime = 0
	po.lastTime = 0
	po.createTime = now
	s.shenMoRankTimeObj = po
	po.SetModified()
}

func (s *shenMoService) GetMyRank(rankType types.RankTimeType, servrId int32, allianceId int64) (pos int32, rankTime int64) {
	pos = 0
	var rankList []*ShenMoRankData
	if rankType == types.RankTimeTypeLast {
		rankList = s.lastRankList
		rankTime = s.lastTime
	} else {
		rankList = s.thisRankList
		rankTime = s.thisTime
	}
	for index, rankObj := range rankList {
		if rankObj.allianceId == allianceId {
			pos = int32(index + 1)
		}
	}
	return
}

func (s *shenMoService) GetRankList(rankType types.RankTimeType, page int32) (dataList []*ShenMoRankData, rankTime int64) {
	var rankList []*ShenMoRankData
	if rankType == types.RankTimeTypeLast {
		rankList = s.lastRankList
		rankTime = s.lastTime
	} else {
		rankList = s.thisRankList
		rankTime = s.thisTime
	}
	pageIndex := page * ranktypes.PageLimit
	len := len(rankList)
	if pageIndex >= int32(len) {
		return
	}
	addLen := pageIndex + ranktypes.PageLimit
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	dataList = rankList[pageIndex:addLen]
	return
}

func (s *shenMoService) GetRankTime(rankType types.RankTimeType) (rankTime int64) {
	if rankType == types.RankTimeTypeLast {
		rankTime = s.lastTime
	} else {
		rankTime = s.thisTime
	}
	return
}

func (s *shenMoService) GetJiFenNum(allianceId int64) (jiFenNum int32) {
	if len(s.thisShenMoRankList) == 0 {
		return
	}
	for _, shenMoRank := range s.thisShenMoRankList {
		if shenMoRank.allianceId == allianceId {
			return shenMoRank.jiFenNum
		}
	}
	return
}

func (s *shenMoService) AddJiFenNum(serverId int32, allianceId int64, allianceName string, addNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj, flag := s.isExistAlliance(allianceId)
	if !flag {
		obj = initShenMoRankObject(serverId, allianceId, allianceName, addNum)
		s.thisShenMoRankList = append(s.thisShenMoRankList, obj)
		s.lastShenMoRankList = append(s.lastShenMoRankList, obj)
	} else {
		obj.jiFenNum += addNum
	}
	obj.SetModified()
	sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
}

func (s *shenMoService) AddLocalJiFenNum(allianceId int64, allianceName string, addNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	serverId := global.GetGame().GetServerIndex()
	obj, flag := s.isExistAlliance(allianceId)
	if !flag {
		obj = initShenMoRankObject(serverId, allianceId, allianceName, addNum)
		s.thisShenMoRankList = append(s.thisShenMoRankList, obj)
		s.lastShenMoRankList = append(s.lastShenMoRankList, obj)
	} else {
		obj.jiFenNum += addNum
	}
	obj.SetModified()
	sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
}

func (s *shenMoService) isExistAlliance(allianceId int64) (shenMoRankObj *ShenMoRankObject, flag bool) {
	if len(s.thisShenMoRankList) == 0 {
		return
	}
	for _, shenMoRank := range s.thisShenMoRankList {
		if shenMoRank.allianceId == allianceId {
			return shenMoRank, true
		}
	}
	return
}

func (s *shenMoService) Star() (err error) {
	if s.activityType == activitytypes.ActivityTypeLocalShenMoWar {
		s.syncRemoteRankList()
		s.syncLocalRankList()
	}
	return
}

//定时同步排行榜列表
func (s *shenMoService) syncRemoteRankList() (err error) {
	if s.shenMoClient == nil {
		err = s.resetClient()
		if err != nil {
			return
		}
	}
	//TODO 超时
	ctx := context.TODO()
	_, err = s.shenMoClient.GetShenMoRankList(ctx)
	if err != nil {
		return
	}
	// s.rwm.Lock()
	// defer s.rwm.Unlock()
	// s.lastTime = resp.LastRankData.RankTime
	// s.thisTime = resp.ThisRankData.RankTime
	// s.thisRankList = convertFromRankInfoList(resp.ThisRankData.RankInfoList)
	// s.lastRankList = convertFromRankInfoList(resp.LastRankData.RankInfoList)
	return nil
}

//定时同步本服排行榜列表
func (s *shenMoService) syncLocalRankList() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.lastTime = s.shenMoRankTimeObj.lastTime
	s.thisTime = s.shenMoRankTimeObj.thisTime
	s.thisRankList = convertFromRankObjectList(s.thisShenMoRankList, true)
	s.lastRankList = convertFromRankObjectList(s.lastShenMoRankList, false)
	return
}

func (s *shenMoService) resetClient() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypePlatform)
	if conn == nil {
		return fmt.Errorf("shenmo:跨服连接不存在")
	}

	//TODO 修改可能连接变化了
	s.shenMoClient = shenmoclient.NewShenMoClient(conn)
	return
}

func (s *shenMoService) Heartbeat() {
	s.hbRunner.Heartbeat()

	//
	s.checkShenMoActivity()
}

func (s *shenMoService) checkRefreshWeekRank() {
	now := global.GetGame().GetTimeService().Now()
	if s.shenMoRankTimeObj.thisTime == 0 {
		timeStamp, _ := timeutils.MondayFivePointTime(now)
		s.shenMoRankTimeObj.thisTime = timeStamp
		s.shenMoRankTimeObj.updateTime = now
		s.shenMoRankTimeObj.SetModified()
		return
	}

	diffTime := now - s.shenMoRankTimeObj.thisTime
	if diffTime < 0 {
		diffTime *= -1 //绝对值
	}
	if diffTime < int64(7*timeutils.DAY) {
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.refreshWeekRank(now)
}

func (s *shenMoService) refreshWeekRank(now int64) {
	timeStamp, _ := timeutils.MondayFivePointTime(now)
	s.shenMoRankTimeObj.lastTime = timeStamp - int64(7*timeutils.DAY)
	s.shenMoRankTimeObj.thisTime = timeStamp
	s.shenMoRankTimeObj.updateTime = now
	s.shenMoRankTimeObj.SetModified()

	for _, rankObj := range s.thisShenMoRankList {
		rankObj.lastJiFenNum = rankObj.jiFenNum
		rankObj.jiFenNum = 0
		rankObj.lastTime = now
		rankObj.updateTime = now
		rankObj.SetModified()
	}

	if len(s.thisShenMoRankList) != 0 {
		sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
	}
	if len(s.lastShenMoRankList) != 0 {
		sort.Sort(sort.Reverse(LastShenMoRankObjectList(s.lastShenMoRankList)))
	}
}

// -----------------------------------------

func (s *shenMoService) checkShenMoActivity() (err error) {
	if s.sceneData != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(s.activityType)
	activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
	if err != nil {
		return
	}
	if activityTimeTemplate == nil {
		return
	}
	endTime, err := activityTimeTemplate.GetEndTime(now)
	if err != nil {
		return
	}

	if s.sceneData != nil {
		return
	}

	s.CreateShenMoScene(activityTemplate.Mapid, endTime)
	return nil
}

func (s *shenMoService) CreateShenMoScene(mapId int32, endTime int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.sceneData != nil {
		return s.sceneData.GetScene()
	}

	s.shenMoDataObj.totalNum = shenmotemplate.GetShenMoTemplateService().GetShenMoConstantTemplate().PlayerLimitCount
	s.shenMoDataObj.lineList = make([]int64, 0, 8)

	sd := shenmoscene.CreateShenMoSceneData()
	sms := shenmoscene.CreateShenMoScene(mapId, endTime, sd)
	if sms == nil {
		return nil
	}
	s.sceneData = sd
	return sms
}

func (s *shenMoService) GetShenMoScene() scene.Scene {
	if s.sceneData == nil {
		return nil
	}

	return s.sceneData.GetScene()
}

func (s *shenMoService) ShenMoSceneFinish() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.shenMoDataObj.init()
	s.sceneData = nil
}

func (s *shenMoService) Attend(playerId int64) (lineUpPos int32, lineUpFlag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	lineLen := int32(len(s.shenMoDataObj.lineList))
	if s.shenMoDataObj.num >= s.shenMoDataObj.totalNum {
		lineUpFlag = true
		s.shenMoDataObj.lineList = append(s.shenMoDataObj.lineList, playerId)
		lineUpPos = lineLen
		return
	}
	s.shenMoDataObj.attendMap[playerId] = struct{}{}
	s.shenMoDataObj.num++
	return 0, false
}

func (s *shenMoService) getHasLineUp(playerId int64) (lineUpPos int32, flag bool) {
	for index, curPlayerId := range s.shenMoDataObj.lineList {
		if playerId == curPlayerId {
			lineUpPos = int32(index)
			flag = true
			break
		}
	}
	return
}

func (s *shenMoService) GetHasLineUp(playerId int64) (lineUpPos int32, flag bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	lineUpPos, flag = s.getHasLineUp(playerId)
	return
}

func (s *shenMoService) CancleLineUp(playerId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	index, flag := s.getHasLineUp(playerId)
	if !flag {
		return
	}
	if index == 0 && len(s.shenMoDataObj.lineList) <= 1 {
		s.shenMoDataObj.lineList = make([]int64, 0, 8)
		return true
	}
	s.shenMoDataObj.lineList = append(s.shenMoDataObj.lineList[:index], s.shenMoDataObj.lineList[index+1:]...)
	gameevent.Emit(shenmoeventtypes.EventTypeShenMoCancleLineUp, s.shenMoDataObj.lineList, int32(index))
	flag = true
	return
}

func (s *shenMoService) SyncSceneNum(scenePlayerNum int32) {
	if scenePlayerNum < 0 || scenePlayerNum > s.shenMoDataObj.totalNum {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.shenMoDataObj.num = scenePlayerNum
}

func (s *shenMoService) GetAllLineUpList() (lineUpList []int64) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.shenMoDataObj.lineList
}

func (s *shenMoService) IsShenMoActivityTime() (flag bool) {
	return s.sceneData != nil
}

func (s *shenMoService) RemoveFirstLineUpPlayer(scenePlayerNum int32) {
	if scenePlayerNum < 0 || scenePlayerNum > s.shenMoDataObj.totalNum {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()

	playerId := int64(0)
	s.shenMoDataObj.num = scenePlayerNum
	lineLen := int32(len(s.shenMoDataObj.lineList))
	if s.shenMoDataObj.num < s.shenMoDataObj.totalNum && lineLen != 0 {
		playerId = s.shenMoDataObj.lineList[0]
		if lineLen == 1 {
			s.shenMoDataObj.lineList = make([]int64, 0, 8)
		} else {
			s.shenMoDataObj.lineList = s.shenMoDataObj.lineList[1:]
		}
	}

	if playerId != 0 {
		s.shenMoDataObj.attendMap[playerId] = struct{}{}
		gameevent.Emit(shenmoeventtypes.EventTypeShenMoPlayerLineUpFinish, s.shenMoDataObj.lineList, playerId)
	}
	return
}

var (
	once sync.Once
	cs   *shenMoService
)

func Init(activityType activitytypes.ActivityType) (err error) {
	once.Do(func() {
		cs = &shenMoService{}
		err = cs.init(activityType)
	})
	return err
}

func GetShenMoService() ShenMoService {
	return cs
}
