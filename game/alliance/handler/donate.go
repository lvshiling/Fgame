package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DONATE_TYPE), dispatch.HandlerFunc(handleAllianceDonate))
}

//处理仙盟捐献
func handleAllianceDonate(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟捐献")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceDonate := msg.(*uipb.CSAllianceDonate)
	typ := csAllianceDonate.GetTyp()
	allianceType := alliancetypes.AllianceJuanXianType(typ)
	if !allianceType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Warn("alliance:处理仙盟捐献,错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = allianceDonate(tpl, allianceType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("alliance:处理仙盟捐献,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Debug("alliance:处理仙盟捐献,完成")
	return nil

}

//仙盟捐献
func allianceDonate(pl player.Player, typ alliancetypes.AllianceJuanXianType) (err error) {
	//获取模板
	unionDonateTemplate := alliancetemplate.GetAllianceTemplateService().GetUnionDonateTemplate(typ)
	if unionDonateTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Warn("alliance:处理仙盟捐献,错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	flag := allianceManager.IfCanDonate(typ)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Warn("alliance:处理仙盟捐献,捐献次数最大")
		playerlogic.SendSystemMessage(pl, lang.AllianceDonateMaxTimes)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch typ {
	case alliancetypes.AllianceJuanXianTypeLingPai:
		//判断是物品足够
		flag = inventoryManager.HasEnoughItem(unionDonateTemplate.DonateItemId, unionDonateTemplate.DonateItemCount)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
					"error":    err,
				}).Warn("alliance:处理仙盟捐献,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		break
	case alliancetypes.AllianceJuanXianTypeGold:
		//判断是元宝足够
		needGold := int64(unionDonateTemplate.DonateGold)
		flag = propertyManager.HasEnoughGold(needGold, false)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
					"error":    err,
				}).Warn("alliance:处理仙盟捐献,元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
		break
	case alliancetypes.AllianceJuanXianTypeSilver:
		//判断是银两足够
		flag = propertyManager.HasEnoughSilver(int64(unionDonateTemplate.DonateSilver))
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
					"error":    err,
				}).Warn("alliance:处理仙盟捐献,银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
		break
	}

	//捐献
	memObj, err := alliance.GetAllianceService().Donate(pl.GetId(), typ)
	if err != nil {
		return
	}

	allianceId := memObj.GetAlliance().GetAllianceId()
	switch typ {
	case alliancetypes.AllianceJuanXianTypeLingPai:
		reason := commonlog.InventoryLogReasonAllianceDonate
		reasonText := fmt.Sprintf(reason.String(), allianceId)
		//消耗物品
		flag = inventoryManager.UseItem(unionDonateTemplate.DonateItemId, unionDonateTemplate.DonateItemCount, reason, reasonText)
		if !flag {
			panic(fmt.Errorf("使用物品应该成功"))
		}
		break
	case alliancetypes.AllianceJuanXianTypeGold:
		reason := commonlog.GoldLogReasonAllianceDonate
		reasonText := fmt.Sprintf(reason.String(), allianceId)
		//消耗元宝
		needGold := int64(unionDonateTemplate.DonateGold)
		flag = propertyManager.CostGold(needGold, false, reason, reasonText)
		if !flag {
			panic(fmt.Errorf("使用元宝应该成功"))
		}
		break
	case alliancetypes.AllianceJuanXianTypeSilver:
		reason := commonlog.SilverLogReasonAllianceDonate
		reasonText := fmt.Sprintf(reason.String(), allianceId)
		//消耗银两
		flag = propertyManager.CostSilver(int64(unionDonateTemplate.DonateSilver), reason, reasonText)
		if !flag {
			panic(fmt.Errorf("使用银两应该成功"))
		}
		break
	}

	flag = allianceManager.Donate(typ)
	if !flag {
		panic(fmt.Errorf("捐献应该成功"))
	}

	//同步属性
	propertylogic.SnapChangedProperty(pl)
	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	scAllianceDonate := pbutil.BuildSCAllianceDonate(typ)
	pl.SendMsg(scAllianceDonate)

	return
}
