package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	moonlogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/moonlove/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOONLOVE_VIEW_DOUBLE_TYPE), dispatch.HandlerFunc(handlerMoonlveViewDouble))
}

//双人赏月
func handlerMoonlveViewDouble(s session.Session, msg interface{}) (err error) {
	log.Debug("moonlove:处理双人赏月请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSMoonloveViewDouble)
	findPlId := csMsg.GetPlayerId()

	err = moonlveViewDouble(tpl, findPlId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("moonlove:处理双人赏月请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("moonlove:处理双人赏月请求完成")

	return
}

//双人赏月逻辑
func moonlveViewDouble(pl player.Player, findPlId int64) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("月下情缘：场景不存在")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("月下情缘：玩家不在月下情缘场景中")
		playerlogic.SendSystemMessage(pl, lang.MoonloveNotInScene)
		return
	}

	findPl := s.GetPlayer(findPlId)
	if findPl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"findPlId": findPlId,
			}).Warn("月下情缘：邀请的玩家不在月下情缘")
		playerlogic.SendSystemMessage(pl, lang.MoonloveFindFail)
		return
	}

	if pl.GetSex() == findPl.GetSex() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"findPlId": findPlId,
			}).Warn("月下情缘：仅能选择异性玩家进行双人赏月")
		playerlogic.SendSystemMessage(pl, lang.MoonloveSameSex)
		return
	}

	sceneData := pl.GetScene().SceneDelegate().(moonlogic.MoonloveSceneData)
	if sceneData.IsCouple(findPlId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("月下情缘：该玩家已和他人赏月中")
		playerlogic.SendSystemMessage(pl, lang.MoonloveNotSinglePlayer)
		return
	}

	target := findPl.GetPosition()
	scViewDouble := pbutil.BuildMoonloveViewDouble(pl.GetId(), findPlId, target)
	pl.SendMsg(scViewDouble)

	return
}
