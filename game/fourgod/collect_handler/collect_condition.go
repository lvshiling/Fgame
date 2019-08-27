package collect_handler

import (
	"fgame/fgame/common/lang"
	collectnpc "fgame/fgame/game/collect/npc"
	playerfourgod "fgame/fgame/game/fourgod/player"
	"fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	collectnpc.RegisterSceneCollectConditionHandler(scenetypes.SceneTypeFourGodWar, collectnpc.SceneCollectConditionHandlerFunc(collectCondition))
}

//四神遗迹 采集条件
func collectCondition(spl scene.Player, cn scene.CollectNPC) (flag bool) {
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	fourManager := pl.GetPlayerDataManager(playertypes.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	keyNum := fourManager.GetKeyNum()
	biologyId := int32(cn.GetBiologyTemplate().TemplateId())
	boxTemplate := template.GetFourGodTemplateService().GetFourGodBoxTemplateByBiologyId(biologyId)
	if boxTemplate == nil {
		return
	}

	if keyNum < boxTemplate.UseItemCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    biologyId,
				"keyNum":   keyNum,
			}).Warn("fourgod:钥匙数不够")
		keyNumStr := fmt.Sprintf("%d", boxTemplate.UseItemCount)
		playerlogic.SendSystemMessage(pl, lang.FourGodKeyNoEnough, keyNumStr)
		return
	}

	flag = true
	return
}
