package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/hongbao/hongbao"
	"fgame/fgame/game/hongbao/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HONGBAO_GET_TYPE), dispatch.HandlerFunc(handleHongBaoGet))

}

//红包查询
func handleHongBaoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("hongbao:红包查询")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csHongbaoGet := msg.(*uipb.CSHongbaoGet)
	hongBaoId := csHongbaoGet.GetHongBaoId()
	channelTypeInt := csHongbaoGet.GetChannel()

	channelType := chattypes.ChannelType(channelTypeInt)
	if !channelType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"channelType": int32(channelType),
			}).Warn("hongbao:红包查询,频道参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = hongBaoGet(tpl, hongBaoId, channelType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": int32(hongBaoId),
				"error":     err,
			}).Error("hongbao:红包查询,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("hongbao:红包查询完成")
	return nil

}

//红包查询逻辑
func hongBaoGet(pl player.Player, hongBaoId int64, channelType chattypes.ChannelType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeHongBao) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:红包查询，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	hongBaoService := hongbao.GetHongBaoService()
	hongBaoObj := hongBaoService.GetHongBaoObj(hongBaoId)
	if hongBaoObj == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"hongBaoId": hongBaoId,
			}).Warn("hongbao:红包查询，红包id错误")
		playerlogic.SendSystemMessage(pl, lang.HongBaoExpire)
		return
	}

	keepTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeHongBaoKeepTime))
	endTime := hongBaoObj.GetCreateTime() + keepTime
	//返回前端
	scMsg := pbutil.BuildSCHongBaoGet(hongBaoObj, endTime, channelType)
	pl.SendMsg(scMsg)
	return
}
