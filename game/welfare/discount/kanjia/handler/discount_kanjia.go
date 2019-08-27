package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	discountkanjiatemplate "fgame/fgame/game/welfare/discount/kanjia/template"
	discountkanjiatypes "fgame/fgame/game/welfare/discount/kanjia/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_KANJIA_TYPE), dispatch.HandlerFunc(handlerKanJia))
}

//处理砍价礼包
func handlerKanJia(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理砍价礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityKanJia)
	groupId := csMsg.GetGroupId()
	giftType := csMsg.GetType()

	err = kanJia(tpl, groupId, giftType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理砍价礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理砍价礼包请求完成")

	return
}

//砍价礼包请求逻辑
func kanJia(pl player.Player, groupId int32, giftType int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeKanJia

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	bargainGroupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountBargainShopGroupTemplate(groupId)
	if bargainGroupTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:砍价礼包请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:砍价礼包请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	discountTemp := bargainGroupTemp.GetDiscountBargainTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
	if discountTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"giftType": giftType,
			}).Warn("welfare:砍价礼包请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*discountkanjiatypes.DiscountKanJiaInfo)
	curKanJiaTimes, _ := info.GetKanJiaInfo(giftType)
	newDiscount := discountTemp.GetDiscount(curKanJiaTimes + 1)
	if newDiscount < 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"giftType":    giftType,
				"kanjiaTimes": curKanJiaTimes + 1,
			}).Warn("welfare:砍价礼包请求,折扣模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//次数
	groupTemp := groupInterface.(*discountkanjiatemplate.GroupTemplateDiscountKanJia)
	hadTimes := info.CountRewTimes(groupTemp.GetBargainRewTimesNeedGold())
	if hadTimes < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"giftType": giftType,
			}).Warn("welfare:砍价礼包请求,没有砍价次数")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityDiscountKanJiaNotEnoughTimes)
		return
	}

	//砍价记录
	info.UpdateKanJiaRecord(giftType, newDiscount)
	welfareManager.UpdateObj(obj)

	scMsg := pbutil.BuildSCOpenActivityKanJia(groupId, info.KanJiaRecord)
	pl.SendMsg(scMsg)
	return
}
