package template

import (
	"fgame/fgame/core/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fmt"
)

func init() {
	template.Register((*PkRewardCdTemplate)(nil))
}

type PkRewardCdTemplate struct {
	*PkRewardCdTemplateVO
	activityType activitytypes.ActivityType
}

func (t *PkRewardCdTemplate) TemplateId() int {
	return t.Id
}

func (t *PkRewardCdTemplate) FileName() string {
	return "tb_pk_reward_cd.json"
}

//组合成需要的数据
func (t *PkRewardCdTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *PkRewardCdTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.activityType = activitytypes.ActivityType(t.ActivityType)
	if !t.activityType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.ActivityType)
		return template.NewTemplateFieldError("ActivityType", err)
	}

	return nil
}

func (t *PkRewardCdTemplate) GetActivityType() activitytypes.ActivityType {
	return t.activityType
}

//检验后组合
func (t *PkRewardCdTemplate) PatchAfterCheck() {
}
