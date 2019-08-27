package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TITLE_WEAR_TYPE), dispatch.HandlerFunc(handleTitleWear))
}

//处理称号穿戴信息
func handleTitleWear(s session.Session, msg interface{}) (err error) {
	log.Debug("title:处理称号穿戴信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTitleWear := msg.(*uipb.CSTitleWear)
	titleId := csTitleWear.GetTitleId()

	err = titleWear(tpl, titleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
				"error":    err,
			}).Error("title:处理称号穿戴信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Debug("title:处理称号穿戴信息完成")
	return nil
}

//处理称号穿戴信息逻辑
func titleWear(pl player.Player, titleId int32) (err error) {
	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	flag := titleManager.IsWearValid(titleId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = titleManager.HasedWeared(titleId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:该称号已穿戴,无需再次穿戴")
		playerlogic.SendSystemMessage(pl, lang.TitleRepeatWear)
		return
	}

	flag = titleManager.IfTitleExist(titleId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:还没有该称号,请先获取")
		playerlogic.SendSystemMessage(pl, lang.TitleNotHas)
		return
	}

	flag = titleManager.TitleWear(titleId)
	if !flag {
		panic(fmt.Errorf("title: titleWear  should be ok"))
	}
	scTitleWear := pbutil.BuildSCTitleWear(titleId)
	pl.SendMsg(scTitleWear)
	return
}
