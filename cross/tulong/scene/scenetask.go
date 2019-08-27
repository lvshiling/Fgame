package scene

import (
	crosstulongtypes "fgame/fgame/cross/tulong/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/template"
	tulongtemplate "fgame/fgame/game/tulong/template"
	tulongtypes "fgame/fgame/game/tulong/types"
	"time"
)

const (
	tuLongTaskTime = time.Second
)

type TuLongSceneTask struct {
	sceneData TuLongSceneData
	//屠龙活动开始时间
	starTime int64
	//boss模板
	bossTemplate *template.TuLongTemplate
	//boss刷新标识
	bossBorn bool
}

func (t *TuLongSceneTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	tuLongConstTempalte := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
	if tuLongConstTempalte == nil {
		return
	}
	//龙蛋采集
	t.collectEggFinish(now, int64(tuLongConstTempalte.CaiJiTime))
	//刷新大boss
	t.refreshBoss(now, int64(tuLongConstTempalte.BossTime))
	//重新刷新小龙蛋
	t.smallEggReborn()
	return
}

//重新刷新小龙蛋
func (t *TuLongSceneTask) smallEggReborn() {
	t.sceneData.SmallEggRebornCheck()
}

//龙蛋采集
func (t *TuLongSceneTask) collectEggFinish(now int64, caiJiTime int64) {
	//龙蛋采集
	collectEggMap := t.sceneData.GetCollectEgg()
	for curNpcId, collectEgg := range collectEggMap {
		if collectEgg.allianceId == 0 {
			continue
		}
		collectEggTime := caiJiTime
		collectStarTime := collectEgg.collectStarTime
		if now-collectStarTime >= collectEggTime {
			t.sceneData.FinishCollectEgg(curNpcId)
		}
	}
}

//刷新大boss
func (t *TuLongSceneTask) refreshBoss(now int64, bossTime int64) {
	//刷boss
	if t.bossBorn {
		return
	}

	tuLongTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongBigBossTemplate()
	if tuLongTemplate == nil {
		return
	}

	escapeTime := now - t.starTime
	if escapeTime < bossTime {
		return
	}
	sceneData, ok := t.sceneData.(*tuLongSceneData)
	if !ok {
		return
	}
	bigEgg := sceneData.bigEgg
	if bigEgg.status != crosstulongtypes.EggStatusTypeInit {
		return
	}
	bigEgg.status = crosstulongtypes.EggStatusTypeBoss
	//移除大龙蛋
	sceneData.s.RemoveSceneObject(bigEgg.npc, false)
	bigEgg.npc = nil

	//刷大Boss
	biologyTemplate := tuLongTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}

	tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypeBoss, bigEgg.bornBiaoShi)
	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), biologyTemplate, tuLongPosTemplate.GetPos(), 0, 0)
	if n != nil {
		//设置场景
		sceneData.s.AddSceneObject(n)
		bigEgg.npc = n
		t.bossBorn = true
	}
	return
}

func (t *TuLongSceneTask) ElapseTime() time.Duration {
	return tuLongTaskTime
}

func CreateTuLongTask(sceneData TuLongSceneData, starTime int64) *TuLongSceneTask {
	tuLongtask := &TuLongSceneTask{
		sceneData: sceneData,
		starTime:  starTime,
		bossBorn:  false,
	}
	return tuLongtask
}
