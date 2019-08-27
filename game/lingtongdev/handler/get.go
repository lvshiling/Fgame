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
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_GET_TYPE), dispatch.HandlerFunc(handleLingTongDevGet))

}

//处理灵童养成类信息
func handleLingTongDevGet(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理获取灵童养成类消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevGet := msg.(*uipb.CSLingTongDevGet)
	classType := csLingTongDevGet.GetClassType()
	err = lingTongDevGet(tpl, types.LingTongDevSysType(classType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"error":     err,
			}).Error("lingtongdev:处理获取灵童养成类消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Debug("lingtongdev:处理获取灵童养成类消息完成")
	return nil

}

//获取灵童养成类界面信息逻辑
func lingTongDevGet(pl player.Player, classType types.LingTongDevSysType) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("LingTongDev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("lingtongdev:请先激活灵童养成类系统")
		//playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		return
	}
	container := manager.GetLingTongDevOtherMap(classType)
	scLingTongDevGet := pbutil.BuildSCLingTongDevGet(int32(classType), lingTongDevInfo, container)
	pl.SendMsg(scLingTongDevGet)
	return
}
