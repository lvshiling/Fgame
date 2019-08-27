package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/center/center"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/huiyuan/pbutil"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

// 玩家加载后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	//下发会员信息
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	err = huiyuanManager.RefresHuiYuanRewards()
	if err != nil {
		return
	}

	isHuiyuanInterim := huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	isReceiveInterim := huiyuanManager.IsReceiveRewards(huiyuantypes.HuiYuanTypeInterim)
	expireTime := huiyuanManager.GetHuiYuanExpireTiem()
	isBuyTodayInterim := huiyuanManager.IsFirstRew(huiyuantypes.HuiYuanTypeInterim)

	isReceivePlus := huiyuanManager.IsReceiveRewards(huiyuantypes.HuiYuanTypePlus)
	isHuiyuanPlus := huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	isBuyTodayPlus := huiyuanManager.IsFirstRew(huiyuantypes.HuiYuanTypePlus)

	houtaiType := center.GetCenterService().GetZhiZunType()
	scHuiYuanInfo := pbutil.BuildSCHuiYuanInfo(isHuiyuanInterim, isReceiveInterim, isReceivePlus, isHuiyuanPlus, isBuyTodayInterim, isBuyTodayPlus, expireTime, houtaiType)
	pl.SendMsg(scHuiYuanInfo)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
