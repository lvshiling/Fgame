package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/huiyuan/pbutil"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantemplate "fgame/fgame/game/huiyuan/template"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HUIYUAN_RECEIVE_RWE_TYPE), dispatch.HandlerFunc(handlerHuiYuanReceiveRew))
}

//处理领取会员奖励
func handlerHuiYuanReceiveRew(s session.Session, msg interface{}) (err error) {
	log.Debug("huiyuan:处理领取会员奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csHuiYuanReceiveRew := msg.(*uipb.CSHuiYuanReceiveRew)
	typ := csHuiYuanReceiveRew.GetHuiyuanType()

	huiyuanType := huiyuantypes.HuiYuanType(typ)
	if !huiyuanType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"huiyuanType": huiyuanType,
			}).Warn("huiyuan:领取会员奖励请求，参数错误")

		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = huiYuanReceive(tpl, huiyuanType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("huiyuan:处理领取会员奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("huiyuan:处理购买HuiYuan请求完成")

	return
}

//领取会员奖励请求逻辑
func huiYuanReceive(pl player.Player, huiyuanType huiyuantypes.HuiYuanType) (err error) {
	houtaiType := center.GetCenterService().GetZhiZunType()
	temp := huiyuantemplate.GetHuiYuanTemplateService().GetHuiYuanTemplate(houtaiType, huiyuanType)
	if temp == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"huiyuanType": huiyuanType,
			}).Warn("huiyuan:领取会员奖励请求，模板不存在")

		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)

	//是否会员
	if !huiyuanManager.IsHuiYuan(huiyuanType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("huiyuan:领取会员奖励请求，不是会员")

		playerlogic.SendSystemMessage(pl, lang.HuiYuanNotHuiYuan)
		return
	}

	//是否领取
	if huiyuanManager.IsReceiveRewards(huiyuanType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("huiyuan:领取会员奖励请求，已领取奖励")

		playerlogic.SendSystemMessage(pl, lang.HuiYuanHadReceiveRewards)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//背包空间
	var rewItemMap map[int32]int32
	if huiyuanManager.IsFirstRew(huiyuanType) {
		rewItemMap = temp.GetRewFirstItemMap()
	} else {
		rewItemMap = temp.GetRewItemMap()
	}
	if !inventoryManager.HasEnoughSlots(rewItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("huiyuan:领取会员奖励请求，背包空间不足")

		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonHuiYuanRew
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), huiyuanType)
	flag := inventoryManager.BatchAdd(rewItemMap, itemGetReason, itemGetReasonText)
	if !flag {
		panic("huiyuan:huiyuan rewards add item should be ok")
	}

	reasonGold := commonlog.GoldLogReasonHuiYuanRew
	reasonSilver := commonlog.SilverLogReasonHuiYuanRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), huiyuanType)
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), huiyuanType)

	rewSilver := temp.RewSilver
	rewBindGold := temp.RewBindGold
	rewGold := temp.RewGold
	rewExp := int32(0)
	rewExpPoint := int32(0)
	flag = propertyManager.AddMoney(int64(rewBindGold), int64(rewGold), reasonGold, reasonGoldText, int64(rewSilver), reasonSilver, reasonSilverText)
	if !flag {
		panic("huiyuan:huiyuan rewards add RewData should be ok")
	}

	//更新
	huiyuanManager.ReceiveHuiYuanRewards(huiyuanType)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	scHuiYuanReceiveRew := pbutil.BuildSCHuiYuanReceiveRew(totalRewData, rewItemMap, houtaiType)
	pl.SendMsg(scHuiYuanReceiveRew)
	return
}
