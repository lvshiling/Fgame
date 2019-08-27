package activity_handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/activity/activity"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancescene "fgame/fgame/game/alliance/scene"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeAlliance, activity.ActivityAttendHandlerFunc(playerEnterAllianceScene))
}

func playerEnterAllianceScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("alliance:九霄城战,玩家不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAllianceForHegemon)
		return
	}

	sd := alliance.GetAllianceService().GetAllianceSceneData()
	if sd == nil {
		act := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeAlliance)
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		timeTemplate, err := act.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			return false, err
		}
		if timeTemplate == nil {
			log.WithFields(
				log.Fields{
					"id": pl.GetId(),
				}).Warn("alliance:九霄城战,城战未开始")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			return false, nil
		}
		endTime, err := timeTemplate.GetEndTime(now)
		if err != nil {
			return false, err
		}

		sd = alliance.GetAllianceService().CreateAllianceSceneData(act.Mapid, endTime, 0)
		if sd == nil {
			log.WithFields(
				log.Fields{
					"id": pl.GetId(),
				}).Warn("alliance:九霄城战,未开启")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			return false, nil
		}
	}

	manager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	manager.RestReliveTime()

	chengWaiScene := sd.GetScene()
	sceneData := chengWaiScene.SceneDelegate().(alliancescene.AllianceSceneData)
	currentReliveAllianceId := sceneData.GetCurrentReliveAllianceId()
	reliveFlag := sceneData.GetReliveFlag()
	pos := chengWaiScene.MapTemplate().GetBornPos()
	if allianceId == currentReliveAllianceId {
		pos = reliveFlag.GetPosition()
	}

	if !scenelogic.PlayerEnterScene(pl, chengWaiScene, pos) {
		return
	}
	flag = true
	return
}
