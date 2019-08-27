package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENQI_INFO_GET_TYPE), dispatch.HandlerFunc(handleShenQiInfoGet))

}

//处理神器信息
func handleShenQiInfoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("shenqi:处理获取神器消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = shenQiInfoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shenqi:处理获取神器消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenqi:处理获取神器消息完成")
	return nil

}

//获取神器界面信息逻辑
func shenQiInfoGet(pl player.Player) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQi) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenqi:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	debrisMap := manager.GetShenQiDebrisMap()
	smeltMap := manager.GetShenQiSmeltMap()
	qiLingMap := manager.GetShenQiQiLingMap()
	shenQiOjb := manager.GetShenQiOjb()
	scMsg := pbutil.BuildSCShenQiInfoGet(qiLingMap, debrisMap, smeltMap, shenQiOjb.LingQiNum)
	pl.SendMsg(scMsg)
	return
}
