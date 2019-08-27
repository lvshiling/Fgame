package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}

// 玩家加载
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	//同步状态
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieyi.GetJieYiService().PlayerLogin(pl.GetId())
	isJieYi := int32(0)
	name := ""
	daoJu := jieyitypes.JieYiDaoJuTypeLow
	token := jieyitypes.JieYiTokenTypeInvalid
	jieYiId := int64(0)
	obj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	rank := int32(0)
	if obj != nil {
		jieYiId = obj.GetJieYiId()
		isJieYi = 1
		daoJu = obj.GetDaoJuType()
		token = obj.GetTokenType()
		rank = obj.GetRank()
		name = obj.GetJieYi().GetJieYiName()
		jieyiObj := obj.GetJieYi().GetJieYiObject()

		// 推送结义成员信息
		memberObjList := jieyi.GetJieYiService().GetJieYiMemberList(jieYiId)
		scMsg := pbutil.BuildSCJieYiMemberInfo(jieyiObj, memberObjList, pl.GetId(), jieYiManager.GetPlayerJieYiObj().GetTokenPro(), jieYiManager.GetPlayerJieYiObj().GetShengWeiZhi())
		pl.SendMsg(scMsg)
	}

	jieYiManager.SyncJieYi(daoJu, token, jieYiId, name, rank)
	scMsg := pbutil.BuildSCJieYiPlayerInfo(isJieYi, name)
	pl.SendMsg(scMsg)

	return
}
