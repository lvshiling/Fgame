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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_UNREAL_TYPE), dispatch.HandlerFunc(handleShenfaUnreal))
}

//处理身法幻化信息
func handleShenfaUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("shenfa:处理身法幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenfaUnreal := msg.(*uipb.CSShenfaUnreal)
	shenfaId := csShenfaUnreal.GetShenfaId()
	err = shenfaUnreal(tpl, int(shenfaId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shenfaId": shenfaId,
				"error":    err,
			}).Error("shenfa:处理身法幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"shenfaId": shenfaId,
		}).Debug("shenfa:处理身法幻化信息完成")
	return nil

}

//身法幻化的逻辑
func shenfaUnreal(pl player.Player, shenfaId int) (err error) {
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenfaId))
	//校验参数
	if shenfaTemplate == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"shenfaId": shenfaId,
		}).Warn("Shenfa:幻化advancedId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否已幻化
	flag := shenfaManager.IsUnrealed(shenfaId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := shenfaTemplate.GetMagicParamIMap()
		if len(paramItemsMap) > 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerid": pl.GetId(),
					"shenfaId": shenfaId,
				}).Warn("Shenfa:还有幻化条件未达成，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealCondNotReached)
				return
			}
		}

		flag = shenfaManager.IsCanUnreal(shenfaId)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"shenfaId": shenfaId,
			}).Warn("Shenfa:还有幻化条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonShenfaUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), shenfaId)
		if len(paramItemsMap) > 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("shenfa:use item should be ok"))
			}
		}

		shenfaManager.AddUnrealInfo(shenfaId)

		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
		shenfalogic.ShenfaPropertyChanged(pl)
	}
	flag = shenfaManager.Unreal(shenfaId)

	if !flag {
		panic(fmt.Errorf("shenfa:幻化应该成功"))
	}

	scShenfaUnreal := pbutil.BuildSCShenfaUnreal(int32(shenfaId))
	pl.SendMsg(scShenfaUnreal)
	return
}
