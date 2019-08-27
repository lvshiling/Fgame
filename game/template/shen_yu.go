package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//神域配置
type ShenYuTemplate struct {
	*ShenYuTemplateVO
	nextTemplate *ShenYuTemplate
	luckItemMap  map[int32]int32
}

func (t *ShenYuTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenYuTemplate) GetNextTemp() *ShenYuTemplate {
	return t.nextTemplate
}

func (t *ShenYuTemplate) GetLuckItemMap() map[int32]int32 {
	return t.luckItemMap
}

func (t *ShenYuTemplate) IsResetkey() bool {
	return t.ResetKeyFlag != 0
}

func (t *ShenYuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//幸运奖物品
	t.luckItemMap = make(map[int32]int32)
	luckyItemArr, err := coreutils.SplitAsIntArray(t.LuckyItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LuckyItemId)
		err = template.NewTemplateFieldError("LuckyItemId", err)
		return
	}
	luckyCountArr, err := coreutils.SplitAsIntArray(t.LuckyItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LuckyItemCount)
		err = template.NewTemplateFieldError("LuckyItemCount", err)
		return
	}
	if len(luckyItemArr) != len(luckyCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.LuckyItemId, t.LuckyItemCount)
		err = template.NewTemplateFieldError("LuckyItemId or LuckyItemCount", err)
		return
	}
	for index, itemId := range luckyItemArr {
		t.luckItemMap[itemId] += luckyCountArr[index]
	}

	//下一阶强化
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(t.NextId, (*ShenYuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemplate = to.(*ShenYuTemplate)
		if t.nextTemplate.RoundType-t.RoundType != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	return nil
}

func (t *ShenYuTemplate) PatchAfterCheck() {

}

func (t *ShenYuTemplate) Check() (err error) {
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
		return template.NewTemplateFieldError("NextId", err)
	}
	mapTemp, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return template.NewTemplateFieldError("NextId", err)
	}
	if mapTemp.GetMapType() != scenetypes.SceneTypeShenYu {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("NextId", err)
	}

	//
	for itemId, num := range t.luckItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.LuckyItemId)
			err = template.NewTemplateFieldError("LuckyItemId", err)
			return
		}

		//数量
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.LuckyItemCount)
			return template.NewTemplateFieldError("LuckyItemCount", err)
		}
	}

	return nil
}

func (t *ShenYuTemplate) FileName() string {
	return "tb_shenyu.json"
}

func init() {
	template.Register((*ShenYuTemplate)(nil))
}
