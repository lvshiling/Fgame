package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/worldboss/dao"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/pkg/idutil"
)

//玩家周卡管理器
type PlayerWorldbossManager struct {
	p             player.Player
	bossReliveMap map[worldbosstypes.BossType]*PlayerBossReliveObject //周卡数据
}

func (m *PlayerWorldbossManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerWorldbossManager) Load() (err error) {
	m.bossReliveMap = make(map[worldbosstypes.BossType]*PlayerBossReliveObject)
	//加载周卡数据
	bossReliveEntityList, err := dao.GetWorldbossDao().GetPlayerBossReliveList(m.p.GetId())
	if err != nil {
		return
	}

	for _, bossReliveEntity := range bossReliveEntityList {
		obj := newPlayerBossReliveObject(m.p)
		obj.FromEntity(bossReliveEntity)
		m.addBossObj(obj)
	}

	return nil
}

//加载
func (m *PlayerWorldbossManager) addBossObj(obj *PlayerBossReliveObject) {
	m.bossReliveMap[obj.bossType] = obj
}

//加载
func (m *PlayerWorldbossManager) getBossObj(bossType worldbosstypes.BossType) *PlayerBossReliveObject {
	obj, ok := m.bossReliveMap[bossType]
	if !ok {
		return nil
	}
	return obj
}

//加载后
func (m *PlayerWorldbossManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerWorldbossManager) Heartbeat() {
}

func (m *PlayerWorldbossManager) initBossReliveObj(bossType worldbosstypes.BossType) *PlayerBossReliveObject {
	now := global.GetGame().GetTimeService().Now()
	obj := newPlayerBossReliveObject(m.p)
	obj.id, _ = idutil.GetId()
	obj.reliveTime = 0
	obj.createTime = now
	obj.bossType = bossType
	obj.SetModified()
	m.addBossObj(obj)
	return obj
}

//领取奖励
func (m *PlayerWorldbossManager) BossSync(bossType worldbosstypes.BossType, reliveTime int32) {
	now := global.GetGame().GetTimeService().Now()
	obj := m.getBossObj(bossType)
	if obj == nil {
		if reliveTime == 0 {
			return
		}
		obj = m.initBossReliveObj(bossType)
		obj.reliveTime = reliveTime
		obj.updateTime = now
		obj.SetModified()
		return
	}
	obj.reliveTime = reliveTime
	obj.updateTime = now
	obj.SetModified()
	return
}

func (m *PlayerWorldbossManager) GetBossReliveTime(bossType worldbosstypes.BossType) int32 {
	obj := m.getBossObj(bossType)
	if obj == nil {
		return 0
	}
	return obj.reliveTime
}

func (m *PlayerWorldbossManager) GetBossReliveList() (reliveList []*scene.PlayerBossReliveData) {
	for bossType, reliveObj := range m.bossReliveMap {
		reliveData := scene.CreatePlayerBossReliveData(bossType, reliveObj.reliveTime)
		reliveList = append(reliveList, reliveData)
	}
	return reliveList
}

func createPlayerWorldbossDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerWorldbossManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerWorldbossManagerType, player.PlayerDataManagerFactoryFunc(createPlayerWorldbossDataManager))
}
