package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	marrylogic "fgame/fgame/game/marry/logic"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_RING_REPLACE_TYPE), dispatch.HandlerFunc(handleMarryRingTypeChange))
}

//处理婚戒替换信息
//元宝替换
func handleMarryRingTypeChange(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚戒替换消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryRingReplace := msg.(*uipb.CSMarryRingReplace)
	ring := csMarryRingReplace.GetRingType()
	err = marryRingTypeChange(tpl, marrytypes.MarryRingType(ring))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"ring":     ring,
				"error":    err,
			}).Error("marry:处理婚戒替换消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("marry:处理婚戒替换消息完成")
	return nil
}

//处理婚戒替换信息逻辑
func marryRingTypeChange(pl player.Player, ringType marrytypes.MarryRingType) (err error) {
	if !ringType.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if ringType.HoutaiType() != marrytemplate.GetMarryTemplateService().GetHouTaiType() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"ringType":      ringType,
			"currentHouTai": marrytemplate.GetMarryTemplateService().GetHouTaiType(),
		}).Warn("marry:后台类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	spouseId := marryInfo.SpouseId
	// 判断有没有真正结婚
	if marryInfo.Status != marrytypes.MarryStatusTypeMarried {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:未婚,无法替换")
		playerlogic.SendSystemMessage(pl, lang.MarryRingReplaceNoMarried)
		return
	}

	curRingType := marryInfo.Ring
	if !ringType.BetterThan(curRingType) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:婚戒替换需要更高级")
		playerlogic.SendSystemMessage(pl, lang.MarryRingReplaceNeedSenior)
		return
	}

	//获取婚烟对戒
	ringItem := marrytemplate.GetMarryTemplateService().GetMarryBanquetTemplateByHouTai(ringType.HoutaiType(), marrytypes.MarryBanquetTypeRing, ringType.BanquetSubTypeRing())
	if ringItem == nil {
		log.WithFields(log.Fields{
			"playerId":             pl.GetId(),
			"ringType":             ringType,
			"MarryBanquetTypeRing": marrytypes.MarryBanquetTypeRing,
			"merrySubRing":         ringType.BanquetSubTypeRing(),
		}).Warn("marry: 婚戒模板不存在")
		playerlogic.SendSystemMessage(pl, lang.MarryRingTemplateNotExist)
		return
	}

	buyGold := int64(ringItem.UseGold)
	bindGold := int64(ringItem.UseBinggold)
	silver := int64(ringItem.UseSilver)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	if !propertyManager.HasEnoughCost(bindGold, buyGold, silver) {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//扣除钱
	goldUseReason := commonlog.GoldLogReasonRingTypeChangeCost
	goldReason := fmt.Sprintf(goldUseReason.String(), curRingType.String(), ringType.String())
	silverUseReason := commonlog.SilverLogReasonRingTypeChangeCost
	silverReason := fmt.Sprintf(silverUseReason.String(), curRingType.String(), ringType.String())

	flag := propertyManager.Cost(bindGold, buyGold, goldUseReason, goldReason, silver, silverUseReason, silverReason)
	if !flag {
		panic(fmt.Errorf("marry:花费钱应该成功"))
	}
	//同步物品
	propertylogic.SnapChangedProperty(pl)

	//婚戒替换
	//获取婚烟对戒
	itemTempalte := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
	itemId := int32(itemTempalte.TemplateId())

	manager.RingReplace(ringType)
	marrylogic.MarryPropertyChanged(pl)
	scMarryRingReplace := pbuitl.BuildSCMarryRingReplace(itemId)
	pl.SendMsg(scMarryRingReplace)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return
	}
	scMarryRingChange := pbuitl.BuildSCMarryRingChange(pl.GetId(), itemId)
	spl.SendMsg(scMarryRingChange)
	return
}
