package player

import (
	"fgame/fgame/core/storage"
	livenessentity "fgame/fgame/game/liveness/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//活跃度任务任务对象
type PlayerLivenessQuestObject struct {
	player     player.Player
	id         int64
	playerId   int64
	questId    int32
	num        int32
	lastTime   int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerLivenessQuestObject(pl player.Player) *PlayerLivenessQuestObject {
	pmo := &PlayerLivenessQuestObject{
		player: pl,
	}
	return pmo
}

func convertObjectQuestToEntity(psco *PlayerLivenessQuestObject) (*livenessentity.PlayerLivenessQuestEntity, error) {

	e := &livenessentity.PlayerLivenessQuestEntity{
		Id:         psco.id,
		PlayerId:   psco.playerId,
		QuestId:    psco.questId,
		Num:        psco.num,
		LastTime:   psco.lastTime,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerLivenessQuestObject) GetPlayerId() int64 {
	return psco.playerId
}

func (psco *PlayerLivenessQuestObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerLivenessQuestObject) GetQuestId() int32 {
	return psco.questId
}

func (psco *PlayerLivenessQuestObject) GetNum() int32 {
	return psco.num
}

func (psco *PlayerLivenessQuestObject) GetLastTime() int64 {
	return psco.lastTime
}

func (psco *PlayerLivenessQuestObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectQuestToEntity(psco)
	return e, err
}

func (psco *PlayerLivenessQuestObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*livenessentity.PlayerLivenessQuestEntity)

	psco.id = pse.Id
	psco.playerId = pse.PlayerId
	psco.questId = pse.QuestId
	psco.num = pse.Num
	psco.lastTime = pse.LastTime
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerLivenessQuestObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("LivenessQuest: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
