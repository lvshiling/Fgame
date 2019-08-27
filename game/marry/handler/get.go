package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_GET_TYPE), dispatch.HandlerFunc(handleMarryGet))
}

//处理结婚界面信息
func handleMarryGet(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理获取结婚界面消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = marryGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理获取结婚界面消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理获取结婚界面消息完成")
	return nil
}

//处理结婚界面界面信息逻辑
func marryGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	period := marry.GetMarryService().GetWeddingPeriod(pl.GetId())
	//定情信物
	playerSuitMap := marry.GetMarryService().GetMarryDingQing(pl.GetId())
	spoutId := marry.GetMarryService().GetSpouseId(pl.GetId())
	var spoutSuiteMap map[int32]map[int32]int32
	if spoutId > 0 {
		spoutSuiteMap = marry.GetMarryService().GetMarryDingQing(spoutId)
	}

	scMarryGet := pbuitl.BuildSCMarryGet(pl, marryInfo, period, false, playerSuitMap, spoutSuiteMap)
	pl.SendMsg(scMarryGet)
	return
}
