package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenxi/dao"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家阵法管理器
type PlayerZhenXiDataManager struct {
	p                player.Player
	playerZhenXiBoss *PlayerZhenXiBossObject
}

func (m *PlayerZhenXiDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerZhenXiDataManager) Load() (err error) {
	err = m.loadZhenXiBoss()
	if err != nil {
		return
	}

	return nil
}

//加载
func (m *PlayerZhenXiDataManager) loadZhenXiBoss() (err error) {
	//加载玩家仙盟信息
	zhenXiEntity, err := dao.GetZhenXiDao().GetPlayerZhenXiBoss(m.p.GetId())
	if err != nil {
		return
	}
	if zhenXiEntity == nil {
		m.initPlayerZhenXiObject()
	} else {
		m.playerZhenXiBoss = NewPlayerZhenXiBossObject(m.p)
		m.playerZhenXiBoss.FromEntity(zhenXiEntity)
	}
	return
}

//第一次初始化
func (m *PlayerZhenXiDataManager) initPlayerZhenXiObject() {
	o := NewPlayerZhenXiBossObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.reliveTime = 0
	o.createTime = now
	o.SetModified()
	m.playerZhenXiBoss = o
}

//加载后
func (m *PlayerZhenXiDataManager) AfterLoad() (err error) {
	m.refreshEnterTimes()
	return
}

//心跳
func (m *PlayerZhenXiDataManager) Heartbeat() {
}

// 刷新进入次数
func (m *PlayerZhenXiDataManager) refreshEnterTimes() {
	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameFive(now, m.playerZhenXiBoss.updateTime)

	if !isSame {
		m.playerZhenXiBoss.enterTimes = 0
		m.playerZhenXiBoss.updateTime = now
		m.playerZhenXiBoss.SetModified()
	}
}

func (m *PlayerZhenXiDataManager) GetReliveTime() int32 {
	return m.playerZhenXiBoss.reliveTime
}

func (m *PlayerZhenXiDataManager) GetPlayerZhenXiObject() *PlayerZhenXiBossObject {
	m.refreshEnterTimes()
	return m.playerZhenXiBoss
}

func (m *PlayerZhenXiDataManager) EnterZhenXiBoss() {
	now := global.GetGame().GetTimeService().Now()
	m.playerZhenXiBoss.enterTimes += 1
	m.playerZhenXiBoss.updateTime = now
	m.playerZhenXiBoss.SetModified()
}

func (m *PlayerZhenXiDataManager) Relive() {
	m.playerZhenXiBoss.reliveTime += 1
}

func CreatePlayerZhenXiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerZhenXiDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerZhenXiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerZhenXiDataManager))
}
