package listener

import (
	"fgame/fgame/core/event"
	playeractivity "fgame/fgame/game/activity/player"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家进入九霄
func playerEnterAllianceScene(target event.EventTarget, data event.EventData) (err error) {
	subSceneData := target.(alliancescene.AllianceSceneData)
	p := data.(player.Player)
	allianceSceneData := alliance.GetAllianceService().GetAllianceSceneData()
	if allianceSceneData == nil {
		return
	}

	endTime := allianceSceneData.GetEndTime()
	initDefendAllianceId := allianceSceneData.GetFirstDefendAllianceId()
	currentDefendAllianceId := subSceneData.GetCurrentDefendAllianceId()
	currentDefendAllianceName := subSceneData.GetCurrentDefendAllianceName()
	currentDefendAllianceHuFu := subSceneData.GetCurrentDefendAllianceHuFu()
	currentDoor := subSceneData.GetCurrentDoor()
	currentReliveAllianceId := subSceneData.GetCurrentReliveAllianceId()
	tp := p.(player.Player)
	allianceManager := tp.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	//刷新
	allianceManager.RefreshAllianceScene(endTime)
	rewardList := allianceManager.GetRewardList()
	warPoint := allianceManager.GetWarPoint()
	collectReliveStartTime := subSceneData.GetCollectReliveStartTime()
	collectReliveAllianceId := subSceneData.GetCollectReliveAllianceId()
	yuXiNpc := subSceneData.GetCollectYuXi()
	scAllianceSceneInfo := pbutil.BuildSCAllianceSceneInfo(initDefendAllianceId,
		currentDefendAllianceId,
		currentDefendAllianceName,
		currentDefendAllianceHuFu,
		currentDoor,
		endTime,
		rewardList,
		currentReliveAllianceId,
		collectReliveAllianceId,
		collectReliveStartTime,
		warPoint,
		yuXiNpc)
	p.SendMsg(scAllianceSceneInfo)

	//更新进入时间
	activityMananger := tp.GetPlayerDataManager(playertypes.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityMananger.UpdateEnterTime(activitytypes.ActivityTypeAlliance, endTime)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerEnterAllianceScene, event.EventListenerFunc(playerEnterAllianceScene))
}
