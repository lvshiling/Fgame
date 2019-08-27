package scene

import (
	coretemplate "fgame/fgame/core/template"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	fourgodtemplate "fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/global"
	scenescene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"time"
)

const (
	fourGodWarTaskTime = time.Second
)

type FourGodWarSceneTask struct {
	sceneData FourGodWarSceneData
	//四神遗迹活动开始时间
	starTime int64
	//常量模板
	fourGodTemplate *template.FourGodTemplate
	//boss刷新标识
	bossBorn bool
	//第一只特殊怪刷新标识
	specialFirstBorn bool
	//能否刷特殊怪
	specialRefresh bool
	//特殊怪最后刷新
	specialLastTime int64
}

func (fg *FourGodWarSceneTask) Run() {
	now := global.GetGame().GetTimeService().Now()

	//宝箱采集
	collectBoxTime := int64(fg.fourGodTemplate.BoxTime)
	collectBoxMap := fg.sceneData.GetCollectBox()
	for _, collectBox := range collectBoxMap {
		if collectBox.playerId == 0 {
			continue
		}
		collectStarTime := collectBox.collectStarTime
		if now-collectStarTime >= collectBoxTime {
			gameevent.Emit(fourgodeventtypes.EventTypeFourGodCollectBoxFinish, collectBox.playerId, collectBox.boxId)
		}
	}

	//刷boss
	if !fg.bossBorn {
		bossTime := int64(fg.fourGodTemplate.BossTime)
		if now-fg.starTime >= bossTime {
			to := coretemplate.GetTemplateService().Get(int(fg.fourGodTemplate.BossId), (*template.BiologyTemplate)(nil))
			bornPos, _ := fourgodtemplate.GetFourGodTemplateService().GetFourGodBossPos()
			bossTemplate := to.(*template.BiologyTemplate)
			n := scenescene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), bossTemplate, bornPos, 0, 0)
			if n != nil {
				//设置场景
				fg.sceneData.GetScene().AddSceneObject(n)
				fg.bossBorn = true
			}
		}
	}

	//刷特殊怪
	if !fg.specialFirstBorn {
		specialTime := int64(fg.fourGodTemplate.SpecialTime)
		if now-fg.starTime >= specialTime {
			fg.specialFirstBorn = true
			fg.specialRefresh = true
		}
	} else {
		curSpecialNum := fg.sceneData.GetSpecialNum()
		specialProTime := int64(fg.fourGodTemplate.SpecialProbabilityTime)
		//数量和时间都满足
		if now-fg.specialLastTime >= specialProTime &&
			curSpecialNum < fg.fourGodTemplate.SpecialNum {
			if hit := mathutils.RandomHit(common.MAX_RATE, int(fg.fourGodTemplate.SpecialProbability)); hit {
				fg.specialRefresh = true
			}
		}
	}
	if fg.specialRefresh {
		to := coretemplate.GetTemplateService().Get(int(fg.fourGodTemplate.SpecialId), (*template.BiologyTemplate)(nil))
		specialTemplate := to.(*template.BiologyTemplate)
		n := scenescene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, int32(0), specialTemplate, fg.sceneData.GetSpecialPos(), 0, 0)
		if n != nil {
			//设置场景
			fg.sceneData.GetScene().AddSceneObject(n)
			fg.sceneData.AddSpecial()
			fg.specialLastTime = now
			fg.specialRefresh = false
		}
	}

}

func (ft *FourGodWarSceneTask) ElapseTime() time.Duration {
	return fourGodWarTaskTime
}

func CreateFourGodWarTask(sceneData FourGodWarSceneData, starTime int64, fourGodTemplate *template.FourGodTemplate) *FourGodWarSceneTask {
	fourGodtask := &FourGodWarSceneTask{
		sceneData:        sceneData,
		starTime:         starTime,
		bossBorn:         false,
		specialFirstBorn: false,
		specialRefresh:   false,
		fourGodTemplate:  fourGodTemplate,
	}
	return fourGodtask
}
