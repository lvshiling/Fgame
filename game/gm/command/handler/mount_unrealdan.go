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
	command.Register(gmcommandtypes.CommandTypeMountDan, command.CommandHandlerFunc(handleMountDan))

}

func handleMountDan(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	unrealDanStr := c.Args[0]
	unrealDanLevel, err := strconv.ParseInt(unrealDanStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置坐骑食幻化丹等级,unrealDanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := mount.GetMountService().GetMountHuanHuaTemplate(int32(unrealDanLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置坐骑食幻化丹等级,unrealDanLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	manager.GmSetMountUnrealDanLevel(int32(unrealDanLevel))

	//同步属性
	mountlogic.MountPropertyChanged(pl)

	scMountUnrealDan := pbutil.BuildSCMountUnrealDan(int32(unrealDanLevel), 0)
	pl.SendMsg(scMountUnrealDan)
	return
}
