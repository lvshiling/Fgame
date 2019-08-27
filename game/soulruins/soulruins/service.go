package soulruins

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/soulruins/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

type chapterMaxInfo struct {
	chapter int32
	level   int32
}

func newChapterMaxInfo(chapter int32, level int32) *chapterMaxInfo {
	d := &chapterMaxInfo{
		chapter: chapter,
		level:   level,
	}
	return d
}

//帝陵遗迹接口处理
type SoulRuinsService interface {
	//获取帝陵遗迹配置通关id
	GetSoulRuinsTemplateById(id int32) *gametemplate.SoulRuinsTemplate
	//获取帝陵遗迹配置
	GetSoulRuinsTemplate(chapter int32, typ types.SoulRuinsType, level int32) *gametemplate.SoulRuinsTemplate
	//获取帝陵遗迹星级奖励配置
	GetSoulRuinsStarTemplate(chapter int32, typ types.SoulRuinsType) *gametemplate.SoulRuinsStarTemplate
	//获取帝陵遗迹触发特殊事件
	GetSoulRuinsTriggerSpecialEvent(chapter int32, typ types.SoulRuinsType, level int32) types.SoulRuinsEventType
	//获取挑战用时获取星数
	GetSoulRuinsStarByTime(chapter int32, typ types.SoulRuinsType, level int32, usedTime int32) int32
	//帝陵遗迹每日默认闯关次数
	GetSoulRuinsChallengeNum() int32
	//帝陵遗迹每日购买挑战次数
	GetSoulRuinsBuyChallengeNum() int32
	//购买1次次数所需要消耗的元宝
	GetSoulRuinsBuyChallengeNumCostGold() int32
	//一键完成消耗元宝
	GetSoulRuinsFinishCostGold() int32
	//获取下一关
	GetSoulRuinsNextLevel(chapter int32, typ types.SoulRuinsType, level int32) (nextChapter int32, nextTyp types.SoulRuinsType, nextLevel int32, flag bool)
}

type soulRuinsService struct {
	soulRuinsTemplateByIdMap map[int32]*gametemplate.SoulRuinsTemplate
	//帝陵遗迹模板
	soulRuinsTemplateMap map[int32]map[types.SoulRuinsType]map[int32]*gametemplate.SoulRuinsTemplate
	//帝陵遗迹星级奖励模板
	soulRuinsStarTemplateMap map[int32]map[types.SoulRuinsType]*gametemplate.SoulRuinsStarTemplate
	//章节、类型
	chapterTypListMap map[int32][]int32
	//类型最大章节和等级
	chapterMaxInfoMap map[types.SoulRuinsType]map[int32]int32
}

//初始化
func (srs *soulRuinsService) init() (err error) {
	srs.chapterMaxInfoMap = make(map[types.SoulRuinsType]map[int32]int32)
	srs.soulRuinsTemplateByIdMap = make(map[int32]*gametemplate.SoulRuinsTemplate)
	srs.soulRuinsTemplateMap = make(map[int32]map[types.SoulRuinsType]map[int32]*gametemplate.SoulRuinsTemplate)
	srs.soulRuinsStarTemplateMap = make(map[int32]map[types.SoulRuinsType]*gametemplate.SoulRuinsStarTemplate)
	srs.chapterTypListMap = make(map[int32][]int32)
	//赋值soulRuinsTemplateMap
	srTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulRuinsTemplate)(nil))
	for _, templateObject := range srTemplateMap {
		soulRuinsTemplate, _ := templateObject.(*gametemplate.SoulRuinsTemplate)

		srs.soulRuinsTemplateByIdMap[int32(soulRuinsTemplate.TemplateId())] = soulRuinsTemplate

		chapter := soulRuinsTemplate.Chapter
		typ := soulRuinsTemplate.GetType()
		level := soulRuinsTemplate.Level
		chapterMap, exist := srs.chapterMaxInfoMap[typ]
		if !exist {
			chapterMap = make(map[int32]int32)
			srs.chapterMaxInfoMap[typ] = chapterMap
		}
		oldLevel := chapterMap[chapter]
		if level > oldLevel {
			chapterMap[chapter] = level
		}

		soulRuinsTypMap, ok := srs.soulRuinsTemplateMap[chapter]
		if !ok {
			soulRuinsTypMap = make(map[types.SoulRuinsType]map[int32]*gametemplate.SoulRuinsTemplate)
			srs.soulRuinsTemplateMap[chapter] = soulRuinsTypMap
		}
		soulRuinsLevelMap, ok := soulRuinsTypMap[typ]
		if !ok {
			soulRuinsLevelMap = make(map[int32]*gametemplate.SoulRuinsTemplate)
			soulRuinsTypMap[typ] = soulRuinsLevelMap
		}
		soulRuinsLevelMap[level] = soulRuinsTemplate

		rtyp := int32(typ)
		typList := srs.chapterTypListMap[chapter]
		existFlag := false
		for _, ruinsType := range typList {
			if ruinsType == rtyp {
				existFlag = true
				break
			}
		}
		if !existFlag {
			typList = append(typList, rtyp)
			srs.chapterTypListMap[chapter] = typList
		}
	}

	//赋值soulRuinsStarTemplateMap
	srsTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SoulRuinsStarTemplate)(nil))
	for _, templateObject := range srsTemplateMap {
		soulRuinsStarTemplate, _ := templateObject.(*gametemplate.SoulRuinsStarTemplate)

		//校验chapter 和 type是否有效
		chapter := soulRuinsStarTemplate.Chapter
		typ := soulRuinsStarTemplate.GetType()
		typList, ok := srs.chapterTypListMap[chapter]
		if !ok {
			err = fmt.Errorf("soulruinsService: soul_ruins_star field chapter:%d and type:%d is invalid", chapter, typ)
			return
		}
		flag := utils.ContainInt32(typList, int32(typ))
		if !flag {
			err = fmt.Errorf("soulruinsService: soul_ruins_star field chapter:%d and type:%d is invalid", chapter, typ)
			return
		}

		soulRuinsStarTypeMap, ok := srs.soulRuinsStarTemplateMap[chapter]
		if !ok {
			soulRuinsStarTypeMap = make(map[types.SoulRuinsType]*gametemplate.SoulRuinsStarTemplate)
			srs.soulRuinsStarTemplateMap[chapter] = soulRuinsStarTypeMap
		}
		soulRuinsStarTypeMap[typ] = soulRuinsStarTemplate
	}

	for chapter, soulRuinsTypMap := range srs.soulRuinsTemplateMap {
		for typ, soulRuinsLevelMap := range soulRuinsTypMap {
			//校验 1级
			_, ok := soulRuinsLevelMap[1]
			if !ok {
				err = fmt.Errorf("soulruinsService: soul_ruins field chapter:%d and type:%d  level= 1 should be exist", chapter, typ)
				return
			}

			//校验soul_ruins_star 是否少配
			soulRuinsStarTypeMap, ok := srs.soulRuinsStarTemplateMap[chapter]
			if !ok {
				err = fmt.Errorf("soulruinsService: soul_ruins_star field chapter:%d is miss", chapter)
				return
			}
			_, ok = soulRuinsStarTypeMap[typ]
			if !ok {
				err = fmt.Errorf("soulruinsService: soul_ruins_star field chapter:%d and type:%d is miss", chapter, typ)
				return
			}
		}
	}

	return nil
}

//获取帝陵遗迹配置通关id
func (srs *soulRuinsService) GetSoulRuinsTemplateById(id int32) *gametemplate.SoulRuinsTemplate {
	to, ok := srs.soulRuinsTemplateByIdMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取帝陵遗迹配置
func (srs *soulRuinsService) GetSoulRuinsTemplate(chapter int32, typ types.SoulRuinsType, level int32) *gametemplate.SoulRuinsTemplate {
	soulRuinsTypMap, ok := srs.soulRuinsTemplateMap[chapter]
	if !ok {
		return nil
	}
	soulRuinsLevelMap, ok := soulRuinsTypMap[typ]
	if !ok {
		return nil
	}
	soulRuins, ok := soulRuinsLevelMap[level]
	if !ok {
		return nil
	}
	return soulRuins
}

//获取帝陵遗迹星级奖励配置
func (srs *soulRuinsService) GetSoulRuinsStarTemplate(chapter int32, typ types.SoulRuinsType) *gametemplate.SoulRuinsStarTemplate {
	soulRuinsStarTypeMap, ok := srs.soulRuinsStarTemplateMap[chapter]
	if !ok {
		return nil
	}
	soulRuinsStarTemplate, ok := soulRuinsStarTypeMap[typ]
	if !ok {
		return nil
	}
	return soulRuinsStarTemplate
}

//获取帝陵遗迹触发特殊事件(优先触发boss,其次帝魂降临,最后马贼)
func (srs *soulRuinsService) GetSoulRuinsTriggerSpecialEvent(chapter int32, typ types.SoulRuinsType, level int32) types.SoulRuinsEventType {
	to := srs.GetSoulRuinsTemplate(chapter, typ, level)
	if to == nil {
		return types.SoulRuinsEventTypeNot
	}
	specialRateMap := to.GetSpecialRateMap()
	//触发boss
	bossRate := specialRateMap[types.SoulRuinsEventTypeBoss]
	flag := mathutils.RandomHit(common.MAX_RATE, int(bossRate))
	if flag {
		return types.SoulRuinsEventTypeBoss
	}
	//触发帝魂降临
	soulRate := specialRateMap[types.SoulRuinsEventTypeSoul]
	flag = mathutils.RandomHit(common.MAX_RATE, int(soulRate))
	if flag {
		return types.SoulRuinsEventTypeSoul
	}
	//触发马贼
	robberRate := specialRateMap[types.SoulRuinsEventTypeRobber]
	flag = mathutils.RandomHit(common.MAX_RATE, int(robberRate))
	if flag {
		return types.SoulRuinsEventTypeRobber
	}
	return types.SoulRuinsEventTypeNot
}

//获取挑战用时获取星数
func (srs *soulRuinsService) GetSoulRuinsStarByTime(chapter int32, typ types.SoulRuinsType, level int32, usedTime int32) int32 {
	if usedTime < 0 {
		panic(fmt.Errorf("soulruinsService:usedTime should be more 0"))
	}
	to := srs.GetSoulRuinsTemplate(chapter, typ, level)
	if to == nil {
		return 0
	}
	starMap := to.GetStarTimeMap()
	if usedTime > starMap[types.SoulRuinsStarNumTypeOne] {
		return int32(types.SoulRuinsStarNumTypeZero)
	} else if usedTime <= starMap[types.SoulRuinsStarNumTypeThree] {
		return int32(types.SoulRuinsStarNumTypeThree)
	} else if usedTime > starMap[types.SoulRuinsStarNumTypeThree] && usedTime <= starMap[types.SoulRuinsStarNumTypeTwo] {
		return int32(types.SoulRuinsStarNumTypeTwo)
	} else {
		return int32(types.SoulRuinsStarNumTypeOne)
	}
}

//获取下一关
func (srs *soulRuinsService) GetSoulRuinsNextLevel(chapter int32, typ types.SoulRuinsType, level int32) (nextChapter int32, nextTyp types.SoulRuinsType, nextLevel int32, flag bool) {
	chapterMap, exist := srs.chapterMaxInfoMap[typ]
	if !exist {
		return
	}
	maxLevel, exist := chapterMap[chapter]
	if !exist {
		return
	}
	if level < maxLevel {
		return chapter, typ, level + 1, true
	}
	if typ >= types.SoulRuinsTypeMax {
		return
	}
	nextTyp = types.SoulRuinsType(typ + 1)
	chapterMap, exist = srs.chapterMaxInfoMap[nextTyp]
	if !exist {
		return
	}
	return chapter + 1, nextTyp, 1, true

}

//帝陵遗迹每日默认闯关次数
func (srs *soulRuinsService) GetSoulRuinsChallengeNum() int32 {
	challengeNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSoulRuinsChallengeNum)
	return challengeNum
}

//帝陵遗迹每日购买挑战次数
func (srs *soulRuinsService) GetSoulRuinsBuyChallengeNum() int32 {
	buyChallengeNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSoulRuinsBuyNum)
	return buyChallengeNum
}

//购买1次次数所需要消耗的元宝
func (srs *soulRuinsService) GetSoulRuinsBuyChallengeNumCostGold() int32 {
	buyChallengeNumCostGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSoulRuinsBuyNumCostGold)
	return buyChallengeNumCostGold
}

//一键完成消耗元宝
func (srs *soulRuinsService) GetSoulRuinsFinishCostGold() int32 {
	finishCostGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSoulRuinsFinishCostGold)
	return finishCostGold
}

var (
	once sync.Once
	cs   *soulRuinsService
)

func Init() (err error) {
	once.Do(func() {
		cs = &soulRuinsService{}
		err = cs.init()
	})
	return err
}

func GetSoulRuinsService() SoulRuinsService {
	return cs
}
