package gem

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gemtypes "fgame/fgame/game/gem/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	"fmt"
	"sort"
	"sync"
)

//宝石接口处理
type GemService interface {
	//获取矿工模板通过等级
	GetMineTemplateByLevel(level int32) *gametemplate.MiningTemplate
	//获取赌石模板通过赌石类型
	GetGambleTemplateByTyp(typ gemtypes.GemGambleType) *gametemplate.GamblingTemplate
	//获取掉落id通过赌石次数
	GetDropIdByNum(pl player.Player, typ gemtypes.GemGambleType, num int32) (int32, bool)
	//功能开启时直接给予玩家的原石数量
	GetFuncOpenGiveStone() (giveStone int32)
}

type gemService struct {
	//矿工等级配置
	mineLevelMap map[int32]*gametemplate.MiningTemplate
	//赌石map
	gambleTypMap map[gemtypes.GemGambleType]*gametemplate.GamblingTemplate
	//赌石掉落id
	dropIdMap map[gemtypes.GemGambleType]map[int32]int32
	//赌石掉落id
	dropIdIntervalListMap map[gemtypes.GemGambleType][]int
}

//初始化
func (gs *gemService) init() error {
	gs.mineLevelMap = make(map[int32]*gametemplate.MiningTemplate)
	gs.gambleTypMap = make(map[gemtypes.GemGambleType]*gametemplate.GamblingTemplate)
	gs.dropIdMap = make(map[gemtypes.GemGambleType]map[int32]int32)
	gs.dropIdIntervalListMap = make(map[gemtypes.GemGambleType][]int)

	//赋值mineLevelMap
	templateMap := template.GetTemplateService().GetAll((*gametemplate.MiningTemplate)(nil))
	for _, templateObject := range templateMap {
		mineTemplate, _ := templateObject.(*gametemplate.MiningTemplate)

		_, ok := gs.mineLevelMap[mineTemplate.Level]
		if ok {
			return fmt.Errorf("gemService: level no should repeat")
		}
		gs.mineLevelMap[mineTemplate.Level] = mineTemplate
	}

	//赋值gambleTypMap
	templateMap = template.GetTemplateService().GetAll((*gametemplate.GamblingTemplate)(nil))
	for _, templateObject := range templateMap {
		gambleTemplate, _ := templateObject.(*gametemplate.GamblingTemplate)
		typ := gemtypes.GemGambleType(gambleTemplate.Type)
		intervalDropMap := gambleTemplate.GetIntervalDropMap()

		dropIdListMap, ok := gs.dropIdMap[typ]
		if !ok {
			dropIdListMap = make(map[int32]int32)
			gs.dropIdMap[typ] = dropIdListMap
		}
		for interval, dropId := range intervalDropMap {
			dropIdListMap[interval] = dropId
		}

		dropIntervalList := gs.dropIdIntervalListMap[typ]
		for interval, _ := range intervalDropMap {
			dropIntervalList = append(dropIntervalList, int(interval))
		}
		//降序排序dropIntervalList
		sort.Sort(sort.Reverse(sort.IntSlice(dropIntervalList)))
		gs.dropIdIntervalListMap[typ] = dropIntervalList

		_, ok = gs.gambleTypMap[typ]
		if ok {
			return fmt.Errorf("gemService: type no should repeat")
		}
		gs.gambleTypMap[typ] = gambleTemplate
	}

	//验证 interval_num 是否有配置1
	for _, dropIdList := range gs.dropIdMap {
		_, ok := dropIdList[1]
		if !ok {
			return fmt.Errorf("gemService: interval_num  should have 1")
		}
	}

	return nil
}

//获取矿工模板通过等级
func (gs *gemService) GetMineTemplateByLevel(level int32) *gametemplate.MiningTemplate {
	to, ok := gs.mineLevelMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取赌石模板通过赌石类型
func (gs *gemService) GetGambleTemplateByTyp(typ gemtypes.GemGambleType) *gametemplate.GamblingTemplate {
	to, ok := gs.gambleTypMap[typ]
	if !ok {
		return nil
	}
	return to
}

//获取掉落id通过赌石次数
func (gs *gemService) GetDropIdByNum(pl player.Player, typ gemtypes.GemGambleType, num int32) (int32, bool) {
	if num <= 0 {
		return 0, false
	}
	//首次赌石特殊掉落包
	if num == 1 {
		gambleTemplate, ok := gs.gambleTypMap[typ]
		if !ok {
			return 0, false
		}
		return gambleTemplate.FirstDrop, true
	}
	dropIdIntervalList, ok := gs.dropIdIntervalListMap[typ]
	if !ok {
		return 0, false
	}

	dropIdListMap, ok := gs.dropIdMap[typ]
	if !ok {
		return 0, false
	}

	ruleTimesMap := viplogic.CountDropTimesWithCostLevel(pl, viptypes.CostLevelRuleTypeGamble, dropIdIntervalList)
	for _, value := range dropIdIntervalList {
		ruleTimes := ruleTimesMap[value]
		ret := num % int32(ruleTimes)
		if ret == 0 {
			return dropIdListMap[int32(value)], true
		}
	}
	return 0, false
}

//功能开启时直接给予玩家的原石数量
func (gs *gemService) GetFuncOpenGiveStone() (giveStone int32) {
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGemStoneGive)
}

var (
	once sync.Once
	cs   *gemService
)

func Init() (err error) {
	once.Do(func() {
		cs = &gemService{}
		err = cs.init()
	})
	return err
}

func GetGemService() GemService {
	return cs
}
