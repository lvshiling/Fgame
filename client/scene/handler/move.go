package handler

import (
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	"fgame/fgame/client/scene/pbutil"
	playerscene "fgame/fgame/client/scene/player"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_SC_OBJECT_MOVE_TYPE), dispatch.HandlerFunc(handleMove))
}

func handleMove(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:移动")
	scObjectMove := msg.(*scenepb.SCObjectMove)
	moveData := scObjectMove.GetMoveData()
	uId := moveData.GetUid()
	pos := pbutil.ConvertFromPos(moveData.GetPos())
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)

	err = objectMove(pl, uId, pos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
				"pos":      pos,
			}).Debug("scene:移动,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
			"pos":      pos,
		}).Debug("scene:移动完成")
	return nil
}

func objectMove(pl *player.Player, uId int64, pos coretypes.Position) (err error) {
	if pl.GetPlayerId() == uId {
		sceneManager := pl.GetManager(player.PlayerDataKeyScene).(*playerscene.PlayerSceneDataManager)
		sceneManager.Move(pos)
	} else {

	}
	return
}
