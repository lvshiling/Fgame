package scene

import (
	"fgame/fgame/game/global"
	lianyutemplate "fgame/fgame/game/lianyu/template"
	lianyutypes "fgame/fgame/game/lianyu/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"time"
)

const (
	lianYuTaskTime = time.Second
)

type LianYuSceneTask struct {
	sceneData LianYuSceneData
	//无间炼狱活动开始时间
	starTime int64
	//boss刷新标识
	bossBorn bool
}

func (t *LianYuSceneTask) Run() {
	//刷boss
	if t.bossBorn {
		return
	}
	constantTemplate := lianyutemplate.GetLianYuTemplateService().GetConstantTemplate(t.sceneData.GetAcitvityType())
	if constantTemplate == nil {
		return
	}
	biologyTemplate := constantTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}

	bornType, ok := lianyutypes.ActivityTypeToBornType(t.sceneData.GetAcitvityType())
	if !ok {
		return
	}
	lianYuPosTemplate := lianyutemplate.GetLianYuTemplateService().GetBornPosTemplate(bornType)
	if lianYuPosTemplate == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	escapeTime := now - t.starTime
	if escapeTime < int64(constantTemplate.BossTime) {
		return
	}
	sceneData, ok := t.sceneData.(*lianYuSceneData)
	if !ok {
		return
	}
	if sceneData.boss == nil {
		return
	}
	if sceneData.boss.bossStatus != lianyutypes.LianYuBossStatusTypeInit {
		return
	}
	//刷大Boss
	bornPos := lianYuPosTemplate.GetPos()
	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), biologyTemplate, bornPos, 0, 0)
	if n != nil {
		//设置场景
		sceneData.s.AddSceneObject(n)
		t.bossBorn = true

	}
	return

}

func (t *LianYuSceneTask) ElapseTime() time.Duration {
	return lianYuTaskTime
}

func CreateLianYuTask(sceneData LianYuSceneData, starTime int64) *LianYuSceneTask {
	tuLongtask := &LianYuSceneTask{
		sceneData: sceneData,
		starTime:  starTime,
		bossBorn:  false,
	}
	return tuLongtask
}
