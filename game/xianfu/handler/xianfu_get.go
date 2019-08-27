package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xianfu/pbutil"
	xianfuplayer "fgame/fgame/game/xianfu/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANFU_GET_TYPE), dispatch.HandlerFunc(handlerXianfuGet))
}

//秘境仙府信息处理
func handlerXianfuGet(s session.Session, msg interface{}) (err error) {
	log.Debug("xianfu:处理秘境仙府信息获取请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err = xianfuGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("xianfu:处理秘境仙府信息获取请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("xianfu：处理秘境仙府信息获取请求完成")

	return
}

//仙府信息处理逻辑
func xianfuGet(pl player.Player) (err error) {
	xianfuManager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	now := global.GetGame().GetTimeService().Now()

	//刷新数据
	err = xianfuManager.RefreshData(now)
	if err != nil {
		return
	}
	//获取信息List
	xianfuArr := xianfuManager.GetPlayerXianfuInfoList()

	scXianfuGet := pbutil.BuildSCXianfuGet(xianfuArr)
	pl.SendMsg(scXianfuGet)
	return
}
