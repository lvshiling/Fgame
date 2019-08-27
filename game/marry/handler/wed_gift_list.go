package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	marryscene "fgame/fgame/game/marry/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GIFT_LIST_TYPE), dispatch.HandlerFunc(handleMarryWedGiftList))
}

//处理贺礼榜单处理信息
func handleMarryWedGiftList(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理贺礼榜单处理消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = marryWedGiftList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理贺礼榜单处理消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理贺礼榜单处理消息完成")
	return nil
}

//处理贺礼榜单处理信息逻辑
func marryWedGiftList(pl player.Player) (err error) {
	s := pl.GetScene()
	marryScene := marry.GetMarryService().GetScene()
	if s != marryScene {
		return
	}

	sd := marry.GetMarryService().GetMarrySceneData()
	if sd.Status == marryscene.MarrySceneStatusTypeInit {
		return
	}

	marryDelegate := marryScene.SceneDelegate().(marryscene.MarrySceneData)
	heroismList := marryDelegate.GetHeroismList()
	scMarryWedGiftList := pbuitl.BuildSCMarryWedGiftList(heroismList)
	pl.SendMsg(scMarryWedGiftList)
	return
}
