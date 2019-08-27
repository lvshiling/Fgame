package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xianfu/pbutil"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANFU_ACCELERATE_TYPE), dispatch.HandlerFunc(handlerXianfuAccelerate))
}

//秘境仙府升级加速请求
func handlerXianfuAccelerate(s session.Session, msg interface{}) (err error) {
	log.Debug("xianfu:处理秘境仙府升级加速请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csXianfuAccelerate := msg.(*uipb.CSXianfuAccelerate)
	typ := csXianfuAccelerate.GetXianfuType()

	xianfuType := xianfutypes.XianfuType(typ)
	if !xianfuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府加速升级请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.XianfuArgumentInvalid)
		return
	}

	err = xianfuAccelerate(tpl, xianfuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   tpl.GetId(),
				"xianfuType": xianfuType,
				"err":        err,
			}).Error("xianfu:处理秘境仙府升级加速请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   tpl.GetId(),
			"xianfuType": xianfuType,
		}).Debug("xianfu：处理秘境仙府升级加速请求完成")

	return
}

//仙府加速请求逻辑
func xianfuAccelerate(pl player.Player, xianfuType xianfutypes.XianfuType) (err error) {
	xianfuManager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	now := global.GetGame().GetTimeService().Now()

	if !xianfuManager.IsUpgrading(xianfuType) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府升级请求，仙府不是升级状态")
		playerlogic.SendSystemMessage(pl, lang.XianfuArgumentInvalid)
		return
	}

	//计算加速所需元宝
	totalNeedGold := xianfuManager.GetAccelerateNeedGold(xianfuType)

	//元宝是否足够
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if totalNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(totalNeedGold), false) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府加速升级请求，当前元宝不足，无法升级")
			playerlogic.SendSystemMessage(pl, lang.XianfuNotEnoughGold)
			return
		}
	}

	//消耗元宝
	if totalNeedGold > 0 {
		goldReason := commonlog.GoldLogReasonXianfuAccelerate
		goldReasonText := fmt.Sprintf(goldReason.String(), xianfuId, xianfuType)
		flag := propertyManager.CostGold(int64(totalNeedGold), false, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("xianfu: xianfuAccelerateUpgrade use gold should be ok"))
		}
	}

	//完成升级
	xianfuManager.FinishAccelerateUpgrade(xianfuType, now)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scXianfuAccelerate := pbutil.BuildSCXianfuAccelerate(xianfuId, xianfuType)
	pl.SendMsg(scXianfuAccelerate)
	return
}
