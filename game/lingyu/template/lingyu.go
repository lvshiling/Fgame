package lingyu

import (
	"fgame/fgame/core/template"
	lingyutypes "fgame/fgame/game/lingyu/types"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

//领域接口处理
type LingyuTemplateService interface {
	//获取领域配置
	GetLingyu(id int) *gametemplate.LingyuTemplate
	//获取领域进阶配置
	GetLingyuByNumber(number int32) *gametemplate.LingyuTemplate
	//获取幻化配置
	GetLingyuHuanHuaTemplate(level int32) *gametemplate.LingyuHuanHuaTemplate
	//随机领域
	RandomLingYuTemplate() *gametemplate.LingyuTemplate
	//获取领域技能
	GetLingyuSkill(advanceId int32) int32
	//吃幻化丹升级
	GetLingYuEatHuanHuaTemplate(curLevel int32, num int32) (*gametemplate.LingyuHuanHuaTemplate, bool)
}

type lingyuTemplateService struct {
	//领域配置
	lingyuMap map[int]*gametemplate.LingyuTemplate
	//领域进阶配置
	lingyuNumberMap map[int32]*gametemplate.LingyuTemplate
	//领域幻化配置
	huanHuaMap map[int32]*gametemplate.LingyuHuanHuaTemplate

	lingYuList []*gametemplate.LingyuTemplate
}

//初始化
func (s *lingyuTemplateService) init() error {
	s.lingyuMap = make(map[int]*gametemplate.LingyuTemplate)
	s.lingyuNumberMap = make(map[int32]*gametemplate.LingyuTemplate)
	s.huanHuaMap = make(map[int32]*gametemplate.LingyuHuanHuaTemplate)

	//领域
	lingyuTempMap := template.GetTemplateService().GetAll((*gametemplate.LingyuTemplate)(nil))
	for _, temp := range lingyuTempMap {
		lingyuTemplate, _ := temp.(*gametemplate.LingyuTemplate)
		s.lingyuMap[lingyuTemplate.TemplateId()] = lingyuTemplate

		typ := lingyuTemplate.GetTyp()
		if typ == lingyutypes.LingyuTypeAdvanced {
			s.lingyuNumberMap[lingyuTemplate.Number] = lingyuTemplate
		}
		s.lingYuList = append(s.lingYuList, lingyuTemplate)
	}

	//身法幻化
	huanHuaTempMap := template.GetTemplateService().GetAll((*gametemplate.LingyuHuanHuaTemplate)(nil))
	for _, temp := range huanHuaTempMap {
		huanHuaTemplate, _ := temp.(*gametemplate.LingyuHuanHuaTemplate)
		s.huanHuaMap[huanHuaTemplate.Level] = huanHuaTemplate

	}
	return nil
}

//获取领域配置
func (s *lingyuTemplateService) GetLingyu(id int) *gametemplate.LingyuTemplate {
	to, ok := s.lingyuMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取身法进阶配置
func (s *lingyuTemplateService) GetLingyuByNumber(number int32) *gametemplate.LingyuTemplate {
	to, ok := s.lingyuNumberMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取幻化配置
func (s *lingyuTemplateService) GetLingyuHuanHuaTemplate(level int32) *gametemplate.LingyuHuanHuaTemplate {
	to, ok := s.huanHuaMap[level]
	if !ok {
		return nil
	}
	return to
}

// 获取领域技能
func (s *lingyuTemplateService) GetLingyuSkill(advanceId int32) (skillId int32) {

	temp, ok := s.lingyuNumberMap[advanceId]
	if !ok {
		return
	}
	skillId = temp.Skill
	return
}

//吃幻化丹升级
func (s *lingyuTemplateService) GetLingYuEatHuanHuaTemplate(curLevel int32, num int32) (lingYuHuanHuaTemplate *gametemplate.LingyuHuanHuaTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		lingYuHuanHuaTemplate, flag = s.huanHuaMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= lingYuHuanHuaTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

//随机领域配置
func (s *lingyuTemplateService) RandomLingYuTemplate() *gametemplate.LingyuTemplate {
	num := len(s.lingYuList)
	index := rand.Intn(num)
	return s.lingYuList[index]
}

var (
	once sync.Once
	cs   *lingyuTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &lingyuTemplateService{}
		err = cs.init()
	})
	return err
}

func GetLingyuTemplateService() LingyuTemplateService {
	return cs
}
