package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/yuxi/dao"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家仙盟管理器
type PlayerYuXiDataManager struct {
	p player.Player
	//玩家玉玺对象
	playerYuXiObject *PlayerAllianceYuXiObject
}

func (m *PlayerYuXiDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerYuXiDataManager) Load() (err error) {
	//加载玩家仙盟信息
	playerYuXiEntity, err := dao.GetYuXiDao().GetPlayerAllianceYuXi(m.p.GetId())
	if err != nil {
		return
	}
	if playerYuXiEntity == nil {
		m.initPlayerYuXiObject()
	} else {
		m.playerYuXiObject = newPlayerAllianceYuXiObject(m.p)
		m.playerYuXiObject.FromEntity(playerYuXiEntity)
	}
	return
}

//第一次初始化
func (m *PlayerYuXiDataManager) initPlayerYuXiObject() {
	o := newPlayerAllianceYuXiObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.isReceive = 0
	o.createTime = now
	o.SetModified()

	m.playerYuXiObject = o
}

//心跳
func (m *PlayerYuXiDataManager) Heartbeat() {
}

//加载后
func (m *PlayerYuXiDataManager) AfterLoad() (err error) {

	//刷新捐献次数
	err = m.RefreshTimes()
	if err != nil {
		return err
	}

	return nil
}

//刷新领取状态
func (m *PlayerYuXiDataManager) RefreshTimes() error {
	now := global.GetGame().GetTimeService().Now()
	flag, err := timeutils.IsSameDay(m.playerYuXiObject.updateTime, now)
	if err != nil {
		return err
	}
	if !flag {
		m.playerYuXiObject.isReceive = 0
		m.playerYuXiObject.updateTime = now
		m.playerYuXiObject.SetModified()
	}
	return nil
}

//刷新领取状态
func (m *PlayerYuXiDataManager) ReceiveDayRew() (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	if m.playerYuXiObject.IsReceive() {
		return
	}

	m.playerYuXiObject.isReceive = 1
	m.playerYuXiObject.updateTime = now
	m.playerYuXiObject.SetModified()
	flag = true
	return
}

//获取玩家玉玺信息
func (m *PlayerYuXiDataManager) GetPlayerYuXiInfo() *PlayerAllianceYuXiObject {
	return m.playerYuXiObject
}

func createPlayerYuXiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerYuXiDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerYuXiDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerYuXiDataManager))
}
