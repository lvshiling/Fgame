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
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_SC_ENTER_SCENE_TYPE), dispatch.HandlerFunc(handleEnterScene))
}

func handleEnterScene(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理进入场景")
	scEnterScene := msg.(*scenepb.SCEnterScene)
	mapId := scEnterScene.GetMapId()
	// scEnterScene.GetPlayerData()
	pos := scEnterScene.GetPos()
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)
	tPos := pbutil.ConvertFromPos(pos)

	err = playerEnterScene(pl, mapId, tPos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
				"mapId":    mapId,
				"pos":      tPos,
			}).Debug("scene:处理进入场景,错误s")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
			"mapId":    mapId,
			"pos":      tPos,
		}).Debug("scene:处理进入场景完成")
	return nil
}

func playerEnterScene(pl *player.Player, mapId int32, pos coretypes.Position) (err error) {
	pl.Game()
	sceneManager := pl.GetManager(player.PlayerDataKeyScene).(*playerscene.PlayerSceneDataManager)
	if mapId == 0 {
		panic(fmt.Errorf("mapid 不应该为0"))
	}
	sceneManager.EnterScene(mapId, pos)
	//随机场景和随机位置

	// pl.StartStrategy(strategy.CreateMoveStrategy(pl))
	return
}
