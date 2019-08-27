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
	shenqitypes "fgame/fgame/game/shenqi/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENQI_KIND_INFO_GET_TYPE), dispatch.HandlerFunc(handleShenQiKindInfoGet))

}

//处理获取一种神器信息
func handleShenQiKindInfoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("shenqi:处理获取一种神器消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenqiKindInfoGet := msg.(*uipb.CSShenqiKindInfoGet)
	shenQiTypeId := csShenqiKindInfoGet.GetShenQiType()
	shenQiType := shenqitypes.ShenQiType(shenQiTypeId)

	//参数不对
	if !shenQiType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"shenQiType": shenQiType.String(),
			}).Warn("shenqi:神器类型,错误")
		return
	}

	err = shenQiKindInfoGet(tpl, shenQiType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shenqi:处理获取一种神器消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenqi:处理获取一种神器消息完成")
	return nil

}

//获取一种神器界面信息逻辑
func shenQiKindInfoGet(pl player.Player, typ shenqitypes.ShenQiType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQi) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenqi:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	debrisMap := manager.GetShenQiDebrisMapByShenQi(typ)
	smeltMap := manager.GetShenQiSmeltMapByShenQi(typ)
	qiLingMap := manager.GetShenQiQiLingMapByShenQi(typ)
	scMsg := pbutil.BuildSCShenQiKindInfoGet(qiLingMap, debrisMap, smeltMap, typ)
	pl.SendMsg(scMsg)
	return
}
