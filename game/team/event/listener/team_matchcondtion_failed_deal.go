package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	teameventtypes "fgame/fgame/game/team/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	playerteam "fgame/fgame/game/team/player"
	"fgame/fgame/game/team/team"
)

//队伍匹配条件不足队长决策
func teamMatchCondtionFailedDeal(target event.EventTarget, data event.EventData) (err error) {
	teamObj, ok := target.(*team.TeamObject)
	if !ok {
		return
	}
	eventData, ok := data.(*teameventtypes.TeamMatchCondtionFailedDealEventData)
	if !ok {
		return
	}
	result := eventData.GetResult()
	memberIdList := eventData.GetMemberIdList()
	if len(memberIdList) == 0 {
		return
	}

	captainId := teamObj.GetCaptain().GetPlayerId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(captainId)
	if pl == nil {
		return
	}
	scTeamMatchCondtionFailedDeal := pbutil.BuildSCTeamMatchCondtionFailedDeal(result)
	pl.SendMsg(scTeamMatchCondtionFailedDeal)

	if result {
		for _, playerId := range memberIdList {
			spl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
			if spl == nil {
				continue
			}

			ctx := scene.WithPlayer(context.Background(), spl)
			spl.Post(message.NewScheduleMessage(onTeamMatchCondtionFailedToMember, ctx, nil, nil))
		}
	}
	scTeamMatchCondtionFailedBroadcast := pbutil.BuildSCTeamMatchCondtionFailedBroadcast(memberIdList)
	teamlogic.BroadcastMsg(teamObj, scTeamMatchCondtionFailedBroadcast)
	return nil
}

func onTeamMatchCondtionFailedToMember(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	teamId := pl.GetTeamId()
	if teamId == 0 {
		return nil
	}

	if pl.IsCross() {

	} else {
		//再过滤一遍
		s := pl.GetScene()
		if s == nil {
			return nil
		}
		if !pl.IsCross() && !s.MapTemplate().IsFuBen() {
			return nil
		}
	}
	mananger := pl.GetPlayerDataManager(types.PlayerTeamDataManagerType).(*playerteam.PlayerTeamDataManager)
	mananger.SetMatchCondtionFailed()

	//推送
	scTeamMatchCondtionFailedToPrepare := pbutil.BuildSCTeamMatchCondtionFailedToPrepare()
	pl.SendMsg(scTeamMatchCondtionFailedToPrepare)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamMatchCondtionFailedDeal, event.EventListenerFunc(teamMatchCondtionFailedDeal))
}
