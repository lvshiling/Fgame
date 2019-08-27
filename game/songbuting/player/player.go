package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/songbuting/dao"
	songbutingeventtypes "fgame/fgame/game/songbuting/event/types"
)

//玩家送不停管理器
type PlayerSongBuTingDataManager struct {
	p player.Player
	//玩家送不停对象
	songBuTingObject *PlayerSongBuTingObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (psdm *PlayerSongBuTingDataManager) Player() player.Player {
	return psdm.p
}

//加载
func (psdm *PlayerSongBuTingDataManager) Load() (err error) {
	//加载玩家送不停
	songBuTingEntity, err := dao.GetSongBuTingDao().GetSongBuTingEntity(psdm.p.GetId())
	if err != nil {
		return
	}

	if songBuTingEntity == nil {
		psdm.initPlayerSongBuTingObject()
	} else {
		psdm.songBuTingObject = NewPlayerSongBuTingObject(psdm.p)
		psdm.songBuTingObject.FromEntity(songBuTingEntity)
	}
	return nil
}

//第一次初始化
func (psdm *PlayerSongBuTingDataManager) initPlayerSongBuTingObject() {
	pwo := NewPlayerSongBuTingObject(psdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.id = id

	//生成id
	pwo.playerId = psdm.p.GetId()
	pwo.isReceive = 0
	pwo.times = 0
	pwo.lastTime = now
	pwo.createTime = now
	psdm.songBuTingObject = pwo
	pwo.SetModified()
}

//刷新数据
func (psdm *PlayerSongBuTingDataManager) refresh() (err error) {
	now := global.GetGame().GetTimeService().Now()
	if !psdm.songBuTingObject.GetIsReceive() {
		return
	}
	lastTime := psdm.songBuTingObject.lastTime
	flag, err := timeutils.IsSameDay(lastTime, now)
	if err != nil {
		return err
	}
	if !flag {
		gameevent.Emit(songbutingeventtypes.EventTypeSongBuTingCrossFive, psdm.p, psdm.songBuTingObject)
	}
	return nil
}

//加载后
func (psdm *PlayerSongBuTingDataManager) AfterLoad() (err error) {
	err = psdm.refresh()
	psdm.heartbeatRunner.AddTask(CreateSongBuTingTask(psdm.p))
	return
}

//心跳
func (pmdm *PlayerSongBuTingDataManager) Heartbeat() {
	pmdm.heartbeatRunner.Heartbeat()
}

func (pmdm *PlayerSongBuTingDataManager) GetSongBuTingObj() *PlayerSongBuTingObject {
	pmdm.refresh()
	return pmdm.songBuTingObject
}

func (pmdm *PlayerSongBuTingDataManager) SetIsReceive() {
	now := global.GetGame().GetTimeService().Now()
	pmdm.songBuTingObject.isReceive = 1
	pmdm.songBuTingObject.lastTime = now
	pmdm.songBuTingObject.updateTime = now
	pmdm.songBuTingObject.SetModified()
}

func (pmdm *PlayerSongBuTingDataManager) Receive() {
	now := global.GetGame().GetTimeService().Now()
	pmdm.songBuTingObject.lastTime = now
	pmdm.songBuTingObject.times = 1
	pmdm.songBuTingObject.updateTime = now
	pmdm.songBuTingObject.SetModified()
}

func (pmdm *PlayerSongBuTingDataManager) CrossFiveReset() *PlayerSongBuTingObject {
	now := global.GetGame().GetTimeService().Now()
	pmdm.songBuTingObject.times = 0
	pmdm.songBuTingObject.lastTime = now
	pmdm.songBuTingObject.SetModified()
	return pmdm.songBuTingObject
}

//仅gm使用
func (pmdm *PlayerSongBuTingDataManager) GmClearNum() {
	if !pmdm.songBuTingObject.GetIsReceive() {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pmdm.songBuTingObject.times = 0
	pmdm.songBuTingObject.updateTime = now
	pmdm.songBuTingObject.SetModified()
}

func CreatePlayerSongBuTingDataManager(p player.Player) player.PlayerDataManager {
	psdm := &PlayerSongBuTingDataManager{}
	psdm.p = p
	psdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return psdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSongBuTingDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSongBuTingDataManager))
}
