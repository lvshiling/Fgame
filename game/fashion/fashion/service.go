package fashion

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//时装接口处理
type FashionService interface {
	//获取时装配置
	GetFashionTemplate(id int) *gametemplate.FashionTemplate
	RandomFashionTemplate() *gametemplate.FashionTemplate
}

type fashionService struct {
	//时装配置
	fashionMap map[int]*gametemplate.FashionTemplate
	//时装列表
	fashionList []*gametemplate.FashionTemplate
}

//初始化
func (s *fashionService) init() error {
	s.fashionMap = make(map[int]*gametemplate.FashionTemplate)
	//时装
	templateMap := template.GetTemplateService().GetAll((*gametemplate.FashionTemplate)(nil))
	for _, templateObject := range templateMap {
		fashionTemplate, _ := templateObject.(*gametemplate.FashionTemplate)
		s.fashionMap[fashionTemplate.TemplateId()] = fashionTemplate
		s.fashionList = append(s.fashionList, fashionTemplate)
	}

	return nil
}

//获取时装配置
func (s *fashionService) GetFashionTemplate(id int) *gametemplate.FashionTemplate {
	to, ok := s.fashionMap[id]
	if !ok {
		return nil
	}
	return to
}

//随机
func (s *fashionService) RandomFashionTemplate() *gametemplate.FashionTemplate {
	num := len(s.fashionList)
	index := rand.Intn(num)
	return s.fashionList[index]
}

var (
	once sync.Once
	cs   *fashionService
)

func Init() (err error) {
	once.Do(func() {
		cs = &fashionService{}
		err = cs.init()
	})
	return err
}

func GetFashionService() FashionService {
	return cs
}
