package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type WushuangWeaponTemplateService interface {
	GetWushuangWeaponBaseTemplate(itemId int32) *gametemplate.WushuangWeaponBaseTemplate
	GetWushuangWeaponBuchangTemplate() *gametemplate.WushuangWeaponBuchangTemplate
}

type wushuangWeaponTemplateService struct {
	wushuangBaseMap map[int32]*gametemplate.WushuangWeaponBaseTemplate
	buchangTemp     *gametemplate.WushuangWeaponBuchangTemplate
}

func (st *wushuangWeaponTemplateService) init() (err error) {
	st.wushuangBaseMap = make(map[int32]*gametemplate.WushuangWeaponBaseTemplate)
	tempMap := template.GetTemplateService().GetAll((*gametemplate.WushuangWeaponBaseTemplate)(nil))
	for itemId, temp := range tempMap {
		baseTemp, _ := temp.(*gametemplate.WushuangWeaponBaseTemplate)
		st.wushuangBaseMap[int32(itemId)] = baseTemp
	}

	buchangTempList := template.GetTemplateService().GetAll((*gametemplate.WushuangWeaponBuchangTemplate)(nil))
	//目前只有一条
	if len(buchangTempList) != 1 {
		panic("wushuangweapon:无双神器补偿配置应该只有一条")
	}
	for _, buchangtemp := range buchangTempList {
		temp, _ := buchangtemp.(*gametemplate.WushuangWeaponBuchangTemplate)
		st.buchangTemp = temp
	}

	return
}

func (st *wushuangWeaponTemplateService) GetWushuangWeaponBuchangTemplate() *gametemplate.WushuangWeaponBuchangTemplate {
	return st.buchangTemp
}

func (st *wushuangWeaponTemplateService) GetWushuangWeaponBaseTemplate(itemId int32) *gametemplate.WushuangWeaponBaseTemplate {
	temp, ok := st.wushuangBaseMap[itemId]
	if !ok {
		return nil
	}
	return temp
}

var (
	once sync.Once
	st   *wushuangWeaponTemplateService
)

func Init() (err error) {
	once.Do(func() {
		st = &wushuangWeaponTemplateService{}
		err = st.init()
	})

	return
}

func GetWushuangWeaponTemplateService() WushuangWeaponTemplateService {
	return st
}
