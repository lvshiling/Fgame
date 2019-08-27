package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_CLICK_CAR_TYPE), dispatch.HandlerFunc(handleMarryClickCar))
}

//处理点击婚车信息
func handleMarryClickCar(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理获取点击婚车消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = marryClickCar(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理获取点击婚车消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理获取点击婚车消息完成")
	return nil
}

//处理点击婚车信息逻辑
func marryClickCar(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	hunCheNpc := marry.GetMarryService().GetHunCheNpc()
	if hunCheNpc == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:处理获取点击婚车消息,婚车不存在")
		return
	}
	hunCheScene := hunCheNpc.GetScene()
	playerScene := pl.GetScene()
	if playerScene != hunCheScene {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:处理获取点击婚车消息,场景不在婚车场景")
		return
	}

	//只有夫妻两人可以点击
	hunCheObj := hunCheNpc.GetHunCheObject()
	playerId := hunCheObj.PlayerId
	spouseId := hunCheObj.SpouseId
	if pl.GetId() != playerId && pl.GetId() != spouseId {
		log.WithFields(
			log.Fields{
				"playerId":              pl.GetId(),
				"currentMarryPlayerId":  playerId,
				"currentSpousePlayerId": spouseId,
			}).Warn("marry:处理获取点击婚车消息,不是结婚的人")
		return
	}

	banquetTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeHunChe, hunCheObj.HunCheGrade)
	cdTime := int64(banquetTemplate.DropTime)
	now := global.GetGame().GetTimeService().Now()
	clickTime := manager.GetClickTime()
	if now-clickTime < cdTime {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("marry:操作过于频繁")
		playerlogic.SendSystemMessage(pl, lang.MarryClickCarFrequent)
		return
	} else {
		manager.SetClickTime(now)
	}

	marrylogic.HunCheDrop(pl.GetPosition(), 0, hunCheObj.SugarGrade)

	scMaryClickCar := pbuitl.BuildSCMaryClickCar(now, int32(hunCheObj.HunCheGrade))
	pl.SendMsg(scMaryClickCar)
	return
}
