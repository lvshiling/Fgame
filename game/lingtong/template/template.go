package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//灵童接口处理
type LingTongTemplateService interface {
	//获取灵童配置
	GetLingTongTemplate(id int32) *gametemplate.LingTongTemplate
	RandomLingTongTemplate() *gametemplate.LingTongTemplate
	//获取灵童时装配置
	GetLingTongFashionTemplate(id int32) *gametemplate.LingTongFashionTemplate
	RandomLingTongFashionTemplate() *gametemplate.LingTongFashionTemplate
	//灵童培养升级
	GetLingTongPeiYangUpgrade(lingTongId int32, curLevel int32, num int32) (lingTongPeiYangTemplate *gametemplate.LingTongPeiYangTemplate, flag bool)
	//是否是出生时装
	IsBornFashion(fashionId int32) bool
	GetLingTongTemplateList() []*gametemplate.LingTongTemplate
}

type lingTongTemplateService struct {
	//灵童配置
	lingTongTemplateMap map[int32]*gametemplate.LingTongTemplate
	//灵童列表
	lingTongTemplateList []*gametemplate.LingTongTemplate
	//出生时装
	bornMap map[int32]struct{}

	//灵童时装
	lingTongFashionMap map[int32]*gametemplate.LingTongFashionTemplate
	//灵童时装列表
	lingTongFashionTemplateList []*gametemplate.LingTongFashionTemplate
}

//初始化
func (s *lingTongTemplateService) init() error {
	s.lingTongTemplateMap = make(map[int32]*gametemplate.LingTongTemplate)
	s.lingTongFashionMap = make(map[int32]*gametemplate.LingTongFashionTemplate)
	s.bornMap = make(map[int32]struct{})
	//灵童
	templateMap := template.GetTemplateService().GetAll((*gametemplate.LingTongTemplate)(nil))
	for _, templateObject := range templateMap {
		lingTongTemplate, _ := templateObject.(*gametemplate.LingTongTemplate)
		s.lingTongTemplateMap[int32(lingTongTemplate.TemplateId())] = lingTongTemplate
		s.lingTongTemplateList = append(s.lingTongTemplateList, lingTongTemplate)
		s.bornMap[lingTongTemplate.LingTongFashionId] = struct{}{}
	}

	//灵童时装
	fashionMap := template.GetTemplateService().GetAll((*gametemplate.LingTongFashionTemplate)(nil))
	for _, templateObject := range fashionMap {
		fashionTemplate, _ := templateObject.(*gametemplate.LingTongFashionTemplate)
		s.lingTongFashionMap[int32(fashionTemplate.TemplateId())] = fashionTemplate
		s.lingTongFashionTemplateList = append(s.lingTongFashionTemplateList, fashionTemplate)
	}

	return nil
}

//获取灵童配置
func (s *lingTongTemplateService) GetLingTongTemplate(id int32) *gametemplate.LingTongTemplate {
	to, ok := s.lingTongTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

//随机
func (s *lingTongTemplateService) RandomLingTongTemplate() *gametemplate.LingTongTemplate {
	num := len(s.lingTongTemplateList)
	index := rand.Intn(num)
	return s.lingTongTemplateList[index]
}

//获取灵童时装配置
func (s *lingTongTemplateService) GetLingTongFashionTemplate(id int32) *gametemplate.LingTongFashionTemplate {
	to, ok := s.lingTongFashionMap[id]
	if !ok {
		return nil
	}
	return to
}

//随机
func (s *lingTongTemplateService) RandomLingTongFashionTemplate() *gametemplate.LingTongFashionTemplate {
	num := len(s.lingTongFashionTemplateList)
	index := rand.Intn(num)
	return s.lingTongFashionTemplateList[index]
}

func (s *lingTongTemplateService) IsBornFashion(fashionId int32) (flag bool) {
	_, ok := s.bornMap[fashionId]
	if !ok {
		return
	}
	flag = true
	return
}

//灵童养成培养
func (s *lingTongTemplateService) GetLingTongPeiYangUpgrade(lingTongId int32, curLevel int32, num int32) (lingTongPeiYangTemplate *gametemplate.LingTongPeiYangTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	lingTongTemplate, ok := s.lingTongTemplateMap[lingTongId]
	if !ok {
		return nil, false
	}
	for level := curLevel + 1; leftNum > 0; level++ {
		lingTongPeiYangTemplate = lingTongTemplate.GetLingTongPeiYangByLevel(level)
		if lingTongPeiYangTemplate == nil {
			return nil, false
		}
		leftNum -= lingTongPeiYangTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

func (s *lingTongTemplateService) GetLingTongTemplateList() []*gametemplate.LingTongTemplate {
	return s.lingTongTemplateList
}

var (
	once sync.Once
	cs   *lingTongTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &lingTongTemplateService{}
		err = cs.init()
	})
	return err
}

func GetLingTongTemplateService() LingTongTemplateService {
	return cs
}
