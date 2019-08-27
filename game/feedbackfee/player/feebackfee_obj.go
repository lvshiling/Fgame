package player

import (
	"fgame/fgame/core/storage"
	feedbackfeeentity "fgame/fgame/game/feedbackfee/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//逆付费
type PlayerFeedbackFeeObject struct {
	player        player.Player
	id            int64
	totalGetMoney int64
	money         int32
	todayUseNum   int32
	useTime       int64
	cashMoney     int64
	goldMoney     int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func newPlayerFeedbackFeeObject(pl player.Player) *PlayerFeedbackFeeObject {
	o := &PlayerFeedbackFeeObject{
		player: pl,
	}
	return o
}

func convertPlayerFeedbackFeeObjectToEntity(o *PlayerFeedbackFeeObject) (e *feedbackfeeentity.PlayerFeedbackFeeEntity, err error) {
	e = &feedbackfeeentity.PlayerFeedbackFeeEntity{
		Id:          o.id,
		PlayerId:    o.player.GetId(),
		Money:       o.money,
		TodayUseNum: o.todayUseNum,
		UseTime:     o.useTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}

	return e, nil
}

func (o *PlayerFeedbackFeeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFeedbackFeeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFeedbackFeeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFeedbackFeeObjectToEntity(o)
	return e, err
}

func (o *PlayerFeedbackFeeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*feedbackfeeentity.PlayerFeedbackFeeEntity)

	o.id = te.Id
	o.money = te.Money
	o.todayUseNum = te.TodayUseNum
	o.useTime = te.UseTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFeedbackFeeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FeedbackFee"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerFeedbackFeeObject) GetMoney() int32 {
	return o.money
}

func (o *PlayerFeedbackFeeObject) GetTodayUseNum() int32 {
	return o.todayUseNum
}
