package chess

import (
	"fgame/fgame/core/storage"
	chessentity "fgame/fgame/game/chess/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//棋局日志列表对象
type ChessLogObject struct {
	id         int64
	serverId   int32
	playerName string
	itemId     int32
	itemNum    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewChessLogObject() *ChessLogObject {
	o := &ChessLogObject{}
	return o
}

func convertNewChessLogObjectToEntity(o *ChessLogObject) (*chessentity.ChessLogEntity, error) {
	e := &chessentity.ChessLogEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		PlayerName: o.playerName,
		ItemId:     o.itemId,
		ItemNum:    o.itemNum,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *ChessLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *ChessLogObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *ChessLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *ChessLogObject) GetItemId() int32 {
	return o.itemId
}

func (o *ChessLogObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *ChessLogObject) GetDBId() int64 {
	return o.id
}

func (o *ChessLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewChessLogObjectToEntity(o)
	return e, err
}

func (o *ChessLogObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chessentity.ChessLogEntity)

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.playerName = pse.PlayerName
	o.itemId = pse.ItemId
	o.itemNum = pse.ItemNum
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *ChessLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ChessLog"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
