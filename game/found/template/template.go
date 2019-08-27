package found

import (
	"fgame/fgame/core/template"
	foundtypes "fgame/fgame/game/found/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type FoundTemplateService interface {
	GetFoundTemplate(id int32) *gametemplate.FoundTemplate
	GetFoundTemplateByType(typ foundtypes.FoundResourceType, resLevel int32) *gametemplate.FoundTemplate
}

type foundServiceTemplate struct {
	//资源找回模板
	foundMap        map[int32]*gametemplate.FoundTemplate
	foundResTypeMap map[foundtypes.FoundResourceType][]*gametemplate.FoundTemplate
}

//初始化找回资源配置
func (st *foundServiceTemplate) init() (err error) {
	//资源找回模板
	st.foundMap = make(map[int32]*gametemplate.FoundTemplate)
	st.foundResTypeMap = make(map[foundtypes.FoundResourceType][]*gametemplate.FoundTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.FoundTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.FoundTemplate)
		st.foundMap[int32(ftem.TemplateId())] = ftem
		st.foundResTypeMap[ftem.GetResType()] = append(st.foundResTypeMap[ftem.GetResType()], ftem)
	}

	return
}

func (st *foundServiceTemplate) GetFoundTemplate(id int32) *gametemplate.FoundTemplate {
	return st.foundMap[id]
}

func (st *foundServiceTemplate) GetFoundTemplateByType(typ foundtypes.FoundResourceType, resLevel int32) *gametemplate.FoundTemplate {
	temArr := st.foundResTypeMap[typ]
	for _, tem := range temArr {
		if resLevel >= tem.LevelMin && resLevel <= tem.LevelMax {
			return tem
		}
	}
	return nil
}

var (
	once sync.Once
	st   *foundServiceTemplate
)

func Init() (err error) {
	once.Do(func() {
		st = &foundServiceTemplate{}
		err = st.init()
	})
	return
}

func GetFoundTemplateService() FoundTemplateService {
	return st
}
