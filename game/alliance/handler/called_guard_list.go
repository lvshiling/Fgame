package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_CALLED_GUARD_LIST_TYPE), dispatch.HandlerFunc(handleCalledGuardList))
}

//处理召唤守卫列表
func handleCalledGuardList(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理召唤守卫列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceCalledGuardList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理召唤守卫列表,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理召唤守卫列表,完成")
	return nil

}

//仙盟
func allianceCalledGuardList(pl player.Player) (err error) {
	s := pl.GetScene()
	sd := s.SceneDelegate()
	//switch ssd := sd.(type) {
	switch sd.(type) {
	case alliancescene.AllianceSceneData:
		callGuardList := alliance.GetAllianceService().GetAllianceSceneData().GetCallGuardList()
		//callGuardList := ssd.GetCallGuardList()
		scAllianceSceneCalledGuardList := pbutil.BuildSCAllianceSceneCalledGuardList(callGuardList)
		pl.SendMsg(scAllianceSceneCalledGuardList)
		break
	default:
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理仙盟召唤守卫,不在城战")
		playerlogic.SendSystemMessage(pl, lang.AllianceSceneNotInAllianceScene)

		break
	}

	return
}
