package template

import (
	"fgame/fgame/core/template"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

//苍龙棋局接口处理
type ChessTemplateService interface {
	//随机获取苍龙棋局配置
	GetChessRandom(typ chesstypes.ChessType, excludeId int32) *gametemplate.ChessTemplate
	//获取苍龙棋局配置
	GetChessByTypAndChessId(typ chesstypes.ChessType, chessId int32) *gametemplate.ChessTemplate
	//获取角色特殊掉落池
	GetChessByRole(role playertypes.RoleType) *gametemplate.ChessTemplate
}

type chessTemplateService struct {
	//苍龙棋局配置
	chessMap map[int32]*gametemplate.ChessTemplate
	//棋局map
	chessByTypAndChessIdMap map[chesstypes.ChessType]map[int32]*gametemplate.ChessTemplate
}

//初始化
func (cs *chessTemplateService) init() error {
	cs.chessMap = make(map[int32]*gametemplate.ChessTemplate)
	cs.chessByTypAndChessIdMap = make(map[chesstypes.ChessType]map[int32]*gametemplate.ChessTemplate)
	//苍龙棋局
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ChessTemplate)(nil))
	for _, templateObject := range templateMap {
		chessTemplate, _ := templateObject.(*gametemplate.ChessTemplate)
		cs.chessMap[chessTemplate.ChessId] = chessTemplate

		cheseByChessIdMap, ok := cs.chessByTypAndChessIdMap[chessTemplate.GetChessType()]
		if !ok {
			cheseByChessIdMap = make(map[int32]*gametemplate.ChessTemplate)
			cs.chessByTypAndChessIdMap[chessTemplate.GetChessType()] = cheseByChessIdMap
		}
		cheseByChessIdMap[chessTemplate.ChessId] = chessTemplate
	}

	//校验 每个类型棋局至少2个
	for chessType, chessMap := range cs.chessByTypAndChessIdMap {
		if len(chessMap) < 2 {
			err := fmt.Errorf("苍龙棋局：每种棋局类型至少2幅棋局;chessType:%d", chessType)
			return err
		}
	}
	return nil
}

//随机获取苍龙棋局配置
func (cs *chessTemplateService) GetChessRandom(typ chesstypes.ChessType, excludeId int32) *gametemplate.ChessTemplate {
	chessMap := cs.chessByTypAndChessIdMap[typ]

	weights := make([]int64, 0, len(chessMap)-1)
	tempChessList := make([]*gametemplate.ChessTemplate, 0, len(chessMap)-1)
	for _, ch := range chessMap {
		if ch.ChessId == excludeId {
			continue
		}
		weights = append(weights, int64(ch.Rate))
		tempChessList = append(tempChessList, ch)
	}
	index := mathutils.RandomWeights(weights)
	if index == -1 {
		return nil
	}
	ch := tempChessList[index]
	return ch
}

//随机获取苍龙棋局配置
func (cs *chessTemplateService) GetChessByTypAndChessId(typ chesstypes.ChessType, chessId int32) *gametemplate.ChessTemplate {
	chessByDrop, ok := cs.chessByTypAndChessIdMap[typ]
	if !ok {
		return nil
	}
	chess, ok := chessByDrop[chessId]
	if !ok {
		return nil
	}

	return chess
}

//随机获取苍龙棋局配置
func (cs *chessTemplateService) GetChessByRole(role playertypes.RoleType) *gametemplate.ChessTemplate {
	chessId := int32(0)
	switch role {
	case playertypes.RoleTypeKaiTian:
		chessId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeKaiTianSpecialChessId))
	case playertypes.RoleTypeYiJian:
		chessId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeYiJianSpecialChessId))
	case playertypes.RoleTypePoYue:
		chessId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePoYueSpecialChessId))
	}

	return cs.chessMap[chessId]
}

var (
	once sync.Once
	cs   *chessTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &chessTemplateService{}
		err = cs.init()
	})
	return err
}

func GetChessTemplateService() ChessTemplateService {
	return cs
}
