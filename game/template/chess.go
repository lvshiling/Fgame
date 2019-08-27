package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	chesstypes "fgame/fgame/game/chess/types"
	"fmt"
	"sort"
)

//苍龙棋局配置
type ChessTemplate struct {
	*ChessTemplateVO
	chessType      chesstypes.ChessType
	dropByTimesMap map[int32]int32 //按次数必定掉落map
	timesList      []int           //循环掉落
}

func (t *ChessTemplate) TemplateId() int {
	return t.Id
}

func (t *ChessTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//按次数必定掉落map
	t.dropByTimesMap = make(map[int32]int32)
	if t.MustAmount1 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount1]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount1)
			return template.NewTemplateFieldError("MustAmount1", err)
		}
		t.dropByTimesMap[t.MustAmount1] = t.MustGet1
		t.timesList = append(t.timesList, int(t.MustAmount1))
	}
	if t.MustAmount2 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount2]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount2)
			return template.NewTemplateFieldError("MustAmount2", err)
		}
		t.dropByTimesMap[t.MustAmount2] = t.MustGet2
		t.timesList = append(t.timesList, int(t.MustAmount2))
	}
	if t.MustAmount3 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount3]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount3)
			return template.NewTemplateFieldError("MustAmount3", err)
		}
		t.dropByTimesMap[t.MustAmount3] = t.MustGet3
		t.timesList = append(t.timesList, int(t.MustAmount3))
	}
	if t.MustAmount4 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount4]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount4)
			return template.NewTemplateFieldError("MustAmount4", err)
		}
		t.dropByTimesMap[t.MustAmount4] = t.MustGet4
		t.timesList = append(t.timesList, int(t.MustAmount4))
	}
	if t.MustAmount5 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount5]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount5)
			return template.NewTemplateFieldError("MustAmount5", err)
		}
		t.dropByTimesMap[t.MustAmount5] = t.MustGet5
		t.timesList = append(t.timesList, int(t.MustAmount5))
	}
	if t.MustAmount6 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount6]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount6)
			return template.NewTemplateFieldError("MustAmount6", err)
		}
		t.dropByTimesMap[t.MustAmount6] = t.MustGet6
		t.timesList = append(t.timesList, int(t.MustAmount6))
	}
	if t.MustAmount7 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount7]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount7)
			return template.NewTemplateFieldError("MustAmount7", err)
		}
		t.dropByTimesMap[t.MustAmount7] = t.MustGet7
		t.timesList = append(t.timesList, int(t.MustAmount7))
	}
	if t.MustAmount8 > 0 {
		if _, ok := t.dropByTimesMap[t.MustAmount8]; ok {
			err = fmt.Errorf("[%d] invalid", t.MustAmount8)
			return template.NewTemplateFieldError("MustAmount8", err)
		}
		t.dropByTimesMap[t.MustAmount8] = t.MustGet8
		t.timesList = append(t.timesList, int(t.MustAmount8))
	}

	return
}

func (t *ChessTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//棋局概率
	err = validator.MinValidate(float64(t.Rate), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("Rate", err)
	}

	//类型
	typ := chesstypes.ChessType(t.Type)
	if !typ.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.chessType = typ

	// 银两消耗
	err = validator.MinValidate(float64(t.SilverUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("SilverUse", err)
	}
	// 元宝消耗
	err = validator.MinValidate(float64(t.GoldUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("GoldUse", err)
	}
	// 绑元消耗
	err = validator.MinValidate(float64(t.BindGoldUse), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BindGoldUse", err)
	}

	// 赠送物品id
	if t.GiftItem > 0 {
		to := template.GetTemplateService().Get(int(t.GiftItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.GiftItem)
			return template.NewTemplateFieldError("GiftItem", err)
		}

		// 赠送物品数量
		err = validator.MinValidate(float64(t.GiftItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("GiftItemCount", err)
			return
		}
	}

	return nil
}

func (t *ChessTemplate) PatchAfterCheck() {
}

func (t *ChessTemplate) FileName() string {
	return "tb_chess.json"
}

func (t *ChessTemplate) GetChessType() chesstypes.ChessType {
	return t.chessType
}

func (t *ChessTemplate) GetRewDropMap() map[int32]int32 {
	return t.dropByTimesMap
}

func (t *ChessTemplate) GetDropTimesDescList() []int {
	newList := make([]int, len(t.timesList))
	copy(newList, t.timesList)
	sort.Sort(sort.Reverse(sort.IntSlice(newList)))
	return newList
}

func init() {
	template.Register((*ChessTemplate)(nil))
}
