package player

import (
	"fgame/fgame/game/densewat/dao"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

//玩家金银密窟管理器
type PlayerDenseWatDataManager struct {
	p player.Player
	//玩家金银密窟对象
	playerDenseWatObject *PlayerDenseWatObject
}

func (pwdm *PlayerDenseWatDataManager) Player() player.Player {
	return pwdm.p
}

//加载
func (pwdm *PlayerDenseWatDataManager) Load() (err error) {
	//加载玩家金银密窟信息
	denseWatEntity, err := dao.GetDenseWatDao().GetDenseWatEntity(pwdm.p.GetId())
	if err != nil {
		return
	}
	if denseWatEntity == nil {
		pwdm.initPlayerDenseWatObject()
	} else {
		pwdm.playerDenseWatObject = NewPlayerDenseWatObject(pwdm.p)
		pwdm.playerDenseWatObject.FromEntity(denseWatEntity)
	}
	return nil
}

//第一次初始化
func (pwdm *PlayerDenseWatDataManager) initPlayerDenseWatObject() {
	pwo := NewPlayerDenseWatObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.id = id
	//生成id
	pwo.playerId = pwdm.p.GetId()
	pwo.num = 0
	pwo.endTime = now
	pwo.createTime = now
	pwdm.playerDenseWatObject = pwo
	pwo.SetModified()
}

//加载后
func (pwdm *PlayerDenseWatDataManager) AfterLoad() (err error) {
	return nil
}

//金银密窟信息对象
func (pwdm *PlayerDenseWatDataManager) GetDenseWatInfo() *PlayerDenseWatObject {
	return pwdm.playerDenseWatObject
}

//心跳
func (pwdm *PlayerDenseWatDataManager) Heartbeat() {

}

func (pwdm *PlayerDenseWatDataManager) Save() {
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerDenseWatObject.num = pwdm.p.GetDenseWatNum()
	pwdm.playerDenseWatObject.endTime = pwdm.p.GetDenseWatEndTime()
	pwdm.playerDenseWatObject.updateTime = now
	pwdm.playerDenseWatObject.SetModified()
}

func (pwdm *PlayerDenseWatDataManager) SetEndTime() {
	if pwdm.playerDenseWatObject.endTime == pwdm.p.GetDenseWatEndTime() {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerDenseWatObject.num = pwdm.p.GetDenseWatNum()
	pwdm.playerDenseWatObject.endTime = pwdm.p.GetDenseWatEndTime()
	pwdm.playerDenseWatObject.updateTime = now
	pwdm.playerDenseWatObject.SetModified()
}

func CreatePlayerDenseWatDataManager(p player.Player) player.PlayerDataManager {
	pwdm := &PlayerDenseWatDataManager{}
	pwdm.p = p
	return pwdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerDenseWatDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerDenseWatDataManager))
}
