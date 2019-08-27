package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	coredirty "fgame/fgame/core/dirty"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_INVITE_TYPE), dispatch.HandlerFunc(handleJieYiInvite))
}

func handleJieYiInvite(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理结义成员请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiInvite)
	daoJuType := jieyitypes.JieYiDaoJuType(csMsg.GetJieYiDaoJu())
	if !daoJuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"daoJuType": daoJuType.String(),
			}).Warn("jieyi: 结义道具类型不符")
		return
	}

	name := csMsg.GetName()
	inviteeId := csMsg.GetPlayerId()

	err = jieYiInvite(tpl, daoJuType, name, inviteeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理结义成员消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理结义成员消息,成功")

	return
}

const (
	minJieYiNameLen = 2
	maxJieYiNameLen = 6
)

func jieYiInvite(pl player.Player, daoJuType jieyitypes.JieYiDaoJuType, name string, inviteeId int64) (err error) {
	playerId := pl.GetId()
	// 验证对方功能开启
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeJieYi) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("jieyi: 结义功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return nil
	}

	//是否是老大
	if !jieyi.GetJieYiService().IsJieYiLaoDa(playerId) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"err":      err,
			}).Warn("jieyi: 自己不是结义老大，无法邀请")
		playerlogic.SendSystemMessage(pl, lang.JieYiCanNotInvite)
		return
	}
	//是否是满人
	if jieyi.GetJieYiService().IsFullMember(playerId) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"err":      err,
			}).Warn("jieyi: 结义满人")
		playerlogic.SendSystemMessage(pl, lang.JieYiMemberAlreadyFull)
		return
	}

	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	plObj := jieYiManager.GetPlayerJieYiObj()
	inviteDaoJu := plObj.GetDaoJuType()
	inviteToken := plObj.GetTokenType()
	inviteTokenLev := plObj.GetTokenLevel()
	nameLev := plObj.GetNameLev()
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	//判断上次解除结义时间
	if !jieYiManager.IsCanJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"err":      err,
			}).Warn("jieyi: 自己解除结义时间cd中")
		playerlogic.SendSystemMessage(pl, lang.JieYiJieChuJieYiCD, fmt.Sprintf("%d", constantTemp.JieChuCD/int64(common.HOUR)))
		return
	}

	//判断自身邀请是否处于cd时间
	if !jieYiManager.IsCanInvite() {
		log.WithFields(
			log.Fields{
				"playerId":  playerId,
				"daojuType": int32(daoJuType),
			}).Warn("jieyi: 邀请CD中")
		playerlogic.SendSystemMessage(pl, lang.JieYiInviteTimeCD, fmt.Sprintf("%d", constantTemp.YaoQingCD/int64(common.MINUTE)))
		return
	}

	// 判断名字是否合法
	name = strings.TrimSpace(name)
	nameLen := len([]rune(name))
	if !jieyi.GetJieYiService().IsAlreadyJieYi(playerId) {
		if nameLen < minJieYiNameLen || nameLen > maxJieYiNameLen {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"name":     name,
				}).Warn("jieyi:处理结义创建,名字不合法")
			playerlogic.SendSystemMessage(pl, lang.JieYiNameIllegal)
			return
		}

		flag := coredirty.GetDirtyService().IsLegal(name)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"name":     name,
				}).Warn("jieyi:处理结义创建,名字含有脏字")
			playerlogic.SendSystemMessage(pl, lang.JieYiNameDirty)
			return
		}
		//判断名字是否重复
		if !jieyi.GetJieYiService().IsNameRepetitive(name) {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"name":     name,
				}).Warn("jieyi:处理结义创建,威名重复")
			playerlogic.SendSystemMessage(pl, lang.JieYiNameRepetitive)
			return
		}
	}

	

	// 判断对方
	invitee := player.GetOnlinePlayerManager().GetPlayerById(inviteeId)
	if invitee == nil {
		log.WithFields(
			log.Fields{
				"playerId":  playerId,
				"inviteeId": inviteeId,
			}).Warn("jieyi: 对方已经离线")
		playerlogic.SendSystemMessage(pl, lang.PlayerOffline)
		return
	}

	if jieyi.GetJieYiService().IsAlreadyJieYi(inviteeId) {
		log.WithFields(
			log.Fields{
				"playerId":  playerId,
				"inviteeId": inviteeId,
			}).Warn("jieyi: 对方已经结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiDuiFangAlreadyJieYi)
		return
	}

	//判断物品足够不
	daoJuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(daoJuType)
	if daoJuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"daojuType": int32(daoJuType),
			}).Warn("jieyi: 结义道具模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		return
	}

	itemMap := daoJuTemp.GetNeedItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"daojuType": int32(daoJuType),
				}).Warn("jieyi: 背包内结义道具不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	// 添加结义邀请数据
	flag := jieyi.GetJieYiService().AddJieYiInviteData(playerId, inviteeId, name, daoJuType, inviteDaoJu, inviteToken, inviteTokenLev, nameLev)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"daojuType": int32(daoJuType),
			}).Warn("jieyi: 邀请失败")
		playerlogic.SendSystemMessage(pl, lang.JieYiInivteFail)
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":  playerId,
			"inviteeId": inviteeId,
		}).Info("jieyi: 邀请成功")
	jieYiManager.InviteSucess()

	reason := commonlog.InventoryLogReasonInviteJieYiUse
	reasonText := fmt.Sprintf(reason.String(), daoJuType.String(), jieyitypes.JieYiItemUseTypeInvite.String())
	flag = inventoryManager.BatchRemove(itemMap, reason, reasonText)
	if !flag {
		panic("jieyi: 消耗物品应该成功")
	}

	// 推送变化
	inventorylogic.SnapInventoryChanged(pl)

	data := &jieyitypes.JieYiInviteData{
		PlayerId:  playerId,
		DaoJuType: daoJuType,
		Name:      name,
	}

	ctx := scene.WithPlayer(context.Background(), invitee)
	msg := message.NewScheduleMessage(inviteeHandleJieYi, ctx, data, nil)
	invitee.Post(msg)

	return
}

// 验证被邀请人
func inviteeHandleJieYi(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	playerId := pl.GetId()
	data := result.(*jieyitypes.JieYiInviteData)
	inviteId := data.PlayerId
	invitePl := player.GetOnlinePlayerManager().GetPlayerById(inviteId)
	if invitePl == nil {
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return nil
	}

	// 验证对方功能开启
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeJieYi) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("jieyi: 对方结义功能未开启")
		playerlogic.SendSystemMessage(invitePl, lang.CommonFuncNoOpen)
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return nil
	}

	//验证对方离开上一个结义时间是否足够长
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	if !jieYiManager.IsCanJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"inviteId": inviteId,
				"err":      err,
			}).Warn("jieyi: 对方解除结义时间cd中")
		playerlogic.SendSystemMessage(invitePl, lang.JieYiDuiFangJieChuJieYiCD, fmt.Sprintf("%d", constantTemp.JieChuCD/int64(common.HOUR)))
		jieyi.GetJieYiService().JieYiInviteFail(inviteId, playerId)
		return nil
	}

	memberNum := int32(0)
	jieYiObj := jieyi.GetJieYiService().GetJieYiMemberInfo(inviteId)
	if jieYiObj != nil {
		jieYiId := jieYiObj.GetJieYiId()
		memberNum = jieyi.GetJieYiService().GetJieYiMemberNum(jieYiId)
	}

	scJieYiInviteNotice := pbutil.BuildSCJieYiInviteNotice(inviteId, invitePl.GetName(), int32(invitePl.GetSex()), int32(invitePl.GetRole()), invitePl.GetForce(), data.Name, memberNum, int32(data.DaoJuType))
	pl.SendMsg(scJieYiInviteNotice)

	//通知被邀请人
	scJieYiInvite := pbutil.BuildSCJieYiInvite(int32(data.DaoJuType), data.Name)
	invitePl.SendMsg(scJieYiInvite)

	return nil
}
