package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fmt"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_QIU_YUAN_TYPE), dispatch.HandlerFunc(handlePlayerQiuYuan))
}

func handlePlayerQiuYuan(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家求援请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	p := gcs.Player()
	pl := p.(player.Player)

	err = playerQiuYuan(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家求援请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("jieyi: 处理玩家求援请求消息,成功")

	return
}

func playerQiuYuan(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi:玩家求援，不在场景中")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}

	if pl.IsCross() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi:玩家处于跨服中，无法求援")
		playerlogic.SendSystemMessage(pl, lang.PlayerInCross)
		return
	}

	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !jieYiManager.IsJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi:玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	obj := jieYiManager.GetPlayerJieYiObj()
	lastQiuYuanTime := obj.GetLastQiuYuanTime()
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()
	if constantTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi:模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		return
	}

	now := global.GetGame().GetTimeService().Now()
	jianGetime := now - lastQiuYuanTime
	if jianGetime < constantTemp.QiuYuanCD {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"lastQiuYuanTime": lastQiuYuanTime,
			}).Warn("jieyi:求援CD中")
		playerlogic.SendSystemMessage(pl, lang.JieYiQiuYuanTimeCD, fmt.Sprintf("%d", constantTemp.QiuYuanCD/(1000*60)))
		return
	}

	// 推送给结义兄弟
	memberList := jieyi.GetJieYiService().GetJieYiMemberList(obj.GetJieYiId())
	scNoticeMsg := pbutil.BuildSCJieYiQiuYuanNotice(pl.GetName(), s.MapId(), pl.GetPos())
	for _, member := range memberList {
		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
		if memberPl != nil {
			// 不推送给求援玩家
			if memberPl.GetId() == pl.GetId() {
				continue
			}
			memberPl.SendMsg(scNoticeMsg)
		}
	}

	scMsg := pbutil.BuildSCJieYiQiuYuan(lastQiuYuanTime)
	pl.SendMsg(scMsg)

	return
}
