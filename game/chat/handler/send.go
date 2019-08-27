package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/chat/chat"
	chateventtypes "fgame/fgame/game/chat/event/types"
	chatlogic "fgame/fgame/game/chat/logic"
	"fgame/fgame/game/chat/pbutil"
	playerchat "fgame/fgame/game/chat/player"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHAT_SEND_TYPE), dispatch.HandlerFunc(handleChatSend))
}

//处理聊天信息
func handleChatSend(s session.Session, msg interface{}) (err error) {
	log.Debug("chat:处理聊天信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csChatSend := msg.(*uipb.CSChatSend)
	channel := chattypes.ChannelType(csChatSend.GetChannel())
	recvId := csChatSend.GetRecvId()
	msgType := chattypes.MsgType(csChatSend.GetMsgType())
	content := csChatSend.GetContent()
	args := csChatSend.GetArgs()

	if !channel.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"channel":  csChatSend.GetChannel(),
				"recvId":   recvId,
				"msgType":  csChatSend.GetMsgType(),
			}).Warn("chat:处理聊天信息,频道错误")
		return
	}
	if !msgType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"channel":  csChatSend.GetChannel(),
				"recvId":   recvId,
				"msgType":  csChatSend.GetMsgType(),
			}).Warn("chat:处理聊天信息,消息类型错误")

		return
	}

	err = chatSend(tpl, channel, recvId, msgType, content, args)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"channel":  channel.String(),
				"error":    err,
			}).Error("chat:处理聊天信息,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chat:处理聊天信息完成")
	return nil
}

//处理消息发送
func chatSend(pl player.Player, channel chattypes.ChannelType, recvId int64, msgType chattypes.MsgType, content []byte, args string) (err error) {
	// 禁言时段
	if chatlogic.IsForbiddenTime() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"channel":  channel.String(),
				"recvId":   recvId,
			}).Warn("chat:处理聊天信息,系统当前时段禁言中")
		playerlogic.SendSystemMessage(pl, lang.ChatForbidden)
		return
	}

	if pl.IsForbidChat() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"channel":  channel.String(),
				"recvId":   recvId,
			}).Warn("chat:处理聊天信息,玩家正在禁言中")
		playerlogic.SendSystemMessage(pl, lang.PlayerInForbidChat)
		return
	}
	chatRecv := pbutil.BuildSCChatRecvWithCliArgs(pl.GetId(), pl.GetName(), channel, recvId, msgType, content, args)
	if pl.IsIgnoreChat() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"channel":  channel.String(),
				"recvId":   recvId,
			}).Warn("chat:处理聊天信息,玩家正在禁默中")
		chatSend := pbutil.BuildSCChatSendWithChatCount(channel, recvId, msgType, content, args, 0)
		pl.SendMsg(chatSend)
		return
	}
	vip := pl.GetVip()
	level := pl.GetLevel()

	recvName := ""
	switch channel {
	case chattypes.ChannelTypeWorld:
		//判断等级和vip
		minLevel := chat.GetChatService().GetWorldLevel()
		minVip := chat.GetChatService().GetWorldVipLevel()
		if vip < minVip && level < minLevel {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"channel":  channel.String(),
					"recvId":   recvId,
					"level":    level,
					"vip":      vip,
				}).Warn("chat:处理聊天信息,等级不够或vip等级不够")

			playerlogic.SendSystemMessage(pl, lang.ChatLevelOrVipLow, fmt.Sprintf("%d", minLevel), fmt.Sprintf("%d", minVip))
			return
		}
		//TODO 验证cd和金钱
		excludeIdList := []int64{pl.GetId()}
		player.GetOnlinePlayerManager().BroadcastMsgExclude(excludeIdList, chatRecv)
		chat.GetChatService().AddWorldChat(pl.GetId(), content, args, msgType)
		break
	case chattypes.ChannelTypeBangPai:
		minLevel := chat.GetChatService().GetAllianceLevel()
		minVip := chat.GetChatService().GetAllianceVipLevel()
		if vip < minVip && level < minLevel {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"channel":  channel.String(),
					"recvId":   recvId,
					"level":    level,
					"vip":      vip,
				}).Warn("chat:处理聊天信息,等级不够或vip等级不够")
			playerlogic.SendSystemMessage(pl, lang.ChatLevelOrVipLow, fmt.Sprintf("%d", minLevel), fmt.Sprintf("%d", minVip))
			return
		}
		//TODO 优化
		allianceId := pl.GetAllianceId()
		if allianceId == 0 {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"channel":  channel.String(),
					"recvId":   recvId,
				}).Warn("chat:处理聊天信息,不在帮派")
			playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
			return
		}
		recvName = pl.GetAllianceName()
		chatlogic.BroadcastAllianceExcludeSelf(allianceId, pl.GetId(), pl.GetName(), msgType, content, args)
		break
	case chattypes.ChannelTypeTeam:
		teamId := pl.GetTeamId()
		if teamId == 0 {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"channel":  channel.String(),
					"recvId":   recvId,
				}).Warn("chat:处理聊天信息,不在队伍")
			playerlogic.SendSystemMessage(pl, lang.TeamPlayerInTeam)
			return
		}
		//判断等级和vip
		minLevel := chat.GetChatService().GetTeamLevel()
		minVip := chat.GetChatService().GetTeamVipLevel()
		if vip < minVip && level < minLevel {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"channel":  channel.String(),
					"recvId":   recvId,
					"level":    level,
					"vip":      vip,
				}).Warn("chat:处理聊天信息,等级不够或vip等级不够")
			playerlogic.SendSystemMessage(pl, lang.ChatLevelOrVipLow, fmt.Sprintf("%d", minLevel), fmt.Sprintf("%d", minVip))
			return
		}

		recvName = pl.GetName()
		chatlogic.BroadcastTeamExcludeSelf(teamId, pl.GetId(), pl.GetName(), msgType, content, args)
		break
	case chattypes.ChannelTypePerson:
		minLevel := chat.GetChatService().GetPrivateLevel()
		minVip := chat.GetChatService().GetPrivateVipLevel()
		if vip < minVip && level < minLevel {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"channel":  channel.String(),
					"recvId":   recvId,
					"level":    level,
					"vip":      vip,
				}).Warn("chat:处理聊天信息,等级不够或vip等级不够")
			playerlogic.SendSystemMessage(pl, lang.ChatLevelOrVipLow, fmt.Sprintf("%d", minLevel), fmt.Sprintf("%d", minVip))
			return
		}
		recvPlayer := player.GetOnlinePlayerManager().GetPlayerById(recvId)
		if recvPlayer != nil {
			recvPlayer.SendMsg(chatRecv)
			recvName = recvPlayer.GetName()
		}

		break
	}

	//发言奖励
	chatCount := int32(0)
	if channel == chattypes.ChannelTypeWorld || channel == chattypes.ChannelTypeBangPai {
		playerChatManager := pl.GetPlayerDataManager(playertypes.PlayerChatDataManagerType).(*playerchat.PlayerChatDataManager)
		chatCount = playerChatManager.GetChatCount()
		if !playerChatManager.IsChatCountReachLimit() {
			chatCount = playerChatManager.AddChatCount()
			num := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatAwardSilver)
			propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
			goldReason := commonlog.GoldLogReasonChatAward
			silverReason := commonlog.SilverLogReasonChatAward
			goldReasonText := goldReason.String()
			silverReasonText := silverReason.String()
			flag := propertyManager.AddMoney(0, 0, goldReason, goldReasonText, int64(num), silverReason, silverReasonText)
			if !flag {
				return fmt.Errorf("chat: AddMoney should be ok ")
			}
			//同步
			propertylogic.SnapChangedProperty(pl)
		}
	}

	eventData := chat.CreateChatEventData(channel, recvId, recvName, msgType, content)
	//发送事件
	gameevent.Emit(chateventtypes.EventTypeChat, pl, eventData)

	chatSend := pbutil.BuildSCChatSendWithChatCount(channel, recvId, msgType, content, args, chatCount)
	pl.SendMsg(chatSend)

	return nil
}
