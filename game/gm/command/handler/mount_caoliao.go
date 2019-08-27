package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMountCaoLiao, command.CommandHandlerFunc(handleMountCaoLiao))

}

func handleMountCaoLiao(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	caoLiaoStr := c.Args[0]
	caoLiaoLevel, err := strconv.ParseInt(caoLiaoStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"caoLiaoLevel": caoLiaoLevel,
				"error":        err,
			}).Warn("gm:处理设置坐骑食草料等级,caoLiaoLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := mount.GetMountService().GetMountCaoLiaoTemplate(int32(caoLiaoLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"caoLiaoLevel": caoLiaoLevel,
				"error":        err,
			}).Warn("gm:处理设置坐骑食草料等级,caoLiaoLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	manager.GmSetMountCaoLiaoLevel(int32(caoLiaoLevel))

	//同步属性
	mountlogic.MountPropertyChanged(pl)

	scMountShiDan := pbutil.BuildSCMountCulDan(int32(caoLiaoLevel), 0)
	pl.SendMsg(scMountShiDan)
	return
}
