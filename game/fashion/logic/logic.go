package logic

import (
	"context"
	"fgame/fgame/common/message"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
)

//时装属性加成
func FashionPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeFashion.Mask())
	return
}

//时装升星判断
func FashionUpstar(curTimesNum int32, curBless int32, fashionStarTemplate *gametemplate.FashionUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := fashionStarTemplate.TimesMin
	timesMax := fashionStarTemplate.TimesMax
	updateRate := fashionStarTemplate.UpstarRate
	blessMax := fashionStarTemplate.ZhufuMax
	addMin := fashionStarTemplate.AddMin
	addMax := fashionStarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

func ActivateCampFashion(pl player.Player, campType chuangshitypes.ChuangShiCampType) {
	ctx := scene.WithPlayer(context.Background(), pl)
	campJoinMsg := message.NewScheduleMessage(onCampFashionActivate, ctx, campType, nil)
	pl.Post(campJoinMsg)
}

func onCampFashionActivate(ctx context.Context, result interface{}, err error) error {
	// tpl := scene.PlayerInContext(ctx)
	// pl := tpl.(player.Player)
	// campType := result.(chuangshitypes.ChuangShiCampType)

	// campTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCampTemp(campType)
	// if campTemp == nil {
	// 	return nil
	// }

	// fashionManager := pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	// if !fashionManager.CampFashionActivate(campTemp.FashionId) {
	// 	return nil
	// }

	return nil
}
