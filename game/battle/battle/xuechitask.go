package battle

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"math"
	"time"
)

const (
	xueChiTaskTime = time.Second
)

type XueChiTask struct {
	pl scene.Player
}

func (t *XueChiTask) Run() {
	s := t.pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().IsXueChi == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	intervalTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeRecoverIntervalTime)
	if t.pl.IsDead() {
		return
	}
	curHp := t.pl.GetHP()
	if t.pl.GetBlood() <= 0 {
		return
	}
	diffTime := now - t.pl.GetLastBloodTime()
	if diffTime < int64(intervalTime) {
		return
	}

	maxHp := t.pl.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)

	hpPercent := int32(math.Floor(float64(curHp) / float64(maxHp) * 100))
	if hpPercent >= t.pl.GetBloodLine() {
		return
	}
	//补血
	recoverBase := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSupplementrHpFixed))
	percent := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSupplementrHpPercent)
	percenBase := int64(math.Ceil(float64(percent) / float64(100) * float64(maxHp)))
	recover := recoverBase + percenBase

	//debuff影响
	effectNum := t.pl.GetEffectNum(scenetypes.BuffEffectTypeXueChiDeBuff)
	effectCover := int64(math.Ceil(float64(recover) * float64(effectNum) / float64(common.MAX_RATE)))
	recover -= effectCover

	loseHp := maxHp - curHp
	if recover > loseHp {
		recover = loseHp
	}
	t.pl.RecoverHp(recover)
}

func (t *XueChiTask) ElapseTime() time.Duration {
	return xueChiTaskTime
}

func CreateXueChiTask(p scene.Player) *XueChiTask {
	xueChiTask := &XueChiTask{
		pl: p,
	}
	return xueChiTask
}
