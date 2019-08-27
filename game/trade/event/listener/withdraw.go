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
	"fgame/fgame/game/trade/pbutil"
	"fgame/fgame/game/trade/trade"
)

//撤销
func tradeWithdraw(target event.EventTarget, data event.EventData) (err error) {
	refundTradeItemObj, ok := target.(*trade.TradeItemObject)
	if !ok {
		return
	}
	playerId := refundTradeItemObj.GetPlayerId()
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	now := global.GetGame().GetTimeService().Now()
	withdrawTitle := lang.GetLangService().ReadLang(lang.TradeUploadWithdrawTitle)
	withdrawContent := lang.GetLangService().ReadLang(lang.TradeUploadWithdrawContent)
	if refundTradeItemObj.IsSystem() {
		withdrawContent = lang.GetLangService().ReadLang(lang.TradeUploadWithdrawBySystemContent)
	}
	if p == nil {
		dropItemData := inventorylogic.ConverToItemData(refundTradeItemObj.GetItemId(), refundTradeItemObj.GetNum(), refundTradeItemObj.GetLevel(), itemtypes.ItemBindTypeUnBind, refundTradeItemObj.GetPropertyData())
		emaillogic.AddOfflineEmailItemLevel(playerId, withdrawTitle, withdrawContent, now, []*droptemplate.DropItemData{dropItemData})
	} else {
		//发送消息
		scTradeWithDrawItem := pbutil.BuildSCTradeWithDrawItem(refundTradeItemObj.GetId())
		p.SendMsg(scTradeWithDrawItem)
		ctx := scene.WithPlayer(context.Background(), p)
		playerTradeWithdraw := message.NewScheduleMessage(onPlayerTradeWithdraw, ctx, refundTradeItemObj, nil)
		p.Post(playerTradeWithdraw)
	}
	return
}

func onPlayerTradeWithdraw(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	refundTradeItemObj := result.(*trade.TradeItemObject)
	withdrawContent := lang.GetLangService().ReadLang(lang.TradeUploadWithdrawContent)
	if refundTradeItemObj.IsSystem() {
		withdrawContent = lang.GetLangService().ReadLang(lang.TradeUploadWithdrawBySystemContent)
	}
	withdrawTitle := lang.GetLangService().ReadLang(lang.TradeUploadWithdrawTitle)

	now := global.GetGame().GetTimeService().Now()
	dropItemData := inventorylogic.ConverToItemData(refundTradeItemObj.GetItemId(), refundTradeItemObj.GetNum(), refundTradeItemObj.GetLevel(), itemtypes.ItemBindTypeUnBind, refundTradeItemObj.GetPropertyData())
	emaillogic.AddEmailItemLevel(pl, withdrawTitle, withdrawContent, now, []*droptemplate.DropItemData{dropItemData})
	return nil
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeWithdraw, event.EventListenerFunc(tradeWithdraw))
}
