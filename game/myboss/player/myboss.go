package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/myboss/dao"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家个人BOSS管理器
type PlayerMyBossDataManager struct {
	p                  player.Player
	playerMyBossObject *PlayerMyBossObject
}

func (m *PlayerMyBossDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerMyBossDataManager) Load() (err error) {
	//加载玩家个人BOSS信息
	mybossEntity, err := dao.GetMyBossDao().GetMyBossEntity(m.p.GetId())
	if err != nil {
		return
	}
	if mybossEntity == nil {
		m.initPlayerMyBossObject()
	} else {
		m.playerMyBossObject = NewPlayerMyBossObject(m.p)
		m.playerMyBossObject.FromEntity(mybossEntity)
	}
	return nil
}

//第一次初始化
func (m *PlayerMyBossDataManager) initPlayerMyBossObject() {
	o := NewPlayerMyBossObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.attendMap = map[int32]int32{}
	o.createTime = now
	o.SetModified()

	m.playerMyBossObject = o
}

//加载后
func (m *PlayerMyBossDataManager) AfterLoad() (err error) {
	m.RefreshTimes()
	return
}

//心跳
func (m *PlayerMyBossDataManager) Heartbeat() {
}

//更新次数
func (m *PlayerMyBossDataManager) AddAttendTimes(biologyId int32) {
	m.playerMyBossObject.attendMap[biologyId] += 1
	now := global.GetGame().GetTimeService().Now()
	m.playerMyBossObject.updateTime = now
	m.playerMyBossObject.SetModified()
}

func (m *PlayerMyBossDataManager) GetAttendTimes(biologyId int32) int32 {
	return m.playerMyBossObject.attendMap[biologyId]
}

func (m *PlayerMyBossDataManager) GetAttendTimesAll() map[int32]int32 {
	return m.playerMyBossObject.attendMap
}

func (m *PlayerMyBossDataManager) RefreshTimes() {
	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameFive(now, m.playerMyBossObject.updateTime)
	if !isSame {
		m.playerMyBossObject.attendMap = map[int32]int32{}
		m.playerMyBossObject.updateTime = now
		m.playerMyBossObject.SetModified()
	}
}

func CreatePlayerMyBossDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerMyBossDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMyBossDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMyBossDataManager))
}
