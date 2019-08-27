package shop

import (
	"fgame/fgame/core/template"
	shoptypes "fgame/fgame/game/shop/types"
	gametemplate "fgame/fgame/game/template"
	"sort"
	"sync"
)

//规则排序
type ShopTemplateList []*gametemplate.ShopTemplate

func (stl ShopTemplateList) Len() int {
	return len(stl)
}

func (stl ShopTemplateList) Less(i, j int) bool {
	if stl[i].Order < stl[j].Order {
		return true
	}
	return stl[i].Id < stl[j].Id

}

func (stl ShopTemplateList) Swap(i, j int) {
	stl[i], stl[j] = stl[j], stl[i]
}

type ShopService interface {
	GetShopTemplate(id int) *gametemplate.ShopTemplate
	// GetShopCost(items map[int32]int32) (currencyCost map[shoptypes.ShopConsumeType]int32, flag bool)
	// GetShopTemplateByItem(itemId int32) *gametemplate.ShopTemplate
	//商店是否出售物品
	ShopIsSellItem(itemId int32) (flag bool)
	//物品shopMap
	GetShopItemMap(itemId int32) (shopItemMap map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate)
}

type shopService struct {
	//只给自动进阶使用
	//shopTemplateMap map[int32]*gametemplate.ShopTemplate
	shopItemMap map[int32]map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate
}

func (ss *shopService) init() error {
	//ss.shopTemplateMap = make(map[int32]*gametemplate.ShopTemplate)
	ss.shopItemMap = make(map[int32]map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate)
	shopTemplateList := template.GetTemplateService().GetAll((*gametemplate.ShopTemplate)(nil))

	for _, to := range shopTemplateList {
		shopTemplate := to.(*gametemplate.ShopTemplate)
		itemId := shopTemplate.ItemId
		consumeType := shopTemplate.GetShopConsumeType()

		shopItemTypeMap, exist := ss.shopItemMap[itemId]
		if !exist {
			shopItemTypeMap = make(map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate)
			ss.shopItemMap[itemId] = shopItemTypeMap
		}
		shopItemTypeMap[consumeType] = append(shopItemTypeMap[consumeType], shopTemplate)

		// if consumeType == shoptypes.ShopConsumeTypeSliver {
		// 	continue
		// }
		// if shopTemplate.BuyCount != 1 {
		// 	continue
		// }
		// if old, ok := ss.shopTemplateMap[itemId]; ok {
		// 	flag := shopTemplate.Priority(old)
		// 	if !flag {
		// 		continue
		// 	}
		// }
		// ss.shopTemplateMap[itemId] = shopTemplate
	}

	//自动购买配银两的
	// for _, to := range shopTemplateList {
	// 	shopTemplate := to.(*gametemplate.ShopTemplate)
	// 	itemId := shopTemplate.ItemId
	// 	consumeType := shopTemplate.GetShopConsumeType()
	// 	_, ok := ss.shopTemplateMap[itemId]
	// 	if ok {
	// 		continue
	// 	}

	// 	if shopTemplate.BuyCount != 1 {
	// 		continue
	// 	}
	// 	if consumeType != shoptypes.ShopConsumeTypeSliver {
	// 		continue
	// 	}
	// 	ss.shopTemplateMap[itemId] = shopTemplate
	// }

	//排序处理
	for _, shopItemTypeMap := range ss.shopItemMap {
		for shopConsumeType, shopIdList := range shopItemTypeMap {
			sort.Sort(sort.Reverse(ShopTemplateList(shopIdList)))
			shopItemTypeMap[shopConsumeType] = shopIdList
		}
	}

	return nil
}

func (ss *shopService) GetShopTemplate(id int) *gametemplate.ShopTemplate {
	//获取商店模板
	tempShopTemplate := template.GetTemplateService().Get(id, (*gametemplate.ShopTemplate)(nil))
	if tempShopTemplate == nil {
		return nil
	}

	shopTemplate, ok := tempShopTemplate.(*gametemplate.ShopTemplate)
	if !ok {
		return nil
	}
	return shopTemplate
}

//获取物品消耗货币
// func (ss *shopService) GetShopCost(items map[int32]int32) (currencyCost map[shoptypes.ShopConsumeType]int32, flag bool) {
// 	if len(items) <= 0 {
// 		panic(fmt.Errorf("shopService: items no should be nil"))
// 	}
// 	currencyCost = make(map[shoptypes.ShopConsumeType]int32)
// 	for itemId, needBuyNum := range items {
// 		if needBuyNum <= 0 {
// 			panic(fmt.Errorf("shopService:item num should be more than 0"))
// 		}

// 		shopTemplate := ss.GetShopTemplateByItem(itemId)
// 		if shopTemplate == nil {
// 			return nil, false
// 		}
// 		consumeType := shoptypes.ShopConsumeType(shopTemplate.ConsumeType)
// 		if v, ok := currencyCost[consumeType]; ok {
// 			currencyCost[consumeType] = v + shopTemplate.ConsumeData1*needBuyNum
// 			continue
// 		}
// 		currencyCost[consumeType] = shopTemplate.ConsumeData1 * needBuyNum
// 	}
// 	return currencyCost, true
// }

//给自动进阶使用
// func (ss *shopService) GetShopTemplateByItem(itemId int32) *gametemplate.ShopTemplate {
// 	if v, ok := ss.shopTemplateMap[itemId]; ok {
// 		return v
// 	}
// 	return nil
// }

func (ss *shopService) ShopIsSellItem(itemId int32) (flag bool) {
	_, flag = ss.shopItemMap[itemId]
	return
}

func (ss *shopService) GetShopItemMap(itemId int32) (shopItemMap map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate) {
	return ss.shopItemMap[itemId]
}

var (
	once sync.Once
	cs   *shopService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shopService{}
		err = cs.init()
	})
	return err
}

func GetShopService() ShopService {
	return cs
}
