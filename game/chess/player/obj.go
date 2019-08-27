package player

import (
	"fgame/fgame/core/storage"
	chessentity "fgame/fgame/game/chess/entity"
	chesstemplate "fgame/fgame/game/chess/template"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

const (
	specialHandleChessTimes = 3 //棋局总次数的前三次不刷新
)

//棋局对象
type PlayerChessObject struct {
	player                player.Player
	id                    int64
	chessId               int32
	attendTimes           int32
	totalAttendTimes      int32
	chessType             chesstypes.ChessType
	lastSystemRefreshTime int64
	updateTime            int64
	createTime            int64
	deleteTime            int64
}

func NewPlayerChessObject(pl player.Player) *PlayerChessObject {
	o := &PlayerChessObject{
		player: pl,
	}
	return o
}

func convertNewPlayerChessObjectToEntity(o *PlayerChessObject) (*chessentity.PlayerChessEntity, error) {
	e := &chessentity.PlayerChessEntity{
		Id:                    o.id,
		PlayerId:              o.GetPlayerId(),
		ChessId:               o.chessId,
		AttendTimes:           o.attendTimes,
		TotalAttendTimes:      o.totalAttendTimes,
		ChessType:             int32(o.chessType),
		LastSystemRefreshTime: o.lastSystemRefreshTime,
		UpdateTime:            o.updateTime,
		CreateTime:            o.createTime,
		DeleteTime:            o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChessObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerChessObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChessObject) GetChessId() int32 {
	return o.chessId
}

func (o *PlayerChessObject) GetChessType() chesstypes.ChessType {
	return o.chessType
}

func (o *PlayerChessObject) GetAttendTimes() int32 {
	return o.attendTimes
}

func (o *PlayerChessObject) GetToatalAttendTimes() int32 {
	return o.totalAttendTimes
}

func (o *PlayerChessObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerChessObjectToEntity(o)
	return e, err
}

func (o *PlayerChessObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chessentity.PlayerChessEntity)

	o.id = pse.Id

	o.chessId = pse.ChessId
	o.attendTimes = pse.AttendTimes
	o.totalAttendTimes = pse.TotalAttendTimes
	o.chessType = chesstypes.ChessType(pse.ChessType)
	o.lastSystemRefreshTime = pse.LastSystemRefreshTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChessObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Chess"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerChessObject) randomChess() (newChessDropId int32) {
	now := global.GetGame().GetTimeService().Now()

	chessTemplate := chesstemplate.GetChessTemplateService().GetChessRandom(o.chessType, o.chessId)
	o.chessId = chessTemplate.ChessId
	o.updateTime = now
	o.lastSystemRefreshTime = now
	o.SetModified()

	newChessDropId = chessTemplate.DropId
	return newChessDropId
}

