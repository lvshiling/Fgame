package player

import (
	lingtongdevcommon "fgame/fgame/game/lingtongdev/common"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家灵童养成管理器
type PlayerLingTongDevDataManager struct {
	p player.Player
	//玩家灵童养成对象
	playerLingTongDevMap map[types.LingTongDevSysType]*PlayerLingTongDevObject
	//玩家非进阶灵童养成对象
	playerOtherMap map[types.LingTongDevSysType]*LingTongOtherContainer
	//玩家灵童养成战力
	playerPowerMap map[types.LingTongDevSysType]*PlayerLingTongPowerObject
}

func (m *PlayerLingTongDevDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerLingTongDevDataManager) Load() (err error) {
	m.loadDevelop()
	m.loadOther()
	m.loadPower()
	return nil
}

//加载后
func (m *PlayerLingTongDevDataManager) AfterLoad() (err error) {
	err = m.refreshAllBless()
	return nil
}

//心跳
func (m *PlayerLingTongDevDataManager) Heartbeat() {

}

//灵童养成信息
func (m *PlayerLingTongDevDataManager) ToAllLingTongDevInfo() (allLingTongDevInfo *lingtongdevcommon.AllLingTongDevInfo) {
	allLingTongDevInfo = &lingtongdevcommon.AllLingTongDevInfo{}
	for _, tempLingTongDev := range m.playerLingTongDevMap {
		lingTongDevInfo := &lingtongdevcommon.LingTongDevInfo{
			ClassType:     int32(tempLingTongDev.classType),
			AdvanceId:     tempLingTongDev.advanceId,
			SeqId:         tempLingTongDev.seqId,
			UnrealLevel:   tempLingTongDev.unrealLevel,
			UnrealPro:     tempLingTongDev.unrealPro,
			CulLevel:      tempLingTongDev.culLevel,
			CulPro:        tempLingTongDev.culPro,
			TongLingLevel: tempLingTongDev.tongLingLevel,
			TongLingPro:   tempLingTongDev.tongLingPro,
		}
		otherContainer, ok := m.playerOtherMap[tempLingTongDev.classType]
		if ok {
			for _, typM := range otherContainer.otherMap {
				for _, otherObj := range typM {
					skinInfo := &lingtongdevcommon.LingTongDevSkinInfo{
						SeqId: otherObj.seqId,
						Level: otherObj.level,
						UpPro: otherObj.upPro,
					}
					lingTongDevInfo.SkinList = append(lingTongDevInfo.SkinList, skinInfo)
				}
			}
		}
		allLingTongDevInfo.LingTongDevList = append(allLingTongDevInfo.LingTongDevList, lingTongDevInfo)
	}
	return
}

func CreatePlayerLingTongDevDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerLingTongDevDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerLingTongDevDataManager))
}
