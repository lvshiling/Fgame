package weapon

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//兵魂接口处理
type WeaponService interface {
	//获取兵魂配置
	GetWeaponTemplate(id int) *gametemplate.WeaponTemplate
	RandomWeaponTemplate() *gametemplate.WeaponTemplate
}

type weaponService struct {
	//兵魂配置
	weaponMap map[int]*gametemplate.WeaponTemplate

	weaponList []*gametemplate.WeaponTemplate
}

//初始化
func (ws *weaponService) init() error {
	ws.weaponMap = make(map[int]*gametemplate.WeaponTemplate)
	//兵魂
	templateMap := template.GetTemplateService().GetAll((*gametemplate.WeaponTemplate)(nil))
	for _, templateObject := range templateMap {
		weaponTemplate, _ := templateObject.(*gametemplate.WeaponTemplate)
		ws.weaponMap[weaponTemplate.TemplateId()] = weaponTemplate
		ws.weaponList = append(ws.weaponList, weaponTemplate)
	}

	return nil
}

func (ws *weaponService) RandomWeaponTemplate() *gametemplate.WeaponTemplate {
	num := len(ws.weaponList)
	index := rand.Intn(num)
	return ws.weaponList[index]
}

//获取兵魂配置
func (ws *weaponService) GetWeaponTemplate(id int) *gametemplate.WeaponTemplate {
	to, ok := ws.weaponMap[id]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *weaponService
)

func Init() (err error) {
	once.Do(func() {
		cs = &weaponService{}
		err = cs.init()
	})
	return err
}

func GetWeaponService() WeaponService {
	return cs
}
