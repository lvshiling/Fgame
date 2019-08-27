package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shenmo/dao"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	"fgame/fgame/pkg/idutil"
)

//玩家神魔管理器
type PlayerShenMoDataManager struct {
	p player.Player
	//玩家神魔对象
	playerShenMoObject *PlayerShenMoObject
}

func (m *PlayerShenMoDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerShenMoDataManager) Load() (err error) {
	//加载玩家神魔信息
	shenmoEntity, err := dao.GetShenMoDao().GetShenMoEntity(m.p.GetId())
	if err != nil {
		return
	}
	if shenmoEntity == nil {
		m.initPlayerShenMoObject()
	} else {
		m.playerShenMoObject = NewPlayerShenMoObject(m.p)
		m.playerShenMoObject.FromEntity(shenmoEntity)
	}
	return
}

//第一次初始化
func (m *PlayerShenMoDataManager) initPlayerShenMoObject() {
	pwo := NewPlayerShenMoObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.id = id
	//生成id
	pwo.gongXunNum = 0
	pwo.killNum = 0
	pwo.endTime = 0
	pwo.rewTime = 0
	pwo.createTime = now
	m.playerShenMoObject = pwo
	pwo.SetModified()
}

//加载后
func (m *PlayerShenMoDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerShenMoDataManager) Heartbeat() {

}

func (m *PlayerShenMoDataManager) GetShenMoInfo() *PlayerShenMoObject {
	return m.playerShenMoObject
}

func (m *PlayerShenMoDataManager) Save() {
	now := global.GetGame().GetTimeService().Now()
	m.playerShenMoObject.gongXunNum = m.p.GetShenMoGongXunNum()
	m.playerShenMoObject.endTime = m.p.GetShenMoEndTime()
	m.playerShenMoObject.killNum = m.p.GetShenMoKillNum()
	m.playerShenMoObject.updateTime = now
	m.playerShenMoObject.SetModified()

}

func (m *PlayerShenMoDataManager) SetEndTime() {
	if m.playerShenMoObject.endTime == m.p.GetShenMoEndTime() {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerShenMoObject.killNum = m.p.GetShenMoKillNum()
	m.playerShenMoObject.endTime = m.p.GetShenMoEndTime()
	m.playerShenMoObject.updateTime = now
	m.playerShenMoObject.SetModified()
}

func (m *PlayerShenMoDataManager) IsHasedReward(rankTime int64) (flag bool) {
	flag = false
	if m.playerShenMoObject.rewTime == rankTime {
		flag = true
		return
	}
	return
}

func (m *PlayerShenMoDataManager) RewardGet(rankTime int64, addGongXunNum int32) (flag bool) {
	if m.playerShenMoObject.rewTime == rankTime {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerShenMoObject.rewTime = rankTime
	m.playerShenMoObject.gongXunNum += addGongXunNum
	m.playerShenMoObject.updateTime = now
	m.playerShenMoObject.SetModified()
	gameevent.Emit(shenmoeventtypes.EventTypeShenMoGongXunNumChanged, m.p, addGongXunNum)
	flag = true
	return
}

func (m *PlayerShenMoDataManager) AddGongXun(addGongXunNum int32) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	m.playerShenMoObject.gongXunNum += addGongXunNum
	m.playerShenMoObject.updateTime = now
	m.playerShenMoObject.SetModified()
	gameevent.Emit(shenmoeventtypes.EventTypeShenMoGongXunNumChanged, m.p, addGongXunNum)
	flag = true
	return
}

func CreatePlayerShenMoDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShenMoDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerShenMoWarDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShenMoDataManager))
}
