package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	"fgame/fgame/game/trade/trade"
)

//上传返还
func tradeUploadRefund(target event.EventTarget, data event.EventData) (err error) {
	refundTradeItemObj, ok := target.(*trade.TradeItemObject)
	if !ok {
		return
	}
	playerId := refundTradeItemObj.GetPlayerId()
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	now := global.GetGame().GetTimeService().Now()
	refundTitle := lang.GetLangService().ReadLang(lang.TradeUploadRefundTitle)
	refundContent := lang.GetLangService().ReadLang(lang.TradeUploadRefundContent)
	if p == nil {
		dropItemData := inventorylogic.ConverToItemData(refundTradeItemObj.GetItemId(), refundTradeItemObj.GetNum(), 0, itemtypes.ItemBindTypeUnBind, refundTradeItemObj.GetPropertyData())
		emaillogic.AddOfflineEmailItemLevel(playerId, refundTitle, refundContent, now, []*droptemplate.DropItemData{dropItemData})
	} else {
		ctx := scene.WithPlayer(context.Background(), p)
		playerTradeUploadRefund := message.NewScheduleMessage(onPlayerTradeUploadRefund, ctx, refundTradeItemObj, nil)
		p.Post(playerTradeUploadRefund)
	}
	return
}

func onPlayerTradeUploadRefund(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	refundTradeItemObj := result.(*trade.TradeItemObject)
	//TODO:zrc 判断是否足够物品
	refundTitle := lang.GetLangService().ReadLang(lang.TradeUploadRefundTitle)
	refundContent := lang.GetLangService().ReadLang(lang.TradeUploadRefundContent)
	now := global.GetGame().GetTimeService().Now()
	dropItemData := inventorylogic.ConverToItemData(refundTradeItemObj.GetItemId(), refundTradeItemObj.GetNum(), refundTradeItemObj.GetLevel(), itemtypes.ItemBindTypeUnBind, refundTradeItemObj.GetPropertyData())
	emaillogic.AddEmailItemLevel(pl, refundTitle, refundContent, now, []*droptemplate.DropItemData{dropItemData})
	return nil
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeUploadRefund, event.EventListenerFunc(tradeUploadRefund))
}
