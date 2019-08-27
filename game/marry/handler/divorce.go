package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_DIVORCE_TYPE), dispatch.HandlerFunc(handleMarryDivorce))
}

//处理离婚信息
func handleMarryDivorce(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理离婚消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryDivorce := msg.(*uipb.CSMarryDivorce)
	typ := csMarryDivorce.GetTyp()
	err = marryDivorce(tpl, marrytypes.MarryDivorceType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("marry:处理离婚消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理离婚消息完成")
	return nil
}

//处理离婚信息逻辑
func marryDivorce(pl player.Player, typ marrytypes.MarryDivorceType) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	status := marryInfo.Status
	onlineFlag := true
	if status == marrytypes.MarryStatusTypeUnmarried ||
		status == marrytypes.MarryStatusTypeDivorce || !typ.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("marry:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	marrySceneData := marry.GetMarryService().GetMarrySceneData()
	playerId := pl.GetId()
	if playerId == marrySceneData.PlayerId || playerId == marrySceneData.SpouseId {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("marry:婚礼期间无法离婚")
		playerlogic.SendSystemMessage(pl, lang.MarryDivorceInWeddingTime)
		return
	}
	proposalId := pl.GetId()
	if marryInfo.IsProposal != 1 {
		proposalId = marryInfo.SpouseId
	}
	flag := marry.GetMarryService().MarryPreWedIsExist(proposalId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("marry:当前请求举办婚宴中,无法离婚")
		playerlogic.SendSystemMessage(pl, lang.MarryDivorceInPreWed)
		return
	}

	err = marry.GetMarryService().Divorce(pl, marryInfo.SpouseId, typ)
	if err != nil && err == marry.ErrorMarryDivorceNoOnline {
		onlineFlag = false
	}
	scMarryDivorce := pbuitl.BuildSCMarryDivorce(int32(typ), onlineFlag)
	pl.SendMsg(scMarryDivorce)
	return
}
