package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/huiyuan/pbutil"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HUIYUAN_INFO_TYPE), dispatch.HandlerFunc(handlerHuiYuanGetInfo))
}

//处理获取会员信息
func handlerHuiYuanGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("huiyuan:处理获取会员请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = getHuiYuanInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("huiyuan:处理获取会员请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("huiyuan:处理获取会员请求完成")

	return
}

//获取会员请求逻辑
func getHuiYuanInfo(pl player.Player) (err error) {
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	err = huiyuanManager.RefresHuiYuanRewards()
	if err != nil {
		return
	}

	isHuiyuanInterim := huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	isReceiveInterim := huiyuanManager.IsReceiveRewards(huiyuantypes.HuiYuanTypeInterim)
	expireTime := huiyuanManager.GetHuiYuanExpireTiem()
	isBuyTodayInterim := huiyuanManager.IsFirstRew(huiyuantypes.HuiYuanTypeInterim)

	isReceivePlus := huiyuanManager.IsReceiveRewards(huiyuantypes.HuiYuanTypePlus)
	isHuiyuanPlus := huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	isBuyTodayPlus := huiyuanManager.IsFirstRew(huiyuantypes.HuiYuanTypePlus)
	houtaiType := center.GetCenterService().GetZhiZunType()

	scHuiYuanInfo := pbutil.BuildSCHuiYuanInfo(isHuiyuanInterim, isReceiveInterim, isReceivePlus, isHuiyuanPlus, isBuyTodayInterim, isBuyTodayPlus, expireTime, houtaiType)
	pl.SendMsg(scHuiYuanInfo)
	return
}
