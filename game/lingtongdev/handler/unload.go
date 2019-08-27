package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_UNLOAD_TYPE), dispatch.HandlerFunc(handleLingTongDevUnload))
}

//处理灵童养成类卸下信息
func handleLingTongDevUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevUnload := msg.(*uipb.CSLingTongDevUnload)
	classType := csLingTongDevUnload.GetClassType()

	err = lingTongDevUnload(tpl, types.LingTongDevSysType(classType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"error":     err,
			}).Error("lingtongdev:处理灵童养成类卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Debug("lingtongdev:处理灵童养成类卸下信息完成")
	return nil

}

//灵童养成类卸下的逻辑
func lingTongDevUnload(pl player.Player, classType types.LingTongDevSysType) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("LingTongDev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongInfo := manager.GetLingTongDevInfo(classType)
	if lingTongInfo == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("lingtongdev:请先激活灵童养成类系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		return
	}
	if lingTongInfo.GetSeqId() == 0 {
		playerlogic.SendSystemMessage(pl, lang.LingTongDevUnrealNoExist)
		return
	}
	manager.Unload(classType)
	scLingTongDevUnload := pbutil.BuildSCLingTongDevUnload(int32(classType), 0)
	pl.SendMsg(scLingTongDevUnload)
	return
}
