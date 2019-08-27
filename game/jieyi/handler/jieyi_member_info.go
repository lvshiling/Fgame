package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_MEMBER_INFO_TYPE), dispatch.HandlerFunc(handleJieYiMemberInfo))
}

func handleJieYiMemberInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理结义成员请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = getJieYiMemberInfo(tpl)
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

func getJieYiMemberInfo(pl player.Player) (err error) {
	playerJieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !playerJieYiManager.IsJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	playerJieYiObj := playerJieYiManager.GetPlayerJieYiObj()
	jieYiObj := jieyi.GetJieYiService().GetJieYiInfo(playerJieYiObj.GetJieYiId())
	if jieYiObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}
	memberObjList := jieyi.GetJieYiService().GetJieYiMemberList(playerJieYiObj.GetJieYiId())
	tokenPro := playerJieYiObj.GetTokenPro()
	shengWei := playerJieYiObj.GetShengWeiZhi()

	scMsg := pbutil.BuildSCJieYiMemberInfo(jieYiObj, memberObjList, pl.GetId(), tokenPro, shengWei)
	pl.SendMsg(scMsg)

	return
}
