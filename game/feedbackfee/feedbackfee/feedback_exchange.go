package feedbackfee

import (
	"fgame/fgame/core/storage"
	feedbackfeeentity "fgame/fgame/game/feedbackfee/entity"
	feebackfeetypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//逆付费
type FeedbackExchangeObject struct {
	id          int64
	serverId    int32
	playerId    int64
	exchangeId  int64
	code        string
	money       int32
	status      feebackfeetypes.FeedbackExchangeStatus
	expiredTime int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func newFeedbackExchangeObject() *FeedbackExchangeObject {
	o := &FeedbackExchangeObject{}
	return o
}

func convertFeedbackExchangeObjectToEntity(o *FeedbackExchangeObject) (e *feedbackfeeentity.FeedbackExchangeEntity, err error) {
	e = &feedbackfeeentity.FeedbackExchangeEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		PlayerId:    o.playerId,
		ExchangeId:  o.exchangeId,
		Code:        o.code,
		Status:      int32(o.status),
		Money:       o.money,
		ExpiredTime: o.expiredTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}

	return e, nil
}

func (o *FeedbackExchangeObject) GetId() int64 {
	return o.id
}
func (o *FeedbackExchangeObject) GetDBId() int64 {
	return o.id
}

func (o *FeedbackExchangeObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *FeedbackExchangeObject) GetExchangeId() int64 {
	return o.exchangeId
}

func (o *FeedbackExchangeObject) GetCode() string {
	return o.code
}

func (o *FeedbackExchangeObject) GetStatus() feebackfeetypes.FeedbackExchangeStatus {
	return o.status
}

func (o *FeedbackExchangeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertFeedbackExchangeObjectToEntity(o)
	return e, err
}

func (o *FeedbackExchangeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*feedbackfeeentity.FeedbackExchangeEntity)

	o.id = te.Id
	o.serverId = te.ServerId
	o.playerId = te.PlayerId
	o.exchangeId = te.ExchangeId
	o.expiredTime = te.ExpiredTime
	o.money = te.Money
	o.code = te.Code
	o.status = feebackfeetypes.FeedbackExchangeStatus(te.Status)
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *FeedbackExchangeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FeedbackExchange"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

func (o *FeedbackExchangeObject) CodeGenerate(code string) bool {
	if o.status != feebackfeetypes.FeedbackExchangeStatusInit {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feebackfeetypes.FeedbackExchangeStatusGenerateCode
	o.updateTime = now
	o.code = code
	o.SetModified()
	return true
}

func (o *FeedbackExchangeObject) FillCode() bool {
	if o.status != feebackfeetypes.FeedbackExchangeStatusGenerateCode {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feebackfeetypes.FeedbackExchangeStatusProcess
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *FeedbackExchangeObject) Expired() bool {
	if o.status != feebackfeetypes.FeedbackExchangeStatusProcess {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feebackfeetypes.FeedbackExchangeStatusFailed
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *FeedbackExchangeObject) Finish() bool {
	if o.status != feebackfeetypes.FeedbackExchangeStatusProcess {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feebackfeetypes.FeedbackExchangeStatusFinish
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *FeedbackExchangeObject) Notify() bool {
	if o.status != feebackfeetypes.FeedbackExchangeStatusFailed && o.status != feebackfeetypes.FeedbackExchangeStatusFinish {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	o.status = feebackfeetypes.FeedbackExchangeStatusNotify
	o.updateTime = now
	o.SetModified()
	return true
}
