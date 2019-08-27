package logic

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

// 竞猜结算
func GuessResult(pl player.Player, attendObj *arenapvp.ArenapvpGuessRecordObject) {
	winnerId := attendObj.GetWinnerId()
	guessType := attendObj.GetGuessType()
	guessId := attendObj.GetGuessId()
	raceNum := attendObj.GetRaceNumber()

	if winnerId == 0 {
		return
	}

	pvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(guessType)
	rewItemMap := make(map[int32]int32)
	mailContent := ""
	if winnerId == guessId {
		rewItemMap = pvpTemp.GetWinMailItemMap()
		mailContent = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpGuessSuccessMailContent), coreutils.FormatNoticeStr(guessType.String()))
	} else {
		rewItemMap = pvpTemp.GetLoseMailItemMap()
		mailContent = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpGuessFailedMailContent), coreutils.FormatNoticeStr(guessType.String()))
	}

	title := lang.GetLangService().ReadLang(lang.ArenapvpGuessMailTitle)
	emaillogic.AddEmail(pl, title, mailContent, rewItemMap)

	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	logObj := arenapvpManager.GuessResult(raceNum, guessType, winnerId)

	scArenapvpGuessInfoPush := pbutil.BuildSCArenapvpGuessInfoPush(logObj)
	pl.SendMsg(scArenapvpGuessInfoPush)
}

//退还
func GuessReturn(pl player.Player, attendObj *arenapvp.ArenapvpGuessRecordObject) {
	winnerId := attendObj.GetWinnerId()
	guessType := attendObj.GetGuessType()
	raceNum := attendObj.GetRaceNumber()

	pvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(guessType)
	rewItemMap := make(map[int32]int32)
	rewItemMap[constanttypes.BindGoldItem] = pvpTemp.JingchaiUseBindgold
	title := lang.GetLangService().ReadLang(lang.ArenapvpGuessReturnMailTitle)
	content := lang.GetLangService().ReadLang(lang.ArenapvpGuessReturnMailContent)
	emaillogic.AddEmail(pl, title, content, rewItemMap)

	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.GuessResult(raceNum, guessType, winnerId)
}
