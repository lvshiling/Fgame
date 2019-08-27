package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"math"
	"sync"
)

//抢龙椅处理
type EmperorTemplateService interface {
	//抢龙椅配置
	GetEmperorTemplate() *gametemplate.DragonChairTemplate
	//获取膜拜奖励
	GetEmperorWorshipRew() (rewData *propertytypes.RewData)
	//获取膜拜次数
	GetEmperorWorshipNum() (worshipNum int32)
	//获取帝王称号id
	GetEmperorTitleId() int32
	//获取国库最大上限
	GetEmperorChestMax() (chestMax int64)
	//获取膜拜增加国库
	GetEmperorWorshipChestSilver() (chestSilver int64)
	//获取抢夺龙椅需要的元宝
	GetEmperorRobNeedGold(robNum int64) int64
	//属性加成权重
	GetEmperorRobCoefficientAttr(robNum int64) float64
	//获取定时时间
	GetEmperorAutoSilverTime() (autoSilverTime int32)
	//获取自动增加银两
	GetEmperorAutoSilver() (autoSilver int64)
	//获取增加银两
	GetEmperorChanChuSilver(robElapseTime int64, elapseTime int64) (silver int64, offTime int64)
	//是否增加帝王宝箱
	IfAddEmperorBox(elapseTime int64, robElapseTime int64, boxOutNum int64) (addNum int32, offTime int64, isAdd bool)
}

type emperorTemplateService struct {
	emperorTemplate *gametemplate.DragonChairTemplate
}

//初始化
func (ets *emperorTemplateService) init() (err error) {
	tempTemplate := template.GetTemplateService().Get(1, (*gametemplate.DragonChairTemplate)(nil))
	ets.emperorTemplate, _ = tempTemplate.(*gametemplate.DragonChairTemplate)
	if ets.emperorTemplate == nil {
		return fmt.Errorf("emperor:emperorTemplate 不应该为空")
	}
	return
}

//抢龙椅配置
func (ets *emperorTemplateService) GetEmperorTemplate() *gametemplate.DragonChairTemplate {
	return ets.emperorTemplate
}

//获取膜拜奖励
func (ets *emperorTemplateService) GetEmperorWorshipRew() (rewData *propertytypes.RewData) {
	rewData = ets.emperorTemplate.GetWorshipRewData()
	return
}

//获取膜拜次数
func (ets *emperorTemplateService) GetEmperorWorshipNum() (worshipNum int32) {
	worshipNum = ets.emperorTemplate.WorshipCount
	return
}

//获取帝王称号id
func (ets *emperorTemplateService) GetEmperorTitleId() (titleId int32) {
	titleId = ets.emperorTemplate.TitleId
	return
}

//获取国库最大上限
func (ets *emperorTemplateService) GetEmperorChestMax() (chestMax int64) {
	chestMax = int64(ets.emperorTemplate.ChestMax)
	return
}

//获取膜拜增加国库
func (ets *emperorTemplateService) GetEmperorWorshipChestSilver() (chestSilver int64) {
	chestSilver = int64(ets.emperorTemplate.WorshipChestSilver)
	return
}

//获取抢夺龙椅需要的元宝
func (ets *emperorTemplateService) GetEmperorRobNeedGold(robNum int64) (needGold int64) {
	valueGold := float64(ets.emperorTemplate.ValueGold)
	firstGold := float64(ets.emperorTemplate.FirstGold)
	coefficient := ets.emperorTemplate.CoefficientGold
	value := math.Pow(float64(robNum), coefficient)
	needGold = int64(math.Ceil((firstGold*value+valueGold)/10)) * 10
	return
}

//属性加成系数
func (ets *emperorTemplateService) GetEmperorRobCoefficientAttr(robNum int64) (weight float64) {
	coefficient := ets.emperorTemplate.CoefficientAttr
	weight = math.Pow(float64(robNum), coefficient)
	return
}

//获取定时时间
func (ets *emperorTemplateService) GetEmperorAutoSilverTime() (autoSilverTime int32) {
	autoSilverTime = ets.emperorTemplate.AutoSilverTime
	return
}

//获取自动增加银两
func (ets *emperorTemplateService) GetEmperorAutoSilver() (autoSilver int64) {
	autoSilver = int64(ets.emperorTemplate.AutoSilver)
	return
}

//获取增加银两
func (ets *emperorTemplateService) GetEmperorChanChuSilver(robElaspeTime int64, elaspeTime int64) (silver int64, offTime int64) {
	silverList := ets.emperorTemplate.GetSilverList()
	flag := false
	addNum := int64(0)
	addNum = elaspeTime / int64(common.MINUTE)
	offTime = addNum * int64(common.MINUTE)
	for _, silverObj := range silverList {
		min := silverObj.GetMin()
		max := silverObj.GetMax()
		if robElaspeTime >= min && robElaspeTime < max {
			silver = silverObj.GetSilver() * addNum
			return
		}
	}
	if !flag {
		silver = int64(ets.emperorTemplate.ChanchuSilver0) * addNum
		return
	}
	return
}

func (ets *emperorTemplateService) IfAddEmperorBox(elapseTime int64, robElapseTime int64, boxOutNum int64) (addNum int32, offTime int64, isAdd bool) {
	dropTime1 := int64(ets.emperorTemplate.DropTime1)
	dropTime2 := int64(ets.emperorTemplate.DropTime2)
	dropTime3 := int64(ets.emperorTemplate.DropTime3)
	dropTime4 := int64(ets.emperorTemplate.DropTime4)
	addNum = 0
	offTime = 0

	if robElapseTime < dropTime1 {
		return
	} else if robElapseTime >= dropTime1 && robElapseTime < dropTime2 {
		if boxOutNum == 1 {
			return
		}
		addNum = 1
		offTime = dropTime1
		isAdd = true
		return
	} else if robElapseTime >= dropTime2 && robElapseTime < dropTime3 {
		if boxOutNum == 2 {
			return
		}
		if boxOutNum == 1 {
			addNum += 1
			offTime += (dropTime2 - dropTime1)
			isAdd = true
			return
		}
		if boxOutNum == 0 {
			addNum += 2
			offTime += dropTime2
			isAdd = true
			return
		}
		return
	} else if robElapseTime >= dropTime3 && robElapseTime < dropTime4 {
		if boxOutNum == 3 {
			return
		}
		if boxOutNum == 2 {
			addNum += 1
			offTime += (dropTime3 - dropTime2)
			isAdd = true
			return
		}
		if boxOutNum == 1 {
			addNum += 2
			offTime += (dropTime3 - dropTime1)
			isAdd = true
			return
		}
		if boxOutNum == 0 {
			addNum += 3
			offTime += dropTime3
			isAdd = true
			return
		}
		return
	} else if elapseTime >= dropTime4 {
		if boxOutNum == 0 {
			addNum += 3
			offTime += dropTime3
			elapseTime -= offTime
			isAdd = true
		}
		if boxOutNum == 1 {
			addNum += 2
			offTime += (dropTime3 - dropTime1)
			elapseTime -= offTime
			isAdd = true
		}
		if boxOutNum == 2 {
			addNum += 1
			offTime += (dropTime3 - dropTime2)
			elapseTime -= offTime
			isAdd = true
		}
		if elapseTime >= dropTime4 {
			needNum := int32(elapseTime / dropTime4)
			addNum += needNum
			offTime += int64(needNum) * dropTime4
			isAdd = true
			return
		}
		return
	}
	return
}

var (
	once sync.Once
	cs   *emperorTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &emperorTemplateService{}
		err = cs.init()
	})
	return err
}

func GetEmperorTemplateService() EmperorTemplateService {
	return cs
}
