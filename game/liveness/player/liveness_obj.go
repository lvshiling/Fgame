package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	livenessentity "fgame/fgame/game/liveness/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//活跃度对象
type PlayerLivenessObject struct {
	player      player.Player
	id          int64
	playerId    int64
	liveness    int64
	openBoxList []int32
	lastTime    int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func NewPlayerLivenessObject(pl player.Player) *PlayerLivenessObject {
	pmo := &PlayerLivenessObject{
		player: pl,
	}
	return pmo
}

func convertObjectToEntity(psco *PlayerLivenessObject) (*livenessentity.PlayerLivenessEntity, error) {

	openBoxBytes, err := json.Marshal(psco.openBoxList)
	if err != nil {
		return nil, err
	}

	e := &livenessentity.PlayerLivenessEntity{
		Id:         psco.id,
		PlayerId:   psco.playerId,
		Liveness:   psco.liveness,
		OpenBoxs:   string(openBoxBytes),
		LastTime:   psco.lastTime,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerLivenessObject) GetPlayerId() int64 {
	return psco.playerId
}

func (psco *PlayerLivenessObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerLivenessObject) GetLiveness() int64 {
	return psco.liveness
}

func (psco *PlayerLivenessObject) GetLastTime() int64 {
	return psco.lastTime
}

func (psco *PlayerLivenessObject) GetOpenBoxs() []int32 {
	return psco.openBoxList
}

func (psco *PlayerLivenessObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectToEntity(psco)
	return e, err
}

func (psco *PlayerLivenessObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*livenessentity.PlayerLivenessEntity)
	openBoxList := make([]int32, 0, 8)

	if err := json.Unmarshal([]byte(pse.OpenBoxs), &openBoxList); err != nil {
		return err
	}

	psco.id = pse.Id
	psco.playerId = pse.PlayerId
	psco.liveness = pse.Liveness
	psco.openBoxList = openBoxList
	psco.lastTime = pse.LastTime
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerLivenessObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("Liveness: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
