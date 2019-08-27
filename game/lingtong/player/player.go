package player

import (
	"fgame/fgame/core/heartbeat"
	lingtongcommon "fgame/fgame/game/lingtong/common"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
)

//玩家灵童管理器
type PlayerLingTongDataManager struct {
	p player.Player
	//激活的时装
	fashionMapObj *LingTongFashionContainer
	//试用的时装
	trialFashionObj *PlayerLingTongFashionTrialObject
	//灵童时装
	lingTongFashionMap map[int32]*PlayerLingTongFashionObject
	//灵童
	lingTongMap map[int32]*PlayerLingTongInfoObject
	//出战灵童
	lingTongObj *PlayerLingTongObject
	//战斗属性组
	battlePropertyGroup *propertycommon.BattlePropertyGroup
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (m *PlayerLingTongDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerLingTongDataManager) Load() (err error) {
	m.loadLingTongFashionTrial()
	m.loadFashion()
	m.loadLingTongFashion()
	m.loadLingTongInfo()
	m.loadLingTong()
	return nil
}

//加载后
func (m *PlayerLingTongDataManager) AfterLoad() (err error) {
	m.refreshLingTongFashion()
	m.heartbeatRunner.AddTask(CreateLingTongFashionTask(m.p))
	return nil
}

//心跳
func (m *PlayerLingTongDataManager) Heartbeat() {
	m.heartbeatRunner.Heartbeat()
}

func (m *PlayerLingTongDataManager) GetAllLingTongLevel() int32 {
	totalLevel := int32(0)
	for _, lingTongInfoObj := range m.lingTongMap {
		totalLevel += lingTongInfoObj.GetLevel()
	}
	return totalLevel
}

func (m *PlayerLingTongDataManager) ToLingTongInfo() *lingtongcommon.LingTongInfo {
	lingTongObj := m.GetLingTong()
	lingTongId := lingTongObj.GetLingTongId()
	level := lingTongObj.GetLevel()
	fashionId := int32(0)
	lingTongFashionObj := m.GetLingTongFashionById(lingTongId)
	if lingTongFashionObj != nil {
		fashionId = lingTongFashionObj.GetFashionId()
	}

	info := &lingtongcommon.LingTongInfo{
		LingTongId: lingTongId,
		FashionId:  fashionId,
		Level:      level,
	}

	for _, lingTongInfoObj := range m.lingTongMap {
		lingTongId := lingTongInfoObj.GetLingTongId()
		selfFashionId := int32(0)
		selfLingTongFashionObj := m.lingTongFashionMap[lingTongId]
		selfFashionId = selfLingTongFashionObj.GetFashionId()
		lingTongInfo := &lingtongcommon.LingTongDetail{
			LingTongId:   lingTongId,
			LingTongName: lingTongInfoObj.GetLingTongName(),
			FashionId:    selfFashionId,
			Level:        lingTongInfoObj.GetLevel(),
			LevelPro:     lingTongInfoObj.GetPro(),
			PeiYangLevel: lingTongInfoObj.GetPeiYangLevel(),
			PeiYangPro:   lingTongInfoObj.GetPeiYangPro(),
			StarLevel:    lingTongInfoObj.GetStarLevel(),
			StarPro:      lingTongInfoObj.GetStarPro(),
		}
		info.LingTongList = append(info.LingTongList, lingTongInfo)
	}

	return info
}

func CreatePlayerLingTongDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerLingTongDataManager{}
	m.p = p
	m.battlePropertyGroup = propertycommon.NewBattlePropertyGroup()
	m.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerLingTongDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerLingTongDataManager))
}
