package chuangshi

import (
	"fgame/fgame/core/runner"
	"fgame/fgame/game/chuangshi/dao"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sync"
)

type ChuangShiService interface {
	runner.Task
	Star()

	// 添加报名创世之战的玩家数量
	AddPlayerNum()
	// 获取报名创世之战玩家数量(包括虚假)
	GetBaoMingChuangShiPlayerNum() int64
	// 检查是否在时间段内
	CheckYuGaoTime() (bool, error)

	// // 神王报名列表
	// GetShenWangSignUpList(playerId int64) []*chuangshidata.MemberInfo
	// // 获取神王报名
	// GetShenWangSignUp(playerId int64) *ChuangShiShenWangSignUpObject
	// // 神王报名
	// ShenWangSignUp(playerId int64)
	// // 移除神王报名
	// ShenWangSignUpRemove(playerId int64)

	// // 神王投票列表
	// GetShenWangVoteList(playerId int64) []*chuangshidata.VoteInfo
	// // 获取神王投票
	// GetShenWangVote(playerId int64) *ChuangShiShenWangVoteObject
	// // 神王投票
	// ShenWangVote(playerId, supportId int64)
	// // 移除神王投票
	// ShenWangVoteRemove(playerId int64)

	// // 获取城防建设
	// GetChengFangJianShe(playerId int64) *ChuangShiChengFangJianSheObject
	// // 城防建设
	// ChengFangJianShe(playerId int64, cityId int64, buildType chuangshitypes.ChuangShiCityJianSheType, num int32)
	// // 移除城防建设
	// ChengFangJianSheRemove(playerId int64)

	// // 阵营信息列表
	// GetCampList() []*chuangshidata.CampData
	// GetCamp(playerId int64) *chuangshidata.CampData
	// GetCity(cityId int64) *chuangshidata.CityInfo
	// GetMember(platerId int64) *chuangshidata.MemberInfo

	// //城主任命
	// CityRenMing(playerId int64, cityId, beCommitId int64) (success bool, err error)
	// //城池工资分配
	// CityPaySchedule(playerId int64, paramList []*chuangshidata.CityPayScheduleParam) (err error)
	// //阵营工资分配
	// CampPaySchedule(playerId int64, paramList []*chuangshidata.CamPayScheduleParam) (err error)
	// //阵营工资领取
	// CampPayReceive(playerId int64) (err error)
	// //设置攻城目标
	// GongChengTargetFuShu(playerId, cityId int64) bool
	// //加入阵营
	// JoinCamp(campType chuangshitypes.ChuangShiCampType, memList ...*chuangshidata.MemberInfo) bool
	// //城池天气设置
	// CityTianQiSet(playerId, cityId int64, level int32) bool
}

type chuangShiService struct {
	rwm sync.RWMutex
	// chuangshiClient chuangshiclient.ChuangshiClient
	// hbRunner        heartbeat.HeartbeatTaskRunner

	// 创世之战预告对象
	chuangShiYuGaoObj *ChuangShiYuGaoObject

	// //神王报名记录
	// shenWangSignUpMap map[int64]*ChuangShiShenWangSignUpObject
	// //神王投票记录
	// shenWangVoteMap map[int64]*ChuangShiShenWangVoteObject
	// //城池建设记录
	// chengFangJianSheMap map[int64]*ChuangShiChengFangJianSheObject

	// -------------跨服数据--------------
	//阵营列表
	// campList      []*chuangshidata.CampData
	// campMemberMap map[int64]*chuangshidata.MemberInfo
}

func (s *chuangShiService) init() (err error) {
	// s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	// s.hbRunner.AddTask(heartbeat.HeartbeatTask(CreateArenapvpPlayerDataTask(s)))
	// s.hbRunner.AddTask(heartbeat.HeartbeatTask(createShenWangSignUpTask(s)))
	// s.hbRunner.AddTask(heartbeat.HeartbeatTask(createShenWangVoteTask(s)))
	// s.hbRunner.AddTask(heartbeat.HeartbeatTask(createChengFangJianSheTask(s)))

	// err = s.resetClient()
	// if err != nil {
	// 	return
	// }

	// 创世预告
	err = s.loadChuangShiYuGao()
	if err != nil {
		return
	}

	// //神王报名记录
	// err = s.loadChuangShiShenWangSignUp()
	// if err != nil {
	// 	return
	// }

	// //神王投票记录
	// err = s.loadChuangShiShenWangVote()
	// if err != nil {
	// 	return
	// }

	// //城池建设记录
	// err = s.loadChengFangJianShe()
	// if err != nil {
	// 	return
	// }

	return
}

// func (s *chuangShiService) resetClient() (err error) {
// 	conn := center.GetCenterService().GetCross(centertypes.GameServerTypeAll)
// 	if conn == nil {
// 		return fmt.Errorf("chuangshi:跨服连接不存在")
// 	}

// 	//TODO 修改可能连接变化了
// 	s.chuangshiClient = chuangshiclient.NewChuangshiClient(conn)
// 	return
// }

func (s *chuangShiService) loadChuangShiYuGao() (err error) {
	entity, err := dao.GetChuangShiDao().GetChuangShiYuGaoEntity()
	if err != nil {
		return
	}

	obj := NewChuangShiYuGaoObject()
	if entity == nil {
		now := global.GetGame().GetTimeService().Now()
		obj.createTime = now
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.SetModified()
	} else {
		obj.FromEntity(entity)
	}
	s.chuangShiYuGaoObj = obj
	return
}

// func (s *chuangShiService) loadChuangShiShenWangSignUp() (err error) {
// 	s.shenWangSignUpMap = make(map[int64]*ChuangShiShenWangSignUpObject)
// 	serverId := global.GetGame().GetServerIndex()
// 	entityList, err := dao.GetChuangShiDao().GetChuangShiShenWangSignUpEntityList(serverId)
// 	if err != nil {
// 		return
// 	}

// 	for _, entity := range entityList {
// 		obj := NewChuangShiShenWangSignUpObject()
// 		obj.FromEntity(entity)
// 		s.shenWangSignUpMap[obj.playerId] = obj
// 	}
// 	return
// }

// func (s *chuangShiService) loadChuangShiShenWangVote() (err error) {
// 	s.shenWangVoteMap = make(map[int64]*ChuangShiShenWangVoteObject)
// 	serverId := global.GetGame().GetServerIndex()
// 	entityList, err := dao.GetChuangShiDao().GetChuangShiShenWangVoteEntityList(serverId)
// 	if err != nil {
// 		return
// 	}

// 	for _, entity := range entityList {
// 		obj := NewChuangShiShenWangVoteObject()
// 		obj.FromEntity(entity)
// 		s.shenWangVoteMap[obj.playerId] = obj
// 	}
// 	return
// }

// func (s *chuangShiService) loadChengFangJianShe() (err error) {
// 	s.chengFangJianSheMap = make(map[int64]*ChuangShiChengFangJianSheObject)
// 	serverId := global.GetGame().GetServerIndex()
// 	entityList, err := dao.GetChuangShiDao().GetChuangShiChengFangJianSheEntityList(serverId)
// 	if err != nil {
// 		return
// 	}

// 	for _, entity := range entityList {
// 		obj := NewChuangShiChengFangJianSheObject()
// 		obj.FromEntity(entity)
// 		s.chengFangJianSheMap[obj.playerId] = obj
// 	}
// 	return
// }

func (s *chuangShiService) Heartbeat() {
	// s.rwm.Lock()
	// defer s.rwm.Unlock()

	// s.hbRunner.Heartbeat()
}

func (s *chuangShiService) Star() {
	// s.syncRemoteChuangshiData()
}

// //定时同步排行榜列表
// func (s *chuangShiService) syncRemoteChuangshiData() (err error) {
// 	if s.chuangshiClient == nil {
// 		err = s.resetClient()
// 		if err != nil {
// 			return
// 		}
// 	}
// 	//TODO 超时
// 	ctx := context.TODO()
// 	resp, err := s.chuangshiClient.GetChuangShiWholeInfo(ctx)
// 	if err != nil {
// 		return
// 	}

// 	newCampList, newMemMap := chuangshidata.ConvertToCampList(resp.CampList)
// 	s.checkShenWangChangedNotice(newCampList)
// 	s.campList = newCampList
// 	s.campMemberMap = newMemMap

// 	return nil
// }

// func (s *chuangShiService) checkShenWangChangedNotice(newCampList []*chuangshidata.CampData) {
// 	for _, newCamp := range newCampList {
// 		if newCamp.KingMem == nil {
// 			continue
// 		}
// 		curCamp := s.getCamp(newCamp.CampType)
// 		if curCamp.KingMem == nil {
// 			gameevent.Emit(chuangshieventtypes.EventTypeChuangShiShenWangChanged, newCamp.KingMem, nil)
// 			continue
// 		}

// 		if curCamp.KingMem.PlayerId == newCamp.KingMem.PlayerId {
// 			continue
// 		}

// 		gameevent.Emit(chuangshieventtypes.EventTypeChuangShiShenWangChanged, newCamp, nil)
// 	}

// 	return
// }

// func (s *chuangShiService) getCampMember(playerId int64) *chuangshidata.MemberInfo {
// 	mem, ok := s.campMemberMap[playerId]
// 	if !ok {
// 		return nil
// 	}

// 	return mem
// }

// func (s *chuangShiService) getCamp(campType chuangshitypes.ChuangShiCampType) *chuangshidata.CampData {
// 	for _, camp := range s.campList {
// 		if camp.CampType != campType {
// 			continue
// 		}

// 		return camp
// 	}

// 	return nil
// }

func (s *chuangShiService) AddPlayerNum() {
	flag, err := s.CheckYuGaoTime()
	if err != nil {
		return
	}
	if !flag {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	s.chuangShiYuGaoObj.num++
	s.chuangShiYuGaoObj.updateTime = now
	s.chuangShiYuGaoObj.SetModified()
}

func (s *chuangShiService) GetBaoMingChuangShiPlayerNum() (num int64) {
	now := global.GetGame().GetTimeService().Now()
	num = s.chuangShiYuGaoObj.num
	yuGaoTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiYuGaoTemplate()
	jiaRenTime := yuGaoTemp.JiaRenTime
	jiaRenNum := yuGaoTemp.JiaRenCount
	startTime, err := timeutils.ParseYYYYMMDD(yuGaoTemp.StartTime)
	if err != nil {
		return
	}
	if now < startTime {
		return
	}

	endTime := startTime + yuGaoTemp.GetJianGeTime()
	if now > endTime {
		now = endTime
	}
	jianGeTime := now - startTime
	xuJiaNum := (jianGeTime / jiaRenTime) * int64(jiaRenNum)
	num += xuJiaNum
	return
}

// 检查是否在时间段内
func (s *chuangShiService) CheckYuGaoTime() (bool, error) {
	now := global.GetGame().GetTimeService().Now()
	yuGaoTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiYuGaoTemplate()
	startTime, err := timeutils.ParseYYYYMMDD(yuGaoTemp.StartTime)
	if err != nil {
		return false, err
	}
	endTime := startTime + yuGaoTemp.GetJianGeTime()

	if now < startTime {
		return false, nil
	}
	if now > endTime {
		return false, nil
	}

	return true, nil
}

// //城池建设
// func (s *chuangShiService) ChengFangJianShe(playerId int64, cityId int64, buildType chuangshitypes.ChuangShiCityJianSheType, num int32) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	jianSheObj := s.getChengFangJianSheObj(playerId)
// 	if jianSheObj != nil {
// 		return
// 	}

// 	id, _ := idutil.GetId()
// 	now := global.GetGame().GetTimeService().Now()

// 	obj := NewChuangShiChengFangJianSheObject()
// 	obj.id = id
// 	obj.serverId = global.GetGame().GetServerIndex()
// 	obj.playerId = playerId
// 	obj.num = num
// 	obj.jianSheType = buildType
// 	obj.cityId = cityId
// 	obj.status = chuangshitypes.ChengFangStatusTypeProgressing
// 	obj.createTime = now
// 	obj.SetModified()
// 	s.chengFangJianSheMap[playerId] = obj

// 	//跨服建设
// 	s.syncChengFangJianShe(obj)
// }

// func (s *chuangShiService) syncChengFangJianShe(obj *ChuangShiChengFangJianSheObject) {
// 	go func(obj *ChuangShiChengFangJianSheObject) {
// 		platform := global.GetGame().GetPlatform()

// 		log.WithFields(
// 			log.Fields{
// 				"platform": platform,
// 				"serverId": obj.serverId,
// 				"playerId": obj.playerId,
// 			}).Infoln("chuangshi:城池建设rpc请求")

// 		defer func() {
// 			if r := recover(); r != nil {
// 				debug.PrintStack()
// 				exceptionContent := string(debug.Stack())
// 				log.WithFields(
// 					log.Fields{
// 						"platform": platform,
// 						"serverId": obj.serverId,
// 						"playerId": obj.playerId,
// 						"error":    r,
// 						"stack":    exceptionContent,
// 					}).Error("chuangshi:城池建设,错误")
// 				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
// 			}
// 		}()

// 		ctx := context.TODO()
// 		resp, err := s.chuangshiClient.CityJianShe(ctx, obj.playerId, obj.cityId, int32(obj.jianSheType), obj.num)
// 		if err != nil {
// 			log.WithFields(
// 				log.Fields{
// 					"platform": platform,
// 					"serverId": obj.serverId,
// 					"playerId": obj.playerId,
// 					"err":      err.Error(),
// 				}).Error("chuangshi:城池建设失败，rpc错误")
// 			return
// 		}

// 		if !resp.Success {
// 			log.WithFields(
// 				log.Fields{
// 					"platform": platform,
// 					"serverId": obj.serverId,
// 					"playerId": obj.playerId,
// 					"err":      err.Error(),
// 				}).Infoln("chuangshi:城池建设失败，未成功")
// 			return
// 		}

// 		//建设结果
// 		s.chengFangJianSheFinish(obj)
// 	}(obj)
// }

// func (s *chuangShiService) chengFangJianSheFinish(obj *ChuangShiChengFangJianSheObject) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	if !obj.IfProgressing() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	obj.status = chuangshitypes.ChengFangStatusTypeSuccess
// 	obj.updateTime = now
// 	obj.SetModified()

// 	// camp := s.getCamp(campType)
// 	// camp.VoteList = newVoteList
// 	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiChengFangJianShe, obj, nil)
// }

// func (s *chuangShiService) getChengFangJianSheObj(playerId int64) *ChuangShiChengFangJianSheObject {
// 	obj, ok := s.chengFangJianSheMap[playerId]
// 	if !ok {
// 		return nil
// 	}

// 	return obj
// }

// //城防建设
// func (s *chuangShiService) GetChengFangJianShe(playerId int64) *ChuangShiChengFangJianSheObject {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	return s.getChengFangJianSheObj(playerId)
// }

// //城防建设移除
// func (s *chuangShiService) ChengFangJianSheRemove(playerId int64) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	jianSheObj := s.getChengFangJianSheObj(playerId)
// 	if jianSheObj == nil {
// 		return
// 	}

// 	if jianSheObj.IfProgressing() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	jianSheObj.deleteTime = now
// 	jianSheObj.updateTime = now
// 	delete(s.chengFangJianSheMap, playerId)
// }

// //神王投票
// func (s *chuangShiService) ShenWangVote(playerId, supportId int64) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	voteObj := s.getShenWangVoteObj(playerId)
// 	if voteObj != nil {
// 		return
// 	}

// 	id, _ := idutil.GetId()
// 	now := global.GetGame().GetTimeService().Now()

// 	obj := NewChuangShiShenWangVoteObject()
// 	obj.id = id
// 	obj.serverId = global.GetGame().GetServerIndex()
// 	obj.playerId = playerId
// 	obj.supportId = supportId
// 	obj.status = chuangshitypes.ShenWangVoteTypeVoting
// 	obj.createTime = now
// 	obj.SetModified()
// 	s.shenWangVoteMap[playerId] = obj

// 	//跨服投票
// 	s.syncShenWangVote(obj)
// }

// func (s *chuangShiService) syncShenWangVote(obj *ChuangShiShenWangVoteObject) {
// 	go func(obj *ChuangShiShenWangVoteObject) {
// 		platform := global.GetGame().GetPlatform()

// 		log.WithFields(
// 			log.Fields{
// 				"platform": platform,
// 				"serverId": obj.serverId,
// 				"playerId": obj.playerId,
// 			}).Infoln("chuangshi:神王竞选投票rpc请求")

// 		defer func() {
// 			if r := recover(); r != nil {
// 				debug.PrintStack()
// 				exceptionContent := string(debug.Stack())
// 				log.WithFields(
// 					log.Fields{
// 						"platform": platform,
// 						"serverId": obj.serverId,
// 						"playerId": obj.playerId,
// 						"error":    r,
// 						"stack":    exceptionContent,
// 					}).Error("chuangshi:神王竞选投票,错误")
// 				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
// 			}
// 		}()

// 		ctx := context.TODO()
// 		resp, err := s.chuangshiClient.ShenWangVote(ctx, obj.supportId)
// 		if err != nil {
// 			log.WithFields(
// 				log.Fields{
// 					"platform": platform,
// 					"serverId": obj.serverId,
// 					"playerId": obj.playerId,
// 					"err":      err.Error(),
// 				}).Error("chuangshi:神王投票错误")
// 			return
// 		}

// 		if !resp.Success {
// 			log.WithFields(
// 				log.Fields{
// 					"platform": platform,
// 					"serverId": obj.serverId,
// 					"playerId": obj.playerId,
// 					"err":      err.Error(),
// 				}).Infoln("chuangshi:神王投票失败")
// 			return
// 		}

// 		//投票结果
// 		s.shenWangVoteFinish(obj, chuangshitypes.ChuangShiCampType(resp.CampType), chuangshidata.ConvertToVoteList(resp.VoteList))
// 	}(obj)
// }

// func (s *chuangShiService) shenWangVoteFinish(obj *ChuangShiShenWangVoteObject, campType chuangshitypes.ChuangShiCampType, newVoteList []*chuangshidata.VoteInfo) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	if !obj.IfVoting() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	obj.status = chuangshitypes.ShenWangVoteTypeSuccess
// 	obj.updateTime = now
// 	obj.SetModified()

// 	camp := s.getCamp(campType)
// 	camp.VoteList = newVoteList
// 	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiShenWangVote, obj, newVoteList)
// }

// func (s *chuangShiService) getShenWangVoteObj(playerId int64) *ChuangShiShenWangVoteObject {
// 	obj, ok := s.shenWangVoteMap[playerId]
// 	if !ok {
// 		return nil
// 	}

// 	return obj
// }

// //神王投票
// func (s *chuangShiService) GetShenWangVote(playerId int64) *ChuangShiShenWangVoteObject {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	return s.getShenWangVoteObj(playerId)
// }

// //神王投票列表
// func (s *chuangShiService) GetShenWangVoteList(playerId int64) []*chuangshidata.VoteInfo {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	campMem := s.getCampMember(playerId)
// 	if campMem == nil {
// 		return nil
// 	}

// 	camp := s.getCamp(campMem.CampType)
// 	if camp == nil {
// 		return nil
// 	}

// 	return camp.VoteList
// }

// //神王投票移除
// func (s *chuangShiService) ShenWangVoteRemove(playerId int64) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	voteObj := s.getShenWangVoteObj(playerId)
// 	if voteObj == nil {
// 		return
// 	}

// 	if voteObj.IfVoting() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	voteObj.deleteTime = now
// 	voteObj.updateTime = now
// 	delete(s.shenWangVoteMap, playerId)
// }

// //神王报名
// func (s *chuangShiService) ShenWangSignUp(playerId int64) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	signUpObj := s.getShenWangSignUpObj(playerId)
// 	if signUpObj != nil {
// 		return
// 	}

// 	id, _ := idutil.GetId()
// 	now := global.GetGame().GetTimeService().Now()

// 	obj := NewChuangShiShenWangSignUpObject()
// 	obj.id = id
// 	obj.serverId = global.GetGame().GetServerIndex()
// 	obj.playerId = playerId
// 	obj.status = chuangshitypes.ShenWangSignUpTypeSigning
// 	obj.createTime = now
// 	obj.SetModified()
// 	s.shenWangSignUpMap[playerId] = obj

// 	//跨服报名
// 	s.syncShenWangSignup(obj)
// }

// //获取神王报名
// func (s *chuangShiService) GetShenWangSignUp(playerId int64) *ChuangShiShenWangSignUpObject {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	return s.getShenWangSignUpObj(playerId)
// }

// //神王报名移除
// func (s *chuangShiService) ShenWangSignUpRemove(playerId int64) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	signUpObj := s.getShenWangSignUpObj(playerId)
// 	if signUpObj == nil {
// 		return
// 	}

// 	if signUpObj.status == chuangshitypes.ShenWangSignUpTypeSigning {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	signUpObj.deleteTime = now
// 	signUpObj.updateTime = now
// 	delete(s.shenWangSignUpMap, playerId)
// }

// //神王报名列表
// func (s *chuangShiService) GetShenWangSignUpList(playerId int64) []*chuangshidata.MemberInfo {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	campMem := s.getCampMember(playerId)
// 	if campMem == nil {
// 		return nil
// 	}

// 	camp := s.getCamp(campMem.CampType)
// 	if camp == nil {
// 		return nil
// 	}

// 	return camp.SignUpList
// }

// func (s *chuangShiService) syncShenWangSignup(obj *ChuangShiShenWangSignUpObject) {
// 	go func(obj *ChuangShiShenWangSignUpObject) {
// 		platform := global.GetGame().GetPlatform()

// 		log.WithFields(
// 			log.Fields{
// 				"platform": platform,
// 				"serverId": obj.serverId,
// 				"playerId": obj.playerId,
// 			}).Infoln("chuangshi:神王竞选报名rpc请求")

// 		defer func() {
// 			if r := recover(); r != nil {
// 				debug.PrintStack()
// 				exceptionContent := string(debug.Stack())
// 				log.WithFields(
// 					log.Fields{
// 						"platform": platform,
// 						"serverId": obj.serverId,
// 						"playerId": obj.playerId,
// 						"error":    r,
// 						"stack":    exceptionContent,
// 					}).Error("chuangshi:神王竞选报名,错误")
// 				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
// 			}
// 		}()

// 		ctx := context.TODO()
// 		resp, err := s.chuangshiClient.ShenWangSignUp(ctx, platform, obj.serverId, obj.playerId)
// 		if err != nil {
// 			log.WithFields(
// 				log.Fields{
// 					"platform": platform,
// 					"serverId": obj.serverId,
// 					"playerId": obj.playerId,
// 					"err":      err.Error(),
// 				}).Error("chuangshi:神王报名错误")
// 			return
// 		}

// 		if !resp.Success {
// 			log.WithFields(
// 				log.Fields{
// 					"platform": platform,
// 					"serverId": obj.serverId,
// 					"playerId": obj.playerId,
// 					"err":      err.Error(),
// 				}).Infoln("chuangshi:神王报名失败")
// 			return
// 		}

// 		//报名结果
// 		s.shenWangSignUpFinish(obj, chuangshitypes.ChuangShiCampType(resp.CampType), chuangshidata.ConvertToMemberList(resp.SignList))
// 	}(obj)
// }

// func (s *chuangShiService) shenWangSignUpFinish(obj *ChuangShiShenWangSignUpObject, campType chuangshitypes.ChuangShiCampType, newSignList []*chuangshidata.MemberInfo) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	if !obj.IfSigning() {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	obj.status = chuangshitypes.ShenWangSignUpTypeSuccess
// 	obj.updateTime = now
// 	obj.SetModified()

// 	camp := s.getCamp(campType)
// 	camp.SignUpList = newSignList
// 	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiShenWangSignUp, obj, newSignList)
// }

// func (s *chuangShiService) getShenWangSignUpObj(playerId int64) *ChuangShiShenWangSignUpObject {
// 	obj, ok := s.shenWangSignUpMap[playerId]
// 	if !ok {
// 		return nil
// 	}

// 	return obj
// }

// func (s *chuangShiService) GetCampList() []*chuangshidata.CampData {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	return s.campList
// }

// func (s *chuangShiService) GetCamp(playerId int64) *chuangshidata.CampData {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	mem := s.getCampMember(playerId)
// 	if mem == nil {
// 		return nil
// 	}

// 	return s.getCamp(mem.CampType)
// }
// func (s *chuangShiService) GetMember(playerId int64) *chuangshidata.MemberInfo {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	return s.getCampMember(playerId)
// }

// func (s *chuangShiService) GetCity(cityId int64) *chuangshidata.CityInfo {
// 	s.rwm.RLock()
// 	defer s.rwm.RUnlock()

// 	for _, camp := range s.campList {
// 		for _, city := range camp.CityList {
// 			if city.CityId != cityId {
// 				continue
// 			}

// 			return city
// 		}
// 	}

// 	return nil
// }

// func (s *chuangShiService) CityRenMing(playerId int64, cityId, beCommitId int64) (success bool, err error) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	ctx := context.TODO()
// 	resp, err := s.chuangshiClient.CityRenMing(ctx, playerId, cityId, beCommitId)
// 	if err != nil {
// 		return
// 	}

// 	//TODO xzk27 beCommitId被任命推送  神王名字，城池名字
// 	//TODO xzk27 原城主卸任推送  神王名字，城池名字

// 	success = resp.Success
// 	return
// }

// func (s *chuangShiService) CityPaySchedule(playerId int64, paramList []*chuangshidata.CityPayScheduleParam) (err error) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	mem := s.getCampMember(playerId)
// 	if mem == nil {
// 		return
// 	}

// 	camp := s.getCamp(mem.CampType)
// 	city := camp.GetCityByChengZhuId(playerId)
// 	if city == nil {
// 		return
// 	}

// 	ctx := context.TODO()
// 	_, err = s.chuangshiClient.CityPaySchedule(ctx, playerId, paramList)
// 	if err != nil {
// 		return
// 	}

// 	//TODO xzk27 工资分配后的处理
// 	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiCityPaySchedule, nil, nil)
// 	return
// }

// func (s *chuangShiService) CampPaySchedule(playerId int64, paramList []*chuangshidata.CamPayScheduleParam) (err error) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	ctx := context.TODO()
// 	resp, err := s.chuangshiClient.CampPaySchedule(ctx, playerId, paramList)
// 	if err != nil {
// 		return
// 	}

// 	//TODO xzk27 工资分配后的处理
// 	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiCampPaySchedule, resp.Camp, nil)
// 	return
// }

// func (s *chuangShiService) CampPayReceive(playerId int64) (err error) {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	ctx := context.TODO()
// 	_, err = s.chuangshiClient.CampPayReceive(ctx, playerId)
// 	if err != nil {
// 		return
// 	}

// 	//TODO xzk27 领取阵营工资后的处理
// 	return
// }

// func (s *chuangShiService) GongChengTargetFuShu(playerId, cityId int64) bool {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	ctx := context.TODO()
// 	resp, err := s.chuangshiClient.GongChengTargetFuShu(ctx, playerId, cityId)
// 	if err != nil {
// 		return false
// 	}

// 	return resp.Success

// 	//TODO xzk27 领取阵营工资后的处理
// }

// func (s *chuangShiService) JoinCamp(campType chuangshitypes.ChuangShiCampType, memList ...*chuangshidata.MemberInfo) bool {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	// TODO xzk27 老玩家数据处理
// 	// 2、更新时，按照盘古、女娲、伏羲的阵营的顺序进行随机，先将盘古阵营随机设置到一个仙盟身上，其次将女娲阵营随机设置到一个仙盟身上，依次类推，直至本服内所有仙盟阵营设置完毕
// 	// 3、将仙盟设置阵营后，则本仙盟内所有玩家的阵营均将会变成与该仙盟阵营一致，同时发送一封补偿邮件，内含阵营更换卡提供玩家更换阵营使用

// 	ctx := context.TODO()
// 	resp, err := s.chuangshiClient.JoinCamp(ctx, int32(campType), memList...)
// 	if err != nil {
// 		return false
// 	}

// 	return resp.Success

// 	//TODO xzk27 领取阵营工资后的处理
// }

// func (s *chuangShiService) CityTianQiSet(playerId, cityId int64, level int32) bool {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()

// 	ctx := context.TODO()
// 	resp, err := s.chuangshiClient.CityTianQiSet(ctx, playerId, cityId, level)
// 	if err != nil {
// 		return false
// 	}

// 	return resp.Success
// }

var (
	once sync.Once
	cs   *chuangShiService
)

func Init() (err error) {
	once.Do(func() {
		cs = &chuangShiService{}
		err = cs.init()
	})
	return err
}

func GetChuangShiService() ChuangShiService {
	return cs
}
