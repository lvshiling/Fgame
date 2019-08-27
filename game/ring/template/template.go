package template

import (
	"fgame/fgame/core/template"
	ringtypes "fgame/fgame/game/ring/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type RingTemplateService interface {
	// 获取特戒模板
	GetRingTemplate(itemId int32) *gametemplate.RingTemplate
	// 获取进阶模板
	GetRingAdvanceTemplate(itemId int32, advance int32) *gametemplate.RingAdvanceTemplate
	// 获取强化模板
	GetRingStrengthenTemplate(itemId int32, level int32) *gametemplate.RingStrengthenTemplate
	// 获取净灵模板
	GetRingJingLingTemplate(itemId int32, level int32) *gametemplate.RingJingLingTemplate
	// 获取融合合成模板
	GetRingFuseSynthesisTemplate(itemId int32) *gametemplate.RingFuseSynthesisTemplate
	//获取装备宝库配置
	GetRingBaoKuTemplate(typ ringtypes.BaoKuType) *gametemplate.RingBaoKuTemplate
	//获取进阶最低阶别
	GetRingMinAdvance(itemId int32) int32
}

type ringTemplateService struct {
	// 特戒
	ringTemplateMap map[int32]*gametemplate.RingTemplate
	// 特戒宝库配置
	ringBaoKuMap map[ringtypes.BaoKuType]*gametemplate.RingBaoKuTemplate
}

func (s *ringTemplateService) init() (err error) {
	s.ringTemplateMap = make(map[int32]*gametemplate.RingTemplate)
	ringTempMap := template.GetTemplateService().GetAll((*gametemplate.RingTemplate)(nil))
	for _, temp := range ringTempMap {
		ringTemp, _ := temp.(*gametemplate.RingTemplate)
		s.ringTemplateMap[int32(ringTemp.Id)] = ringTemp
	}

	s.ringBaoKuMap = make(map[ringtypes.BaoKuType]*gametemplate.RingBaoKuTemplate)
	baoKuTempMap := template.GetTemplateService().GetAll((*gametemplate.RingBaoKuTemplate)(nil))
	for _, temp := range baoKuTempMap {
		ringTemp, _ := temp.(*gametemplate.RingBaoKuTemplate)
		s.ringBaoKuMap[ringTemp.GetBaoKuType()] = ringTemp
	}
	return
}

func (s *ringTemplateService) GetRingTemplate(itemId int32) *gametemplate.RingTemplate {
	ringTemp, ok := s.ringTemplateMap[itemId]
	if !ok {
		return nil
	}
	return ringTemp
}

func (s *ringTemplateService) GetRingAdvanceTemplate(itemId int32, advance int32) *gametemplate.RingAdvanceTemplate {
	ringTemp, ok := s.ringTemplateMap[itemId]
	if !ok {
		return nil
	}
	return ringTemp.GetAdvanceTemplate(advance)
}

func (s *ringTemplateService) GetRingStrengthenTemplate(itemId int32, level int32) *gametemplate.RingStrengthenTemplate {
	ringTemp, ok := s.ringTemplateMap[itemId]
	if !ok {
		return nil
	}
	return ringTemp.GetStrengthenTemplate(level)
}

func (s *ringTemplateService) GetRingJingLingTemplate(itemId int32, level int32) *gametemplate.RingJingLingTemplate {
	ringTemp, ok := s.ringTemplateMap[itemId]
	if !ok {
		return nil
	}
	return ringTemp.GetJingLingTemplate(level)
}

func (s *ringTemplateService) GetRingFuseSynthesisTemplate(itemId int32) *gametemplate.RingFuseSynthesisTemplate {
	temp := s.GetRingTemplate(itemId)
	if temp == nil {
		return nil
	}
	return temp.GetFuseSynthesisTemplate()
}

func (s *ringTemplateService) GetRingBaoKuTemplate(typ ringtypes.BaoKuType) *gametemplate.RingBaoKuTemplate {
	ringTemp, ok := s.ringBaoKuMap[typ]
	if !ok {
		return nil
	}
	return ringTemp
}

func (s *ringTemplateService) GetRingMinAdvance(itemId int32) (minAdvance int32) {
	ringTemp, ok := s.ringTemplateMap[itemId]
	if !ok {
		return
	}
	minAdvance = 100
	for _, temp := range ringTemp.GetAdvanceTemplateMap() {
		if minAdvance > temp.Advance {
			minAdvance = temp.Advance
		}
	}
	return
}

var (
	once sync.Once
	ring *ringTemplateService
)

func Init() (err error) {
	once.Do(func() {
		ring = &ringTemplateService{}
		err = ring.init()
	})
	return
}

func GetRingTemplateService() RingTemplateService {
	return ring
}
