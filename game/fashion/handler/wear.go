package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FASHION_WEAR_TYPE), dispatch.HandlerFunc(handlefashionWear))
}

//处理时装穿戴信息
func handlefashionWear(s session.Session, msg interface{}) (err error) {
	log.Debug("fashion:处理时装穿戴信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFashionWear := msg.(*uipb.CSFashionWear)
	fashionId := csFashionWear.GetFashionId()

	err = fashionWear(tpl, fashionId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Error("fashion:处理时装穿戴信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Debug("fashion:处理时装穿戴信息完成")
	return nil
}

//处理时装穿戴信息逻辑
func fashionWear(pl player.Player, fashionId int32) (err error) {
	fashionManager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	flag := fashionManager.IsValid(fashionId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = fashionManager.HasedWeared(fashionId)
	if flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:该时装已穿戴,无需再次穿戴")
		playerlogic.SendSystemMessage(pl, lang.FashionRepeatWear)
		return
	}

	flag = fashionManager.IfFashionWear(fashionId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:还没有该时装,请先获取")
		playerlogic.SendSystemMessage(pl, lang.FashionNotHas)
		return
	}

	flag = fashionManager.FashionWear(fashionId)
	if !flag {
		panic(fmt.Errorf("fashion: fashionWear  should be ok"))
	}

	scFashionWear := pbutil.BuildSCFashionWear(fashionId)
	pl.SendMsg(scFashionWear)
	return
}
