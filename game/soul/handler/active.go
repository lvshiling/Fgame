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
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	soullogic "fgame/fgame/game/soul/logic"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_ACTIVE_TYPE), dispatch.HandlerFunc(handleSoulActive))
}

//处理帝魂激活信息
func handleSoulActive(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理帝魂激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulActive := msg.(*uipb.CSSoulActive)
	soulTag := csSoulActive.GetSoulTag()

	err = soulActive(tpl, soultypes.SoulType(soulTag))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
				"error":    err,
			}).Error("soul:处理帝魂激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Debug("soul:处理帝魂激活信息完成")
	return nil
}

//处理帝魂激活信息逻辑
func soulActive(pl player.Player, soulTag soultypes.SoulType) (err error) {
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := soulTag.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = soulManager.IfSoulTagExist(soulTag)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:该帝魂已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.SoulRepeatActive)
		return
	}

	soulActiveTemplate := soul.GetSoulService().GetSoulActiveTemplate(soulTag)
	//激活的前置帝魂条件
	preSoulCond := soulActiveTemplate.GetPreSoulCond()
	if preSoulCond != nil {
		preSoulTag := preSoulCond.GetSoulType()
		flag := soulManager.IfPreSoul(preSoulTag, preSoulCond.Level)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
			}).Warn("soul:激活该帝魂的前置条件不足")
			playerlogic.SendSystemMessage(pl, lang.SoulActiveNotPreCond)
			return
		}
	}

	//激活需要物品
	items := soulActiveTemplate.GetNeedItemMap()
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
			}).Warn("soul:您的物品不足，无法激活帝魂")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonSoulActive
		reasonText := fmt.Sprintf(inventoryReason.String(), soulTag)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soul: soulActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	soulObj, autoEmbed, flag := soulManager.SoulActive(soulTag)
	if !flag {
		panic(fmt.Errorf("soul: soulActive should be ok"))
	}
	//同步属性
	soullogic.SoulPropertyChanged(pl)
	//同步属性
	propertylogic.SnapChangedProperty(pl)
	scSoulActive := pbutil.BuildSCSoulActive(soulObj)
	pl.SendMsg(scSoulActive)

	if autoEmbed {
		soulId, flag := soulManager.GetSoulIdByOrder(soulTag)
		if !flag {
			panic(fmt.Errorf("soul: soulActive GetSoulIdByOrder should be ok"))
		}
		scSoulEmbed := pbutil.BuildSCSoulEmbed(soulId)
		pl.SendMsg(scSoulEmbed)
	}
	return
}
