package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_GET_TYPE), dispatch.HandlerFunc(handleMountGet))
}

//处理坐骑信息
func handleMountGet(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理获取坐骑消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = mountGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mount:处理获取坐骑消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mount:处理获取坐骑消息完成")
	return nil

}

//获取坐骑界面信息的逻辑
func mountGet(pl player.Player) (err error) {
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	mountOtherMap := mountManager.GetMountOtherMap()
	scMountGet := pbutil.BuildSCMountGet(mountInfo, mountOtherMap)
	pl.SendMsg(scMountGet)
	return
}
