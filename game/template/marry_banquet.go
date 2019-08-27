package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/marry/types"
	scenetypes "fgame/fgame/game/scene/types"

	"fmt"
)

//婚宴档次配置
type MarryBanquetTemplate struct {
	*MarryBanquetTemplateVO
	banquetType    types.MarryBanquetType    //婚礼类型
	banquetSubType types.MarryBanquetSubType //子类型

	houtaiType         types.MarryHoutaiType //后台类型
	isCruise           bool                  //是否巡游
	useItemTemplate    *ItemTemplate         //进阶物品
	battleAttrTemplate *AttrTemplate         //阶别属性
	biologyTemplate    *BiologyTemplate      //婚宴生物模板
	posList            []coretypes.Position  //酒桌位置
	dropList           []int32               //掉落id
	endRewIdMap        map[int32]int32       //婚礼结束后奖励物品
	mapTemplate        *MapTemplate
}

func (mbt *MarryBanquetTemplate) TemplateId() int {
	return mbt.Id
}

func (mbt *MarryBanquetTemplate) GetHoutaiType() types.MarryHoutaiType {
	return mbt.houtaiType
}

func (mbt *MarryBanquetTemplate) GetBanquetType() types.MarryBanquetType {
	return mbt.banquetType
}

func (mbt *MarryBanquetTemplate) GetBanquetSubType() types.MarryBanquetSubType {
	return mbt.banquetSubType
}

func (mbt *MarryBanquetTemplate) GetCruise() bool {
	return mbt.isCruise
}

func (mbt *MarryBanquetTemplate) GetBiologyTemplate() *BiologyTemplate {
	return mbt.biologyTemplate
}

func (mbt *MarryBanquetTemplate) GetUseItemTemplate() *ItemTemplate {
	return mbt.useItemTemplate
}

func (mbt *MarryBanquetTemplate) GetPosList() []coretypes.Position {
	return mbt.posList
}

func (mbt *MarryBanquetTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return mbt.battleAttrTemplate
}

func (mbt *MarryBanquetTemplate) GetDropList() []int32 {
	return mbt.dropList
}

func (mbt *MarryBanquetTemplate) GetEndRewIdMap() map[int32]int32 {
	return mbt.endRewIdMap
}

func (mbt *MarryBanquetTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mbt.FileName(), mbt.TemplateId(), err)
			return
		}
	}()

	mbt.banquetType = types.MarryBanquetType(mbt.Type)
	if !mbt.banquetType.Valid() {
		err = fmt.Errorf("[%d] invalid", mbt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	mbt.houtaiType = types.MarryHoutaiType(mbt.HoutaiType)
	if !mbt.houtaiType.Valid() {
		err = fmt.Errorf("[%d] invalid", mbt.HoutaiType)
		return template.NewTemplateFieldError("HoutaiType", err)
	}

	//验证 subType
	mbt.banquetSubType = types.CreateMarryBanquetSubType(mbt.banquetType, mbt.SubType)
	if mbt.banquetSubType == nil || !mbt.banquetSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", mbt.SubType)
		return template.NewTemplateFieldError("subType", err)
	}

	err = validator.RangeValidate(float64(mbt.IsYouJie), float64(0), true, float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mbt.IsYouJie)
		err = template.NewTemplateFieldError("IsYouJie", err)
		return
	}

	mbt.isCruise = false
	if mbt.IsYouJie == 1 {
		mbt.isCruise = true
	}

	//banquet_biology
	if mbt.banquetType == types.MarryBanquetTypeWed || mbt.banquetType == types.MarryBanquetTypeHunChe {
		to := template.GetTemplateService().Get(int(mbt.BanquetBiology), (*BiologyTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mbt.BanquetBiology)
			return template.NewTemplateFieldError("BanquetBiology", err)
		}
		mbt.biologyTemplate = to.(*BiologyTemplate)

	}

	mbt.dropList = make([]int32, 0, 1)
	if mbt.DropId != "" {
		dropArr, err := utils.SplitAsIntArray(mbt.DropId)
		if err != nil {
			return err
		}

		for _, dropId := range dropArr {
			mbt.dropList = append(mbt.dropList, dropId)
		}
	}

	mbt.posList = make([]coretypes.Position, 0, 3)

	to := template.GetTemplateService().Get(int(mbt.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", mbt.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	mbt.mapTemplate = to.(*MapTemplate)

	if mbt.banquetType == types.MarryBanquetTypeWed {
		if mbt.PosX1 != 0 {
			pos1 := coretypes.Position{
				X: mbt.PosX1,
				Y: mbt.PosY1,
				Z: mbt.PosZ1,
			}

			mbt.posList = append(mbt.posList, pos1)
		}

		if mbt.PosX2 != 0 {
			pos2 := coretypes.Position{
				X: mbt.PosX2,
				Y: mbt.PosY2,
				Z: mbt.PosZ2,
			}

			mbt.posList = append(mbt.posList, pos2)
		}

		if mbt.PosX3 != 0 {
			pos3 := coretypes.Position{
				X: mbt.PosX3,
				Y: mbt.PosY3,
				Z: mbt.PosZ3,
			}

			mbt.posList = append(mbt.posList, pos3)
		}
	}

	mbt.endRewIdMap = make(map[int32]int32)
	if mbt.EndRewId != "" {
		itemArr, err := utils.SplitAsIntArray(mbt.EndRewId)
		if err != nil {
			return err
		}
		numArr, err := utils.SplitAsIntArray(mbt.EndRewCount)
		if err != nil {
			return err
		}
		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s] invalid", mbt.EndRewId)
			return template.NewTemplateFieldError("EndRewId", err)
		}

		for i := int(0); i < len(itemArr); i++ {
			mbt.endRewIdMap[itemArr[i]] = numArr[i]
		}
	}

	return nil
}

func (mbt *MarryBanquetTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mbt.FileName(), mbt.TemplateId(), err)
			return
		}
	}()

	//验证物品
	for itemId, num := range mbt.endRewIdMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", mbt.EndRewId)
			return template.NewTemplateFieldError("EndRewId", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", mbt.EndRewCount)
			err = template.NewTemplateFieldError("EndRewCount", err)
			return
		}
	}

	for index, pos := range mbt.posList {
		mask := mbt.mapTemplate.GetMap().IsMask(pos.X, pos.Z)
		if !mask {
			err = fmt.Errorf("pos[%s] invalid", pos.String())
			err = template.NewTemplateFieldError("pos", err)
			return
		}
		y := mbt.mapTemplate.GetMap().GetHeight(pos.X, pos.Z)
		mbt.posList[index].Y = y
	}

	err = validator.MinValidate(float64(mbt.SugarEach), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mbt.SugarEach)
		err = template.NewTemplateFieldError("SugarEach", err)
		return
	}

	err = validator.MinValidate(float64(mbt.DropTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mbt.DropTime)
		err = template.NewTemplateFieldError("DropTime", err)
		return
	}

	if mbt.banquetType == types.MarryBanquetTypeHunChe {
		if mbt.DropTime <= 0 {
			err = fmt.Errorf("[%d] invalid", mbt.DropTime)
			err = template.NewTemplateFieldError("DropTime", err)
			return
		}

		if mbt.SugarEach <= 0 {
			err = fmt.Errorf("[%d] invalid", mbt.SugarEach)
			err = template.NewTemplateFieldError("SugarEach", err)
			return
		}
	}

	//验证use_silver
	err = validator.MinValidate(float64(mbt.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mbt.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证use_binggold
	err = validator.MinValidate(float64(mbt.UseBinggold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mbt.UseBinggold)
		err = template.NewTemplateFieldError("UseBinggold", err)
		return
	}

	//验证use_gold
	err = validator.MinValidate(float64(mbt.UseGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mbt.UseGold)
		err = template.NewTemplateFieldError("UseGold", err)
		return
	}

	switch mbt.banquetType {
	case types.MarryBanquetTypeWed:
		{
			if mbt.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeWedBanquet {
				err = fmt.Errorf("[%d] invalid", mbt.BanquetBiology)
				err = template.NewTemplateFieldError("BanquetBiology", err)
				return
			}
			to := template.GetTemplateService().Get(int(mbt.AddBuffId), (*BuffTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("[%d] invalid", mbt.AddBuffId)
				err = template.NewTemplateFieldError("AddBuffId", err)
				return
			}
			if mbt.mapTemplate.GetMapType() != scenetypes.SceneTypeMarry {
				err = fmt.Errorf("[%d] invalid", mbt.MapId)
				err = template.NewTemplateFieldError("MapId", err)
				return
			}
			break
		}
	case types.MarryBanquetTypeHunChe:
		{
			if mbt.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeWeddingCar {
				err = fmt.Errorf("[%d] invalid", mbt.BanquetBiology)
				err = template.NewTemplateFieldError("BanquetBiology", err)
				return
			}
			if mbt.mapTemplate.GetMapType() != scenetypes.SceneTypeWorld {
				err = fmt.Errorf("[%d] invalid", mbt.MapId)
				err = template.NewTemplateFieldError("MapId", err)
				return
			}
			break
		}
	case types.MarryBanquetTypeSugar:
		{
			if mbt.mapTemplate.GetMapType() != scenetypes.SceneTypeWorld {
				err = fmt.Errorf("[%d] invalid", mbt.MapId)
				err = template.NewTemplateFieldError("MapId", err)
				return
			}
			break
		}
	}

	return nil
}
func (mbt *MarryBanquetTemplate) PatchAfterCheck() {

}
func (mbt *MarryBanquetTemplate) FileName() string {
	return "tb_marry_banquet.json"
}

func init() {
	template.Register((*MarryBanquetTemplate)(nil))
}
