package player

import (
	"fgame/fgame/core/heartbeat"
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家挂机管理器
type PlayerGuaJiManager struct {
	p             player.Player
	runner        heartbeat.HeartbeatTaskRunner
	guaJiDataList []*guajitypes.GuaJiData //挂机列表

	guaJiIndex       int32 //挂机索引
	currentGuaJiType guajitypes.GuaJiType

	advanceSettingMap map[guajitypes.GuaJiAdvanceType]int32
	globalTypeMap     map[guajitypes.GuaJiGlobalType]int32
	//挂机剩余
	remainSilver int64
}

func (m *PlayerGuaJiManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerGuaJiManager) Load() (err error) {
	return nil
}

//加载后
func (m *PlayerGuaJiManager) AfterLoad() (err error) {
	m.guaJiIndex = -1
	//提升检查任务
	checkTask := CreateCheckTask(m.p)
	m.runner.AddTask(checkTask)
	advanceCheckTask := CreateAdvanceCheckTask(m.p)
	m.runner.AddTask(advanceCheckTask)
	return
}

//心跳
func (m *PlayerGuaJiManager) Heartbeat() {
	//没有在挂机的
	if len(m.guaJiDataList) <= 0 {
		return
	}
	m.runner.Heartbeat()
}

//开始挂机
func (m *PlayerGuaJiManager) StartGuaJiList(guaJiDataList []*guajitypes.GuaJiData, guaJiAdvanceTypeSettingMap map[guajitypes.GuaJiAdvanceType]int32) {
	if len(guaJiDataList) <= 0 {
		return
	}
	m.guaJiDataList = guaJiDataList
	m.advanceSettingMap = guaJiAdvanceTypeSettingMap
}

func (m *PlayerGuaJiManager) StopGuaJiList() {
	if m.guaJiIndex >= 0 {
		m.StopCurrentGuaJi()
	}
	m.guaJiDataList = nil

}

func (m *PlayerGuaJiManager) GetGuaJiTypeList() []*guajitypes.GuaJiData {
	return m.guaJiDataList
}

func (m *PlayerGuaJiManager) GetCurrentGuaJiType() (guaJiData *guajitypes.GuaJiData, index int32) {
	if m.guaJiIndex >= 0 {
		guaJiData = m.guaJiDataList[m.guaJiIndex]
		index = m.guaJiIndex
		return
	}
	index = -1
	return
}

func (m *PlayerGuaJiManager) GetGuaJiType(typ guajitypes.GuaJiType) (guaJiData *guajitypes.GuaJiData) {
	for _, tempGuaJiData := range m.guaJiDataList {
		if tempGuaJiData.GetType() == typ {
			return tempGuaJiData
		}
	}
	return
}

//开始挂机
func (m *PlayerGuaJiManager) StartGuaJi(index int32) bool {
	if m.guaJiIndex >= 0 {
		return false
	}
	if index >= int32(len(m.guaJiDataList)) {
		return false
	}

	m.guaJiIndex = index
	gameevent.Emit(guajieventtypes.GuaJiEventTypeGuaJiStart, m.p, nil)
	return true
}

//停止挂机
func (m *PlayerGuaJiManager) StopCurrentGuaJi() bool {
	if m.guaJiIndex < 0 {
		return false
	}
	m.guaJiIndex = -1
	gameevent.Emit(guajieventtypes.GuaJiEventTypeGuaJiStop, m.p, nil)
	return true
}

//停止挂机
func (m *PlayerGuaJiManager) GetRemainSilver() int64 {
	return m.remainSilver
}

func (m *PlayerGuaJiManager) GetGlobalValue(typ guajitypes.GuaJiGlobalType) int32 {
	return m.globalTypeMap[typ]
}

func (m *PlayerGuaJiManager) GetAdvanceSettingValue(typ guajitypes.GuaJiAdvanceType) int32 {
	return m.advanceSettingMap[typ]
}

const (
	remainSilver = 100000
)

func createPlayerGuaJiManager(p player.Player) player.PlayerDataManager {
	m := &PlayerGuaJiManager{}
	m.p = p
	m.runner = heartbeat.NewHeartbeatTaskRunner()
	m.remainSilver = remainSilver
	m.globalTypeMap = make(map[guajitypes.GuaJiGlobalType]int32)
	m.advanceSettingMap = make(map[guajitypes.GuaJiAdvanceType]int32)
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeMount] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeWing] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeBodyshield] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeAnqi] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeFabao] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeShenfa] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeXianti] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingyu] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeShihunfan] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeTianmoti] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeFeather] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeShield] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeMassacre] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongWeapon] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongMount] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongWing] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongShenFa] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongLingYu] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongFaBao] = 5
	// m.advanceSettingMap[guajitypes.GuaJiAdvanceTypeLingTongXianTi] = 5

	m.globalTypeMap[guajitypes.GuaJiGlobalTypeAdvanceAutoBuy] = 1
	m.globalTypeMap[guajitypes.GuaJiGlobalTypeBagRemainSlots] = 10
	m.globalTypeMap[guajitypes.GuaJiAutoBuyBagLevel] = 70
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerGuaJiManagerType, player.PlayerDataManagerFactoryFunc(createPlayerGuaJiManager))
}
