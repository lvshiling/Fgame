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
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UNREAL_TYPE), dispatch.HandlerFunc(handleLingTongDevUnreal))

}

//处理灵童养成类幻化信息
func handleLingTongDevUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevUnreal := msg.(*uipb.CSLingTongDevUnreal)
	classType := csLingTongDevUnreal.GetClassType()
	seqId := csLingTongDevUnreal.GetSeqId()
	err = lingTongDevUnreal(tpl, types.LingTongDevSysType(classType), int(seqId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"seqId":     seqId,
				"error":     err,
			}).Error("lingtongdev:处理灵童养成类幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Debug("lingtongdev:处理灵童养成类幻化信息完成")
	return nil

}

//灵童养成类幻化的逻辑
func lingTongDevUnreal(pl player.Player, classType types.LingTongDevSysType, seqId int) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("LingTongDev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))
	//校验参数
	if lingTongDevTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:幻化advancedId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	lingTongInfo := manager.GetLingTongDevInfo(classType)
	if lingTongInfo == nil || lingTongInfo.GetAdvancedId() <= 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:请先激活灵童养成类系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		return
	}

	//是否已幻化
	flag := manager.IsUnrealed(classType, seqId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := lingTongDevTemplate.GetMagicParamIMap()
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerId":  pl.GetId(),
					"classType": classType,
					"seqId":     seqId,
				}).Warn("LingTongDev:还有幻化条件未达成，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.LingTongDevUnrealCondNotReached)
				return
			}
		}

		flag = manager.IsCanUnreal(classType, seqId)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"seqId":     seqId,
			}).Warn("LingTongDev:还有幻化条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.LingTongDevUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonLingTongDevUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), classType.String(), seqId)
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("lingtongdev:use item should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
		manager.AddUnrealInfo(classType, seqId)
		lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)
	}

	flag = manager.Unreal(classType, seqId)
	if !flag {
		panic(fmt.Errorf("lingtongdev:幻化应该成功"))
	}
	scLingTongDevUnreal := pbutil.BuildSCLingTongDevUnreal(int32(classType), int32(seqId))
	pl.SendMsg(scLingTongDevUnreal)
	return
}
