package use

import (
	"fgame/fgame/common/lang"
	coredirty "fgame/fgame/core/dirty"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	alliancetemplate "fgame/fgame/game/alliance/template"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAlliance, itemtypes.ItemAllianceSubTypeHeMengGaiMingKa, playerinventory.ItemUseHandleFunc(handlerRename))
}

const (
	minAllianceNameLen = 1
	maxAllianceNameLen = 6
)

// 改名卡
func handlerRename(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if pl.GetAllianceId() == 0 {
		return
	}

	needNum := alliancetemplate.GetAllianceTemplateService().GetAllianceConstantTemp().GaimingItemCount
	if num < needNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"needNum":  needNum,
			}).Warn("alliance:处理仙盟改名,物品不足")
		return
	}

	newName := strings.TrimSpace(args)
	lenOfName := len([]rune(newName))

	if lenOfName < minAllianceNameLen && lenOfName > maxAllianceNameLen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("alliance:处理仙盟改名,名字不合法")
		playerlogic.SendSystemMessage(pl, lang.AllianceNameIllegal)
		return
	}

	tflag := coredirty.GetDirtyService().IsLegal(newName)
	if !tflag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("alliance:处理仙盟改名,名字含有脏字")
		playerlogic.SendSystemMessage(pl, lang.AllianceNameDirty)
		return
	}
	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	if al == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("alliance:处理仙盟改名,用户不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInAlliance)
		return
	}
	beforeName := al.GetAllianceName()
	if beforeName == newName {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("alliance:处理仙盟改名,名字一样")
		playerlogic.SendSystemMessage(pl, lang.AllianceRenameSame)
		return
	}

	err = alliance.GetAllianceService().UpdateAllianceName(pl, newName)
	if err != nil {
		return
	}

	// 系统频道
	playerName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(pl.GetName()))
	oldAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(beforeName))
	newAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(newName))
	systemContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceRenameNotice), playerName, oldAlName, newAlName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(systemContent))

	flag = true
	return
}
