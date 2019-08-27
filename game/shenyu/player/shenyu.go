package player

import (
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shenyu/dao"
	shenyueventtypes "fgame/fgame/game/shenyu/event/types"
	shenyutemplate "fgame/fgame/game/shenyu/template"
	"fgame/fgame/pkg/idutil"
	"math"
)

//玩家神域管理器
type PlayerShenYuDataManager struct {
	p player.Player
	//玩家神域对象
	shenYuObject *PlayerShenYuObject
	lastDeadTime int64
}

func (m *PlayerShenYuDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerShenYuDataManager) Load() (err error) {
	//加载玩家神域
	entity, err := dao.GetShenYuDao().GetShenYuEntity(m.p.GetId())
	if err != nil {
		return
	}
	if entity == nil {
		m.initPlayerShenYuObject()
	} else {
		m.shenYuObject = NewPlayerShenYuObject(m.p)
		m.shenYuObject.FromEntity(entity)
	}

	return nil
}

//第一次初始化
func (m *PlayerShenYuDataManager) initPlayerShenYuObject() {
	pdo := NewPlayerShenYuObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pdo.id = id
	//生成id
	pdo.keyNum = int32(0)
	pdo.round = 0
	pdo.exp = int64(0)
	pdo.itemMap = make(map[int32]int32)
	pdo.endTime = int64(0)
	pdo.createTime = now
	m.shenYuObject = pdo
	pdo.SetModified()
}

//加载后
func (m *PlayerShenYuDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerShenYuDataManager) Heartbeat() {

}

//获取钥匙
func (m *PlayerShenYuDataManager) GetKeyNum() int32 {
	return m.shenYuObject.keyNum
}

//获取参赛轮
func (m *PlayerShenYuDataManager) GetRound() int32 {
	return m.shenYuObject.round
}

//进入神域
func (m *PlayerShenYuDataManager) EnterShenYu(endTime int64, enterRound int32, resetFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	constTemp := shenyutemplate.GetShenYuTemplateService().GetShenYuConstantTemplate()
	// 第一次进入神域
	if !m.IsAttendShenYu(endTime) {
		m.shenYuObject.keyNum = constTemp.InitKeyNum
		m.shenYuObject.round = 1
		m.shenYuObject.exp = 0
		m.shenYuObject.itemMap = make(map[int32]int32)
		m.shenYuObject.endTime = endTime
		m.shenYuObject.updateTime = now
		m.shenYuObject.SetModified()
		gameevent.Emit(shenyueventtypes.EventTypeShenYuKeyChange, m.p, nil)
		return
	}

	// 重登钥匙减半
	curRound := m.shenYuObject.round
	if curRound == enterRound {
		m.keyHalve()
		return
	}

	// 重置钥匙
	if resetFlag {
		m.shenYuObject.keyNum = constTemp.ResetKeyNum
		gameevent.Emit(shenyueventtypes.EventTypeShenYuKeyChange, m.p, nil)
	}

	m.shenYuObject.round += 1
	m.shenYuObject.updateTime = now
	m.shenYuObject.SetModified()
	return
}

//增加钥匙
func (m *PlayerShenYuDataManager) AddKeyNum(num int32) {
	if num <= 0 {
		return
	}
	curNum := m.shenYuObject.keyNum + num
	keyMax := shenyutemplate.GetShenYuTemplateService().GetShenYuConstantTemplate().KeyMax

	if curNum > keyMax {
		curNum = keyMax
	}

	now := global.GetGame().GetTimeService().Now()
	m.shenYuObject.keyNum = curNum
	m.shenYuObject.updateTime = now
	m.shenYuObject.SetModified()

	gameevent.Emit(shenyueventtypes.EventTypeShenYuKeyChange, m.p, nil)
	return
}

// 玩家死亡
func (m *PlayerShenYuDataManager) PlayerDead() (dropNum int32) {
	return m.keyHalve()
}

//钥匙减半
func (m *PlayerShenYuDataManager) keyHalve() (dropNum int32) {
	shenYuConstantTemp := shenyutemplate.GetShenYuTemplateService().GetShenYuConstantTemplate()
	beforeNum := m.shenYuObject.keyNum
	keepMin := shenYuConstantTemp.KeyKeepMin
	if beforeNum <= keepMin {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	deadDropRatio := shenYuConstantTemp.KeyDropPercent
	dropNum = int32(math.Floor(float64(beforeNum) * float64(deadDropRatio) / float64(common.MAX_RATE)))

	curNum := m.shenYuObject.keyNum - dropNum
	if curNum < 0 {
		curNum = 0
	}
	m.shenYuObject.keyNum = curNum
	m.shenYuObject.updateTime = now
	m.shenYuObject.SetModified()

	gameevent.Emit(shenyueventtypes.EventTypeShenYuKeyChange, m.p, nil)
	return
}

//GM 设置钥匙数量
func (m *PlayerShenYuDataManager) GMSetKeyNum(keyNum int32) {
	if keyNum < 0 {
		return
	}
	keyMax := shenyutemplate.GetShenYuTemplateService().GetShenYuConstantTemplate().KeyMax
	if keyNum > keyMax {
		keyNum = keyMax
	}

	now := global.GetGame().GetTimeService().Now()
	m.shenYuObject.keyNum = keyNum
	m.shenYuObject.updateTime = now
	m.shenYuObject.SetModified()
	gameevent.Emit(shenyueventtypes.EventTypeShenYuKeyChange, m.p, nil)
}

// 是否参与
func (m *PlayerShenYuDataManager) IsAttendShenYu(endTime int64) bool {
	if m.shenYuObject.endTime != endTime {
		return false
	}

	return m.shenYuObject.round > 0
}

func CreatePlayerShenYuDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShenYuDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerShenYuDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShenYuDataManager))
}
