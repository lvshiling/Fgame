package qixue

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

//泣血枪接口处理
type QiXueTemplateService interface {
	//泣血枪激活兵魂等级
	GetQiXeuActivateWeaponLev() int32
	//泣血枪对应等级星数
	GetQiXueStarCount(lev int32) int32
	//获取泣血枪配置
	GetQiXueTemplate(id int32) *gametemplate.QiXueTemplate
	GetQiXueTemplateByLevel(lev int32, star int32) *gametemplate.QiXueTemplate
	// 泣血枪常量配置
	GetQiXueConstantTemplate() *gametemplate.QiXueConstantTemplate
}

type qiXueTemplateService struct {
	//泣血枪激活兵魂等级
	weaponActivateLev int32
	//泣血枪配置
	qiXueMap          map[int]*gametemplate.QiXueTemplate
	qiXueLevMap       map[int32]map[int32]*gametemplate.QiXueTemplate
	qiXueStarMap      map[int32]int32
	qiXueConstantTemp *gametemplate.QiXueConstantTemplate
}

//初始化
func (s *qiXueTemplateService) init() error {
	s.qiXueMap = make(map[int]*gametemplate.QiXueTemplate)
	s.qiXueLevMap = make(map[int32]map[int32]*gametemplate.QiXueTemplate)
	s.qiXueStarMap = make(map[int32]int32)
	//泣血枪
	templateMap := template.GetTemplateService().GetAll((*gametemplate.QiXueTemplate)(nil))
	for _, templateObject := range templateMap {
		qiXueTemplate := templateObject.(*gametemplate.QiXueTemplate)
		s.qiXueMap[qiXueTemplate.TemplateId()] = qiXueTemplate

		subMap, ok := s.qiXueLevMap[qiXueTemplate.Level]
		if !ok {
			subMap = make(map[int32]*gametemplate.QiXueTemplate)
			s.qiXueLevMap[qiXueTemplate.Level] = subMap
		}
		subMap[qiXueTemplate.Star] = qiXueTemplate

		// 兵魂激活等级
		if qiXueTemplate.WeaponId != 0 {
			if s.weaponActivateLev == 0 {
				s.weaponActivateLev = qiXueTemplate.Level
			}

			if qiXueTemplate.Level < s.weaponActivateLev {
				s.weaponActivateLev = qiXueTemplate.Level
			}
		}

		//泣血枪等级星数
		star, ok := s.qiXueStarMap[qiXueTemplate.Level]
		if !ok || star < qiXueTemplate.Star {
			s.qiXueStarMap[qiXueTemplate.Level] = qiXueTemplate.Star
		}
	}

	//泣血常量配置
	constantMap := template.GetTemplateService().GetAll((*gametemplate.QiXueConstantTemplate)(nil))
	if len(constantMap) != 1 {
		return fmt.Errorf("泣血枪常量配置应该有且只有一条")
	}

	for _, to := range constantMap {
		temp := to.(*gametemplate.QiXueConstantTemplate)
		s.qiXueConstantTemp = temp
		break
	}

	return nil
}

//获取泣血枪配置
func (s *qiXueTemplateService) GetQiXeuActivateWeaponLev() int32 {
	return s.weaponActivateLev
}

//获取泣血枪等级星数
func (s *qiXueTemplateService) GetQiXueStarCount(lev int32) int32 {
	maxStar, ok := s.qiXueStarMap[lev]
	if !ok {
		return 0
	}

	return maxStar
}

//获取泣血枪配置
func (s *qiXueTemplateService) GetQiXueTemplate(id int32) *gametemplate.QiXueTemplate {
	to, ok := s.qiXueMap[int(id)]
	if !ok {
		return nil
	}
	return to
}

//获取泣血枪配置
func (s *qiXueTemplateService) GetQiXueTemplateByLevel(lev int32, star int32) *gametemplate.QiXueTemplate {
	subMap, ok := s.qiXueLevMap[lev]
	if !ok {
		return nil
	}
	temp, ok := subMap[star]
	if !ok {
		return nil
	}
	return temp
}

//获取泣血枪配置
func (s *qiXueTemplateService) GetQiXueConstantTemplate() *gametemplate.QiXueConstantTemplate {
	return s.qiXueConstantTemp
}

var (
	once sync.Once
	cs   *qiXueTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &qiXueTemplateService{}
		err = cs.init()
	})
	return err
}

func GetQiXueTemplateService() QiXueTemplateService {
	return cs
}
