package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/moonlove/dao"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

//玩家月下情缘数据管理器
type PlayerMoonloveDataManager struct {
	play player.Player
	//月下情缘数据对象
	moonloveObject *PlayerMoonloveObject
	//上次奖励时间
	preRewTime int64
	//进入场景时间
	enterTime int64
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMoonloveDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMoonloveDataManager))
}

func (m *PlayerMoonloveDataManager) Player() player.Player {
	return m.play
}

//加载月下情缘信息
func (m *PlayerMoonloveDataManager) Load() error {
	xfEntity, err := dao.GetMoonloveDao().GetMoonloveInfo(m.play.GetId())
	if err != nil {
		return err
	}

	if xfEntity != nil {
		newObj := CreatePlayerMoonloveObject(m.play)
		newObj.FromEntity(xfEntity)
		m.moonloveObject = newObj
	} else {
		m.initNewPlayerMoonloveObject()
	}

	return nil
}

//加载后处理
func (m *PlayerMoonloveDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerMoonloveDataManager) Heartbeat() {
}

//初始化月下情缘数据
func (m *PlayerMoonloveDataManager) initNewPlayerMoonloveObject() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	newObj := CreatePlayerMoonloveObject(m.play)
	newObj.id = id
	newObj.charmNum = 0
	newObj.generousNum = 0
	newObj.createTime = now

	m.moonloveObject = newObj
	newObj.SetModified()
}

//增加豪气值
func (m *PlayerMoonloveDataManager) AddGenerousNum(num int32) {
	if num < 0 {
		panic("generousNum should be greater than 0")
	}
	m.moonloveObject.generousNum += num
	m.moonloveObject.SetModified()
}

//获取豪气值
func (m *PlayerMoonloveDataManager) GetGenerousNum() int32 {
	return m.moonloveObject.generousNum
}

//增加魅力值
func (m *PlayerMoonloveDataManager) AddCharmNum(num int32) {
	if num < 0 {
		panic("charmNum should be greater than 0")
	}
	m.moonloveObject.charmNum += num
	m.moonloveObject.SetModified()

}

//获取魅力值
func (m *PlayerMoonloveDataManager) GetCharmNum() int32 {
	return m.moonloveObject.charmNum
}

//更新进入场景时间
func (m *PlayerMoonloveDataManager) UpdateEnterTime(endTime int64) {
	now := global.GetGame().GetTimeService().Now()
	m.enterTime = now
	m.preRewTime = 0

	if endTime != m.moonloveObject.preActivityTime {
		m.moonloveObject.preActivityTime = endTime
		m.moonloveObject.charmNum = 0
		m.moonloveObject.generousNum = 0

		m.moonloveObject.SetModified()
	}

}

//更新获取奖励时间
func (m *PlayerMoonloveDataManager) UpdateRewTime() {
	now := global.GetGame().GetTimeService().Now()
	m.preRewTime = now
}

func (m *PlayerMoonloveDataManager) GetEnterTime() int64 {
	return m.enterTime
}

func (m *PlayerMoonloveDataManager) GetPreRewTime() int64 {
	return m.preRewTime
}

func CreatePlayerMoonloveDataManager(pl player.Player) player.PlayerDataManager {
	m := &PlayerMoonloveDataManager{}
	m.play = pl

	return m
}
