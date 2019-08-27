package handler

import (
	"fgame/fgame/client/player/pbutil"
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_PLAYER_INFO_TYPE), dispatch.HandlerFunc(handlerPlayerInfo))
}

func handlerPlayerInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("player:玩家基本信息")
	scPlayerInfo := msg.(*uipb.SCPlayerInfo)
	playerId := scPlayerInfo.GetPlayerId()
	name := scPlayerInfo.GetName()
	role := playertypes.RoleType(scPlayerInfo.GetRole())
	sex := playertypes.SexType(scPlayerInfo.GetSex())

	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)

	err = playerInfo(pl, playerId, name, role, sex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
			}).Debug("player:玩家基本信息,错误s")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
		}).Debug("player:玩家基本信息")
	return nil
}

func playerInfo(pl *player.Player, playerId int64, name string, role playertypes.RoleType, sex playertypes.SexType) (err error) {
	basicManager := pl.GetManager(player.PlayerDataKeyBasic).(*player.PlayerBasicManager)
	basicManager.Load(playerId, name, role, sex)
	mainGuaJi := pbutil.BuildMainGuaJi()
	pl.SendMessage(mainGuaJi)
	return
}
