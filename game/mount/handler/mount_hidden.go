package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_HIDDEN_TYPE), dispatch.HandlerFunc(handleMountHidden))
}

//处理坐骑隐藏展示信息
func handleMountHidden(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑隐藏展示信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMountHidden := msg.(*uipb.CSMountHidden)
	hiddenFlag := csMountHidden.GetHidden()

	err = mountHidden(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("mount:处理坐骑隐藏展示信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("mount:处理坐骑隐藏展示信息完成")
	return nil

}

//坐骑隐藏展示的逻辑
func mountHidden(pl player.Player, hiddenFlag bool) (err error) {

	if pl.GetMountId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
			}).Warn("mount:处理坐骑隐藏展示信息,没有坐骑")
		return
	}

	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
			}).Warn("mount:处理坐骑隐藏展示信息,玩家没有在场景中")
		return
	}

	if s.MapTemplate().LimitRideHorse != 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
			}).Warn("mount:处理坐骑隐藏展示信息,地图禁止上马")
		return
	}

	if pl.IsPvpBattle() && !hiddenFlag {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
			}).Warn("mount:处理坐骑隐藏展示信息,pvp中不能上马")
		return
	}

	if ((pl.GetBattleLimit() & scenetypes.BattleLimitTypeMount.Mask()) != 0) && !hiddenFlag {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
			}).Warn("mount:处理坐骑隐藏展示信息,buff中")
		return
	}

	// if pl.IsMountHidden() == hiddenFlag {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":   pl.GetId(),
	// 			"hiddenFlag": hiddenFlag,
	// 		}).Warn("mount:处理坐骑隐藏展示信息,重复操作")
	// 	return
	// }

	pl.MountHidden(hiddenFlag)

	return
}
