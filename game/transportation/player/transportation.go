package player

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/transportation/dao"
	transportationeventtypes "fgame/fgame/game/transportation/event/types"
	transportationtypes "fgame/fgame/game/transportation/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家镖车管理器
type PlayerTransportationDataManager struct {
	p player.Player
	//玩家镖车信息
	transportationObject *PlayerTransportationObject
}

func (m *PlayerTransportationDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerTransportationDataManager) Load() (err error) {
	transpotEntity, err := dao.GetTransportationDao().GetPlayerTransportInfo(m.p.GetId())
	if err != nil {
		return
	}

	if transpotEntity != nil {
		obj := newPlayerTransportationObject(m.p)
		obj.FromEntity(transpotEntity)
		m.transportationObject = obj
	}
	return
}

//初始化镖车信息
func (m *PlayerTransportationDataManager) iniTransportationObj() {
	obj := newPlayerTransportationObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()

	obj.id = id
	obj.robList = make([]int64, 0)
	obj.personalTransportTimes = 0
	obj.createTime = now
	obj.SetModified()

	m.transportationObject = obj
}

//加载后
func (m *PlayerTransportationDataManager) AfterLoad() (err error) {
	err = m.refreshTimes()
	return
}

func (m *PlayerTransportationDataManager) refreshTimes() (err error) {
	transportationObject := m.getTransportObj()
	now := global.GetGame().GetTimeService().Now()
	lastUpdateTime := transportationObject.updateTime

	flag, err := timeutils.IsSameFive(lastUpdateTime, now)
	if err != nil {
		return err
	}
	if !flag {
		transportationObject.personalTransportTimes = 0
		transportationObject.robList = make([]int64, 0)
		transportationObject.updateTime = now
		transportationObject.SetModified()
	}
	return
}

//心跳
func (m *PlayerTransportationDataManager) Heartbeat() {
}

func createPlayerTransportationDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerTransportationDataManager{}
	m.p = p
	return m
}

func (m *PlayerTransportationDataManager) getTransportObj() *PlayerTransportationObject {
	if m.transportationObject == nil {
		m.iniTransportationObj()
	}

	return m.transportationObject
}

//是否足够押镖次数
func (m *PlayerTransportationDataManager) IsEnoughTransportTimes() bool {
	joinTimes := m.GetTranspotTimes()
	maxTimes := constant.GetConstantService().GetConstant(constanttypes.ConstantTypePersonalTransportationTimes)
	return joinTimes < maxTimes
}

//消耗镖车次数
func (m *PlayerTransportationDataManager) AcceptTransportation(typ transportationtypes.TransportationType) {
	transportationObject := m.getTransportObj()
	now := global.GetGame().GetTimeService().Now()
	transportationObject.updateTime = now
	transportationObject.personalTransportTimes += 1
	transportationObject.SetModified()

	gameevent.Emit(transportationeventtypes.EventTypeTransportationAccept, m.p, typ)
}

//获取玩家镖车次数
func (m *PlayerTransportationDataManager) GetTranspotTimes() int32 {
	m.refreshTimes()
	return m.getTransportObj().personalTransportTimes
}

//获取玩家剩余镖车次数
func (m *PlayerTransportationDataManager) GetLeftTranspotTims() (leftNum int32) {
	joinTimes := m.GetTranspotTimes()
	maxTimes := constant.GetConstantService().GetConstant(constanttypes.ConstantTypePersonalTransportationTimes)
	return maxTimes - joinTimes
}

//GM重置镖车次数
func (m *PlayerTransportationDataManager) GMResetTimes() {
	transportationObject := m.getTransportObj()
	now := global.GetGame().GetTimeService().Now()
	preDay, _ := timeutils.PreDayOfTime(now)
	transportationObject.updateTime = preDay
	transportationObject.SetModified()

	m.refreshTimes()
}

//TODO 修改为攻击过
//劫镖
func (m *PlayerTransportationDataManager) RobTransportation(transportId int64) (flag bool) {
	m.refreshTimes()

	transportationObject := m.getTransportObj()
	maxRobTimes := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRobTransportationTimes)
	curRobTiems := int32(len(transportationObject.robList))
	if curRobTiems >= maxRobTimes {
		return
	}
	//已劫过的镖
	for _, robTransportId := range transportationObject.robList {
		if transportId == robTransportId {
			return
		}
	}

	transportationObject.robList = append(transportationObject.robList, transportId)
	transportationObject.SetModified()

	return true
}

func (m *PlayerTransportationDataManager) GetRobOfTimes() int32 {
	obj := m.getTransportObj()
	return int32(len(obj.robList))
}

func (m *PlayerTransportationDataManager) IfRob(transportId int64) bool {
	obj := m.getTransportObj()
	for _, rob := range obj.robList {
		if rob == transportId {
			return true
		}
	}
	return false
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTransportationType, player.PlayerDataManagerFactoryFunc(createPlayerTransportationDataManager))
}
