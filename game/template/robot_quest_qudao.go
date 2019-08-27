package template

import (
	accounttypes "fgame/fgame/account/login/types"
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//任务机器人
type RobotQuestQudaoTemplate struct {
	*RobotQuestQudaoTemplateVO
	sdkType               accounttypes.SDKType
	mapTemplate           *MapTemplate
	questBeginTemplate    *QuestTemplate
	questOverTemplate     *QuestTemplate
	portalTemplate        *PortalTemplate
	portalBiologyTemplate *BiologyTemplate
}

func (t *RobotQuestQudaoTemplate) TemplateId() int {
	return t.Id
}

func (t *RobotQuestQudaoTemplate) GetSDKType() accounttypes.SDKType {
	return t.sdkType
}
func (t *RobotQuestQudaoTemplate) RandomProperty() map[propertytypes.BattlePropertyType]int64 {
	props := make(map[propertytypes.BattlePropertyType]int64)
	hp := int64(mathutils.RandomRange(int(t.HpMin), int(t.HpMax)))
	attack := int64(mathutils.RandomRange(int(t.AttackMin), int(t.AttackMax)))
	defence := int64(mathutils.RandomRange(int(t.DefenceMin), int(t.DefenceMax)))
	props[propertytypes.BattlePropertyTypeMaxHP] = hp
	props[propertytypes.BattlePropertyTypeAttack] = attack
	props[propertytypes.BattlePropertyTypeDefend] = defence
	return props
}

func (t *RobotQuestQudaoTemplate) GetRefreshTime() int64 {
	return int64(t.RefreshTime)
}

func (t *RobotQuestQudaoTemplate) GetPlayerLimitCount() int32 {
	return t.PlayerLimitCount
}

func (t *RobotQuestQudaoTemplate) GetQuestBeginId() int32 {
	return t.QuestBeginId
}
func (t *RobotQuestQudaoTemplate) GetQuestOverId() int32 {
	return t.QuestOverId
}

func (t *RobotQuestQudaoTemplate) GetPortalTemplate() *PortalTemplate {

	return t.portalTemplate
}

func (t *RobotQuestQudaoTemplate) PatchAfterCheck() {
}

func (t *RobotQuestQudaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.sdkType = accounttypes.SDKType(t.QudaoId)
	tempMapTemplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}
	t.mapTemplate = tempMapTemplate.(*MapTemplate)

	tempQuestBeginTemplate := template.GetTemplateService().Get(int(t.QuestBeginId), (*QuestTemplate)(nil))
	if tempQuestBeginTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.QuestBeginId)
		err = template.NewTemplateFieldError("QuestBeginId", err)
		return
	}
	t.questBeginTemplate = tempQuestBeginTemplate.(*QuestTemplate)
	tempQuestOverTemplate := template.GetTemplateService().Get(int(t.QuestOverId), (*QuestTemplate)(nil))
	if tempQuestOverTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.QuestOverId)
		err = template.NewTemplateFieldError("QuestOverId", err)
		return
	}
	t.questOverTemplate = tempQuestOverTemplate.(*QuestTemplate)

	tempBiologyTemplate := template.GetTemplateService().Get(int(t.ChuansongzhenId), (*BiologyTemplate)(nil))
	if tempBiologyTemplate != nil {
		t.portalBiologyTemplate = tempBiologyTemplate.(*BiologyTemplate)
	}

	return nil
}

func (t *RobotQuestQudaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//验证地图类型
	if !t.sdkType.Valid() {
		err = fmt.Errorf("[%d] 不是合法的sdk", t.QudaoId)
		err = template.NewTemplateFieldError("QudaoId", err)
		return
	}
	//验证地图类型
	if !t.mapTemplate.IsWorld() {
		err = fmt.Errorf("[%d] 不是世界地图", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}
	//验证人数
	err = validator.MinValidate(float64(t.PlayerLimitCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerLimitCount)
		err = template.NewTemplateFieldError("PlayerLimitCount", err)
		return
	}
	//验证任务类型
	if t.questBeginTemplate.GetQuestType() != questtypes.QuestTypeOnce {
		err = fmt.Errorf("[%d] 不是主线任务", t.QuestBeginId)
		err = template.NewTemplateFieldError("QuestBeginId", err)
		return
	}
	if t.questOverTemplate.GetQuestType() != questtypes.QuestTypeOnce {
		err = fmt.Errorf("[%d] 不是主线任务", t.QuestOverId)
		err = template.NewTemplateFieldError("QuestOverId", err)
		return
	}

	//验证血量
	err = validator.MinValidate(float64(t.HpMin), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.HpMin)
		err = template.NewTemplateFieldError("HpMin", err)
		return
	}

	//验证血量
	err = validator.MinValidate(float64(t.HpMax), float64(t.HpMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.HpMin)
		err = template.NewTemplateFieldError("HpMin", err)
		return
	}

	//验证血量
	err = validator.MinValidate(float64(t.HpMin), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.HpMin)
		err = template.NewTemplateFieldError("HpMin", err)
		return
	}

	//攻击
	err = validator.MinValidate(float64(t.AttackMin), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttackMin)
		err = template.NewTemplateFieldError("AttackMin", err)
		return
	}
	//攻击
	err = validator.MinValidate(float64(t.AttackMax), float64(t.AttackMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttackMax)
		err = template.NewTemplateFieldError("AttackMax", err)
		return
	}
	//防御
	err = validator.MinValidate(float64(t.DefenceMin), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DefenceMin)
		err = template.NewTemplateFieldError("DefenceMin", err)
		return
	}
	//防御
	err = validator.MinValidate(float64(t.DefenceMax), float64(t.DefenceMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DefenceMax)
		err = template.NewTemplateFieldError("DefenceMax", err)
		return
	}
	//防御
	err = validator.RangeValidate(float64(t.PlayerLimitCount), float64(0), true, float64(robotMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerLimitCount)
		err = template.NewTemplateFieldError("PlayerLimitCount", err)
		return
	}
	//防御
	err = validator.MinValidate(float64(t.RefreshTime), float64(refreshTimeMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RefreshTime)
		err = template.NewTemplateFieldError("RefreshTime", err)
		return
	}
	if t.portalBiologyTemplate != nil {
		t.portalTemplate = t.portalBiologyTemplate.GetPortalTemplate()
	}

	return nil
}

func (t *RobotQuestQudaoTemplate) FileName() string {
	return "tb_robot_quest_qudao.json"
}

func init() {
	template.Register((*RobotQuestQudaoTemplate)(nil))
}
