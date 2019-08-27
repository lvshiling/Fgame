package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_SUO_YAO_TYPE), dispatch.HandlerFunc(handleJieYiTokenSuoYao))
}

func handleJieYiTokenSuoYao(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理信物索要请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiTokenSuoYao)
	receiverId := csMsg.GetPlayerId()
	leaveWord := csMsg.GetLeaveWord()
	token := jieyitypes.JieYiTokenType(csMsg.GetToken())
	if !token.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"token":    int32(token),
			}).Warn("jieyi: 信物类型不合法")
		return
	}

	err = jieYiTokenSuoYao(tpl, receiverId, token, leaveWord)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理信物索要请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理信物索要请求消息,成功")
	return
}

func jieYiTokenSuoYao(pl player.Player, receiverId int64, token jieyitypes.JieYiTokenType, leaveWord string) (err error) {
	receivePl := player.GetOnlinePlayerManager().GetPlayerById(receiverId)
	if receivePl == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 对方不在线")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotOnline)
		return
	}

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	receiveObj := jieyi.GetJieYiService().GetJieYiMemberInfo(receiverId)
	if plObj == nil || receiveObj == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	if plObj.GetJieYiId() != receiveObj.GetJieYiId() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 不在同一结义阵营")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotIsSameJieYi)
		return
	}

	if receivePl != nil {
		scMsg := pbutil.BuildSCJieYiTokenSuoYaoNotice(int32(token), pl.GetId(), pl.GetName(), leaveWord)
		receivePl.SendMsg(scMsg)
	}

	scMsg := pbutil.BuildSCJieYiTokenSuoYao(receiverId, int32(token), leaveWord)
	pl.SendMsg(scMsg)

	return
}
