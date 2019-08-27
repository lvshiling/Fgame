package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FUBEN_EXIT_TYPE), dispatch.HandlerFunc(handleFuBenExit))
}

//处理副本退出
func handleFuBenExit(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理副本退出场景")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	//TODO 判断
	err = fuBenExit(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("scene:处理副本退出场景,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("scene:处理副本退出场景,完成")

	return nil
}

//玩家退出副本
func fuBenExit(pl player.Player) (err error) {
	//TODO 判断是否在副本
	//判断是否在跨服

	if pl.GetScene() == nil {
		return
	}
	if pl.GetScene().MapTemplate().GetMapType().MapType() == scenetypes.MapTypeWorld {
		return
	}

	if pl.IsPvpBattle() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:处理副本退出场景,完成")
		playerlogic.SendSystemMessage(pl, lang.ScenePlayerBattleStatus)
		return
	}

	scenelogic.PlayerBackLastScene(pl)

	scFuBenExit := pbutil.BuildSCFuBenExit()
	err = pl.SendMsg(scFuBenExit)
	return
}
