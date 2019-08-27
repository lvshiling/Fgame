package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fmt"
)

type ChuangShiYuGaoTemplate struct {
	*ChuangShiYuGaoTemplateVO
	jianGeTime     int64 // 毫秒
	receiveItemMap map[int32]int32
}

func (t *ChuangShiYuGaoTemplate) GetReceiveItemMap() map[int32]int32 {
	return t.receiveItemMap
}

func (t *ChuangShiYuGaoTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiYuGaoTemplate) GetJianGeTime() int64 {
	return t.jianGeTime
}

func (t *ChuangShiYuGaoTemplate) FileName() string {
	return "tb_chuangshi_yugao.json"
}

func (t *ChuangShiYuGaoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 报名获得
	t.receiveItemMap = make(map[int32]int32)
	useItemIdArr, err := utils.SplitAsIntArray(t.BaoMingGet)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BaoMingGet)
		return template.NewTemplateFieldError("BaoMingGet", err)
	}
	useItemCountArr, err := utils.SplitAsIntArray(t.BaoMingGetCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BaoMingGetCount)
		return template.NewTemplateFieldError("BaoMingGetCount", err)
	}
	if len(useItemIdArr) != len(useItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.BaoMingGet, t.BaoMingGetCount)
		return template.NewTemplateFieldError("UseItemId or BaoMingGetCount", err)
	}
	if len(useItemIdArr) > 0 {
		for index, itemId := range useItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.BaoMingGet)
				return template.NewTemplateFieldError("BaoMingGet", err)
			}

			err = validator.MinValidate(float64(useItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UseItemCount", err)
			}

			t.receiveItemMap[itemId] = useItemCountArr[index]
		}
	}

	return
}

func (t *ChuangShiYuGaoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证持续天数
	err = validator.MinValidate(float64(t.DaoJiShiDay), 0, true)
	if err != nil {
		return template.NewTemplateFieldError("DaoJiShi", err)
	}
	// 验证间隔时间
	err = validator.MinValidate(float64(t.JiaRenTime), 0, false)
	if err != nil {
		return template.NewTemplateFieldError("JiaRenTime", err)
	}
	// 验证增加的人数
	err = validator.MinValidate(float64(t.JiaRenCount), 0, true)
	if err != nil {
		return template.NewTemplateFieldError("JiaRenCount", err)
	}
	return
}

func (t *ChuangShiYuGaoTemplate) PatchAfterCheck() {
	t.jianGeTime = t.JiaRenTime * int64(common.DAY)
}

func init() {
	template.Register((*ChuangShiYuGaoTemplate)(nil))
}
