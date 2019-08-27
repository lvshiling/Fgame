package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/mingge/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//命格配置
type MingGeTemplate struct {
	*MingGeTemplateVO
	mingGeType        types.MingGeType
	mingGeSubType     types.MingGeAllSubType
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

func (mt *MingGeTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MingGeTemplate) GetMingGeType() types.MingGeType {
	return mt.mingGeType
}

func (mt *MingGeTemplate) GetMingGeSubType() types.MingGeAllSubType {
	return mt.mingGeSubType
}

func (mt *MingGeTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return mt.battlePropertyMap
}

func (mt *MingGeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	mt.mingGeType = types.MingGeType(mt.Type)
	if !mt.mingGeType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	mt.mingGeSubType = types.MingGeAllSubType(mt.SubType)
	if !mt.mingGeSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.SubType)
		err = template.NewTemplateFieldError("SubType", err)
		return
	}

	//属性
	mt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(mt.Hp)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(mt.Attack)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(mt.Defence)

	return nil
}

func (mt *MingGeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (mt *MingGeTemplate) PatchAfterCheck() {

}

func (mt *MingGeTemplate) FileName() string {
	return "tb_mingge.json"
}

func init() {
	template.Register((*MingGeTemplate)(nil))
}
