package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	moonlovetypes "fgame/fgame/game/moonlove/types"
	"fmt"
)

func init() {
	template.Register((*MoonloveRankTemplate)(nil))
}

//月下情缘排行榜配置
type MoonloveRankTemplate struct {
	*MoonloveRankTemplateVO
	rankType      moonlovetypes.MoonloveRankType
	rewardItemMap map[int32]int32
}

func (t *MoonloveRankTemplate) FileName() string {
	return "tb_moon_love_rank.json"
}

func (t *MoonloveRankTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewardItemMap = make(map[int32]int32)

	//奖励物品
	itemIdArr, err := utils.SplitAsIntArray(t.RewItemId)
	if err != nil {
		return template.NewTemplateFieldError("RewItemId", fmt.Errorf("[%s] invalid", t.RewItemId))
	}
	//奖励数量
	itemCountArr, err := utils.SplitAsIntArray(t.RewItemCount)
	if err != nil {
		return template.NewTemplateFieldError("RewItemCount", fmt.Errorf("[%s] invalid", t.RewItemCount))
	}
	if len(itemIdArr) != len(itemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewItemId, t.RewItemCount)
		return template.NewTemplateFieldError("RewItemId or RewItemCount", err)
	}
	if len(itemIdArr) > 0 {
		for index, itemId := range itemIdArr {

			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("RewItemId", fmt.Errorf("[%s] invalid", t.RewItemId))
			}

			err = validator.MinValidate(float64(itemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RewItemCount", err)
			}

			t.rewardItemMap[itemId] = itemCountArr[index]
		}

	}

	return nil
}

func (t *MoonloveRankTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//类型
	typ := moonlovetypes.MoonloveRankType(t.Type)
	if !typ.Valid() {
		err = template.NewTemplateFieldError("Type", fmt.Errorf("[%d] invalid", t.Type))
		return
	}
	t.rankType = typ
	//排名
	if err = validator.MinValidate(float64(t.Rank), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("Rank", err)
		return
	}
	//奖励经验
	if err = validator.MinValidate(float64(t.RewExp), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}
	//奖励经验点
	if err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}
	//奖励银两
	if err = validator.MinValidate(float64(t.RewSilver), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}
	//奖励元宝
	if err = validator.MinValidate(float64(t.RewGold), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}
	//奖励绑元
	if err = validator.MinValidate(float64(t.RewBindGold), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}

	return nil
}

func (t *MoonloveRankTemplate) PatchAfterCheck() {

}

func (t *MoonloveRankTemplate) TemplateId() int {
	return t.Id
}

func (t *MoonloveRankTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewardItemMap
}

func (t *MoonloveRankTemplate) GetRankType() moonlovetypes.MoonloveRankType {
	return t.rankType
}
