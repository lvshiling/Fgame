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
	soullogic "fgame/fgame/game/soul/logic"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_AWAKEN_TYPE), dispatch.HandlerFunc(handleSoulAwaken))

}

//处理帝魂觉醒信息
func handleSoulAwaken(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理帝魂觉醒信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulAwaken := msg.(*uipb.CSSoulAwaken)
	soulTag := csSoulAwaken.GetSoulTag()

	err = soulAwanken(tpl, soultypes.SoulType(soulTag))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
				"error":    err,
			}).Error("soul:处理帝魂觉醒信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理帝魂觉醒信息完成")
	return nil

}

//帝魂觉醒逻辑
func soulAwanken(pl player.Player, soulTag soultypes.SoulType) (err error) {
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
		}).Warn("soul:未激活的帝魂,无法觉醒")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotAwaken)
		return
	}

	flag = soulManager.IfIsAwaken(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:帝魂已经觉醒过了")
		playerlogic.SendSystemMessage(pl, lang.SoulAwakenIsExist)
		return
	}

	//觉醒阶别物品
	soulAwakenTemplate := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, 1)
	items := soulAwakenTemplate.GetNeedItemMap()
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
		inventoryReason := commonlog.InventoryLogReasonSoulAwaken
		reasonText := fmt.Sprintf(inventoryReason.String(), soulTag)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soul: soulAwanken use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = soulManager.Awaken(soulTag)
	if !flag {
		panic(fmt.Errorf("soul: Awaken  should be ok"))
	}

	//同步属性
	soullogic.SoulPropertyChanged(pl)
	scSoulAwaken := pbutil.BuildSCSoulAwaken(int32(soulTag), true)
	pl.SendMsg(scSoulAwaken)
	return
}
