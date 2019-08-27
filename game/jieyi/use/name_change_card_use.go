package use

import (
	"fgame/fgame/common/lang"
	coredirty "fgame/fgame/core/dirty"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeJieYiItem, itemtypes.ItemJieYiSubTypeChangeName, playerinventory.ItemUseHandleFunc(handlerRename))
}

const (
	minJieYiNameLen = 2
	maxJieYiNameLen = 6
)

// 结义改名卡
func handlerRename(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	needNum := int32(1)
	if num < needNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"needNum":  needNum,
			}).Warn("jieyi:处理结义改名,物品不足")
		return
	}

	memberObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if memberObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi:处理结义改名,玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	newName := strings.TrimSpace(args)
	lenOfName := len([]rune(newName))

	if lenOfName < minJieYiNameLen && lenOfName > maxJieYiNameLen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("jieyi:处理结义改名,名字不合法")
		playerlogic.SendSystemMessage(pl, lang.JieYiNameIllegal)
		return
	}

	if !coredirty.GetDirtyService().IsLegal(newName) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("jieyi:处理结义改名,名字含有脏字")
		playerlogic.SendSystemMessage(pl, lang.JieYiNameDirty)
		return
	}

	jieyi.GetJieYiService().IsJieYiLaoDa(pl.GetId())
	if !jieyi.GetJieYiService().IsJieYiLaoDa(pl.GetId()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("jieyi:处理结义改名,玩家不是结义老大")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotIsLaoDa)
		return
	}

	if !jieyi.GetJieYiService().IsNameRepetitive(newName) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("jieyi:处理结义改名,名字一样")
		playerlogic.SendSystemMessage(pl, lang.JieYiNameRepetitive)
		return
	}

	flag = jieyi.GetJieYiService().SetJieYiName(memberObj.GetJieYiId(), newName)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("jieyi:处理结义改名,使用失败")
		playerlogic.SendSystemMessage(pl, lang.JieYiUseGaiMingKaFail)
		return
	}
	// 推送给结义兄弟
	jieYi := memberObj.GetJieYi()
	jieyilogic.JieYiMemberChanged(jieYi)

	// 改名成功
	playerlogic.SendSystemMessage(pl, lang.JieYiChangeNameSuccess)

	flag = true
	return
}
