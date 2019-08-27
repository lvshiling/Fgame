package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fmt"
)

//配置
type ChuangShiWarTemplate struct {
	*ChuangShiWarTemplateVO
	firstXianZhiTemplate  *XianZhiQuYuTemplate
	secondXianZhiTemplate *XianZhiQuYuTemplate
	firstXianZhiArea      []coretypes.Position
	secondXianZhiArea     []coretypes.Position
	fixPos                []coretypes.Position

	yuXiPos                coretypes.Position //地图位置
	protectPos             coretypes.Position //保护罩位置
	yuXiBiologyTemp        *BiologyTemplate
	protectBiologyTemp     *BiologyTemplate
	protectXianZhiTemplate *XianZhiQuYuTemplate
	protectArea            []coretypes.Position
	protectFixPos          coretypes.Position

	campType chuangshitypes.ChuangShiCampType
}

func (t *ChuangShiWarTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiWarTemplate) GetCampType() chuangshitypes.ChuangShiCampType {
	return t.campType
}

func (t *ChuangShiWarTemplate) GetYuXiPos() coretypes.Position {
	return t.yuXiPos
}

func (t *ChuangShiWarTemplate) GetProtectPos() coretypes.Position {
	return t.protectPos
}

func (t *ChuangShiWarTemplate) GetYuXiBiologyTemp() *BiologyTemplate {
	return t.yuXiBiologyTemp
}
func (t *ChuangShiWarTemplate) GetProtectBiologyTemp() *BiologyTemplate {
	return t.protectBiologyTemp
}

func (t *ChuangShiWarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// tempFirstXianZhiQuYuTemplate := template.GetTemplateService().Get(int(t.FirstXianzhi), (*XianZhiQuYuTemplate)(nil))
	// if tempFirstXianZhiQuYuTemplate == nil {
	// 	err = fmt.Errorf("[%s] invalid", t.FirstXianzhi)
	// 	return template.NewTemplateFieldError("FirstXianzhi", err)
	// }
	// t.firstXianZhiTemplate = tempFirstXianZhiQuYuTemplate.(*XianZhiQuYuTemplate)
	// tempSecondXianZhiQuYuTemplate := template.GetTemplateService().Get(int(t.SecondXianzhi), (*XianZhiQuYuTemplate)(nil))
	// if tempSecondXianZhiQuYuTemplate == nil {
	// 	err = fmt.Errorf("[%s] invalid", t.SecondXianzhi)
	// 	return template.NewTemplateFieldError("SecondXianzhi", err)
	// }
	// t.secondXianZhiTemplate = tempSecondXianZhiQuYuTemplate.(*XianZhiQuYuTemplate)

	// // 防护罩限制
	// protectQuYuTemplate := template.GetTemplateService().Get(int(t.ProtectQuYuPos), (*XianZhiQuYuTemplate)(nil))
	// if protectQuYuTemplate == nil {
	// 	err = fmt.Errorf("[%s] invalid", t.ProtectQuYuPos)
	// 	return template.NewTemplateFieldError("ProtectQuYuPos", err)
	// }
	// t.protectXianZhiTemplate = protectQuYuTemplate.(*XianZhiQuYuTemplate)

	// //
	// pos1Arr, err := coreutils.SplitAsFloatArray(t.LahuiPos1)
	// if err != nil {
	// 	return
	// }
	// if len(pos1Arr) != 3 {
	// 	err = fmt.Errorf("[%s] invalid", t.LahuiPos1)
	// 	return template.NewTemplateFieldError("LahuiPos1", err)
	// }
	// pos1 := coretypes.Position{
	// 	X: pos1Arr[0],
	// 	Y: pos1Arr[1],
	// 	Z: pos1Arr[2],
	// }
	// t.fixPos = append(t.fixPos, pos1)

	// // //
	// // pos2Arr, err := coreutils.SplitAsFloatArray(t.LahuiPos2)
	// // if err != nil {
	// // 	return
	// // }
	// // if len(pos2Arr) != 3 {
	// // 	err = fmt.Errorf("[%s] invalid", t.LahuiPos2)
	// // 	return template.NewTemplateFieldError("LahuiPos2",  err)
	// // }

	// // pos2 := coretypes.Position{
	// // 	X: pos2Arr[0],
	// // 	Y: pos2Arr[1],
	// // 	Z: pos2Arr[2],
	// // }
	// // t.fixPos = append(t.fixPos, pos2)

	// //
	// t.yuXiPos = coretypes.Position{
	// 	X: t.YuXiPosX,
	// 	Y: t.YuXiPosY,
	// 	Z: t.YuXiPosZ,
	// }
	// t.protectPos = coretypes.Position{
	// 	X: t.ProtectPosX,
	// 	Y: t.ProtectPosY,
	// 	Z: t.ProtectPosZ,
	// }

	// //防护墙驱逐位置
	// protectFixPosArr, err := coreutils.SplitAsFloatArray(t.ProtectLaHuiPos)
	// if err != nil {
	// 	return
	// }
	// if len(protectFixPosArr) != 3 {
	// 	err = fmt.Errorf("[%s] invalid", t.ProtectLaHuiPos)
	// 	return template.NewTemplateFieldError("ProtectLaHuiPos", err)
	// }
	// protectFixPos := coretypes.Position{
	// 	X: protectFixPosArr[0],
	// 	Y: protectFixPosArr[1],
	// 	Z: protectFixPosArr[2],
	// }
	// t.protectFixPos = protectFixPos

	// //阵营类型
	// t.campType = chuangshitypes.ChuangShiCampType(t.CampType)
	// if !t.campType.Valid() {
	// 	err = fmt.Errorf("[%d] invalid", t.CampType)
	// 	return template.NewTemplateFieldError("CampType", err)
	// }
	return
}

func (t *ChuangShiWarTemplate) PatchAfterCheck() {

}

func (t *ChuangShiWarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//原地复活次数
	err = validator.MinValidate(float64(t.RebornSitu), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RebornSitu", fmt.Errorf("[%s] invalid", t.RebornSitu))
	}

	//进入后多少毫秒获得奖励
	if err = validator.MinValidate(float64(t.FristTiem), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("FristTiem", err)
		return
	}
	//奖励时间间隔
	if err = validator.MinValidate(float64(t.RewTiem), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewTiem", err)
		return
	}
	//单次发放的经验
	if err = validator.MinValidate(float64(t.RewExp), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}
	//单次发放的经验点
	if err = validator.MinValidate(float64(t.RewExpPoint), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}
	//奖励银两
	if err = validator.MinValidate(float64(t.RewSilver), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}
	//单次安全区发放的经验
	if err = validator.MinValidate(float64(t.RewExpSafeArea), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExpSafeArea", err)
		return
	}
	//单次安全区发放的经验点
	if err = validator.MinValidate(float64(t.RewExpPointSafeArea), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewExpPointSafeArea", err)
		return
	}
	//安全区奖励银两
	if err = validator.MinValidate(float64(t.RewSilverSafeArea), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("RewSilverSafeArea", err)
		return
	}

	// t.firstXianZhiArea = make([]coretypes.Position, 0, 1)
	// t.firstXianZhiArea = append(t.firstXianZhiArea, t.firstXianZhiTemplate.GetPos())
	// currentXainZhiQuYuTemplate := t.firstXianZhiTemplate
	// for {
	// 	currentXainZhiQuYuTemplate = currentXainZhiQuYuTemplate.GetNext()
	// 	if currentXainZhiQuYuTemplate == nil {
	// 		break
	// 	}
	// 	t.firstXianZhiArea = append(t.firstXianZhiArea, currentXainZhiQuYuTemplate.GetPos())
	// }
	// if len(t.firstXianZhiArea) < 3 {
	// 	err = fmt.Errorf("[%d]  不是多边形", t.firstXianZhiArea)
	// 	return template.NewTemplateFieldError("FirstXianzhi", err)
	// }

	// t.secondXianZhiArea = make([]coretypes.Position, 0, 1)
	// t.secondXianZhiArea = append(t.secondXianZhiArea, t.secondXianZhiTemplate.GetPos())
	// currentXainZhiQuYuTemplate = t.secondXianZhiTemplate
	// for {
	// 	currentXainZhiQuYuTemplate = currentXainZhiQuYuTemplate.GetNext()
	// 	if currentXainZhiQuYuTemplate == nil {
	// 		break
	// 	}
	// 	t.secondXianZhiArea = append(t.secondXianZhiArea, currentXainZhiQuYuTemplate.GetPos())
	// }
	// if len(t.secondXianZhiArea) < 3 {
	// 	err = fmt.Errorf("[%d]  不是多边形", t.secondXianZhiArea)
	// 	return template.NewTemplateFieldError("SecondXianzhi", err)
	// }

	// // 校验位置
	// to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	// if to == nil {
	// 	err = fmt.Errorf("[%d] invalid", t.MapId)
	// 	return template.NewTemplateFieldError("MapId", err)
	// }
	// mapTemp := to.(*MapTemplate)
	// if !mapTemp.GetMap().IsMask(t.yuXiPos.X, t.yuXiPos.Z) {
	// 	err = fmt.Errorf("pos[%s] invalid", t.yuXiPos.String())
	// 	err = template.NewTemplateFieldError("pos", err)
	// 	return
	// }
	// yuxiY := mapTemp.GetMap().GetHeight(t.yuXiPos.X, t.yuXiPos.Z)
	// t.yuXiPos.Y = yuxiY

	// // 校验位置
	// if !mapTemp.GetMap().IsMask(t.protectPos.X, t.protectPos.Z) {
	// 	err = fmt.Errorf("pos[%s] invalid", t.yuXiPos.String())
	// 	err = template.NewTemplateFieldError("pos", err)
	// 	return
	// }
	// protectY := mapTemp.GetMap().GetHeight(t.protectPos.X, t.protectPos.Z)
	// t.protectPos.Y = protectY

	// //
	// yuxiTo := template.GetTemplateService().Get(int(t.YuxiId), (*BiologyTemplate)(nil))
	// if yuxiTo == nil {
	// 	err = fmt.Errorf("[%d] invalid", t.YuxiId)
	// 	return template.NewTemplateFieldError("YuxiId", err)
	// }
	// t.yuXiBiologyTemp = yuxiTo.(*BiologyTemplate)

	// //
	// protectTo := template.GetTemplateService().Get(int(t.ProtectId), (*BiologyTemplate)(nil))
	// if protectTo == nil {
	// 	err = fmt.Errorf("[%d] invalid", t.ProtectId)
	// 	return template.NewTemplateFieldError("ProtectId", err)
	// }
	// t.protectBiologyTemp = protectTo.(*BiologyTemplate)

	// //保护罩区域
	// t.protectArea = append(t.protectArea, t.protectXianZhiTemplate.GetPos())
	// currentProtectQuYuTemplate := t.protectXianZhiTemplate
	// for {
	// 	currentProtectQuYuTemplate = currentProtectQuYuTemplate.GetNext()
	// 	if currentProtectQuYuTemplate == nil {
	// 		break
	// 	}
	// 	t.protectArea = append(t.protectArea, currentProtectQuYuTemplate.GetPos())
	// }
	// if len(t.protectArea) < 3 {
	// 	err = fmt.Errorf("[%d]  不是多边形", t.protectArea)
	// 	return template.NewTemplateFieldError("ZhaoZiXianzhi", err)
	// }
	// //
	// if !mapTemp.GetMap().IsMask(t.protectFixPos.X, t.protectFixPos.Z) {
	// 	err = fmt.Errorf("pos[%s] invalid", t.protectFixPos.String())
	// 	err = template.NewTemplateFieldError("pos", err)
	// 	return
	// }
	// protectFixPosY := mapTemp.GetMap().GetHeight(t.protectFixPos.X, t.protectFixPos.Z)
	// t.protectFixPos.Y = protectFixPosY
	return nil
}

func (t *ChuangShiWarTemplate) GetArea(pos coretypes.Position) int32 {
	if coreutils.PointInPolygon(pos, t.firstXianZhiArea) {
		return 0
	}
	return -1
}

func (t *ChuangShiWarTemplate) IsOnProtectArea(pos coretypes.Position) bool {
	return coreutils.PointInPolygon(pos, t.protectArea)
}

func (t *ChuangShiWarTemplate) GetFixPos(area int32) (pos coretypes.Position, flag bool) {
	if len(t.fixPos) <= int(area) {
		return
	}
	pos = t.fixPos[area]
	flag = true
	return
}

func (t *ChuangShiWarTemplate) GetProtectFixPos() (pos coretypes.Position) {
	return t.protectFixPos
}

func (tt *ChuangShiWarTemplate) FileName() string {
	return "tb_chuangshi_war_constant.json"
}

func init() {
	template.Register((*ChuangShiWarTemplate)(nil))
}
