package logic

import (
	"fgame/fgame/common/lang"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/myboss/pbutil"
	mybosstemplate "fgame/fgame/game/myboss/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func PlayerEnterBoss(pl player.Player, biologyId int32) (flag bool) {
	mybossTemplate := mybosstemplate.GetMyBossTemplateService().GetMyBossTemplate(biologyId)
	if mybossTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("myboss:处理跳转个人BOSS,boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := createMyBossSceneData(pl.GetId(), mybossTemplate.BiologyId)
	s := scene.CreateFuBenScene(mybossTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("myboss:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)

	flag = true
	return
}

//挑战结束
func onMyBossFinish(p player.Player, itemList []*droptemplate.DropItemData, isSuccess bool) {
	scMsg := pbutil.BuildSCMyBossChallengeResult(isSuccess, itemList)
	p.SendMsg(scMsg)
}

//下发场景信息
func onPushSceneInfo(p player.Player, startTime int64, bossId int32) {
	scMsg := pbutil.BuildSCMyBossSceneInfo(startTime, bossId)
	p.SendMsg(scMsg)
}
