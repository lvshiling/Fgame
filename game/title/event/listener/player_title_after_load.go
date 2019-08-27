package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/emperor/emperor"
	emperortemplate "fgame/fgame/game/emperor/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
	titletypes "fgame/fgame/game/title/types"
)

//TODO 称号是否分到各模块
//玩家登录成功后下发称号
func playerTitleAfterLogin(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	playerId := p.GetId()
	manager := p.GetPlayerDataManager(playertypes.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)

	tempTitleList := make([]int32, 0, 3)
	//判断是否功能开启
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMarry) && !manager.HasTitle() {
		playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(p.GetRole(), p.GetSex())
		tempTitleList = append(tempTitleList, playerCreateTemplate.TitleId)
	}

	//获取我是大皇帝(抢龙椅)
	emperorId, _ := emperor.GetEmperorService().GetEmperorIdAndRobNum()
	emperorTitleId := emperortemplate.GetEmperorTemplateService().GetEmperorTitleId()
	titleWear := manager.GetTitleId()
	if emperorId == playerId {
		tempTitleList = append(tempTitleList, emperorTitleId)
	} else {
		if titleWear == emperorTitleId {
			manager.TitleNoWear()
		}
	}

	//天劫塔称号
	// realmTitleId, _ := title.GetTitleService().GetTitleId(titletypes.TitleTypeRank, titletypes.TitleRankSubTypeRealm)
	// realmFirstId := realm.GetRealmRankService().GetTianJieTaFirstId()
	// if realmFirstId == playerId {
	// 	tempTitleList = append(tempTitleList, realmTitleId)
	// } else {
	// 	if titleWear == realmTitleId {
	// 		manager.TitleNoWear()
	// 	}
	// }

	//排行榜称号
	titleRankMap := title.GetTitleService().GetTitleTypeMap(titletypes.TitleTypeRank)
	for titleSubType, titleId := range titleRankMap {
		titleRankSubType := titleSubType.(titletypes.TitleRankSubType)
		// if titleRankSubType == titletypes.TitleRankSubTypeRealm {
		// 	continue
		// }
		rankType := titleRankSubType.TitleRankSubTypeToRankType()
		rankFirstId := rank.GetRankService().GetRankFirstId(ranktypes.RankClassTypeLocal, 0, rankType)
		if rankFirstId == playerId {
			tempTitleList = append(tempTitleList, titleId)
		} else {
			if titleWear == titleId {
				manager.TitleNoWear()
			}
		}
	}

	//添加临时称号
	if len(tempTitleList) != 0 {
		manager.AddTempTitleIdList(tempTitleList)
	}

	titleWear = manager.GetTitleWear().TitleWear
	titleIdMap := manager.GetTitleIdMap()
	scSkillGet := pbutil.BuildSCTitleGet(p, titleWear, titleIdMap)
	p.SendMsg(scSkillGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerTitleAfterLogin))
}
