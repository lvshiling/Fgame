package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/common/common"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"math/rand"
	"strconv"
)

type ArenaConstantTemplate struct {
	*ArenaConstantTemplateVO
	arenaMapTemplate    *MapTemplate
	qingLongMapTemplate *MapTemplate
	baiHuMapTemplate    *MapTemplate
	zhuQueMapTemplate   *MapTemplate
	xuanWuMapTemplate   *MapTemplate
	team1BornPos        coretypes.Position
	team2BornPos        coretypes.Position
	equipBoxRateMap     map[arenatypes.FourGodType]int32
	equipBoxMap         map[arenatypes.FourGodType]int32
	beginTime           int64
	endTime             int64
}

func (t *ArenaConstantTemplate) GetFourGodMapTemplate(fourGodType arenatypes.FourGodType) *MapTemplate {
	switch fourGodType {
	case arenatypes.FourGodTypeBaiHu:
		return t.baiHuMapTemplate
	case arenatypes.FourGodTypeQingLong:
		return t.qingLongMapTemplate
	case arenatypes.FourGodTypeXuanWu:
		return t.xuanWuMapTemplate
	case arenatypes.FourGodTypeZhuQue:
		return t.zhuQueMapTemplate
	}
	return nil
}

func (t *ArenaConstantTemplate) IsOnArenaTime(now int64) bool {
	if now < t.GetBeginTime(now) {
		return false
	}

	if now > t.GetEndTime(now) {
		return false
	}

	return true
}

func (t *ArenaConstantTemplate) GetEndTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return beginDay + t.endTime
}

func (t *ArenaConstantTemplate) GetBeginTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return beginDay + t.beginTime
}

func (t *ArenaConstantTemplate) GetArenaMapTemplate() *MapTemplate {
	return t.arenaMapTemplate
}

func (t *ArenaConstantTemplate) GetTeam1BornPos() coretypes.Position {
	return t.team1BornPos
}

func (t *ArenaConstantTemplate) GetTeam2BornPos() coretypes.Position {
	return t.team2BornPos
}

func (t *ArenaConstantTemplate) GetEquipBoxRate(fourGodType arenatypes.FourGodType) int32 {
	rate := t.equipBoxRateMap[fourGodType]
	return rate
}

func (t *ArenaConstantTemplate) GetEquipBox(fourGodType arenatypes.FourGodType) int32 {
	boxId := t.equipBoxMap[fourGodType]
	return boxId
}

func (t *ArenaConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	tempMapTeamplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTeamplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.arenaMapTemplate = tempMapTeamplate.(*MapTemplate)

	t.team1BornPos = coretypes.Position{
		X: t.PosX1,
		Y: t.PosY1,
		Z: t.PosZ1,
	}

	t.team2BornPos = coretypes.Position{
		X: t.PosX2,
		Y: t.PosY2,
		Z: t.PosZ2,
	}
	tempMap1Teamplate := template.GetTemplateService().Get(int(t.MapId1), (*MapTemplate)(nil))
	if tempMap1Teamplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId1)
		return template.NewTemplateFieldError("MapId1", err)
	}
	t.qingLongMapTemplate = tempMap1Teamplate.(*MapTemplate)

	tempMap2Teamplate := template.GetTemplateService().Get(int(t.MapId2), (*MapTemplate)(nil))
	if tempMap2Teamplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId2)
		return template.NewTemplateFieldError("MapId2", err)
	}
	t.baiHuMapTemplate = tempMap2Teamplate.(*MapTemplate)

	tempMap3Teamplate := template.GetTemplateService().Get(int(t.MapId3), (*MapTemplate)(nil))
	if tempMap3Teamplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId3)
		return template.NewTemplateFieldError("MapId3", err)
	}
	t.zhuQueMapTemplate = tempMap3Teamplate.(*MapTemplate)

	tempMap4Teamplate := template.GetTemplateService().Get(int(t.MapId4), (*MapTemplate)(nil))
	if tempMap4Teamplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId4)
		return template.NewTemplateFieldError("MapId4", err)
	}
	t.xuanWuMapTemplate = tempMap4Teamplate.(*MapTemplate)

	t.equipBoxRateMap = make(map[arenatypes.FourGodType]int32)
	rateArr, err := coreutils.SplitAsIntArray(t.EquipBoxRate)
	if err != nil {
		return template.NewTemplateFieldError("EquipBoxRate", err)
	}

	for i, rate := range rateArr {
		t.equipBoxRateMap[arenatypes.FourGodType(i)] = rate
	}
	t.equipBoxMap = make(map[arenatypes.FourGodType]int32)
	boxIdArr, err := coreutils.SplitAsIntArray(t.EquipBoxId)
	if err != nil {
		return template.NewTemplateFieldError("EquipBoxId", err)
	}
	for i, boxId := range boxIdArr {
		t.equipBoxMap[arenatypes.FourGodType(i)] = boxId
	}

	return
}
func (t *ArenaConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if err = validator.MinValidate(float64(t.RebornAmountMax), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RebornAmountMax", err)
		return
	}

	if !t.arenaMapTemplate.GetMap().IsMask(t.team1BornPos.X, t.team1BornPos.Z) {
		err = fmt.Errorf("team1 born pos invalid")
		return template.NewTemplateFieldError("Team1BornPos", err)
	}
	t.team1BornPos.Y = t.arenaMapTemplate.GetMap().GetHeight(t.team1BornPos.X, t.team1BornPos.Z)
	if !t.arenaMapTemplate.GetMap().IsMask(t.team2BornPos.X, t.team2BornPos.Z) {
		err = fmt.Errorf("team2 born pos invalid")
		return template.NewTemplateFieldError("Team2BornPos", err)
	}
	t.team2BornPos.Y = t.arenaMapTemplate.GetMap().GetHeight(t.team2BornPos.X, t.team2BornPos.Z)

	//判断是否是采集宝箱
	for _, boxId := range t.equipBoxMap {
		tempBiologyTemplate := template.GetTemplateService().Get(int(boxId), (*BiologyTemplate)(nil))
		if tempBiologyTemplate == nil {
			err = fmt.Errorf("EquipBoxId[%s] invalid", t.EquipBoxId)
			return template.NewTemplateFieldError("EquipBoxId", err)
		}
		biologyTemplate := tempBiologyTemplate.(*BiologyTemplate)
		if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeArenaTreasure {
			err = fmt.Errorf("EquipBoxId[%s] 不是宝箱", t.EquipBoxId)
			return template.NewTemplateFieldError("EquipBoxId", err)
		}
	}

	if err = validator.MinValidate(float64(t.TeamCount), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("TeamCount", err)
		return
	}

	//复活最小
	if err = validator.MinValidate(float64(t.ReviveMin), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("ReviveMin", err)
		return
	}

	if err = validator.MinValidate(float64(t.ReviveMax), float64(t.ReviveMin), true); err != nil {
		err = template.NewTemplateFieldError("ReviveMax", err)
		return
	}

	//属性
	if err = validator.MinValidate(float64(t.AttrMin), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("AttrMin", err)
		return
	}

	if err = validator.MinValidate(float64(t.AttrMax), float64(t.AttrMin), true); err != nil {
		err = template.NewTemplateFieldError("AttrMax", err)
		return
	}

	err = validator.MinValidate(float64(t.RiZhiTimeMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RiZhiTimeMin)
		err = template.NewTemplateFieldError("RiZhiTimeMin", err)
		return
	}

	err = validator.MinValidate(float64(t.RiZhiTimeMax), float64(t.RiZhiTimeMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RiZhiTimeMax)
		err = template.NewTemplateFieldError("RiZhiTimeMax", err)
		return
	}

	err = validator.MinValidate(float64(t.RiZhiMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RiZhiMax)
		err = template.NewTemplateFieldError("RiZhiMax", err)
		return
	}

	//匹配时间
	//活动开始时间
	beginTime, err := timeutils.ParseDayOfHHMM(t.BeginTime)
	if err != nil {
		return template.NewTemplateFieldError("BeginTime", fmt.Errorf("[%s] invalid", t.BeginTime))
	}
	t.beginTime = beginTime

	//活动结束时间
	endInt, err := strconv.ParseInt(t.EndTime, 10, 64)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.EndTime)
		return template.NewTemplateFieldError("EndTime", err)
	}
	if endInt == 2400 {
		t.endTime = int64(common.DAY)
	} else {
		endTime, err := timeutils.ParseDayOfHHMM(t.EndTime)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.EndTime)
			return template.NewTemplateFieldError("EndTime", err)
		}
		t.endTime = endTime
	}

	//每日积分
	if err = validator.MinValidate(float64(t.JiFenMaxDay), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("JiFenMaxDay", err)
		return
	}

	return
}

func (t *ArenaConstantTemplate) PatchAfterCheck() {
}

func (t *ArenaConstantTemplate) RandomPropertyPercent() int32 {
	n := t.AttrMax - t.AttrMin
	if n < 0 {
		panic(fmt.Errorf("template:不可能"))
	}
	if n == 0 {
		return t.AttrMin
	}
	randomN := rand.Int31n(n)
	return randomN + t.AttrMin
}

func (t *ArenaConstantTemplate) RandomRevive() int32 {
	n := t.ReviveMax - t.ReviveMin
	if n < 0 {
		panic(fmt.Errorf("template:不可能"))
	}
	if n == 0 {
		return t.ReviveMin
	}
	randomN := rand.Int31n(n)
	return randomN + t.ReviveMin
}

func (t *ArenaConstantTemplate) TemplateId() int {
	return t.Id
}

func (at *ArenaConstantTemplate) FileName() string {
	return "tb_arena_constant.json"
}

func init() {
	template.Register((*ArenaConstantTemplate)(nil))
}
