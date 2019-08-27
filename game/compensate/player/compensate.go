package player

import (
	compensatetypes "fgame/fgame/game/compensate/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"fgame/fgame/game/compensate/dao"
)

//玩家补偿管理器
type PlayerCompensateDataManager struct {
	p                         player.Player
	playerCompensateObjectMap map[int64]*PlayerCompensateObject
}

func (m *PlayerCompensateDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerCompensateDataManager) Load() (err error) {
	//加载玩家补偿信息
	compensateList, err := dao.GetCompensateDao().GetPlayerCompensateEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range compensateList {
		obj := NewPlayerCompensateObject(m.p)
		obj.FromEntity(entity)

		m.playerCompensateObjectMap[obj.compensateId] = obj
	}
	return
}

//加载后
func (m *PlayerCompensateDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerCompensateDataManager) Heartbeat() {

}

// 是否领取补偿
func (m *PlayerCompensateDataManager) IsHadCompensate(compensateId int64) bool {
	obj := m.getCompensateObj(compensateId)
	if obj == nil {
		return false
	}

	return true
}

// 添加补偿记录
func (m *PlayerCompensateDataManager) AddCompensate(compensateId int64, state compensatetypes.CompensateRecordSate) {
	obj := m.getCompensateObj(compensateId)
	if obj != nil {
		return
	}

	obj = NewPlayerCompensateObject(m.p)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.compensateId = compensateId
	obj.state = state
	obj.createTime = now
	obj.SetModified()

	m.playerCompensateObjectMap[compensateId] = obj
}

func (m *PlayerCompensateDataManager) getCompensateObj(compensateId int64) *PlayerCompensateObject {
	obj, ok := m.playerCompensateObjectMap[compensateId]
	if !ok {
		return nil
	}

	return obj
}

func CreatePlayerCompensateDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerCompensateDataManager{}
	m.p = p
	m.playerCompensateObjectMap = make(map[int64]*PlayerCompensateObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerCompensateDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerCompensateDataManager))
}
