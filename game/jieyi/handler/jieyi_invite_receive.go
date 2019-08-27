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
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_HELP_INVITE_TYPE), dispatch.HandlerFunc(handleJieYiInviteReceive))
}

func handleJieYiInviteReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理接收结义请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiHandleInvite)
	inviteId := csMsg.GetPlayerId()
	agree := csMsg.GetAgree()
	name := csMsg.GetJieYiName()
	daoJu := jieyitypes.JieYiDaoJuType(csMsg.GetJieYiDaoJu())
	if !daoJu.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"daoJuType": int32(daoJu),
			}).Warn("jieyi: 结义道具类型不符")
		return
	}

	err = jieYiInviteReceive(tpl, inviteId, agree, daoJu, name)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"daoJuType": int32(daoJu),
				"err":       err,
			}).Error("jieyi: 开始处理接收结义请求消息,失败")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 开始处理接收结义请求消息,成功")

	return
}

func jieYiInviteReceive(pl player.Player, inviteId int64, agree bool, daoJu jieyitypes.JieYiDaoJuType, name string) (err error) {
	playerId := pl.GetId()
	inviteObj := jieyi.GetJieYiService().GetInviteData(inviteId)
	if inviteObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"inviteId": inviteId,
			}).Warn("jieyi: 邀请过期")
		playerlogic.SendSystemMessage(pl, lang.JieYiInivteExprise)
		return
	}

	name = inviteObj.GetJieYiName()
	daoJu = inviteObj.GetJieYiDaoJu()
	inviteeId := inviteObj.GetInviteeId()

	// 与本地数据进行验证
	if inviteeId != playerId {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"inviteId": inviteId,
				"daoJu":    daoJu.String(),
				"name":     name,
			}).Warn("jieyi: 参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	invitePl := player.GetOnlinePlayerManager().GetPlayerById(inviteId)
	// 判断对方
	// 判断对方是否在线
	if invitePl == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"inviteId": inviteId,
			}).Warn("jieyi: 对方不在线")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotOnline)
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return
	}

	// 被邀请者不同意
	if !agree {
		scMsg := pbutil.BuildSCJieYiHandleInvite(inviteId, false, 0)
		pl.SendMsg(scMsg)

		scNoticeMsg := pbutil.BuildSCJieYiHandleInviteNotice(pl.GetId(), pl.GetName(), false, invitePl.GetId(), int32(invitePl.GetSex()))
		invitePl.SendMsg(scNoticeMsg)

		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return
	}

	// 判断自己是否能结义
	if jieyi.GetJieYiService().IsAlreadyJieYi(playerId) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"inviteId": inviteId,
			}).Warn("jieyi: 自己已经结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiAlreadyJieYi)
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return
	}

	daoJuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(daoJu)
	if daoJuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"daojuType": int32(daoJu),
			}).Warn("jieyi: 结义道具模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return
	}

	itemMap := daoJuTemp.GetNeedItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"daojuType": int32(daoJu),
				}).Warn("jieyi: 背包内结义道具不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
			return
		}
	}
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	plObj := jieYiManager.GetPlayerJieYiObj()
	inviteeDaoJu := plObj.GetDaoJuType()
	inviteeToken := plObj.GetTokenType()
	inviteeTokenLev := plObj.GetTokenLevel()
	nameLev := plObj.GetNameLev()
	// 消耗物品
	jieYi, inviteMemObj, flag := jieyi.GetJieYiService().JieYiSuccess(inviteId, inviteeId, inviteeDaoJu, inviteeToken, inviteeTokenLev, nameLev)
	if !flag {
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return
	}

	reason := commonlog.InventoryLogReasonInviteJieYiUse
	reasonText := fmt.Sprintf(reason.String(), daoJu.String(), jieyitypes.JieYiItemUseTypeInvite.String())
	flag = inventoryManager.BatchRemove(itemMap, reason, reasonText)
	if !flag {
		panic("jieyi: 消耗物品应该成功")
	}

	// 推送变化
	inventorylogic.SnapInventoryChanged(pl)

	// 获取结义兄弟列表
	scJieYiHandleInviteNotice := pbutil.BuildSCJieYiHandleInviteNotice(pl.GetId(), pl.GetName(), agree, inviteId, int32(pl.GetSex()))
	jieyilogic.BroadcastJieYi(jieYi, scJieYiHandleInviteNotice)

	//结义成员变化
	jieyilogic.JieYiMemberChanged(jieYi)

	scMsg := pbutil.BuildSCJieYiHandleInvite(inviteId, agree, int32(inviteMemObj.GetSex()))
	pl.SendMsg(scMsg)

	return
}
