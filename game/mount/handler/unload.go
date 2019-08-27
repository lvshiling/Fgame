package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_UNLOAD_TYPE), dispatch.HandlerFunc(handleMountUnload))
}

//处理坐骑卸下信息
func handleMountUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = mountUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mount:处理坐骑卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mount:处理坐骑卸下信息完成")
	return nil

}

//坐骑卸下的逻辑
func mountUnload(pl player.Player) (err error) {
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	obj := mountManager.GetMountInfo()
	if obj.MountId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("mount:处理坐骑卸下,没有坐骑")
		playerlogic.SendSystemMessage(pl, lang.MountUnrealNoExist)
		return
	}

	mountManager.Unload()
	scMountUnload := pbutil.BuildSCMountUnload(0)
	pl.SendMsg(scMountUnload)
	return
}
