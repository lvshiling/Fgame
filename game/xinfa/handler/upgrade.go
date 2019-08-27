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
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xinfa/pbutil"
	playerxinfa "fgame/fgame/game/xinfa/player"
	xinfatypes "fgame/fgame/game/xinfa/types"
	"fgame/fgame/game/xinfa/xinfa"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XINFA_UPGRADE_TYPE), dispatch.HandlerFunc(handleXinFaUpgrade))
}

//处理心法升级信息
func handleXinFaUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("xinfa:处理心法升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXinFaUpgrade := msg.(*uipb.CSXinFaUpgrade)
	typ := csXinFaUpgrade.GetTyp()

	err = xinfaUpgrade(tpl, xinfatypes.XinFaType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("xinfa:处理心法升级信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xinfa:处理心法升级完成")
	return nil
}

//心法升级的逻辑
func xinfaUpgrade(pl player.Player, typ xinfatypes.XinFaType) (err error) {
	xfManager := pl.GetPlayerDataManager(types.PlayerXinFaDataManagerType).(*playerxinfa.PlayerXinFaDataManager)
	flag := xfManager.IfXinFaExist(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("xinfa:未激活的心法,无法升级")
		playerlogic.SendSystemMessage(pl, lang.XinFaNotActiveNotUpgrade)
		return
	}

	flag = xfManager.IfCanUpgrade(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("xinfa:心法已达最高级")
		playerlogic.SendSystemMessage(pl, lang.XinFaReacheFullUpgrade)
		return
	}

	curLevel := xfManager.GetXinFaLevelByTyp(typ)
	nextLevel := curLevel + 1
	xinFaTemplate := xinfa.GetXinFaService().GetXinFaByTypeAndLevel(typ, nextLevel)
	useItem := xinFaTemplate.GetUseItemTemplate()
	needYinLiang := int64(xinFaTemplate.NeedYinLiang)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//银两判断
	if needYinLiang != 0 {
		flag = propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("xinfa:银两不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	if useItem != nil {
		needItem := xinFaTemplate.NeedItemId
		needCount := xinFaTemplate.NeedItemNum
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItem(needItem, needCount)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("xinfa:道具不足，无法升级")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		reasonText := commonlog.InventoryLogReasonXinFaUpgrade.String()
		flag = inventoryManager.UseItem(needItem, needCount, commonlog.InventoryLogReasonXinFaUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("xinfa: xinfaUpgrade use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//消耗银两
	if needYinLiang != 0 {
		reasonText := commonlog.SilverLogReasonXinFaUpgrade.String()
		flag = propertyManager.CostSilver(needYinLiang, commonlog.SilverLogReasonXinFaUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("xinfa: xinfaUpgrade CostSilver  should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	flag = xfManager.Upgrade(typ)
	if !flag {
		panic(fmt.Errorf("xinfa: Upgrade should be ok"))
	}

	xinfaId := int32(xinFaTemplate.TemplateId())
	scXinFaUpgrade := pbutil.BuildSCXinFaUpgrade(xinfaId)
	pl.SendMsg(scXinFaUpgrade)
	return
}
