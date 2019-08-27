package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playerfriend "fgame/fgame/game/friend/player"
	moonlovelogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/moonlove/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOONLOVE_PLAYER_LIST_TYPE), dispatch.HandlerFunc(handlerMoonlvePlayerList))
}

//月下情缘玩家列表
func handlerMoonlvePlayerList(s session.Session, msg interface{}) (err error) {
	log.Debug("moonlove:处理场景玩家列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err = moonlvePlayerList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("moonlove:处理场景玩家列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("moonlove:处理场景玩家列表请求完成")

	return
}

//玩家列表逻辑
func moonlvePlayerList(pl player.Player) (err error) {
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("月下情缘：玩家不在月下情缘场景中")

		playerlogic.SendSystemMessage(pl, lang.MoonloveNotInScene)
		return
	}

	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	sd := s.SceneDelegate().(moonlovelogic.MoonloveSceneData)
	coupleMap := sd.GetCoupleMap()
	friendMap := friendManager.GetFriends()
	allPlayer := s.GetAllPlayers()
	scMoonlovePlayerList := pbutil.BuildMoonlovePlayerList(pl.GetId(), allPlayer, friendMap, coupleMap)
	pl.SendMsg(scMoonlovePlayerList)

	return
}
