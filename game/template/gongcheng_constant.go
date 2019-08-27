package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	scenetypes "fgame/fgame/game/scene/types"

	"fmt"
)

//神兽攻城常量配置
type GongChengConstantTemplate struct {
	*GongChengConstantTemplateVO
	biologyTemplateMap map[int32]*BiologyTemplate
	bossIdMap          map[int32]int32
	systemNameMap      map[int32]string
	collectIdList      []int32 //金银密窟特殊采集物id
}

func (t *GongChengConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *GongChengConstantTemplate) GetMapId(godType godsiegetypes.GodSiegeType) (mapId int32) {
	switch godType {
	case godsiegetypes.GodSiegeTypeQiLin,
		godsiegetypes.GodSiegeTypeLocalQiLin:
		{
			return t.MapId1
		}
	case godsiegetypes.GodSiegeTypeHuoFeng:
		{
			return t.MapId2
		}
	case godsiegetypes.GodSiegeTypeDuLong:
		{
			return t.MapId3
		}
	case godsiegetypes.GodSiegeTypeDenseWat:
		{
			return t.MapId4
		}
	}
	return
}

func (t *GongChengConstantTemplate) GetBiologyTemplate(mapId int32) *BiologyTemplate {
	return t.biologyTemplateMap[mapId]
}

func (t *GongChengConstantTemplate) GetSystemName(mapId int32) string {
	return t.systemNameMap[mapId]
}

func (t *GongChengConstantTemplate) GetDensewatCollectList() []int32 {
	return t.collectIdList
}

func (t *GongChengConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.bossIdMap = make(map[int32]int32)
	t.bossIdMap[t.MapId1] = t.BossId1
	t.bossIdMap[t.MapId2] = t.BossId2
	t.bossIdMap[t.MapId3] = t.BossId3

	t.systemNameMap = make(map[int32]string)
	t.systemNameMap[t.MapId1] = t.SystemName1
	t.systemNameMap[t.MapId2] = t.SystemName2
	t.systemNameMap[t.MapId3] = t.SystemName3

	//采集物id
	t.collectIdList, err = utils.SplitAsIntArray(t.CaijiBiologyId)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CaijiBiologyId)
		return template.NewTemplateFieldError("CaijiBiologyId", err)
	}

	return nil
}

func (t *GongChengConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.biologyTemplateMap = make(map[int32]*BiologyTemplate)
	for mapId, bossId := range t.bossIdMap {
		to := template.GetTemplateService().Get(int(mapId), (*MapTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mapId)
			return template.NewTemplateFieldError("MapId", err)
		}

		mapTemplate := to.(*MapTemplate)

		if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossGodSiege {
			err = fmt.Errorf("[%d] invalid", mapId)
			return template.NewTemplateFieldError("MapId", err)
		}

		bto := template.GetTemplateService().Get(int(bossId), (*BiologyTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", bossId)
			return template.NewTemplateFieldError("BossId", err)
		}
		biologyTemplate := bto.(*BiologyTemplate)

		if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeGodSiegeBoss {
			err = fmt.Errorf("[%d] invalid", bossId)
			err = template.NewTemplateFieldError("BossId", err)
			return
		}
		t.biologyTemplateMap[mapId] = biologyTemplate
	}

	err = validator.MinValidate(float64(t.BossTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossTime)
		err = template.NewTemplateFieldError("BossTime", err)
		return
	}

	err = validator.MinValidate(float64(t.PlayerLimitCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerLimitCount)
		err = template.NewTemplateFieldError("PlayerLimitCount", err)
		return
	}

	err = validator.MinValidate(float64(t.MoneyLimitCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MoneyLimitCount)
		err = template.NewTemplateFieldError("MoneyLimitCount", err)
		return
	}

	err = validator.MinValidate(float64(t.CaiJiCountLimit), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CaiJiCountLimit)
		err = template.NewTemplateFieldError("CaiJiCountLimit", err)
		return
	}

	for _, collectId := range t.collectIdList {
		to := template.GetTemplateService().Get(int(collectId), (*BiologyTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.CaijiBiologyId)
			return template.NewTemplateFieldError("CaijiBiologyId", err)
		}

		collect := to.(*BiologyTemplate)
		if collect.GetBiologyScriptType() != scenetypes.BiologyScriptTypeGeneralCollect {
			err = fmt.Errorf("[%d] invalid, ScriptType error", t.CaijiBiologyId)
			return template.NewTemplateFieldError("CaijiBiologyId", err)
		}
	}

	if t.SystemName1 == "" {
		err = fmt.Errorf("[%s] invalid", t.SystemName1)
		err = template.NewTemplateFieldError("SystemName1", err)
		return
	}

	if t.SystemName2 == "" {
		err = fmt.Errorf("[%s] invalid", t.SystemName2)
		err = template.NewTemplateFieldError("SystemName2", err)
		return
	}

	if t.SystemName3 == "" {
		err = fmt.Errorf("[%s] invalid", t.SystemName3)
		err = template.NewTemplateFieldError("SystemName3", err)
		return
	}

	return nil
}
func (t *GongChengConstantTemplate) PatchAfterCheck() {

}
func (t *GongChengConstantTemplate) FileName() string {
	return "tb_gongcheng_constant.json"
}

func init() {
	template.Register((*GongChengConstantTemplate)(nil))
}
