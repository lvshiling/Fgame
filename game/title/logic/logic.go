package logic

import (
	"context"
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
)

//变更称号属性
func TitlePropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeTitle.Mask())
	return
}

func OnTempTitleChangedGet(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	titleId := result.(int32)
	tempTitelGet(pl, titleId)
	return nil
}

func OnTempTitleChangedGetRemove(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)

	pl := tpl.(player.Player)
	titleId := result.(int32)
	tempTitleRemove(pl, titleId)
	return nil
}

//临时称号获得
func tempTitelGet(pl player.Player, titleId int32) (err error) {
	if pl == nil {
		return
	}
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	manager.TempTitleAdd(titleId)

	//同步属性
	attrTemplate := titleTemplate.GetBattleAttrTemplate()
	if attrTemplate != nil {
		TitlePropertyChanged(pl)
	}

	scTitleAdd := pbutil.BuildSCTitleAdd(titleId)
	pl.SendMsg(scTitleAdd)
	return
}

//临时称号移除
func tempTitleRemove(pl player.Player, titleId int32) (err error) {
	if pl == nil {
		return
	}
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	manager.TempTitleRemove(titleId)

	//同步属性
	attrTemplate := titleTemplate.GetBattleAttrTemplate()
	if attrTemplate != nil {
		TitlePropertyChanged(pl)
	}
	scTitleAdd := pbutil.BuildSCTitleRemove(titleId)
	pl.SendMsg(scTitleAdd)
	return
}

//称号升星判断
func TitleUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.TitleUpStarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
