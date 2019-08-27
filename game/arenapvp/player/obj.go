package player

import (
	"fgame/fgame/core/storage"
	arenapvpentity "fgame/fgame/game/arenapvp/entity"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//竞技场对象
type PlayerArenapvpObject struct {
	player      player.Player
	id          int64
	reliveTimes int32
	outStatus   int32
	jiFen       int32
	guessNotice int32
	pvpRecord   arenapvptypes.ArenapvpType
	ticketFlag  int32
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func (o *PlayerArenapvpObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerArenapvpObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerArenapvpObject) FromEntity(e storage.Entity) error {
	te := e.(*arenapvpentity.PlayerArenapvpEntity)
	o.id = te.Id
	o.reliveTimes = te.ReliveTimes
	o.outStatus = te.OutStatus
	o.jiFen = te.JiFen
	o.guessNotice = te.GuessNotice
	o.pvpRecord = arenapvptypes.ArenapvpType(te.PvpRecord)
	o.ticketFlag = te.TicketFlag
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}
func (o *PlayerArenapvpObject) ToEntity() (e storage.Entity, err error) {

	e = &arenapvpentity.PlayerArenapvpEntity{
		Id:          o.id,
		PlayerId:    o.player.GetId(),
		ReliveTimes: o.reliveTimes,
		OutStatus:   o.outStatus,
		JiFen:       o.jiFen,
		GuessNotice: o.guessNotice,
		PvpRecord:   int32(o.pvpRecord),
		TicketFlag:  o.ticketFlag,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *PlayerArenapvpObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Arenapvp"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
}

func (o *PlayerArenapvpObject) GetReliveTimes() int32 {
	return o.reliveTimes
}

func (o *PlayerArenapvpObject) GetOutStatus() int32 {
	return o.outStatus
}

func (o *PlayerArenapvpObject) GetTicketFlag() int32 {
	return o.ticketFlag
}

func (o *PlayerArenapvpObject) GetJiFen() int32 {
	return o.jiFen
}

func (o *PlayerArenapvpObject) GetGuessNotice() int32 {
	return o.guessNotice
}

func (o *PlayerArenapvpObject) GetPvpRecord() arenapvptypes.ArenapvpType {
	return o.pvpRecord
}

func NewPlayerArenapvpObject(pl player.Player) *PlayerArenapvpObject {
	o := &PlayerArenapvpObject{
		player: pl,
	}
	return o
}

func (o *PlayerArenapvpObject) IfEnoughJiFen(num int32) bool {
	if num < 0 {
		return false
	}

	if o.jiFen >= num {
		return true
	}

	return false
}

func (o *PlayerArenapvpObject) IfGuessNotice() bool {
	return o.guessNotice != 0
}

func (o *PlayerArenapvpObject) IfBuyTicket() bool {
	return o.ticketFlag != 0
}

func (o *PlayerArenapvpObject) setPvpRecord(win bool, pvpType arenapvptypes.ArenapvpType) {
	if win {
		if pvpType == arenapvptypes.ArenapvpTypeFinals {
			o.pvpRecord = arenapvptypes.ArenapvpTypeChampion
		} else {
			arenapvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(pvpType)
			o.pvpRecord = arenapvpTemp.GetNextTemp().GetArenapvpType()
		}
		return
	}

	o.pvpRecord = pvpType
}

//竞技场竞猜日志对象
type PlayerArenapvpGuessLogObject struct {
	player     player.Player
	id         int64
	raceNum    int32
	guessType  arenapvptypes.ArenapvpType
	guessId    int64
	winnerId   int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func (o *PlayerArenapvpGuessLogObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerArenapvpGuessLogObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerArenapvpGuessLogObject) FromEntity(e storage.Entity) error {
	te := e.(*arenapvpentity.PlayerArenapvpGuessLogEntity)
	o.id = te.Id
	o.raceNum = te.RaceNum
	o.guessId = te.GuessId
	o.winnerId = te.WinnerId
	o.guessType = arenapvptypes.ArenapvpType(te.GuessType)
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}
func (o *PlayerArenapvpGuessLogObject) ToEntity() (e storage.Entity, err error) {

	e = &arenapvpentity.PlayerArenapvpGuessLogEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		RaceNum:    o.raceNum,
		WinnerId:   o.winnerId,
		GuessType:  int32(o.guessType),
		GuessId:    o.guessId,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerArenapvpGuessLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ArenapvpGuessLog"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
}

func NewPlayerArenapvpGuessLogObject(pl player.Player) *PlayerArenapvpGuessLogObject {
	o := &PlayerArenapvpGuessLogObject{
		player: pl,
	}
	return o
}

func (o *PlayerArenapvpGuessLogObject) GetGuessId() int64 {
	return o.guessId
}

func (o *PlayerArenapvpGuessLogObject) GetRaceNum() int32 {
	return o.raceNum
}

func (o *PlayerArenapvpGuessLogObject) GetWinnerId() int64 {
	return o.winnerId
}

func (o *PlayerArenapvpGuessLogObject) GetGuessType() arenapvptypes.ArenapvpType {
	return o.guessType
}

func (o *PlayerArenapvpGuessLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerArenapvpGuessLogObject) IfAttendGuess(raceNum int32, guessType arenapvptypes.ArenapvpType) bool {
	if o.raceNum != raceNum {
		return false
	}

	if o.guessType != guessType {
		return false
	}

	return true
}
