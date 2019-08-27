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
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UPSTAR_TYPE), dispatch.HandlerFunc(handleLingTongDevUpstar))
}

//处理灵童养成类皮肤升星信息
func handleLingTongDevUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevUpstar := msg.(*uipb.CSLingTongDevUpstar)
	classType := csLingTongDevUpstar.GetClassType()
	seqId := csLingTongDevUpstar.GetSeqId()

	err = lingTongDevUpstar(tpl, types.LingTongDevSysType(classType), seqId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"seqId":    seqId,
				"error":    err,
			}).Error("lingtongdev:处理灵童养成类皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtongdev:处理灵童养成类皮肤升星完成")
	return nil
}

//灵童养成类皮肤升星的逻辑
func lingTongDevUpstar(pl player.Player, classType types.LingTongDevSysType, seqId int32) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))
	if lingTongDevTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if lingTongDevTemplate.GetUpstarBeginId() == 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	container := manager.GetLingTongDevOtherMap(classType)
	if container == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:未激活的灵童养成类皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevSkinUpstarNoActive, classType.String())
		return
	}
	otherObj := container.GetOtherObj(seqId)
	if otherObj == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:未激活的灵童养成类皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevSkinUpstarNoActive, classType.String())
		return
	}

	curLevel := otherObj.GetLevel()
	nextLevel := curLevel + 1
	nextlingTongDevUpstarTemplate := lingTongDevTemplate.GetLingTongDevUpstarByLevel(nextLevel)
	if nextlingTongDevUpstarTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
			"seqId":     seqId,
		}).Warn("lingtongdev:s皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevSkinReacheFullStar, classType.String())
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItems := nextlingTongDevUpstarTemplate.GetItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid":  pl.GetId(),
				"classType": classType,
				"seqId":     seqId,
			}).Warn("lingtongdev:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongDevUpstar.String(), classType.String())
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingTongDevUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtongdev: lingTongDevUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//皮肤升星判断
	pro, _, sucess := lingtongdevlogic.LingTongDevSkinUpstar(otherObj.GetUpNum(), otherObj.GetUpPro(), nextlingTongDevUpstarTemplate)
	flag := manager.Upstar(classType, seqId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("lingtongdev: lingTongDevUpstar should be ok"))
	}
	if sucess {
		lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)
	}
	scLingTongDevUpstar := pbutil.BuildSCLingTongDevUpstar(int32(classType), seqId, otherObj.GetLevel(), otherObj.GetUpPro())
	pl.SendMsg(scLingTongDevUpstar)
	return
}
