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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BUY_HUIYUAN_TYPE), dispatch.HandlerFunc(handlerBuyHuiYuan))
}

//处理购买会员
func handlerBuyHuiYuan(s session.Session, msg interface{}) (err error) {
	log.Debug("huiyuan:处理购买会员请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csBuyHuiYuan := msg.(*uipb.CSBuyHuiYuan)
	typ := csBuyHuiYuan.GetHuiyuanType()

	vipType := huiyuantypes.HuiYuanType(typ)
	if !vipType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"vipType":  vipType,
			}).Warn("huiyuan:购买会员请求，类型错误")

		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = buyHuiYuan(tpl, vipType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("huiyuan:处理购买会员请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("huiyuan:处理购买HuiYuan请求完成")

	return
}

//购买会员请求逻辑
func buyHuiYuan(pl player.Player, huiyuanType huiyuantypes.HuiYuanType) (err error) {
	houtaiType := center.GetCenterService().GetZhiZunType()
	temp := huiyuantemplate.GetHuiYuanTemplateService().GetHuiYuanTemplate(houtaiType, huiyuanType)
	if temp == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"huiyuanType": huiyuanType,
			}).Warn("huiyuan:购买会员请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否会员
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	if huiyuanManager.IsHuiYuan(huiyuanType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("huiyuan:购买会员请求，已经是会员")
		playerlogic.SendSystemMessage(pl, lang.HuiYuanHadBuyHuiYuan)
		return
	}

	needGold := int64(temp.NeedGold)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//元宝是否足够
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("huiyuan:购买会员请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyHuiYuan
	goldReasonText := fmt.Sprintf(goldReason.String(), huiyuanType)
	flag := propertyManager.CostGold(needGold, false, goldReason, goldReasonText)
	if !flag {
		panic(fmt.Errorf("huiyuan: buy huiyuan use gold should be ok"))
	}

	//更新
	expireTime := temp.Duration
	obj := huiyuanManager.BuyHuiYuan(huiyuanType, expireTime)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scBuyHuiYuan := pbutil.BuildSCBuyHuiYuan(obj.GetExpireTime(), houtaiType)
	pl.SendMsg(scBuyHuiYuan)
	return
}
