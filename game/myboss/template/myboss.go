package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//个人BOSS配置服务
type MyBossTemplaterService interface {
	// 个人BOSS配置
	GetMyBossTemplate(biologyId int32) *gametemplate.MyBossTemplate
}

type mybossTemplaterService struct {
	bossMap map[int32]*gametemplate.MyBossTemplate
}

//初始化
func (ts *mybossTemplaterService) init() error {
	ts.bossMap = make(map[int32]*gametemplate.MyBossTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.MyBossTemplate)(nil))
	for _, temp := range templateMap {
		mybossTemplate, _ := temp.(*gametemplate.MyBossTemplate)
		ts.bossMap[mybossTemplate.BiologyId] = mybossTemplate
	}

	return nil
}

func (ts *mybossTemplaterService) GetMyBossTemplate(biologyId int32) *gametemplate.MyBossTemplate {
	temp, ok := ts.bossMap[biologyId]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	cs   *mybossTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &mybossTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetMyBossTemplateService() MyBossTemplaterService {
	return cs
}
