package battle

import (
	"fgame/fgame/game/alliance/alliance"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterRelivePointHandler(scenetypes.SceneTypeChengZhan, scene.RelivePointHandlerFunc(RelivePoint))
}

//复活点复活
func RelivePoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return false
	}
	//清空原地复活次数
	tpl, ok := pl.(player.Player)
	if !ok {
		return false
	}
	manager := tpl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	manager.RestReliveTime()
	//获取城战数据
	allianceSceneData := alliance.GetAllianceService().GetAllianceSceneData()
	if allianceSceneData == nil {
		return
	}
	as := allianceSceneData.GetScene()
	//获取当前联盟
	currentReliveAllianceId := allianceSceneData.GetCurrentReliveAllianceId()

	reliveFlag := allianceSceneData.GetReliveFlag()
	if reliveFlag == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:处理对象复活消息,没有设置复活点")
		return false
	}

	allianceId := pl.GetAllianceId()
	if allianceId == currentReliveAllianceId {
		//获取场景的复活点
		rebornPos := reliveFlag.GetPosition()
		//同一个场景
		if s == as {
			pl.Reborn(rebornPos)
			return true
		} else {
			scenelogic.PlayerEnterScene(pl, as, rebornPos)
			return true
		}
	} else {
		rebornPos := as.MapTemplate().GetRebornPos()
		//同一个场景
		if s == as {
			pl.Reborn(rebornPos)
			return true
		} else {
			scenelogic.PlayerEnterScene(pl, as, rebornPos)
			return true
		}
		return true
	}
	return true
}
