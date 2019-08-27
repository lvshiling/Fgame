package transpotation

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/transportation/dao"
	transportationeventtypes "fgame/fgame/game/transportation/event/types"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	transportationtypes "fgame/fgame/game/transportation/types"
	"sync"
)

type TransportService interface {
	Start()
	//生成个人镖车
	AddPersonalTransportation(playerId int64, plName string, allianceId int64, typ transportationtypes.TransportationType) (biaoche *biaochenpc.BiaocheNPC, err error)
	//生成仙盟镖车
	AddAllianceTransportation(playerId, allianceId int64, owerName string) (biaoche *biaochenpc.BiaocheNPC, err error)
	//删除镖车
	RemoveTransportation(playerId int64) error
	//押镖失败
	TransportFail(playerId int64, robName string) error
	//押镖成功
	TransportFinish(playerId int64) error
	//获取镖车信息
	GetTransportation(playerId int64) *biaochenpc.BiaocheNPC
	//穿云箭
	DistressSignal(playerId int64) error
}

type transportService struct {
	rwm sync.RWMutex
	//玩家镖车信息
	transportationMap map[int64]*biaochenpc.TransportationObject
	//玩家镖车
	transportationNPCMap map[int64]scene.NPC
}

func (s *transportService) init() (err error) {
	s.transportationMap = make(map[int64]*biaochenpc.TransportationObject)
	s.transportationNPCMap = make(map[int64]scene.NPC)

	//加载所有镖车
	err = s.initTransport()
	if err != nil {
		return
	}
	return
}

func (s *transportService) initTransport() (err error) {
	entityList, err := dao.GetTransportationDao().GetTransportList(global.GetGame().GetServerIndex())
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := biaochenpc.CreateTransportationObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}

		s.transportationMap[obj.GetPlayerId()] = obj
	}

	return
}

func (s *transportService) AddPersonalTransportation(playerId int64, plName string, playerAllianceId int64, typ transportationtypes.TransportationType) (biaoChen *biaochenpc.BiaocheNPC, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if typ == transportationtypes.TransportationTypeAlliance {
		return
	}

	if s.isOnTransporting(playerId) {
		err = errorTransportationOnDoing
		return
	}

	biaoChe := biaochenpc.CreateBiaoCheNPC(playerId, 0, playerAllianceId, plName, typ)
	s.transportationMap[playerId] = biaoChe.GetTransportationObject()
	s.transportationNPCMap[playerId] = biaoChe
	return biaoChe, nil
}

func (s *transportService) AddAllianceTransportation(playerId, allianceId int64, owerName string) (biaoChen *biaochenpc.BiaocheNPC, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.isOnTransporting(playerId) {
		err = errorTransportationOnDoing
		return
	}

	biaoChe := biaochenpc.CreateBiaoCheNPC(playerId, allianceId, allianceId, owerName, transportationtypes.TransportationTypeAlliance)
	s.transportationMap[playerId] = biaoChe.GetTransportationObject()
	s.transportationNPCMap[playerId] = biaoChe
	return biaoChe, nil
}

func (s *transportService) RemoveTransportation(playerId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	n := s.getTransportationNpc(playerId)
	if n == nil {
		return
	}
	n.Remove()
	delete(s.transportationMap, playerId)
	delete(s.transportationNPCMap, playerId)

	return
}

func (s *transportService) isOnTransporting(playerId int64) bool {
	if _, ok := s.transportationMap[playerId]; ok {
		return true
	}
	return false
}

func (s *transportService) TransportFinish(playerId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	n := s.getTransportationNpc(playerId)
	if n == nil {
		return
	}
	n.Finish()

	return
}

func (s *transportService) TransportFail(playerId int64, robName string) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	n := s.getTransportationNpc(playerId)
	if n == nil {
		return
	}
	n.Fail(robName)

	return
}

func (s *transportService) DistressSignal(playerId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	biaoChenNpc := s.getTransportationNpc(playerId)
	if biaoChenNpc == nil {
		err = errorTransportationNotExsit
		return
	}
	transportObj := biaoChenNpc.GetTransportationObject()
	if transportObj.GetAllianceId() == 0 {
		err = errorTransportationNotAllianceTransportation
		return
	}

	//穿云箭CD
	if transportObj.IsDistressCD() {
		err = errorTransportationDistressCD
		return
	}

	biaoChenNpc.DistressSignal()

	return
}

func (s *transportService) GetTransportation(playerId int64) *biaochenpc.BiaocheNPC {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	biaoChenNPC := s.getTransportationNpc(playerId)

	return biaoChenNPC
}

func (s *transportService) getTransportationNpc(playerId int64) *biaochenpc.BiaocheNPC {
	n, ok := s.transportationNPCMap[playerId]
	if !ok {
		return nil
	}
	tn := n.(*biaochenpc.BiaocheNPC)
	return tn
}

func (s *transportService) Start() {
	for playerId, transportObj := range s.transportationMap {
		n := biaochenpc.CreateBiaoCheNPCWithObj(transportObj, transportObj.GetAllianceId())
		s.transportationNPCMap[playerId] = n
		//发出事件
		gameevent.Emit(transportationeventtypes.EventTypeTransportationInit, n, nil)
	}
	return
}

var (
	once sync.Once
	as   *transportService
)

func Init() (err error) {
	once.Do(func() {
		as = &transportService{}
		err = as.init()
	})
	return err
}

func GetTransportService() TransportService {
	return as
}
