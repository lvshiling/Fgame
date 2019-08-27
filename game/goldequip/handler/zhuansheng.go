package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ZHUANSHENG_TYPE), dispatch.HandlerFunc(handleZhuanSheng))
}

//处理转生信息
func handleZhuanSheng(s session.Session, msg interface{}) (err error) {
	log.Debug("equip:处理转生消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = zhuansheng(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("equip:处理转生消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("equip:处理转生消息完成")
	return nil
}

//转生界面信息的逻辑
func zhuansheng(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curZhuanShu := manager.GetZhuanSheng()
	nextZhuanShu := curZhuanShu + 1
	nextTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetZhuanShengTemplate(nextZhuanShu)
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:您当前已达最高转数,无法再转生")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipZhuanShengReachLimit)
		return
	}

	needZhuanShu := nextTemplate.NeedZhuanshu
	if curZhuanShu < needZhuanShu {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:转生需要的玩家转数不足")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipZhuanShengNeedZhuanShuNoEnough)
		return
	}

	nextNeedLevel := nextTemplate.NeedLevel
	if pl.GetLevel() < nextNeedLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:您当前等级不足,无法转生")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipZhuanShengLevelNoEnough)
		return
	}

	feishengManager := pl.GetPlayerDataManager(types.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	curFeisheng := feishengManager.GetFeiShengLevel()
	if curFeisheng < nextTemplate.NeedFeisheng {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:您当前飞升等级不足,无法转生")
		playerlogic.SendSystemMessage(pl, lang.FeiShengLevelToLower)
		return
	}

	equipNum := int32(0)
	needEquipCount := nextTemplate.NeedEquipCount
	needEquipZhuanShu := nextTemplate.NeedEquipZhuanshu
	needEquipLevel := nextTemplate.NeedEquipLevel
	needEquipStrenth := nextTemplate.NeedEquipStreng
	needEquipQuality := nextTemplate.NeedEquipQuality
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldEquipObjList := goldequipManager.GetGoldEquipBag().GetAll()
	for _, goldEquipObj := range goldEquipObjList {
		if goldEquipObj.IsEmpty() {
			continue
		}

		propertyData := goldEquipObj.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)
		if propertyData.UpstarLevel < needEquipStrenth {
			continue
		}
		itemId := goldEquipObj.GetItemId()
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		if itemTemplate == nil {
			continue
		}

		if itemTemplate.NeedLevel < needEquipLevel {
			continue
		}
		if itemTemplate.Quality < needEquipQuality {
			continue
		}
		if itemTemplate.NeedZhuanShu < needEquipZhuanShu {
			continue
		}
		equipNum++
		if equipNum >= needEquipCount {
			break
		}
	}

	if equipNum < needEquipCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("goldequip:您当前等级不足,无法转生")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipZhuanShengGoldEquipNoEnough)
		return
	}

	reasonZhuanSheng := commonlog.ZhuanShengLogReasonGoldEquip
	manager.SetZhuanSheng(nextZhuanShu, reasonZhuanSheng, reasonZhuanSheng.String())
	scGoldEquipZhuanSheng := pbutil.BuildSCGoldEquipZhuanSheng(nextZhuanShu)
	pl.SendMsg(scGoldEquipZhuanSheng)
	goldequiplogic.ZhuanShengPropertyChanged(pl)
	propertylogic.SnapChangedProperty(pl)
	return
}
