package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fmt"
)

//结义道具配置
type JieYiDaoJuTemplate struct {
	*JieYiDaoJuTemplateVO
	daoJuType   jieyitypes.JieYiDaoJuType
	daoJuName   string
	needItemMap map[int32]int32
	fashionMap  map[int32]int32
}

func (t *JieYiDaoJuTemplate) TemplateId() int {
	return t.Id
}

func (t *JieYiDaoJuTemplate) GetDaoJuType() jieyitypes.JieYiDaoJuType {
	return t.daoJuType
}

func (t *JieYiDaoJuTemplate) GetDaoJuName() string {
	return t.daoJuName
}

func (t *JieYiDaoJuTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *JieYiDaoJuTemplate) GetFashionMap() map[int32]int32 {
	return t.fashionMap
}

func (t *JieYiDaoJuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证道具类型
	t.daoJuType = jieyitypes.JieYiDaoJuType(t.Type)
	if !t.daoJuType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.needItemMap = make(map[int32]int32)
	//验证物品id
	if t.UseItemId != 0 {
		to := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemId)
			return template.NewTemplateFieldError("UseItemId", err)
		}
		temp := to.(*ItemTemplate)
		t.daoJuName = temp.Name

		err = validator.MinValidate(float64(t.UseItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemCount)
			return template.NewTemplateFieldError("UseItemCount", err)
		}
		t.needItemMap[t.UseItemId] = t.UseItemCount
	}

	t.fashionMap = make(map[int32]int32)
	// 验证时装id
	useItemIdArr, err := utils.SplitAsIntArray(t.GetItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GetItemId)
		return template.NewTemplateFieldError("GetItemId", err)
	}
	useItemCountArr, err := utils.SplitAsIntArray(t.GetItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GetItemCount)
		return template.NewTemplateFieldError("GetItemCount", err)
	}
	if len(useItemIdArr) != len(useItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.GetItemId, t.GetItemCount)
		return template.NewTemplateFieldError("GetItemId or GetItemCount", err)
	}
	if len(useItemIdArr) > 0 {
		for index, itemId := range useItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.GetItemId)
				return template.NewTemplateFieldError("GetItemId", err)
			}

			err = validator.MinValidate(float64(useItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("GetItemCount", err)
			}

			t.fashionMap[itemId] = useItemCountArr[index]
		}
	}

	return
}

func (t *JieYiDaoJuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证时装id
	err = validator.MinValidate(float64(t.FashionId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FashionId)
		return template.NewTemplateFieldError("FashionId", err)
	}

	return nil
}

func (t *JieYiDaoJuTemplate) PatchAfterCheck() {
}

func (t *JieYiDaoJuTemplate) FileName() string {
	return "tb_jieyi_use.json"
}

func init() {
	template.Register((*JieYiDaoJuTemplate)(nil))
}
