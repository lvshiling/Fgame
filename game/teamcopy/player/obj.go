package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	teamtypes "fgame/fgame/game/team/types"
	teamcopyentity "fgame/fgame/game/teamcopy/entity"

	"github.com/pkg/errors"
)

//组队副本对象
type PlayerTeamCopyObject struct {
	player     player.Player
	id         int64
	purpose    teamtypes.TeamPurposeType
	num        int32
	rewTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerTeamCopyObject(pl player.Player) *PlayerTeamCopyObject {
	pmo := &PlayerTeamCopyObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerTeamCopyObjectToEntity(pqo *PlayerTeamCopyObject) (e *teamcopyentity.PlayerTeamCopyEntity, err error) {
	e = &teamcopyentity.PlayerTeamCopyEntity{
		Id:         pqo.id,
		PurPose:    int32(pqo.purpose),
		PlayerId:   pqo.player.GetId(),
		Num:        pqo.num,
		RewTime:    pqo.rewTime,
		UpdateTime: pqo.updateTime,
		CreateTime: pqo.createTime,
		DeleteTime: pqo.deleteTime,
	}
	return
}

func (pqo *PlayerTeamCopyObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerTeamCopyObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerTeamCopyObject) GetNum() int32 {
	return pqo.num
}

func (pqo *PlayerTeamCopyObject) GetPurpose() teamtypes.TeamPurposeType {
	return pqo.purpose
}

func (pqo *PlayerTeamCopyObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerTeamCopyObjectToEntity(pqo)
	return
}

func (pqo *PlayerTeamCopyObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*teamcopyentity.PlayerTeamCopyEntity)

	pqo.id = pqe.Id
	pqo.num = pqe.Num
	pqo.purpose = teamtypes.TeamPurposeType(pqe.PurPose)
	pqo.rewTime = pqe.RewTime
	pqo.updateTime = pqe.UpdateTime
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerTeamCopyObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "team_copy"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pqo.player.AddChangedObject(obj)
	return
}
