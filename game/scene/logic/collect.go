package logic

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/template"
	christmaspbutil "fgame/fgame/game/christmas/pbutil"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
)

type refreshCollectResult struct {
	refreshNum int32
	biologyId  int32
}

//场景刷新采集物
func RefreshCollectOnScene(s scene.Scene, refreshNum, biologyId int32) {
	ctx := scene.WithScene(context.Background(), s)
	result := &refreshCollectResult{
		refreshNum: refreshNum,
		biologyId:  biologyId,
	}
	//异步进入场景
	s.Post(message.NewScheduleMessage(onRefreshCollect, ctx, result, nil))
}

//回调
func onRefreshCollect(ctx context.Context, result interface{}, err error) error {
	s := scene.SceneInContext(ctx)
	tResult := result.(*refreshCollectResult)
	biologyId := int(tResult.biologyId)
	refreshNum := tResult.refreshNum

	templateObj := template.GetTemplateService().Get(biologyId, (*gametemplate.BiologyTemplate)(nil))
	if templateObj == nil {
		return nil
	}
	biologyTemp := templateObj.(*gametemplate.BiologyTemplate)

	// 移除
	remainNum := int32(0)
	oldNPCS := s.GetNPCS(biologyTemp.GetBiologyScriptType())
	for _, npc := range oldNPCS {
		if npc.GetBiologyTemplate().Id == int(biologyId) {
			remainNum += 1
		}
	}

	refreshNum -= remainNum
	//刷新
	for i := int32(0); i < refreshNum; i++ {
		pos := s.MapTemplate().GetMap().RandomPosition()
		n := collectnpc.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, biologyTemp, pos, 0, 0)

		s.AddSceneObject(n)
	}

	scMsg := christmaspbutil.BuildSCChristmasCollectNumNotice(tResult.refreshNum)
	s.BroadcastMsg(scMsg)
	return nil
}

//场景清除采集物
func ClearCollectOnScene(s scene.Scene, biologyId int32) {
	ctx := scene.WithScene(context.Background(), s)
	//异步进入场景
	s.Post(message.NewScheduleMessage(onClearCollect, ctx, biologyId, nil))
}

//回调
func onClearCollect(ctx context.Context, result interface{}, err error) error {
	s := scene.SceneInContext(ctx)
	biologyId := result.(int32)
	templateObj := template.GetTemplateService().Get(int(biologyId), (*gametemplate.BiologyTemplate)(nil))
	if templateObj == nil {
		return nil
	}
	biologyTemp := templateObj.(*gametemplate.BiologyTemplate)
	// 移除
	oldNPCS := s.GetNPCS(biologyTemp.GetBiologyScriptType())
	for _, npc := range oldNPCS {
		if npc.GetBiologyTemplate().Id == int(biologyId) {
			s.RemoveSceneObject(npc, false)
		}
	}

	return nil
}
