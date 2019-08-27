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
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_LAO_DA_TI_REN_TYPE), dispatch.HandlerFunc(handleJieYiTiRen))
}

func handleJieYiTiRen(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理结义老大踢人请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiLaoDaTiRen)
	receiverId := csMsg.GetPlayerId()

	err = jieYiTiRen(tpl, receiverId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理结义老大踢人请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理结义老大踢人请求消息,成功")

	return
}

func jieYiTiRen(pl player.Player, receiverId int64) (err error) {
	playerId := pl.GetId()
	if !jieyi.GetJieYiService().IsJieYiLaoDa(playerId) {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"receiverId": receiverId,
			}).Warn("jieyi: 玩家不是老大")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotIsLaoDa)
		return
	}

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(playerId)
	receiveObj := jieyi.GetJieYiService().GetJieYiMemberInfo(receiverId)
	if plObj.GetJieYiId() != receiveObj.GetJieYiId() {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"receiverId": receiverId,
			}).Warn("jieyi: 不在同一结义阵营")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotIsSameJieYi)
		return
	}
	jieYi, tiMemList, flag := jieyi.GetJieYiService().JieYiTiRen(playerId, receiverId)
	if !flag {
		return
	}

	// 推送给被踢的人
	scJieYiLaoDaTiRenOtherNotice := pbutil.BuildSCJieYiLaoDaTiRenOtherNotice(pl.GetId(), pl.GetName())
	if len(tiMemList) > 1 {
		//解散

	} else {
		tiPlayerName := ""
		for _, mem := range tiMemList {
			tiPlayerName = mem.GetPlayerName()
		}
		// 推送给兄弟
		scJieYiLaoDaTiRenNotice := pbutil.BuildSCJieYiLaoDaTiRenNotice(receiverId, tiPlayerName, pl.GetId(), pl.GetName())
		jieyilogic.BroadcastJieYiExculdeSelf(jieYi, nil, scJieYiLaoDaTiRenNotice)
	}
	// 推送给结义兄弟并刷新属性
	jieyilogic.JieYiMemberChanged(jieYi)

	//更新被踢的数据
	for _, mem := range tiMemList {
		tiPlayer := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if tiPlayer != nil {
			ctx := scene.WithPlayer(context.Background(), tiPlayer)
			msg := message.NewScheduleMessage(onJieYiJieChu, ctx, nil, nil)
			tiPlayer.Post(msg)
			if tiPlayer == pl {
				continue
			}
			tiPlayer.SendMsg(scJieYiLaoDaTiRenOtherNotice)
		}
	}

	scMsg := pbutil.BuildSCJieYiLaoDaTiRen(receiverId)
	pl.SendMsg(scMsg)

	return
}

func onJieYiJieChu(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)

	// 刷新被邀请人数据
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieYiManager.JieChuSuccess()
	//属性变化
	jieyilogic.JieYiPropertyChange(pl)
	return nil
}
