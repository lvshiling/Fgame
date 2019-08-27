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
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_FEED_TYPE), dispatch.HandlerFunc(handleSoulFeed))
}

//处理帝魂喂养信息
func handleSoulFeed(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理帝魂喂养信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulFeed := msg.(*uipb.CSSoulFeed)
	soulTag := csSoulFeed.GetSoulTag()

	err = soulFeed(tpl, soultypes.SoulType(soulTag))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
				"error":    err,
			}).Error("soul:处理帝魂喂养信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理帝魂喂养完成")
	return nil
}

//帝魂喂养的逻辑
func soulFeed(pl player.Player, soulTag soultypes.SoulType) (err error) {
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := soulManager.IfSoulTagExist(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:未激活的帝魂,无法喂养")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotFeed)
		return
	}

	flag = soulManager.IfCanFeed(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:帝魂已满级")
		playerlogic.SendSystemMessage(pl, lang.SoulReacheFullLevel)
		return
	}

	//帝魂信息
	level := soulManager.GetSoulLevelByTag(soulTag)
	//吞噬物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	devourItems := soul.GetSoulService().GetSoulDevourTemplateByLevel(soulTag, level)
	totalExp := int32(0)
	needItems := make(map[int32]int32)
	if len(devourItems) != 0 {
		for itemId, exp := range devourItems {
			num := inventoryManager.NumOfItems(itemId)
			if num <= 0 {
				continue
			}
			totalExp += exp * num
			needItems[itemId] = num
		}
	}

	if len(needItems) == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:当前没有可以喂养的物品")
		playerlogic.SendSystemMessage(pl, lang.SoulFeedNotItem)
		return
	}

	//消耗吞噬物品
	inventoryReason := commonlog.InventoryLogReasonSoulFeed
	reasonText := fmt.Sprintf(inventoryReason.String(), soulTag)
	flag = inventoryManager.BatchRemove(needItems, inventoryReason, reasonText)
	if !flag {
		panic(fmt.Errorf("soul: soulFeed use item should be ok"))
	}
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)
	//喂养升级
	flag = soulManager.UpgradesByFeed(soulTag, totalExp)
	if !flag {
		panic(fmt.Errorf("soul: soulFeed UpgradesByFeed should be ok"))
	}

	//喂养有升级
	nowLevel := soulManager.GetSoulLevelByTag(soulTag)
	if level != nowLevel {
		//同步属性推送
		soullogic.SoulPropertyChanged(pl)
	}

	nowExp := soulManager.GetSoulExpByTag(soulTag)
	scSoulFeed := pbutil.BuildSCSoulFeed(int32(soulTag), nowLevel, nowExp)
	pl.SendMsg(scSoulFeed)
	return
}
