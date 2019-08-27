package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arena/pbutil"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_COLLECT_EXP_TREE_TYPE), dispatch.HandlerFunc(handleArenaCollectExpTree))
}

//处理获得经验树
func handleArenaCollectExpTree(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理获得经验树")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaCollectExpTree(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理获得经验树,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理获得经验树,完成")
	return nil

}

//3v3匹配
func arenaCollectExpTree(pl player.Player) (err error) {
	exp := int64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().TreeExp)
	expPoint := int64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().TreeExpPoint)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if expPoint > 0 {
		propertyManager.AddExpPoint(expPoint, commonlog.LevelLogReasonArenaExpTree, commonlog.LevelLogReasonArenaExpTree.String())
	}
	if exp > 0 {
		propertyManager.AddExp(exp, commonlog.LevelLogReasonArenaExpTree, commonlog.LevelLogReasonArenaExpTree.String())
	}
	propertylogic.SnapChangedProperty(pl)

	siArenaCollectExpTree := pbutil.BuildSIArenaCollectExpTree()
	pl.SendMsg(siArenaCollectExpTree)
	return
}
