package template

import (
	"fgame/fgame/core/template"
	propertytypes "fgame/fgame/game/property/types"
)

//天赋等级配置
type BuffDongTaiTemplate struct {
	*BuffDongTaiTemplateVO
	//属性加成
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//属性万分比加成
	battlePropertyPercentMap map[propertytypes.BattlePropertyType]int64
}

func (t *BuffDongTaiTemplate) TemplateId() int {
	return t.Id
}

func (t *BuffDongTaiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)

	t.battlePropertyMap[propertytypes.BattlePropertyTypeTough] = int64(t.ToughAdd)

	t.battlePropertyPercentMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMoveSpeed] = int64(t.SpeedMovePercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeTough] = int64(t.ToughPercent)

	return nil
}

func (t *BuffDongTaiTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *BuffDongTaiTemplate) GetBattlePropertyPercentMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyPercentMap
}

func (t *BuffDongTaiTemplate) PatchAfterCheck() {

}

func (t *BuffDongTaiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *BuffDongTaiTemplate) FileName() string {
	return "tb_buff_dongtai.json"
}

func init() {
	template.Register((*BuffDongTaiTemplate)(nil))
}
