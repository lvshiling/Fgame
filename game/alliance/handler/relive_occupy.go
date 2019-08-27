package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	alliancescene "fgame/fgame/game/alliance/scene"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_RELIVE_OCCUPYING_TYPE), dispatch.HandlerFunc(handleAllianceSceneReliveOccupying))
}

//复活点占领
func handleAllianceSceneReliveOccupying(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理复活点占领")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = allianceSceneReliveOccupying(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理复活点占领,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理复活点占领,完成")
	return nil

}

//处理复活点占领
func allianceSceneReliveOccupying(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}
	sceneData, ok := sd.(alliancescene.AllianceSceneData)
	if !ok {
		return
	}

	allianceId := pl.GetAllianceId()

	if sceneData.GetCurrentReliveAllianceId() == allianceId {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理复活点占领,复活点已经属于自己联盟")
		return
	}

	if sceneData.GetCollectReliveAllianceId() != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理复活点占领,复活点自己联盟正在占领")
		return
	}

	//判断采集距离
	reliveFlag := sceneData.GetReliveFlag()
	distance := coreutils.Distance(reliveFlag.GetPosition(), pl.GetPos())
	collectDistance := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectDistance)) / float64(1000)
	if distance > float64(collectDistance) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("alliance:不在采集范围内")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectNoDistance)
		return
	}

	flag := sceneData.ReliveOccupy(allianceId, pl.GetId())
	if !flag {
		panic("alliance:占领复活点应该成功")
	}
	return
}
