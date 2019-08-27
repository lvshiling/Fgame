package template

import (
	"fgame/fgame/core/template"
	centertypes "fgame/fgame/game/center/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type HuiYuanTemplateService interface {
	//获取会员配置
	GetHuiYuanTemplate(houtaiType centertypes.ZhiZunType, typ huiyuantypes.HuiYuanType) *gametemplate.HuiYuanTemplate
}

type huiyuanTemplateService struct {
	//会员特权配置
	huiyuanMap map[centertypes.ZhiZunType]map[huiyuantypes.HuiYuanType]*gametemplate.HuiYuanTemplate
}

func (st *huiyuanTemplateService) init() (err error) {
	// 会员特权配置
	st.huiyuanMap = make(map[centertypes.ZhiZunType]map[huiyuantypes.HuiYuanType]*gametemplate.HuiYuanTemplate)
	tempMap := template.GetTemplateService().GetAll((*gametemplate.HuiYuanTemplate)(nil))
	for _, temp := range tempMap {
		huiyuanTemp, _ := temp.(*gametemplate.HuiYuanTemplate)
		subMap, ok := st.huiyuanMap[huiyuanTemp.GetHoutaiType()]
		if !ok {
			subMap = make(map[huiyuantypes.HuiYuanType]*gametemplate.HuiYuanTemplate)
		}
		subMap[huiyuanTemp.GetHuiYuanType()] = huiyuanTemp
		st.huiyuanMap[huiyuanTemp.GetHoutaiType()] = subMap
	}

	return
}

func (st *huiyuanTemplateService) GetHuiYuanTemplate(houtaiType centertypes.ZhiZunType, typ huiyuantypes.HuiYuanType) *gametemplate.HuiYuanTemplate {
	//暂定豪华版本
	subMap, ok := st.huiyuanMap[houtaiType]
	if !ok {
		return nil
	}
	temp, ok := subMap[typ]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	st   *huiyuanTemplateService
)

func Init() (err error) {
	once.Do(func() {
		st = &huiyuanTemplateService{}
		err = st.init()
	})

	return
}

func GetHuiYuanTemplateService() HuiYuanTemplateService {
	return st
}
