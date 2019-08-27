package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/team/types"
	"fmt"
	"math/rand"
)

//组队副本配置
type TeamCopyTemplate struct {
	*TeamCopyTemplateVO
	purposeType     types.TeamPurposeType  //类型
	biologyTemplate *BiologyTemplate       //生物模板
	rewData         *propertytypes.RewData //奖励属性
	rewItemMap      map[int32]int32        //奖励物品
	mapTempate      *MapTemplate           //地图
}

func (t *TeamCopyTemplate) TemplateId() int {
	return t.Id
}

func (t *TeamCopyTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *TeamCopyTemplate) GetPurposeType() types.TeamPurposeType {
	return t.purposeType
}

func (t *TeamCopyTemplate) GetRewData() *propertytypes.RewData {
	return t.rewData
}

func (t *TeamCopyTemplate) GetBornPos() coretypes.Position {
	return t.mapTempate.GetBornPos()
}

func (t *TeamCopyTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTempate
}

func (t *TeamCopyTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证类型
	t.purposeType = types.TeamPurposeType(t.Type)
	if !t.purposeType.Vaild() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.rewItemMap = make(map[int32]int32)
	if t.ItemId != "" {
		rewItemIdArr, err := utils.SplitAsIntArray(t.ItemId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.ItemId)
			return template.NewTemplateFieldError("ItemId", err)
		}
		rewItemCountArr, err := utils.SplitAsIntArray(t.ItemCount)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}
		if len(rewItemIdArr) == 0 || len(rewItemCountArr) != len(rewItemIdArr) {
			err = fmt.Errorf("[%s] invalid", t.ItemId)
			return template.NewTemplateFieldError("RewItemId", err)
		}
		for i := 0; i < len(rewItemIdArr); i++ {
			t.rewItemMap[rewItemIdArr[i]] = rewItemCountArr[i]
		}
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(t.RewardSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardSilver)
		return template.NewTemplateFieldError("RewardSilver", err)
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(t.RewardGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardGold)
		return template.NewTemplateFieldError("RewardGold", err)
	}

	//验证 rew_bindgold
	err = validator.MinValidate(float64(t.RewardBindgold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardBindgold)
		return template.NewTemplateFieldError("RewardBindgold", err)
	}

	if t.RewardSilver > 0 || t.RewardGold > 0 || t.RewardBindgold > 0 {
		t.rewData = propertytypes.CreateRewData(0, 0, t.RewardSilver, t.RewardGold, t.RewardBindgold)
	}

	return nil
}

func (t *TeamCopyTemplate) PatchAfterCheck() {

}

func (t *TeamCopyTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	for itemId, num := range t.rewItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.ItemId)
			err = template.NewTemplateFieldError("ItemId", err)
			return
		}
		if num <= 0 {
			err = fmt.Errorf("[%s] invalid", t.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//验证 reward_number
	err = validator.MinValidate(float64(t.RewardNumber), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewardNumber)
		return template.NewTemplateFieldError("RewardNumber", err)
	}

	//验证 resurrection_number
	err = validator.MinValidate(float64(t.ResurrectionNumber), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ResurrectionNumber)
		return template.NewTemplateFieldError("ResurrectionNumber", err)
	}

	to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	t.mapTempate = to.(*MapTemplate)

	if t.mapTempate.GetMapType() != scenetypes.SceneTypeCrossTeamCopy {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	return nil
}

func (t *TeamCopyTemplate) RandomRevive() int32 {
	n := t.ResurrectionNumber
	if n < 0 {
		panic(fmt.Errorf("template:不可能"))
	}
	if n == 0 {
		return t.ResurrectionNumber
	}
	randomN := rand.Int31n(n)
	return randomN
}

func (t *TeamCopyTemplate) FileName() string {
	return "tb_zuduifuben.json"
}

func init() {
	template.Register((*TeamCopyTemplate)(nil))
}
