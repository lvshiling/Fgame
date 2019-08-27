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
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_UPGRADE_TYPE), dispatch.HandlerFunc(handleSoulUpgrade))

}

//处理帝魂升级信息
func handleSoulUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理帝魂升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulUpgrade := msg.(*uipb.CSSoulUpgrade)
	soulTag := csSoulUpgrade.GetSoulTag()

	err = soulUpgrade(tpl, soultypes.SoulType(soulTag))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
				"error":    err,
			}).Error("soul:处理帝魂升级信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理帝魂升级信息完成")
	return nil

}

//帝魂升级逻辑
func soulUpgrade(pl player.Player, soulTag soultypes.SoulType) (err error) {
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
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:未激活的帝魂,无法升级")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotUpgrade)
		return
	}

	flag = soulManager.IfCanUpgrade(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:帝魂升级阶别已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.SoulUpgradeReackLimit)
		return
	}

	//升级阶别物品
	order := soulManager.GetSoulOrderByTag(soulTag)
	soulAwakenTemplate := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, order+1)
	items := soulAwakenTemplate.GetUpLevelItemMap()
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
			}).Warn("soul:物品数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonSoulUpgrade
		reasonText := fmt.Sprintf(inventoryReason.String(), soulTag, order)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soul: soulUpgrade use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = soulManager.Upgrade(soulTag)
	if !flag {
		panic(fmt.Errorf("soul: Upgrade  should be ok"))
	}
	curOrder := soulManager.GetSoulOrderByTag(soulTag)
	scSoulUpgrade := pbutil.BuildSCSoulUpgrade(int32(soulTag), curOrder)
	pl.SendMsg(scSoulUpgrade)
	return
}
