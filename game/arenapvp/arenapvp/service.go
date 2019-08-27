package arenapvp

import (
	"context"
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/core/heartbeat"
	arenapvpclient "fgame/fgame/cross/arenapvp/client"
	"fgame/fgame/game/arenapvp/dao"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/center/center"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sync"
)

type ArenapvpService interface {
	Heartbeat()
	Star() (err error)
	//获取pvp赛程情况
	GetArenapvpPlayerDataList() (dataList []*arenapvpdata.PvpPlayerInfo)
	//获取海选赛程情况
	GetArenapvpElectionList() []*arenapvpdata.ElectionData
	//历届霸主信息
	GetAreanapvpBaZhuList(page, pageNum int32) (dataList []*arenapvpdata.BaZhuData, totalPage int32)
	//最近霸主信息
	GetLastAreanapvpBaZhu() *arenapvpdata.BaZhuData
	//当前竞猜
	GetArenapvpGuessData() *arenapvpdata.GuessData
	//参与竞猜
	AttendGuess(playerId int64, raceNumber int32, guessId int64, guessType arenapvptypes.ArenapvpType)
	//获取玩家竞猜记录
	GetPlayerGuessRecordList(playerId int64) []*ArenapvpGuessRecordObject
	//移除玩家竞猜记录
	RemovePlayerGuessRecord(playerId int64, raceNumber int32, guessType arenapvptypes.ArenapvpType)
}

type arenapvpService struct {
	rwm            sync.RWMutex
	arenapvpClient arenapvpclient.ArenapvpClient
	hbRunner       heartbeat.HeartbeatTaskRunner

	// ------------跨服数据----------------
	//pvp赛程列表
	arenapvpBattlePlayerList []*arenapvpdata.PvpPlayerInfo
	//海选赛程列表
	arenapvpElectionList []*arenapvpdata.ElectionData
	//历届霸主
	arenapvpBaZhuList []*arenapvpdata.BaZhuData
	//竞猜信息
	arenapvpGuessDataList []*arenapvpdata.GuessData
	//晋级推送记录
	winnerNoticeMap map[arenapvptypes.ArenapvpType]struct{}

	// ------------------------------
	guessAttendMap map[int64][]*ArenapvpGuessRecordObject
}

func (s *arenapvpService) init() (err error) {
	s.arenapvpBattlePlayerList = make([]*arenapvpdata.PvpPlayerInfo, 0, 32)
	s.guessAttendMap = make(map[int64][]*ArenapvpGuessRecordObject)
	s.winnerNoticeMap = make(map[arenapvptypes.ArenapvpType]struct{})

	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateArenapvpPlayerDataTask(s))

	err = s.resetClient()
	if err != nil {
		return
	}

	//加载竞猜记录
	err = s.loadGuessRecord()
	if err != nil {
		return err
	}

	return nil
}

func (s *arenapvpService) loadGuessRecord() (err error) {
	serverId := global.GetGame().GetServerIndex()

	//竞猜记录列表
	entityList, err := dao.GetArenapvpDao().GetArenapvpGuessRecordList(serverId)
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := NewArenapvpGuessRecordObject()
		obj.FromEntity(entity)
		s.addGuessRecord(obj)
	}

	return
}

func (s *arenapvpService) GetArenapvpGuessData() *arenapvpdata.GuessData {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.getLastGuessData()
}

func (s *arenapvpService) getLastGuessData() *arenapvpdata.GuessData {
	len := len(s.arenapvpGuessDataList)
	if len == 0 {
		return nil
	}
	return s.arenapvpGuessDataList[len-1]
}

func (s *arenapvpService) AttendGuess(playerId int64, raceNumber int32, guessId int64, guessType arenapvptypes.ArenapvpType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	serverId := global.GetGame().GetServerIndex()
	id, _ := idutil.GetId()

	recordObj := NewArenapvpGuessRecordObject()
	recordObj.id = id
	recordObj.serverId = serverId
	recordObj.playerId = playerId
	recordObj.raceNumber = raceNumber
	recordObj.guessType = guessType
	recordObj.guessId = guessId
	recordObj.status = arenapvptypes.ArenapvpGuessStateInit
	recordObj.createTime = now
	recordObj.SetModified()

	s.addGuessRecord(recordObj)
}

func (s *arenapvpService) GetPlayerGuessRecordList(playerId int64) (recordList []*ArenapvpGuessRecordObject) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.guessAttendMap[playerId]
}

func (s *arenapvpService) RemovePlayerGuessRecord(playerId int64, raceNumber int32, guessType arenapvptypes.ArenapvpType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.delGuessRecord(playerId, raceNumber, guessType)
	return
}

func (s *arenapvpService) addGuessRecord(recordObj *ArenapvpGuessRecordObject) {
	s.guessAttendMap[recordObj.playerId] = append(s.guessAttendMap[recordObj.playerId], recordObj)
}

func (s *arenapvpService) getGuessRecord(playerId int64, raceNum int32, guessType arenapvptypes.ArenapvpType) (int, *ArenapvpGuessRecordObject) {
	for index, attendObj := range s.guessAttendMap[playerId] {
		if attendObj.raceNumber != raceNum {
			continue
		}

		if attendObj.guessType != guessType {
			continue
		}
		return index, attendObj
	}

	return -1, nil

}

func (s *arenapvpService) delGuessRecord(playerId int64, raceNumber int32, guessType arenapvptypes.ArenapvpType) {
	index, attendObj := s.getGuessRecord(playerId, raceNumber, guessType)
	if attendObj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	attendObj.updateTime = now
	attendObj.deleteTime = now
	attendObj.SetModified()

	s.guessAttendMap[playerId] = append(s.guessAttendMap[playerId][:index], s.guessAttendMap[playerId][index+1:]...)
}

func (s *arenapvpService) GetArenapvpPlayerDataList() (dataList []*arenapvpdata.PvpPlayerInfo) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.arenapvpBattlePlayerList
}

func (s *arenapvpService) GetArenapvpElectionList() (dataList []*arenapvpdata.ElectionData) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.arenapvpElectionList
}

func (s *arenapvpService) GetAreanapvpBaZhuList(page, pageNum int32) (dataList []*arenapvpdata.BaZhuData, totalPage int32) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	if pageNum <= 0 {
		return
	}

	len := int32(len(s.arenapvpBaZhuList))
	if len == 0 {
		return
	}

	pageStart := page * pageNum
	pageEnd := pageStart + pageNum
	if pageStart >= len {
		return
	}

	if pageEnd > len {
		pageEnd = len
	}
	totalPage = len / pageNum
	if len%pageNum != 0 {
		totalPage += 1
	}
	return s.arenapvpBaZhuList[pageStart:pageEnd], totalPage
}

func (s *arenapvpService) GetLastAreanapvpBaZhu() *arenapvpdata.BaZhuData {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	len := len(s.arenapvpBaZhuList)
	if len == 0 {
		return nil
	}

	return s.arenapvpBaZhuList[len-1]
}

func (s *arenapvpService) Star() (err error) {
	err = s.syncRemoteArenapvpData()
	if err != nil {
		return
	}
	return
}

//定时同步排行榜列表
func (s *arenapvpService) syncRemoteArenapvpData() (err error) {
	if s.arenapvpClient == nil {
		err = s.resetClient()
		if err != nil {
			return
		}
	}
	//TODO 超时
	ctx := context.TODO()
	resp, err := s.arenapvpClient.GetArenapvpData(ctx)
	if err != nil {
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.arenapvpBaZhuList = convertToBaZhuList(resp.ArenapvpData.BaZhuList)
	s.arenapvpBattlePlayerList = convertFromAreanapvpPlayerList(resp.ArenapvpData.PlayerList)

	newElectionList := convertFromArenapvpElectionDataList(resp.ArenapvpData.ElectionDataList)
	s.electionLuckyNotice(newElectionList)
	s.arenapvpElectionList = newElectionList

	newGuessDataList := convertToGuessDataList(resp.ArenapvpData.GuessDataList)
	s.noticeGuessData(newGuessDataList)
	s.checkGuessData(newGuessDataList)
	s.arenapvpGuessDataList = newGuessDataList

	//晋级公告推送
	s.noticeWinnerList()
	return nil
}

//晋级公告推送
func (s *arenapvpService) noticeWinnerList() {
	lastGuessData := s.getLastGuessData()
	if lastGuessData == nil {
		return
	}

	// 是否公告过
	noticeType := lastGuessData.PvpType
	if noticeType == arenapvptypes.ArenapvpTypeFinals && lastGuessData.GetWinnerId() != 0 {
		noticeType = arenapvptypes.ArenapvpTypeChampion
	}
	_, ok := s.winnerNoticeMap[noticeType]
	if ok {
		return
	}

	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpWinnerNotice, s.arenapvpBattlePlayerList, noticeType)
	s.winnerNoticeMap[noticeType] = struct{}{}
}

//竞猜提醒推送
func (s *arenapvpService) noticeGuessData(newGuessDataList []*arenapvpdata.GuessData) {
	lenNew := len(newGuessDataList)
	if lenNew == 0 {
		return
	}

	lastNewGuessData := newGuessDataList[lenNew-1]
	if lastNewGuessData.GetWinnerId() != 0 {
		return
	}

	pvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(lastNewGuessData.PvpType)
	if pvpTemp == nil {
		return
	}

	//没结果超过战斗开始时间
	now := global.GetGame().GetTimeService().Now()
	if !pvpTemp.IfCanGuess(now) {
		return
	}

	lastGuessData := s.getLastGuessData()
	if lastGuessData == nil {
		gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpGuessBroadcast, lastNewGuessData, nil)
		return
	}

	// 与上一次类型相等
	if lastGuessData.PvpType == lastNewGuessData.PvpType {
		return
	}
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpGuessBroadcast, lastNewGuessData, nil)
}

//竞猜退还/结算
func (s *arenapvpService) checkGuessData(newGuessDataList []*arenapvpdata.GuessData) {
	var returnList []*ArenapvpGuessRecordObject
	var resultList []*ArenapvpGuessRecordObject
	now := global.GetGame().GetTimeService().Now()
	for playerId, attendObjList := range s.guessAttendMap {
		for _, attendObj := range attendObjList {
			//竞猜退还
			isReturn := true
			for _, newGuessData := range newGuessDataList {
				if attendObj.raceNumber != newGuessData.RaceNumber {
					continue
				}
				if attendObj.guessType != newGuessData.PvpType {
					continue
				}
				isReturn = false

				//结算
				if newGuessData.GetWinnerId() == 0 {
					continue
				}
				attendObj.winnerId = newGuessData.GetWinnerId()
				attendObj.status = arenapvptypes.ArenapvpGuessStateResult
				attendObj.updateTime = now
				attendObj.SetModified()

				pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
				if pl != nil {
					resultList = append(resultList, attendObj)
				}
			}

			// 退还
			if isReturn {
				attendObj.status = arenapvptypes.ArenapvpGuessStateReturn
				attendObj.updateTime = now
				attendObj.SetModified()
				pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
				if pl != nil {
					returnList = append(returnList, attendObj)
				}
			}
		}
	}

	//结算
	if len(resultList) > 0 {
		gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpGuessResult, nil, resultList)
		for _, attendObj := range resultList {
			s.delGuessRecord(attendObj.playerId, attendObj.raceNumber, attendObj.guessType)
		}
	}

	//退还
	if len(returnList) > 0 {
		gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpGuessReturn, nil, returnList)
		for _, attendObj := range returnList {
			s.delGuessRecord(attendObj.playerId, attendObj.raceNumber, attendObj.guessType)
		}
	}

	return
}

//海选幸运奖推送
func (s *arenapvpService) electionLuckyNotice(newElectionList []*arenapvpdata.ElectionData) {
	for _, newElection := range newElectionList {
		if len(newElection.LuckyNameText) == 0 {
			continue
		}

		electionData := s.getElectionData(newElection.ElectionIndex)
		if electionData == nil {
			// 发事件
			gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpElectionLuckyRew, newElection, nil)
			continue
		}

		if electionData.LastLuckyTime == newElection.LastLuckyTime {
			continue
		}

		// 发事件
		gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpElectionLuckyRew, newElection, nil)
	}
	return
}
func (s *arenapvpService) getElectionData(electionIndex int32) *arenapvpdata.ElectionData {
	for _, electionData := range s.arenapvpElectionList {
		if electionData.ElectionIndex != electionIndex {
			continue
		}

		return electionData
	}

	return nil
}

func (s *arenapvpService) resetClient() (err error) {
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypeAll)
	if conn == nil {
		return fmt.Errorf("arenapvp:跨服连接不存在，type:%s", centertypes.GameServerTypeAll.String())
	}

	//TODO 修改可能连接变化了
	s.arenapvpClient = arenapvpclient.NewArenapvpClient(conn)
	return
}

func (s *arenapvpService) Heartbeat() {
	s.hbRunner.Heartbeat()
}

var (
	once sync.Once
	cs   *arenapvpService
)

func Init() (err error) {
	once.Do(func() {
		cs = &arenapvpService{}
		err = cs.init()
	})
	return err
}

func GetArenapvpService() ArenapvpService {
	return cs
}
