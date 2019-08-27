package handler

import (
	"fgame/fgame/client/player/pbutil"
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	playerproperty "fgame/fgame/client/property/player"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_PLAYER_PROPERTY_TYPE), dispatch.HandlerFunc(handlePlayerPropertyChanged))
}

func handlePlayerPropertyChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("player:玩家属性变化")
	scPlayerProperty := msg.(*uipb.SCPlayerProperty)

	battlePropertyMap := pbutil.ConvertFromBattleProperty(scPlayerProperty.GetBattlePropertyList())
	basePropertyMap := pbutil.ConvertFromBaseProperty(scPlayerProperty.GetBasePropertyList())
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)

	err = playerPropertyChanged(pl, basePropertyMap, battlePropertyMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
			}).Debug("player:玩家基本信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
		}).Debug("player:玩家基本信息")
	return nil
}

func playerPropertyChanged(pl *player.Player, basePropertyMap map[int32]int64, battlePropertyMap map[int32]int64) (err error) {
	propertyManager := pl.GetManager(player.PlayerDataKeyProperty).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateSystemBattleProperty(battlePropertyMap)

	return
}
