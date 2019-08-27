package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/coupon/coupon"
	couponeventtypes "fgame/fgame/game/coupon/event/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/welfare/pbutil"

	log "github.com/Sirupsen/logrus"
)

//兑换成功
func exchangeFinish(target event.EventTarget, data event.EventData) (err error) {
	playerId := target.(int64)
	eventData := data.(*coupon.ExchangeFinishEventData)
	//验证附件格式
	attachment, err := droplogic.ParseAttachmentList(eventData.GetAttachment())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"err":      err,
			}).Warn("coupon:兑换兑换码,格式错误")
		err = nil
		return
	}
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		//TODO 发送离线
		now := global.GetGame().GetTimeService().Now()
		emaillogic.AddOfflineEmailItemLevel(playerId, eventData.GetTitle(), eventData.GetContent(), now, attachment)
		return
	}

	//发送在线
	ctx := scene.WithPlayer(context.Background(), p)
	r := &exchangeResult{
		title:      eventData.GetTitle(),
		content:    eventData.GetContent(),
		attachment: attachment,
	}
	msg := message.NewScheduleMessage(onCouponExchange, ctx, r, nil)
	p.Post(msg)
	return
}

type exchangeResult struct {
	title      string
	content    string
	attachment []*droptemplate.DropItemData
}

//兑换成功
func onCouponExchange(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	r := result.(*exchangeResult)
	now := global.GetGame().GetTimeService().Now()

	emaillogic.AddEmailItemLevel(pl, r.title, r.content, now, r.attachment)
	totalRewData := &propertytypes.RewData{}
	itemMap := make(map[int32]int32)
	for _, dropItem := range r.attachment {
		itemMap[dropItem.ItemId] = dropItem.Num
	}
	scOpenActivityGiftCode := pbutil.BuildSCOpenActivityGiftCode(totalRewData, itemMap)
	pl.SendMsg(scOpenActivityGiftCode)
	return nil
}

func init() {
	gameevent.AddEventListener(couponeventtypes.CouponEventTypeExchangeFinish, event.EventListenerFunc(exchangeFinish))
}
