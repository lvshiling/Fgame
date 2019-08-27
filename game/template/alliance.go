package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fmt"
)

//活动模板配置
type AllianceTemplate struct {
	*AllianceTemplateVO
	allianceVersion           alliancetypes.AllianceVersionType
	alliancePositionNumMap    map[alliancetypes.AlliancePosition]int32
	nextLevelAllianceTemplate *AllianceTemplate
}

func (at *AllianceTemplate) GetAlliancePositionNum(pos alliancetypes.AlliancePosition) int32 {
	return at.alliancePositionNumMap[pos]
}

func (t *AllianceTemplate) GetNextLevelAllianceTemplate() *AllianceTemplate {
	return t.nextLevelAllianceTemplate
}

func (t *AllianceTemplate) GetAllianceVersion() alliancetypes.AllianceVersionType {
	return t.allianceVersion
}

func (at *AllianceTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()

	at.alliancePositionNumMap = make(map[alliancetypes.AlliancePosition]int32)

	at.alliancePositionNumMap[alliancetypes.AlliancePositionFuMengZhu] = at.UnionPost2
	at.alliancePositionNumMap[alliancetypes.AlliancePositionTangZhu] = at.UnionPost3
	at.alliancePositionNumMap[alliancetypes.AlliancePositionJingYing] = at.UnionPost4
	if at.NextLevelId != 0 {
		tempNextLevelAllianceTemplate := template.GetTemplateService().Get(int(at.NextLevelId), (*AllianceTemplate)(nil))
		if tempNextLevelAllianceTemplate != nil {
			at.nextLevelAllianceTemplate = tempNextLevelAllianceTemplate.(*AllianceTemplate)
		}
	}
	return
}
func (at *AllianceTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()

	//建设度
	err = validator.MinValidate(float64(at.UnionBuild), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.UnionBuild)
		return template.NewTemplateFieldError("UnionBuild", err)
	}
	//下一个等级
	if at.NextLevelId != 0 {
		if at.nextLevelAllianceTemplate == nil {
			err = fmt.Errorf("[%d] invalid", at.NextLevelId)
			return template.NewTemplateFieldError("NextLevelId", err)
		}
	}

	version := alliancetypes.AllianceVersionType(at.UnionType)
	if !version.Valid() {
		err = fmt.Errorf("[%d] invalid", at.UnionType)
		return template.NewTemplateFieldError("UnionType", err)
	}
	at.allianceVersion = version

	//仙盟最高人数
	err = validator.MaxValidate(float64(at.UnionMax), float64(100), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.UnionMax)
		return template.NewTemplateFieldError("UnionMax", err)
	}

	//职位人数
	for _, count := range at.alliancePositionNumMap {
		err = validator.MinValidate(float64(count), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", count)
			return template.NewTemplateFieldError("UnionPost", err)
		}
	}

	//仙盟仓库数量
	err = validator.MinValidate(float64(at.UnionStorage), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.UnionStorage)
		return template.NewTemplateFieldError("UnionStorage", err)
	}

	return
}

func (at *AllianceTemplate) PatchAfterCheck() {

}
func (at *AllianceTemplate) TemplateId() int {
	return at.Id
}

func (at *AllianceTemplate) FileName() string {
	return "tb_union.json"
}

func init() {
	template.Register((*AllianceTemplate)(nil))
}
