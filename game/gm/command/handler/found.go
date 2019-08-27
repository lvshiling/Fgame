package handler

import (
	"fgame/fgame/common/lang"
	playerfound "fgame/fgame/game/found/player"
	foundtemplate "fgame/fgame/game/found/template"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/timeutils"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeFoundResBack, command.CommandHandlerFunc(handleFoundRes))
}

func handleFoundRes(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理资源找回设置")
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	foundResArr := c.Args[0:]
	if len(foundResArr)%2 != 0 {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"foundResArr": foundResArr,
			}).Warn("gm:处理资源找回,资源参数错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	var resIdArr []int32
	var countArr []int32
	for index, number := range foundResArr {
		num, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id":          pl.GetId(),
					"foundResArr": foundResArr,
					"error":       err,
				}).Warn("gm:处理资源找回，资源设置不是数字")
			playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
			return err
		}

		if index%2 == 0 || index%2 == 2 {
			tem := foundtemplate.GetFoundTemplateService().GetFoundTemplate(int32(num))
			if tem == nil {
				log.WithFields(
					log.Fields{
						"id":          pl.GetId(),
						"foundResArr": foundResArr,
					}).Warn("gm:处理资源找回，资源不存在")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
				return err
			}
			resIdArr = append(resIdArr, int32(num))
		} else {
			countArr = append(countArr, int32(num))
		}
	}

	resInfoMap := make(map[int32]int32)
	for index, resId := range resIdArr {
		if _, ok := resInfoMap[resId]; ok {
			resInfoMap[resId] += countArr[index]
		} else {
			resInfoMap[resId] = countArr[index]
		}
	}

	err = addFoundRecord(pl, resInfoMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"foundResArr": foundResArr,
				"error":       err,
			}).Warn("gm:处理资源找回,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":          pl.GetId(),
			"foundResArr": foundResArr,
		}).Debug("gm:处理资源找回,完成")
	return
}

func addFoundRecord(player player.Player, resMap map[int32]int32) (err error) {
	manager := player.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)

	for typ, count := range resMap {
		for count > 0 {
			count -= 1
			resType := foundtypes.FoundResourceType(typ)
			manager.IncreFoundResJoinTimes(resType)
		}
	}

	now := global.GetGame().GetTimeService().Now()
	preDay, err := timeutils.PreDayOfTime(now)
	for _, obj := range manager.GetCurDayResRecordList() {
		obj.SetUpdateTime(preDay)
		obj.SetModified()
	}

	return

}
