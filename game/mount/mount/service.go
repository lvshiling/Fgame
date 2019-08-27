package mount

import (
	"fgame/fgame/core/template"
	mounttypes "fgame/fgame/game/mount/types"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//坐骑接口处理
type MountService interface {
	//获取坐骑进阶配置
	GetMountNumber(number int32) *gametemplate.MountTemplate
	//获取坐骑配置
	GetMount(id int) *gametemplate.MountTemplate
	//获取坐骑幻化配置
	GetMountHuanHuaTemplate(level int32) *gametemplate.MountHuanHuaTemplate
	//获取坐骑草料配置
	GetMountCaoLiaoTemplate(level int32) *gametemplate.MountCaoLiaoTemplate
	RandomMountTemplate() *gametemplate.MountTemplate
	//吃草料升级
	GetMountEatCaoLiaoTemplate(curCulLevel int32, num int32) (*gametemplate.MountCaoLiaoTemplate, bool)
	//吃幻化丹升级
	GetMountEatHuanHuanTemplate(curLevel int32, num int32) (*gametemplate.MountHuanHuaTemplate, bool)
}

type mountService struct {
	//进阶map
	mountNumberMap map[int32]*gametemplate.MountTemplate
	//坐骑配置
	mountMap map[int]*gametemplate.MountTemplate
	//坐骑幻化配置
	huanHuaMap map[int32]*gametemplate.MountHuanHuaTemplate
	//坐骑草料配置
	caoLiaoMap map[int32]*gametemplate.MountCaoLiaoTemplate
	mountList  []*gametemplate.MountTemplate
}

//初始化
func (ms *mountService) init() error {
	ms.mountNumberMap = make(map[int32]*gametemplate.MountTemplate)
	ms.mountMap = make(map[int]*gametemplate.MountTemplate)
	ms.huanHuaMap = make(map[int32]*gametemplate.MountHuanHuaTemplate)
	ms.caoLiaoMap = make(map[int32]*gametemplate.MountCaoLiaoTemplate)
	//坐骑
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MountTemplate)(nil))
	for _, templateObject := range templateMap {
		mountTemplate, _ := templateObject.(*gametemplate.MountTemplate)
		ms.mountMap[mountTemplate.TemplateId()] = mountTemplate

		typ := mountTemplate.GetTyp()
		if typ == mounttypes.MountTypeAdvanced {
			ms.mountNumberMap[mountTemplate.Number] = mountTemplate
		}
		ms.mountList = append(ms.mountList, mountTemplate)
	}

	//坐骑幻化
	huanHuaTemplateMap := template.GetTemplateService().GetAll((*gametemplate.MountHuanHuaTemplate)(nil))
	for _, templateObject := range huanHuaTemplateMap {
		mountHuanHuaTemplate, _ := templateObject.(*gametemplate.MountHuanHuaTemplate)
		ms.huanHuaMap[mountHuanHuaTemplate.Level] = mountHuanHuaTemplate
	}
	//坐骑草料
	caoLiaoTemplateMap := template.GetTemplateService().GetAll((*gametemplate.MountCaoLiaoTemplate)(nil))
	for _, templateObject := range caoLiaoTemplateMap {
		mountCaoLiaoTemplate, _ := templateObject.(*gametemplate.MountCaoLiaoTemplate)
		ms.caoLiaoMap[mountCaoLiaoTemplate.Level] = mountCaoLiaoTemplate

	}

	return nil
}

//获取坐骑进阶配置
func (ms *mountService) GetMountNumber(number int32) *gametemplate.MountTemplate {
	to, ok := ms.mountNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取坐骑配置
func (ms *mountService) GetMount(id int) *gametemplate.MountTemplate {
	to, ok := ms.mountMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取坐骑幻化配置
func (ms *mountService) GetMountHuanHuaTemplate(level int32) *gametemplate.MountHuanHuaTemplate {
	to, ok := ms.huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取坐骑草料配置
func (ms *mountService) GetMountCaoLiaoTemplate(level int32) *gametemplate.MountCaoLiaoTemplate {
	to, ok := ms.caoLiaoMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取坐骑配置
func (ms *mountService) RandomMountTemplate() *gametemplate.MountTemplate {
	num := len(ms.mountList)
	index := rand.Intn(num)
	return ms.mountList[index]
}

func (ms *mountService) GetMountEatCaoLiaoTemplate(curCulLevel int32, num int32) (caoLiaoTemplate *gametemplate.MountCaoLiaoTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curCulLevel + 1; leftNum > 0; level++ {
		caoLiaoTemplate, flag = ms.caoLiaoMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= caoLiaoTemplate.ItemCount
	}
	//次数要满足刚好升级升级
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

func (ms *mountService) GetMountEatHuanHuanTemplate(curLevel int32, num int32) (huanHuaTemplate *gametemplate.MountHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		huanHuaTemplate, flag = ms.huanHuaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= huanHuaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *mountService
)

func Init() (err error) {
	once.Do(func() {
		cs = &mountService{}
		err = cs.init()
	})
	return err
}

func GetMountService() MountService {
	return cs
}
