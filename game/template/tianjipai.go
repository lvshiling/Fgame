package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	funcopentypes "fgame/fgame/game/funcopen/types"
	questtypes "fgame/fgame/game/quest/types"
	secretcardtypes "fgame/fgame/game/secretcard/types"
	"fmt"
	"sort"
)

//天机牌配置
type TianJiPaiTemplate struct {
	*TianJiPaiTemplateVO
	poolTyp          secretcardtypes.SecretCardPoolType //任务池类型
	funcOpenTyp      funcopentypes.FuncOpenType         //模块功能开启类型
	intervalDropMap  map[int32]int32                    //间隔掉落包
	dropIntervalList []int
	speDropList      []int32 //角色前5次掉落

}

func (tt *TianJiPaiTemplate) TemplateId() int {
	return tt.Id
}

func (tt *TianJiPaiTemplate) GetPoolTyp() secretcardtypes.SecretCardPoolType {
	return tt.poolTyp
}

func (tt *TianJiPaiTemplate) GetFuncOpenTyp() funcopentypes.FuncOpenType {
	return tt.funcOpenTyp
}

func (tt *TianJiPaiTemplate) GetIntervalDropMap() map[int32]int32 {
	return tt.intervalDropMap
}

func (tt *TianJiPaiTemplate) GetSpeDropList() []int32 {
	return tt.speDropList
}

func (tt *TianJiPaiTemplate) GetDropIntervalList() []int {
	return tt.dropIntervalList
}

func (tt *TianJiPaiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//pool_type
	tt.poolTyp = secretcardtypes.SecretCardPoolType(tt.PoolType)
	if !tt.poolTyp.Valid() {
		err = fmt.Errorf("[%d] invalid", tt.PoolType)
		return template.NewTemplateFieldError("PoolType", err)
	}

	tt.intervalDropMap = make(map[int32]int32)
	tt.dropIntervalList = make([]int, 0, 4)
	//验证 drop_count_1
	err = validator.MinValidate(float64(tt.DropCount1), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tt.DropCount1)
		err = template.NewTemplateFieldError("DropCount1", err)
		return
	}
	tt.intervalDropMap[tt.DropCount1] = tt.DropId1

	//验证 drop_count_2
	err = validator.MinValidate(float64(tt.DropCount2), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tt.DropCount2)
		err = template.NewTemplateFieldError("DropCount2", err)
		return
	}
	tt.intervalDropMap[tt.DropCount2] = tt.DropId2

	//验证 drop_count_3
	err = validator.MinValidate(float64(tt.DropCount3), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tt.DropCount3)
		err = template.NewTemplateFieldError("DropCount3", err)
		return
	}
	//验证 drop_id_3
	tt.intervalDropMap[tt.DropCount3] = tt.DropId3

	//验证 drop_count_4
	err = validator.MinValidate(float64(tt.DropCount4), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tt.DropCount4)
		err = template.NewTemplateFieldError("DropCount4", err)
		return
	}
	//验证 drop_id_4
	tt.intervalDropMap[tt.DropCount4] = tt.DropId4

	dropIdArr, err := utils.SplitAsIntArray(tt.SpeDrop)
	if err != nil || len(dropIdArr) != 5 {
		err = fmt.Errorf("[%s] invalid", tt.SpeDrop)
		err = template.NewTemplateFieldError("SpeDrop", err)
		return
	}

	for _, dropId := range dropIdArr {
		tt.speDropList = append(tt.speDropList, dropId)
	}

	for interval, _ := range tt.intervalDropMap {
		tt.dropIntervalList = append(tt.dropIntervalList, int(interval))
	}
	//降序排序dropIntervalList
	sort.Sort(sort.Reverse(sort.IntSlice(tt.dropIntervalList)))

	return nil
}

func (tt *TianJiPaiTemplate) PatchAfterCheck() {

}

func (tt *TianJiPaiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//module_opened_id
	if tt.ModuleOpenedId != 0 {
		to := template.GetTemplateService().Get(int(tt.ModuleOpenedId), (*ModuleOpenedTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tt.ModuleOpenedId)
			return template.NewTemplateFieldError("ModuleOpenedId", err)
		}

		tt.funcOpenTyp = to.(*ModuleOpenedTemplate).GetFuncOpenType()
	}

	//quest_id
	if tt.QuestId != 0 {
		to := template.GetTemplateService().Get(int(tt.QuestId), (*QuestTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tt.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}

		questTemplate := to.(*QuestTemplate)
		if questTemplate.GetQuestType() != questtypes.QuestTypeTianJiPai {
			err = fmt.Errorf("[%d] invalid", tt.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}
	}

	//star_min
	err = validator.MinValidate(float64(tt.StarMin), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("StarMin", err)
		return
	}

	err = validator.MinValidate(float64(tt.StarMax), float64(tt.StarMin), true)
	if err != nil {
		err = template.NewTemplateFieldError("StarMax", err)
		return
	}

	//rew_silver
	err = validator.MinValidate(float64(tt.RewSilver), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	return nil
}

func (tt *TianJiPaiTemplate) FileName() string {
	return "tb_tianjipai.json"
}

func init() {
	template.Register((*TianJiPaiTemplate)(nil))
}
