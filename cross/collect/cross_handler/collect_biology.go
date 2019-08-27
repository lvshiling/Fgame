package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	droptemplate "fgame/fgame/game/drop/template"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_COLLECT_FINISH_TYPE), dispatch.HandlerFunc(handleCollectFinish))
}

func handleCollectFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:采集完成")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siCollectFinish := msg.(*crosspb.SICollectFinish)
	itemInfoList := siCollectFinish.GetItemList()

	ps := tpl.GetScene()
	if ps != nil {
		mapType := ps.MapTemplate().GetMapType()
		switch mapType {
		case scenetypes.SceneTypeCrossDenseWat:
			for _, itemInfo := range itemInfoList {
				itemId := itemInfo.GetItemId()
				num := itemInfo.GetNum()
				itemData := droptemplate.CreateItemData(itemId, num, 0, 0)
				ps.OnPlayerGetItem(tpl, itemData)
			}
		}
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("collect:采集完成,完成")
	return nil
}
