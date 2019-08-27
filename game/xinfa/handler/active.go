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
	processor.Register(codec.MessageType(uipb.MessageType_CS_XINFA_ACTIVE_TYPE), dispatch.HandlerFunc(handleXinFaActive))
}

//处理心法激活信息
func handleXinFaActive(s session.Session, msg interface{}) (err error) {
	log.Debug("xinfa:处理心法激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXinFaActive := msg.(*uipb.CSXinFaActive)
	typ := csXinFaActive.GetTyp()

	err = xinfaActive(tpl, xinfatypes.XinFaType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("xinfa:处理心法激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Debug("xinfa:处理心法激活信息完成")
	return nil
}

//处理心法激活信息逻辑
func xinfaActive(pl player.Player, typ xinfatypes.XinFaType) (err error) {
	xfManager := pl.GetPlayerDataManager(types.PlayerXinFaDataManagerType).(*playerxinfa.PlayerXinFaDataManager)
	flag := typ.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("xinfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = xfManager.IfXinFaExist(typ)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("xinfa:该心法已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.XinFaRepeatActive)
		return
	}

	xinFaTemplate := xinfa.GetXinFaService().GetXinFaByTypeAndLevel(typ, 1)
	useItem := xinFaTemplate.GetUseItemTemplate()
	needYinLiang := int64(xinFaTemplate.NeedYinLiang)

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//银两判断
	if needYinLiang != 0 {
		flag = propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("xinfa:银两不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//物品判断
	if useItem != nil {
		needItem := xinFaTemplate.NeedItemId
		needCount := xinFaTemplate.NeedItemNum
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItem(needItem, needCount)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("xinfa:道具不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		inventoryReason := commonlog.InventoryLogReasonXinFaActive
		reasonText := fmt.Sprintf(inventoryReason.String(), typ)
		flag = inventoryManager.UseItem(needItem, needCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("xinfa: xinfaActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//消耗银两
	if needYinLiang != 0 {
		reasonText := commonlog.SilverLogReasonXinFaActive.String()
		flag = propertyManager.CostSilver(needYinLiang, commonlog.SilverLogReasonXinFaActive, reasonText)
		if !flag {
			panic(fmt.Errorf("xinfa: xinfaActive CostSilver  should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	flag = xfManager.XinFaActive(typ)
	if !flag {
		panic(fmt.Errorf("xinfa: XinFaActive should be ok"))
	}

	xinFaId := int32(xinFaTemplate.TemplateId())
	scXinFaActive := pbutil.BuildSCXinFaActive(xinFaId)
	pl.SendMsg(scXinFaActive)
	return
}
