package listener

import (
	"fgame/fgame/core/event"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	teameventtypes "fgame/fgame/game/team/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家准备决策
func teamMatchCondtionPrepareDeal(target event.EventTarget, data event.EventData) (err error) {
	teamObj, ok := target.(*team.TeamObject)
	if !ok {
		return
	}
	eventData, ok := data.(*teameventtypes.TeamMatchCondtionPrepareDealEventData)
	if !ok {
		return
	}
	result := eventData.GetResult()
	pl := eventData.GetPlayer()
	dealResult := result
	//在跨服
	s := pl.GetScene()

	noNeedExitScene := false
	if s != nil {
		//再次校验 玩家是否符合准备
		if !result {
			if s.MapTemplate().IsWorld() {
				result = true
				noNeedExitScene = true
			}
		}
	}

	scTeamMatchCondtionPrepareDeal := pbutil.BuildSCTeamMatchCondtionPrepareDeal(dealResult)
	pl.SendMsg(scTeamMatchCondtionPrepareDeal)

	if result {
		scTeamMatchCondtionPrepareBroadcast := pbutil.BuildSCTeamMatchCondtionPrepareBroadcast(pl.GetId())
		teamlogic.BroadcastMsg(teamObj, scTeamMatchCondtionPrepareBroadcast)
	}

	//临时处理，以防死锁
	if result && !noNeedExitScene {
		if pl.IsCross() {
			crosslogic.AsyncPlayerExitCross(pl)
		} else {
			scenelogic.AsyncPlayerBackLastScene(pl)
		}
	}
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamMatchCondtionPrepareDeal, event.EventListenerFunc(teamMatchCondtionPrepareDeal))
}
