package template

import (
	"fgame/fgame/core/template"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fmt"
)

func init() {
	template.Register((*GuideReplicaTemplate)(nil))
}

type GuideReplicaTemplate struct {
	*GuideReplicaTemplateVO
	guideType   guidereplicatypes.GuideReplicaType
	bossTemp    *BiologyTemplate
	catDogTemp  *GuideReplicaCatDogTemplate  //猫狗模板
	moJianTemp  *GuideReplicaMoJianTemplate  //魔剑模板
	rescureTemp *GuideReplicaRescureTemplate //救援模板
}

func (t *GuideReplicaTemplate) TemplateId() int {
	return t.Id
}

func (t *GuideReplicaTemplate) GetGuideType() guidereplicatypes.GuideReplicaType {
	return t.guideType
}

func (t *GuideReplicaTemplate) GetCatDogGuideTemp() *GuideReplicaCatDogTemplate {
	return t.catDogTemp
}

func (t *GuideReplicaTemplate) GetMoJianGuideTemp() *GuideReplicaMoJianTemplate {
	return t.moJianTemp
}

func (t *GuideReplicaTemplate) GetRescureGuideTemp() *GuideReplicaRescureTemplate {
	return t.rescureTemp
}

func (t *GuideReplicaTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}

func (t *GuideReplicaTemplate) FileName() string {
	return "tb_yindao_boss.json"
}

//组合成需要的数据
func (t *GuideReplicaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//引导类型
	t.guideType = guidereplicatypes.GuideReplicaType(t.Type)
	if !t.guideType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//猫狗
	switch t.guideType {
	case guidereplicatypes.GuideReplicaTypeCatDog:
		to := template.GetTemplateService().Get(int(t.SubGuideId), (*GuideReplicaCatDogTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SubGuideId)
			return template.NewTemplateFieldError("SubGuideId", err)
		}
		t.catDogTemp = to.(*GuideReplicaCatDogTemplate)
	case guidereplicatypes.GuideReplicaTypeMoJian:
		to := template.GetTemplateService().Get(int(t.SubGuideId), (*GuideReplicaMoJianTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SubGuideId)
			return template.NewTemplateFieldError("SubGuideId", err)
		}
		t.moJianTemp = to.(*GuideReplicaMoJianTemplate)
	case guidereplicatypes.GuideReplicaTypeRescure:
		to := template.GetTemplateService().Get(int(t.SubGuideId), (*GuideReplicaRescureTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.SubGuideId)
			return template.NewTemplateFieldError("SubGuideId", err)
		}
		t.rescureTemp = to.(*GuideReplicaRescureTemplate)
	}

	return nil
}

//检查有效性
func (t *GuideReplicaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//地图
	mto := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mto == nil {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return err
	}
	_, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return
	}

	//怪物id
	if t.BiologyId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.BiologyId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.BiologyId)
			return err
		}
		bossTemp, ok := bilogyTemp.(*BiologyTemplate)
		if !ok {
			err = fmt.Errorf("BiologyId [%d] no exist", t.BiologyId)
			return
		}
		t.bossTemp = bossTemp
	}

	return nil
}

//检验后组合
func (t *GuideReplicaTemplate) PatchAfterCheck() {
}
