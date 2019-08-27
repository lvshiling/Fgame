package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fmt"
)

func init() {
	template.Register((*GuideReplicaCatDogTemplate)(nil))
}

type GuideReplicaCatDogTemplate struct {
	*GuideReplicaCatDogTemplateVO
	catDropIdList     []int32
	dogDropIdList     []int32
	defaultDropIdList []int32
}

func (t *GuideReplicaCatDogTemplate) TemplateId() int {
	return t.Id
}

func (t *GuideReplicaCatDogTemplate) GetDropId(killMap map[guidereplicatypes.CatDogKillType]int32) []int32 {
	rewKillType := guidereplicatypes.CatDogKillTypeDefault
	initNum := int32(0)
	isSame := true
	for typ, num := range killMap {
		if initNum == 0 {
			initNum = num
			rewKillType = typ
			continue
		}

		if initNum != num {
			isSame = false
		}

		if initNum > num {
			continue
		}

		initNum = num
		rewKillType = typ
	}

	if isSame {
		rewKillType = guidereplicatypes.CatDogKillTypeDefault
	}

	switch rewKillType {
	case guidereplicatypes.CatDogKillTypeCat:
		return t.catDropIdList
	case guidereplicatypes.CatDogKillTypeDog:
		return t.dogDropIdList
	default:
		return t.defaultDropIdList
	}
}

func (t *GuideReplicaCatDogTemplate) FileName() string {
	return "tb_daily_maogou.json"
}

//组合成需要的数据
func (t *GuideReplicaCatDogTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.catDropIdList, err = utils.SplitAsIntArray(t.CatDropId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.CatDropId)
		return template.NewTemplateFieldError("CatDropId", err)
	}
	//
	t.dogDropIdList, err = utils.SplitAsIntArray(t.DogDropId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DogDropId)
		return template.NewTemplateFieldError("DogDropId", err)
	}
	//
	t.defaultDropIdList, err = utils.SplitAsIntArray(t.DefaultDropId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DefaultDropId)
		return template.NewTemplateFieldError("DefaultDropId", err)
	}

	return nil
}

//检查有效性
func (t *GuideReplicaCatDogTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//怪物id
	if t.CatBiologyId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.CatBiologyId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.CatBiologyId)
			return err
		}
		_, ok := bilogyTemp.(*BiologyTemplate)
		if !ok {
			err = fmt.Errorf("BiologyId [%d] no exist", t.CatBiologyId)
			return
		}
	}
	//怪物id
	if t.DogBiologyId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.DogBiologyId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.DogBiologyId)
			return err
		}
		_, ok := bilogyTemp.(*BiologyTemplate)
		if !ok {
			err = fmt.Errorf("BiologyId [%d] no exist", t.DogBiologyId)
			return
		}
	}

	return nil
}

//检验后组合
func (t *GuideReplicaCatDogTemplate) PatchAfterCheck() {
}
