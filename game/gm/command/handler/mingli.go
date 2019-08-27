package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	minggelogic "fgame/fgame/game/mingge/logic"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetemplate "fgame/fgame/game/mingge/template"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommnadTypeMingLi, command.CommandHandlerFunc(handleMingLi))

}

func handleMingLi(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 3 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	mingGongStr := c.Args[0]
	mingLiPosStr := c.Args[1]
	slotStr := c.Args[2]
	propertyStr := c.Args[3]

	mingGongInt64, err := strconv.ParseInt(mingGongStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,mingGongStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	mingLiPosInt64, err := strconv.ParseInt(mingLiPosStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,mingLiPosStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	slotInt64, err := strconv.ParseInt(slotStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,slotStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	propertyInt64, err := strconv.ParseInt(propertyStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,propertyStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	mingGongType := minggetypes.MingGongType(mingGongInt64)
	if !mingGongType.Valid() {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	mingGongPosType := minggetypes.MingGongAllSubType(mingLiPosInt64)
	if !mingGongPosType.Valid() {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	mingLiSlotType := minggetypes.MingLiSlotType(slotInt64)
	if !mingLiSlotType.Vaild() {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	propertyType := minggetypes.MingGePropertyType(propertyInt64)
	if !propertyType.Valid() {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
				"error":        err,
			}).Warn("gm:处理设置命理属性,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	mingLiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingLiTemplate(mingGongType, mingGongPosType)
	if mingLiTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
			}).Warn("gm:处理设置命理属性,number模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	obj := manager.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongPosType)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
			}).Warn("gm:处理设置命理属性,命宫未激活")
		playerlogic.SendSystemMessage(pl, lang.MingGeMingLiBaptize)
		return
	}

	mingLiMap := obj.GetMingLiMap()
	_, ok := mingLiMap[mingLiSlotType]
	if !ok {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"mingGongStr":  mingGongStr,
				"mingLiPosStr": mingLiPosStr,
				"slotStr":      slotStr,
				"propertyStr":  propertyStr,
			}).Warn("gm:处理设置命理属性,洗炼一次才能设置属性")
		playerlogic.SendSystemMessage(pl, lang.MingGeMingLiSetProperty)
		return
	}

	mingGongTypeMap := manager.GmSetMingLi(mingGongType, mingGongPosType, mingLiSlotType, propertyType)
	if len(mingGongTypeMap) != 0 {
		mingLiMap := manager.GetMingLiMap()
		scMingGeMingGongActivate := pbutil.BuildSCMingGeMingGongActivate(mingLiMap, mingGongTypeMap)
		pl.SendMsg(scMingGeMingGongActivate)
	}
	slotList := make([]int32, 0, 1)
	slotList = append(slotList, int32(mingLiSlotType))
	//同步属性
	minggelogic.MingGePropertyChanged(pl)
	scMingGeMingLiBaptize := pbutil.BuildSCMingGeMingLiBaptize(int32(mingGongType), int32(mingGongPosType), obj, slotList)
	pl.SendMsg(scMingGeMingLiBaptize)
	return
}
