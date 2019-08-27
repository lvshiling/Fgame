package template

import (
	"fgame/fgame/core/template"
	shoptypes "fgame/fgame/game/shop/types"
	shopdiscounttypes "fgame/fgame/game/shopdiscount/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//商城促销接口处理
type ShopDiscountTemplateService interface {
	//获取商城促销模板
	GetShopDiscountTemplateByType(typ shopdiscounttypes.ShopDiscountType, subType shoptypes.ShopConsumeType) *gametemplate.ShopDiscountTemplate
	GetShopDiscountTemplate(id int32) *gametemplate.ShopDiscountTemplate
}

type shopDiscountTemplateService struct {
	//商城促销模板
	shopDiscountTemplateMap map[int32]*gametemplate.ShopDiscountTemplate
	//商城促销配置
	shopDiscountTypeTemplateMap map[shopdiscounttypes.ShopDiscountType]map[shoptypes.ShopConsumeType]*gametemplate.ShopDiscountTemplate
}

//初始化
func (s *shopDiscountTemplateService) init() error {
	s.shopDiscountTemplateMap = make(map[int32]*gametemplate.ShopDiscountTemplate)
	s.shopDiscountTypeTemplateMap = make(map[shopdiscounttypes.ShopDiscountType]map[shoptypes.ShopConsumeType]*gametemplate.ShopDiscountTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShopDiscountTemplate)(nil))
	for _, templateObject := range templateMap {
		skTemplate, _ := templateObject.(*gametemplate.ShopDiscountTemplate)

		s.shopDiscountTemplateMap[int32(skTemplate.TemplateId())] = skTemplate

		typ := skTemplate.GetDiscountType()
		subType := skTemplate.GetShopType()

		skTypeTemplateMap, ok := s.shopDiscountTypeTemplateMap[typ]
		if !ok {
			skTypeTemplateMap = make(map[shoptypes.ShopConsumeType]*gametemplate.ShopDiscountTemplate)
			s.shopDiscountTypeTemplateMap[typ] = skTypeTemplateMap
		}
		skTypeTemplateMap[subType] = skTemplate
	}

	return nil
}

//获取商城促销配置
func (s *shopDiscountTemplateService) GetShopDiscountTemplateByType(typ shopdiscounttypes.ShopDiscountType, subType shoptypes.ShopConsumeType) *gametemplate.ShopDiscountTemplate {
	subMap, ok := s.shopDiscountTypeTemplateMap[typ]
	if !ok {
		return nil
	}
	to, ok := subMap[subType]
	if !ok {
		return nil
	}
	return to
}

func (s *shopDiscountTemplateService) GetShopDiscountTemplate(id int32) *gametemplate.ShopDiscountTemplate {
	to, ok := s.shopDiscountTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *shopDiscountTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shopDiscountTemplateService{}
		err = cs.init()
	})
	return err
}

func GetShopDiscountTemplateService() ShopDiscountTemplateService {
	return cs
}
