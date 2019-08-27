package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_UNREAL_TYPE), dispatch.HandlerFunc(handleLingyuUnreal))

}

//处理领域幻化信息
func handleLingyuUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("lingyu:处理领域幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingyuUnreal := msg.(*uipb.CSLingyuUnreal)
	lingyuId := csLingyuUnreal.GetLingyuId()
	err = lingyuUnreal(tpl, int(lingyuId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"lingyuId": lingyuId,
				"error":    err,
			}).Error("lingyu:处理领域幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"lingyuId": lingyuId,
		}).Debug("lingyu:处理领域幻化信息完成")
	return nil

}

//领域幻化的逻辑
func lingyuUnreal(pl player.Player, lingyuId int) (err error) {
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingyuId))
	//校验参数
	if lingyuTemplate == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"lingyuId": lingyuId,
		}).Warn("Lingyu:幻化advancedId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否已幻化
	flag := lingyuManager.IsUnrealed(lingyuId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := lingyuTemplate.GetMagicParamIMap()
		if len(paramItemsMap) > 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerid": pl.GetId(),
					"lingyuId": lingyuId,
				}).Warn("Lingyu:幻化物品不足，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.LingyuUnrealCondNotReached)
				return
			}
		}

		flag = lingyuManager.IsCanUnreal(lingyuId)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"lingyuId": lingyuId,
			}).Warn("Lingyu:还有幻化条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.LingyuUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonLingyuUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), lingyuId)
		if len(paramItemsMap) > 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("lingyu:use item should be ok"))
			}
		}

		lingyuManager.AddUnrealInfo(lingyuId)
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
		lingyulogic.LingyuPropertyChanged(pl)
	}
	flag = lingyuManager.Unreal(lingyuId)
	if !flag {
		panic(fmt.Errorf("lingyu:幻化应该成功"))
	}

	scLingyuUnreal := pbutil.BuildSCLingyuUnreal(int32(lingyuId))
	pl.SendMsg(scLingyuUnreal)
	return
}
