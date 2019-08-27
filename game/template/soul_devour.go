package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fmt"
)

//帝魂吞噬配置
type SoulDevourTemplate struct {
	*SoulDevourTemplateVO
	devourExpMap map[int32]int32 //吞噬物品经验值

}

func (sdt *SoulDevourTemplate) TemplateId() int {
	return sdt.Id
}

func (sdt *SoulDevourTemplate) GetDevourExpMap() map[int32]int32 {
	return sdt.devourExpMap
}

func (sdt *SoulDevourTemplate) PatchAfterCheck() {

}

func (sdt *SoulDevourTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sdt.FileName(), sdt.TemplateId(), err)
			return
		}
	}()

	//devour_id
	if sdt.DevourId != "" {
		itemArr, err := utils.SplitAsIntArray(sdt.DevourId)
		if len(itemArr) <= 0 {
			err = fmt.Errorf("[%s] invalid", sdt.DevourId)
			err = template.NewTemplateFieldError("DevourId", err)
			return err
		}

		sdt.devourExpMap = make(map[int32]int32)
		for _, itemId := range itemArr {
			to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("[%s] invalid", sdt.DevourId)
				err = template.NewTemplateFieldError("DevourId", err)
				return err
			}
			//单个经验值
			sdt.devourExpMap[itemId] = to.(*ItemTemplate).TypeFlag1
		}
	}

	return nil
}

func (sdt *SoulDevourTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sdt.FileName(), sdt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (sdt *SoulDevourTemplate) FileName() string {
	return "tb_soul_devour.json"
}

func init() {
	template.Register((*SoulDevourTemplate)(nil))
}
