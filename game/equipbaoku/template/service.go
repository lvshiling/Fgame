package template

import (
	"fgame/fgame/core/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//装备宝库接口处理
type EquipBaoKuTemplateService interface {
	//获取装备宝库配置
	GetEquipBaoKuByLevAndZhuanNum(lev int32, zhuanNum int32, typ equipbaokutypes.BaoKuType) *gametemplate.EquipBaoKuTemplate
	//获取宝库积分商城配置
	GetBaoKuJiFenTemplate(id int) *gametemplate.BaoKuJiFenTemplate
}

type equipBaoKuTemplateService struct {
	//装备宝库配置
	equipBaoKuMap map[int]*gametemplate.EquipBaoKuTemplate
	//宝库积分商城配置
	baoKuJiFenMap map[int]*gametemplate.BaoKuJiFenTemplate
}

//初始化
func (cs *equipBaoKuTemplateService) init() error {
	cs.equipBaoKuMap = make(map[int]*gametemplate.EquipBaoKuTemplate)
	cs.baoKuJiFenMap = make(map[int]*gametemplate.BaoKuJiFenTemplate)
	//装备宝库
	templateMap := template.GetTemplateService().GetAll((*gametemplate.EquipBaoKuTemplate)(nil))
	for _, templateObject := range templateMap {
		equipBaoKuTemplate, _ := templateObject.(*gametemplate.EquipBaoKuTemplate)
		cs.equipBaoKuMap[equipBaoKuTemplate.TemplateId()] = equipBaoKuTemplate
	}
	//装备宝库
	baoKuJiFenTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BaoKuJiFenTemplate)(nil))
	for _, templateObject := range baoKuJiFenTemplateMap {
		baoKuJiFenTemplate, _ := templateObject.(*gametemplate.BaoKuJiFenTemplate)
		cs.baoKuJiFenMap[baoKuJiFenTemplate.TemplateId()] = baoKuJiFenTemplate
	}
	return nil
}

//获取装备宝库配置
func (cs *equipBaoKuTemplateService) GetEquipBaoKuByLevAndZhuanNum(lev int32, zhuanNum int32, typ equipbaokutypes.BaoKuType) (lastTemplate *gametemplate.EquipBaoKuTemplate) {
	if typ == equipbaokutypes.BaoKuTypeEquip {
		startBoxTemplate := cs.equipBaoKuMap[int(1)]
		nextTemplate := startBoxTemplate.GetNextTemplate()
		lastTemplate = startBoxTemplate
		for nextTemplate != nil {
			if zhuanNum < nextTemplate.ZhuanshuMin || lev < nextTemplate.LevelMin {
				return
			}
			lastTemplate = nextTemplate
			nextTemplate = nextTemplate.GetNextTemplate()
		}
	} else {
		startBoxTemplate := cs.equipBaoKuMap[int(101)]
		nextTemplate := startBoxTemplate.GetNextTemplate()
		lastTemplate = startBoxTemplate
		for nextTemplate != nil {
			if zhuanNum < nextTemplate.ZhuanshuMin || lev < nextTemplate.LevelMin {
				return
			}
			lastTemplate = nextTemplate
			nextTemplate = nextTemplate.GetNextTemplate()
		}
	}
	return
}

//获取宝库积分商城配置
func (cs *equipBaoKuTemplateService) GetBaoKuJiFenTemplate(id int) *gametemplate.BaoKuJiFenTemplate {
	to, ok := cs.baoKuJiFenMap[id]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *equipBaoKuTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &equipBaoKuTemplateService{}
		err = cs.init()
	})
	return err
}

func GetEquipBaoKuTemplateService() EquipBaoKuTemplateService {
	return cs
}
