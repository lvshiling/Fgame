package scene

import (
	"fgame/fgame/game/global"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"time"
)

const (
	godSiegeTaskTime = time.Second
)

type GodSiegeSceneTask struct {
	sceneData GodSiegeSceneData
	//神兽攻城活动开始时间
	starTime int64
	//boss刷新标识
	bossBorn bool
	//地图id
	mapId int32
}

func (t *GodSiegeSceneTask) Run() {
	//刷boss
	if t.bossBorn {
		return
	}
	constantTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetConstantTemplate()
	if constantTemplate == nil {
		return
	}
	biologyTemplate := constantTemplate.GetBiologyTemplate(t.mapId)
	if biologyTemplate == nil {
		return
	}
	godSiegePosTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetBossPosTemplate(t.mapId)
	if godSiegePosTemplate == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	escapeTime := now - t.starTime
	if escapeTime < int64(constantTemplate.BossTime) {
		return
	}
	sceneData, ok := t.sceneData.(*godSiegeSceneData)
	if !ok {
		return
	}
	if sceneData.boss == nil {
		return
	}
	if sceneData.boss.bossStatus != godsiegetypes.GodSiegeBossStatusTypeInit {
		return
	}
	//刷大Boss
	bornPos := godSiegePosTemplate.GetPos()
	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), biologyTemplate, bornPos, 0, 0)
	if n != nil {
		//设置场景
		sceneData.s.AddSceneObject(n)
		t.bossBorn = true
	}
	return

}

func (t *GodSiegeSceneTask) ElapseTime() time.Duration {
	return godSiegeTaskTime
}

func CreateGodSiegeTask(sceneData GodSiegeSceneData, starTime int64, mapId int32) *GodSiegeSceneTask {
	tuLongtask := &GodSiegeSceneTask{
		sceneData: sceneData,
		starTime:  starTime,
		bossBorn:  false,
		mapId:     mapId,
	}
	return tuLongtask
}
