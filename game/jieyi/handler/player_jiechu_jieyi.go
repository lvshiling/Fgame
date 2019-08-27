package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_JIE_CHU_TYPE), dispatch.HandlerFunc(handlePlayerJieChuJieYi))
}

func handlePlayerJieChuJieYi(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家解除结义送请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = playerJieChuJieYi(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家解除结义送请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理玩家解除结义送请求消息,成功")

	return
}

func playerJieChuJieYi(pl player.Player) (err error) {
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !jieYiManager.IsJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	// 刷新数据
	jieYi, tiMemList, flag := jieyi.GetJieYiService().JieChuJieYiSuccess(pl.GetId())
	if !flag {
		return
	}
	//解散
	scJieYiJieChuNotice := pbutil.BuildSCJieYiJieChuNotice(pl.GetId(), pl.GetName())
	if len(tiMemList) > 1 {
		for _, mem := range tiMemList {
			if mem.GetPlayerId() == pl.GetId() {
				continue
			}
			tiPlayer := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
			if tiPlayer != nil {
				tiPlayer.SendMsg(scJieYiJieChuNotice)
			}
		}
	} else {
		jieyilogic.BroadcastJieYiExculdeSelf(jieYi, nil, scJieYiJieChuNotice)
		jieyilogic.JieYiMemberChanged(jieYi)
	}

	//更新被踢的数据
	for _, mem := range tiMemList {
		tiPlayer := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if tiPlayer != nil {
			ctx := scene.WithPlayer(context.Background(), tiPlayer)
			msg := message.NewScheduleMessage(onJieYiJieChu, ctx, nil, nil)
			tiPlayer.Post(msg)
		}
	}

	scMsg := pbutil.BuildSCJieYiJieChu()
	pl.SendMsg(scMsg)
	return
}
