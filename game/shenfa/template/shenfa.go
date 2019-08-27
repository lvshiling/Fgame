package shenfa

import (
	"fgame/fgame/core/template"
	shenfatypes "fgame/fgame/game/shenfa/types"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//身法接口处理
type ShenfaTemplateService interface {
	//获取身法配置
	GetShenfa(id int) *gametemplate.ShenfaTemplate
	//获取身法进阶配置
	GetShenfaByNumber(number int32) *gametemplate.ShenfaTemplate
	//获取幻化配置
	GetShenfaHuanHuaTemplate(level int32) *gametemplate.ShenfaHuanHuaTemplate
	//随机身法
	RandomShenFaTemplate() *gametemplate.ShenfaTemplate
	//获取身法技能
	GetShenfaSkill(advanceId int32) int32
	//吃幻化丹升级
	GetShenFaEatHuanHuaTemplate(curLevel int32, num int32) (*gametemplate.ShenfaHuanHuaTemplate, bool)
}

type shenfaTemplateService struct {
	//身法配置
	shenfaMap map[int]*gametemplate.ShenfaTemplate
	//身法进阶配置
	shenfaNumberMap map[int32]*gametemplate.ShenfaTemplate
	//身法幻化配置
	huanHuaMap map[int32]*gametemplate.ShenfaHuanHuaTemplate

	shenFaList []*gametemplate.ShenfaTemplate
}

//初始化
func (s *shenfaTemplateService) init() error {
	s.shenfaMap = make(map[int]*gametemplate.ShenfaTemplate)
	s.shenfaNumberMap = make(map[int32]*gametemplate.ShenfaTemplate)
	s.huanHuaMap = make(map[int32]*gametemplate.ShenfaHuanHuaTemplate)

	//身法
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShenfaTemplate)(nil))
	for _, temp := range templateMap {
		shenfaTemplate, _ := temp.(*gametemplate.ShenfaTemplate)
		s.shenfaMap[shenfaTemplate.TemplateId()] = shenfaTemplate

		typ := shenfaTemplate.GetTyp()
		if typ == shenfatypes.ShenfaTypeAdvanced {
			s.shenfaNumberMap[shenfaTemplate.Number] = shenfaTemplate
		}
		s.shenFaList = append(s.shenFaList, shenfaTemplate)
	}

	//身法幻化
	huanHuaTempMap := template.GetTemplateService().GetAll((*gametemplate.ShenfaHuanHuaTemplate)(nil))
	for _, temp := range huanHuaTempMap {
		huanHuaTemplate, _ := temp.(*gametemplate.ShenfaHuanHuaTemplate)
		s.huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate
	}

	return nil
}

//获取身法配置
func (s *shenfaTemplateService) GetShenfa(id int) *gametemplate.ShenfaTemplate {
	to, ok := s.shenfaMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取身法进阶配置
func (s *shenfaTemplateService) GetShenfaByNumber(number int32) *gametemplate.ShenfaTemplate {
	to, ok := s.shenfaNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取幻化配置
func (s *shenfaTemplateService) GetShenfaHuanHuaTemplate(level int32) *gametemplate.ShenfaHuanHuaTemplate {
	to, ok := s.huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

// 获取领域技能
func (s *shenfaTemplateService) GetShenfaSkill(advanceId int32) (skillId int32) {

	temp, ok := s.shenfaNumberMap[advanceId]
	if !ok {
		return
	}
	skillId = temp.Skill
	return
}

//随机身法配置
func (s *shenfaTemplateService) RandomShenFaTemplate() *gametemplate.ShenfaTemplate {
	num := len(s.shenFaList)
	index := rand.Intn(num)
	return s.shenFaList[index]
}

//吃幻化丹升级
func (s *shenfaTemplateService) GetShenFaEatHuanHuaTemplate(curLevel int32, num int32) (shenFaHuanHuaTemplate *gametemplate.ShenfaHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		shenFaHuanHuaTemplate, flag = s.huanHuaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= shenFaHuanHuaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *shenfaTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shenfaTemplateService{}
		err = cs.init()
	})
	return err
} 

func GetShenfaTemplateService() ShenfaTemplateService {
	return cs
}
