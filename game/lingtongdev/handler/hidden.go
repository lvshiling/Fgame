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
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_HIDDEN_TYPE), dispatch.HandlerFunc(handleLingTongHidden))
}

//处理灵童养成隐藏展示信息
func handleLingTongHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csLingTongDevHidden := msg.(*uipb.CSLingTongDevHidden)
	classType := csLingTongDevHidden.GetClassType()
	hiddenFlag := csLingTongDevHidden.GetHidden()

	err = lingTongHidden(tpl, types.LingTongDevSysType(classType), hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"classType":  classType,
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("lingtongdev:处理灵童养成隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"classType":  classType,
			"hiddenFlag": hiddenFlag,
		}).Debug("lingtongdev:处理灵童养成隐藏展示信息完成")
	return nil

}

//灵童养成隐藏展示的逻辑
func lingTongHidden(pl player.Player, classType types.LingTongDevSysType, hiddenFlag bool) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("LingTongDev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//灵兵不能隐藏
	if classType == types.LingTongDevSysTypeLingBing {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("LingTongDev:灵兵无法隐藏")
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"classType":  classType,
			"hiddenFlag": hiddenFlag,
		}).Warn("lingtongdev:请先激活灵童养成类系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		return
	}
	manager.Hidden(classType, hiddenFlag)
	scLingTongDevHidden := pbutil.BuildSCLingTongDevHidden(int32(classType), hiddenFlag)
	pl.SendMsg(scLingTongDevHidden)
	return
}
