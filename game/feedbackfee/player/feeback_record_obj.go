package player

import (
	"fgame/fgame/core/storage"
	feedbackfeeentity "fgame/fgame/game/feedbackfee/entity"
	feedbackfeetypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//逆付费
type PlayerFeedbackRecordObject struct {
	player      player.Player
	id          int64
	money       int32
	code        string
	status      feedbackfeetypes.FeedbackStatus
	typ         feedbackfeetypes.FeedbackFeeType
	expiredTime int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func newPlayerFeedbackRecordObject(pl player.Player) *PlayerFeedbackRecordObject {
	o := &PlayerFeedbackRecordObject{
		player: pl,
	}
	return o
}

func convertPlayerFeedbackRecordObjectToEntity(o *PlayerFeedbackRecordObject) (e *feedbackfeeentity.PlayerFeedbackRecordEntity, err error) {
	e = &feedbackfeeentity.PlayerFeedbackRecordEntity{
		Id:          o.id,
		PlayerId:    o.player.GetId(),
		Money:       o.money,
		Code:        o.code,
		Status:      int32(o.status),
		Type:        int32(o.typ),
		ExpiredTime: o.expiredTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}

	return e, nil
}

func (o *PlayerFeedbackRecordObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFeedbackRecordObject) GetId() int64 {
	return o.id
}

func (o *PlayerFeedbackRecordObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFeedbackRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFeedbackRecordObjectToEntity(o)
	return e, err
}

func (o *PlayerFeedbackRecordObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*feedbackfeeentity.PlayerFeedbackRecordEntity)

	o.id = te.Id
	o.money = te.Money
	o.code = te.Code
	o.status = feedbackfeetypes.FeedbackStatus(te.Status)
	o.typ = feedbackfeetypes.FeedbackFeeType(te.Type)
	o.expiredTime = te.ExpiredTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFeedbackRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FeedbackRecord"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerFeedbackRecordObject) GetStatus() feedbackfeetypes.FeedbackStatus {
	return o.status
}

func (o *PlayerFeedbackRecordObject) GetCode() string {
	return o.code
}

func (o *PlayerFeedbackRecordObject) GetMoney() int32 {
	return o.money
}

func (o *PlayerFeedbackRecordObject) GetExpiredTime() int64 {
	return o.expiredTime
}

func (o *PlayerFeedbackRecordObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerFeedbackRecordObject) Expire() bool {
	if o.status != feedbackfeetypes.FeedbackStatusProcess {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feedbackfeetypes.FeedbackStatusFailed
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *PlayerFeedbackRecordObject) Code(code string) bool {
	if o.status != feedbackfeetypes.FeedbackStatusInit {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feedbackfeetypes.FeedbackStatusProcess
	o.code = code
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *PlayerFeedbackRecordObject) Finish() bool {
	if o.status != feedbackfeetypes.FeedbackStatusProcess {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feedbackfeetypes.FeedbackStatusFinish
	o.updateTime = now
	o.SetModified()
	return true
}
