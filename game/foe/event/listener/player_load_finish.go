package listener

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/foe/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	// 玩家的仇人列表
	if err = sendFoeList(p); err != nil {
		return
	}

	// 仇人反馈信息
	if err = sendFoeFeedbackInfo(p); err != nil {
		return
	}

	return
}

func sendFoeList(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	//仇人列表信息
	var foeList []*uipb.FoeInfo
	for _, foeObj := range manager.GetFoeMap() {
		attackId := foeObj.AttackId
		killTime := foeObj.KillTime
		foeInfo, err := player.GetPlayerService().GetPlayerInfo(attackId)
		if err != nil {
			return err
		}
		if foeInfo == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"attackId": attackId,
				}).Info("foe:仇人")
			continue
		}
		tempFoe := pbutil.BuildFoe(attackId, killTime, foeInfo)
		foeList = append(foeList, tempFoe)
	}

	scFoesGet := pbutil.BuildSCFoesGet(foeList)
	pl.SendMsg(scFoesGet)
	return
}

func sendFoeFeedbackInfo(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	expireTime := manager.GetFoeFeedbackProtectExpireTime()
	feedbackList := manager.GetFoeFeedbackList()

	scMsg := pbutil.BuildSCFoeFeedbackInfo(expireTime, feedbackList)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
