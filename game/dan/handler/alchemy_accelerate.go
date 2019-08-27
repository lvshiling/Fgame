package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
	dantypes "fgame/fgame/game/dan/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYACCELERATE_TYPE), dispatch.HandlerFunc(handleAlchemyAccelerate))
}

//处理炼丹加速信息
func handleAlchemyAccelerate(s session.Session, msg interface{}) (err error) {
	log.Debug("dan:处理炼丹加速信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAlchemyAccelerate := msg.(*uipb.CSAlchemyAccelerate)
	kindId := csAlchemyAccelerate.GetKindId()

	err = alchemyAccelerate(tpl, int(kindId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"kindId":   kindId,
				"error":    err,
			}).Error("dan:处理炼丹加速信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Debug("dan:处理炼丹加速信息完成")

	return nil
}

//处理加速炼丹的逻辑
func alchemyAccelerate(pl player.Player, kindId int) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	alchemyObj := danManager.GetAlchemy(kindId)
	if alchemyObj == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Warn("dan:无效的kindId,无法加速炼丹")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if alchemyObj.State != dantypes.AlchemyStateStart {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Warn("dan:丹药已练成,无需加速")
		playerlogic.SendSystemMessage(pl, lang.DanHasedFinshed)
		return
	}

	totalAlchemyMoney, synthetiseId, leftNum := danManager.GetAccelerateNeedGold(kindId)
	if totalAlchemyMoney != 0 {
		//元宝是否足够
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		flag := propertyManager.HasEnoughGold(int64(totalAlchemyMoney), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"kindId":   kindId,
			}).Warn("dan:当前元宝不足，请及时充值")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		//消耗元宝
		goldReason := commonlog.GoldLogReasonDanAccelerate
		goldReasonText := fmt.Sprintf(goldReason.String(), synthetiseId, leftNum)
		flag = propertyManager.CostGold(int64(totalAlchemyMoney), false, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("dan: alchemyAccelerate cost gold should be ok"))
		}
		//同步元宝
		propertylogic.SnapChangedProperty(pl)
	}

	//更新炼丹状态
	danManager.AlchemyAccelerateState(kindId)
	scAlchemyAccelerate := pbuitl.BuildSCAlchemyAccelerate(int32(kindId))
	pl.SendMsg(scAlchemyAccelerate)
	return
}
