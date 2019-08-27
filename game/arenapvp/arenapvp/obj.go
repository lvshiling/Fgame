package arenapvp

import (
	"fgame/fgame/core/storage"
	arenapvpentity "fgame/fgame/game/arenapvp/entity"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//竞技场竞猜日志对象
type ArenapvpGuessRecordObject struct {
	id         int64
	serverId   int32
	playerId   int64
	raceNumber int32
	guessType  arenapvptypes.ArenapvpType
	guessId    int64
	winnerId   int64
	status     arenapvptypes.ArenapvpGuessState
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewArenapvpGuessRecordObject() *ArenapvpGuessRecordObject {
	o := &ArenapvpGuessRecordObject{}
	return o
}

func (o *ArenapvpGuessRecordObject) GetDBId() int64 {
	return o.id
}

func (o *ArenapvpGuessRecordObject) FromEntity(e storage.Entity) error {
	te := e.(*arenapvpentity.ArenapvpGuessRecordEntity)
	o.id = te.Id
	o.serverId = te.ServerId
	o.playerId = te.PlayerId
	o.raceNumber = te.RaceNumber
	o.guessType = arenapvptypes.ArenapvpType(te.GuessType)
	o.guessId = te.GuessId
	o.winnerId = te.WinnerId
	o.status = arenapvptypes.ArenapvpGuessState(te.Status)
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}
func (o *ArenapvpGuessRecordObject) ToEntity() (e storage.Entity, err error) {

	e = &arenapvpentity.ArenapvpGuessRecordEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		PlayerId:   o.playerId,
		RaceNumber: o.raceNumber,
		GuessType:  int32(o.guessType),
		GuessId:    o.guessId,
		WinnerId:   o.winnerId,
		Status:     int32(o.status),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *ArenapvpGuessRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ArenapvpGuessRecord"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
}

func (o *ArenapvpGuessRecordObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *ArenapvpGuessRecordObject) GetGuessId() int64 {
	return o.guessId
}

func (o *ArenapvpGuessRecordObject) GetWinnerId() int64 {
	return o.winnerId
}

func (o *ArenapvpGuessRecordObject) GetRaceNumber() int32 {
	return o.raceNumber
}

func (o *ArenapvpGuessRecordObject) GetGuessType() arenapvptypes.ArenapvpType {
	return o.guessType
}

func (o *ArenapvpGuessRecordObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *ArenapvpGuessRecordObject) GetStatus() arenapvptypes.ArenapvpGuessState {
	return o.status
}
