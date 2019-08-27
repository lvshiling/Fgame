package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	"fmt"
)

//玉玺常量配置
type YuXiConstantTemplate struct {
	*YuXiConstantTemplateVO
	winItemMap           map[int32]int32
	dayItemMap           map[int32]int32
	yuXiBiologyTemp      *BiologyTemplate
	yuXiPos              coretypes.Position //地图位置
	winnerModelPos       coretypes.Position //获胜盟主雕像位置
	winnerModelCouplePos coretypes.Position //获胜盟主配偶雕像位置
}

func (t *YuXiConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *YuXiConstantTemplate) GetYuXiPos() coretypes.Position {
	return t.yuXiPos
}

func (t *YuXiConstantTemplate) GetWinnerModelPos() coretypes.Position {
	return t.winnerModelPos
}

func (t *YuXiConstantTemplate) GetWinnerModelCouplePos() coretypes.Position {
	return t.winnerModelCouplePos
}

func (t *YuXiConstantTemplate) GetWinItemMap() map[int32]int32 {
	return t.winItemMap
}

func (t *YuXiConstantTemplate) GetDayItemMap() map[int32]int32 {
	return t.dayItemMap
}

func (t *YuXiConstantTemplate) GetYuXiBiologyTemp() *BiologyTemplate {
	return t.yuXiBiologyTemp
}

func (t *YuXiConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//获胜获得物品
	t.winItemMap = make(map[int32]int32)
	winItemArr, err := coreutils.SplitAsIntArray(t.MailItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MailItemId)
		err = template.NewTemplateFieldError("MailItemId", err)
		return
	}
	winCountArr, err := coreutils.SplitAsIntArray(t.MailItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MailItemCount)
		err = template.NewTemplateFieldError("MailItemCount", err)
		return
	}
	if len(winItemArr) != len(winCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.MailItemId, t.MailItemCount)
		err = template.NewTemplateFieldError("MailItemId or MailItemCount", err)
		return
	}
	for index, itemId := range winItemArr {
		t.winItemMap[itemId] += winCountArr[index]
	}

	//获胜每日获得物品
	t.dayItemMap = make(map[int32]int32)
	dayItemArr, err := coreutils.SplitAsIntArray(t.WinDayItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.WinDayItemId)
		err = template.NewTemplateFieldError("WinDayItemId", err)
		return
	}
	dayCountArr, err := coreutils.SplitAsIntArray(t.WinDayItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.WinDayItemCount)
		err = template.NewTemplateFieldError("WinDayItemCount", err)
		return
	}
	if len(dayItemArr) != len(dayCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.WinDayItemId, t.WinDayItemCount)
		err = template.NewTemplateFieldError("WinDayItemId or WinDayItemCount", err)
		return
	}
	for index, itemId := range dayItemArr {
		t.dayItemMap[itemId] += dayCountArr[index]
	}

	//位置
	t.yuXiPos = coretypes.Position{
		X: t.PosX,
		Y: t.PosY,
		Z: t.PosZ,
	}
	//位置
	t.winnerModelPos = coretypes.Position{
		X: t.ModelPosX,
		Y: t.ModelPosY,
		Z: t.ModelPosZ,
	}
	//位置
	t.winnerModelCouplePos = coretypes.Position{
		X: t.ModelCouplePosX,
		Y: t.ModelCouplePosY,
		Z: t.ModelCouplePosZ,
	}
	return nil
}

func (t *YuXiConstantTemplate) PatchAfterCheck() {

}

func (t *YuXiConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	// 校验位置
	to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	mapTemp := to.(*MapTemplate)
	mask := mapTemp.GetMap().IsMask(t.yuXiPos.X, t.yuXiPos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", t.yuXiPos.String())
		err = template.NewTemplateFieldError("pos", err)
		return
	}
	y := mapTemp.GetMap().GetHeight(t.yuXiPos.X, t.yuXiPos.Z)
	t.yuXiPos.Y = y

	// 雕像位置
	modelMapTo := template.GetTemplateService().Get(int(t.ModelMapId), (*MapTemplate)(nil))
	if modelMapTo == nil {
		err = fmt.Errorf("[%d] invalid", t.ModelMapId)
		return template.NewTemplateFieldError("ModelMapId", err)
	}
	modelMapTemp := modelMapTo.(*MapTemplate)
	if !modelMapTemp.IsWorld() {
		err = fmt.Errorf("[%d] invalid", t.ModelMapId)
		return template.NewTemplateFieldError("ModelMapId", err)
	}
	// modelMask := modelMapTemp.GetMap().IsMask(t.winnerModelPos.X, t.winnerModelPos.Z)
	// if !modelMask {
	// 	err = fmt.Errorf("pos[%s] invalid", t.winnerModelPos.String())
	// 	err = template.NewTemplateFieldError("pos", err)
	// 	return
	// }

	// // 雕像配偶位置
	// modelCoupleMask := modelMapTemp.GetMap().IsMask(t.winnerModelCouplePos.X, t.winnerModelCouplePos.Z)
	// if !modelCoupleMask {
	// 	err = fmt.Errorf("pos[%s] invalid", t.winnerModelCouplePos.String())
	// 	err = template.NewTemplateFieldError("pos", err)
	// 	return
	// }

	//WinTime
	err = validator.MinValidate(float64(t.WinTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WinTime)
		return template.NewTemplateFieldError("WinTime", err)
	}

	//YuxiBiologyId
	biologyTo := template.GetTemplateService().Get(int(t.YuxiBiologyId), (*BiologyTemplate)(nil))
	if biologyTo == nil {
		err = fmt.Errorf("[%d] invalid", t.YuxiBiologyId)
		return template.NewTemplateFieldError("YuxiBiologyId", err)
	}
	t.yuXiBiologyTemp = biologyTo.(*BiologyTemplate)

	//BuffId
	buffTo := template.GetTemplateService().Get(int(t.BuffId), (*BuffTemplate)(nil))
	if buffTo == nil {
		err = fmt.Errorf("[%d] invalid", t.BuffId)
		return template.NewTemplateFieldError("BuffId", err)
	}

	//
	for itemId, num := range t.winItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.MailItemId)
			err = template.NewTemplateFieldError("MailItemId", err)
			return
		}

		//数量
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.MailItemCount)
			return template.NewTemplateFieldError("MailItemCount", err)
		}
	}

	//
	for itemId, num := range t.dayItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.WinDayItemId)
			err = template.NewTemplateFieldError("WinDayItemId", err)
			return
		}

		//数量
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.WinDayItemCount)
			return template.NewTemplateFieldError("WinDayItemCount", err)
		}
	}

	return nil
}

func (t *YuXiConstantTemplate) FileName() string {
	return "tb_yuxi_constant.json"
}

func init() {
	template.Register((*YuXiConstantTemplate)(nil))
}
