package player

import (
	"fgame/fgame/game/cross/dao"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

//玩家跨服数据管理器
type PlayerCrossDataManager struct {
	p                 player.Player
	playerCrossObject *PlayerCrossObject
}

func (m *PlayerCrossDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerCrossDataManager) Load() (err error) {
	//TODO 数据加载封装
	e, err := dao.GetCrossDao().GetCrossEntity(m.p.GetId())
	if err != nil {
		return
	}
	if e == nil {
		m.initPlayerCrossObject()
	} else {
		m.playerCrossObject = NewPlayerCrossObject(m.p)
		m.playerCrossObject.FromEntity(e)
	}
	//刷新

	return nil
}

func (m *PlayerCrossDataManager) EnterCross(crossType crosstypes.CrossType, args ...string) {
	now := global.GetGame().GetTimeService().Now()
	m.playerCrossObject.crossType = crossType
	m.playerCrossObject.crossArgs = args
	m.playerCrossObject.updateTime = now
	m.playerCrossObject.SetModified()
}

func (m *PlayerCrossDataManager) GetCrossType() crosstypes.CrossType {
	return m.playerCrossObject.crossType
}

func (m *PlayerCrossDataManager) GetCrossArgs() []string {
	return m.playerCrossObject.crossArgs
}

func (m *PlayerCrossDataManager) ExitCross() {
	if m.playerCrossObject.crossType == crosstypes.CrossTypeNone {
		return
	}

	m.playerCrossObject.crossType = crosstypes.CrossTypeNone
	m.playerCrossObject.crossArgs = make([]string, 0, 4)
	now := global.GetGame().GetTimeService().Now()
	m.playerCrossObject.updateTime = now
	m.playerCrossObject.SetModified()

}

func (m *PlayerCrossDataManager) AfterLoad() (err error) {

	return nil
}

func (m *PlayerCrossDataManager) Heartbeat() {

}

//第一次初始化
func (m *PlayerCrossDataManager) initPlayerCrossObject() {
	o := NewPlayerCrossObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.player = m.p
	o.crossType = crosstypes.CrossTypeNone
	o.crossArgs = make([]string, 0, 4)
	o.createTime = now
	m.playerCrossObject = o
	o.SetModified()
}

func CreatePlayerCrossDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerCrossDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerCrossDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerCrossDataManager))
}
