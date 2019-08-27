package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	viptypes "fgame/fgame/game/vip/types"
	"fmt"
)

//消费等级配置
type CostLevelTemplate struct {
	*CostLevelTemplateVO
	nextTemp        *CostLevelTemplate
	advancedRuleMap map[viptypes.CostLevelRuleType]*viptypes.AdvancedRateData //升阶规则
	dropRuleMap     map[viptypes.CostLevelRuleType][]int                      //掉落规则
}

func (t *CostLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *CostLevelTemplate) GetAdvancedRuleMap() map[viptypes.CostLevelRuleType]*viptypes.AdvancedRateData {
	return t.advancedRuleMap
}

func (t *CostLevelTemplate) GetDropRuleMap() map[viptypes.CostLevelRuleType][]int {
	return t.dropRuleMap
}

func (t *CostLevelTemplate) GetNextTemplate() *CostLevelTemplate {
	return t.nextTemp
}

func (t *CostLevelTemplate) PatchAfterCheck() {

}
func (t *CostLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.advancedRuleMap = make(map[viptypes.CostLevelRuleType]*viptypes.AdvancedRateData)
	// 坐骑
	mountRateData := viptypes.CreateAdvancedRateData(t.MountMinCount, t.MountMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeMount] = mountRateData

	// 战翼
	wingRateData := viptypes.CreateAdvancedRateData(t.WingMinCount, t.WingMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeWing] = wingRateData

	// 暗器
	anqiRateData := viptypes.CreateAdvancedRateData(t.AnqiMinCount, t.AnqiMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeAnqi] = anqiRateData

	// 护体盾
	bodyShieldRateData := viptypes.CreateAdvancedRateData(t.BodyShieldMinCount, t.BodyShieldMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeBodyShield] = bodyShieldRateData

	// 仙羽
	featherRateData := viptypes.CreateAdvancedRateData(t.FeatherMinCount, t.FeatherMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeFeather] = featherRateData

	// 盾刺
	shieldRateData := viptypes.CreateAdvancedRateData(t.ShieldMinCount, t.ShieldMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeShield] = shieldRateData

	// 领域
	lingyuRateData := viptypes.CreateAdvancedRateData(t.FieldMinCount, t.FieldMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingyu] = lingyuRateData

	// 身法
	shenfaRateData := viptypes.CreateAdvancedRateData(t.ShenfaMinCount, t.ShenfaMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeShenfa] = shenfaRateData

	// 婚戒
	marryRateData := viptypes.CreateAdvancedRateData(t.MarryRingMinCount, t.MarryRingMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeMarryRing] = marryRateData

	// 兵魂升星
	weaponRateData := viptypes.CreateAdvancedRateData(t.WeaponUpstarMinCount, t.WeaponUpstarMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeWeaponUpstar] = weaponRateData

	// 戮仙刃
	massacreRateData := viptypes.CreateAdvancedRateData(t.LuxianrenMinCount, t.LuxianrenMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeMassacre] = massacreRateData

	//法宝
	faBaoRateData := viptypes.CreateAdvancedRateData(t.FabaoMinCount, t.FabaoMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeFaBao] = faBaoRateData

	//仙体
	xianTiRateData := viptypes.CreateAdvancedRateData(t.XianTiMinCount, t.XianTiMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeXianTi] = xianTiRateData

	//点星
	dianXingRateData := viptypes.CreateAdvancedRateData(t.DianXingMinCount, t.DianXingMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeDianXing] = dianXingRateData

	//神器
	shenQiRateData := viptypes.CreateAdvancedRateData(t.ShenQiMinCount, t.ShenQiMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeShenQi] = shenQiRateData

	//噬魂幡
	shiHunFanRateData := viptypes.CreateAdvancedRateData(t.ShiHunFanMinCount, t.ShiHunFanMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeShiHunFan] = shiHunFanRateData
	//天魔
	tianmoRateData := viptypes.CreateAdvancedRateData(t.TianmotiMinCount, t.TianmotiMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeTianMoTi] = tianmoRateData
	//灵兵
	lingTongWeaponRateData := viptypes.CreateAdvancedRateData(t.LingTongWeaponMinCount, t.LingTongWeaponMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongWeapon] = lingTongWeaponRateData
	//灵骑
	lingTongMountRateData := viptypes.CreateAdvancedRateData(t.LingTongMountMinCount, t.LingTongMountMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongMount] = lingTongMountRateData
	//灵翼
	lingTongWingRateData := viptypes.CreateAdvancedRateData(t.LingTongWingMinCount, t.LingTongWingMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongWing] = lingTongWingRateData
	//灵身
	lingTongShenFaRateData := viptypes.CreateAdvancedRateData(t.LingTongShenFaMinCount, t.LingTongShenFaMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongShenFa] = lingTongShenFaRateData
	//灵域
	lingTongLingYuRateData := viptypes.CreateAdvancedRateData(t.LingTongLingYuMinCount, t.LingTongLingYuMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongLingYu] = lingTongLingYuRateData
	//灵宝
	lingTongFaBaoRateData := viptypes.CreateAdvancedRateData(t.LingTongFaBaoMinCount, t.LingTongFaBaoMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongFaBao] = lingTongFaBaoRateData
	//灵体
	lingTongXianTiRateData := viptypes.CreateAdvancedRateData(t.LingTongXianTiMinCount, t.LingTongXianTiMaxCount)
	t.advancedRuleMap[viptypes.CostLevelRuleTypeLingTongXianTi] = lingTongXianTiRateData

	t.dropRuleMap = make(map[viptypes.CostLevelRuleType][]int)
	//赌石
	var timesData []int
	timesData = append(timesData, int(t.DushiCount1Count), int(t.DushiCount2Count), int(t.DushiCount3Count))
	t.dropRuleMap[viptypes.CostLevelRuleTypeGamble] = timesData

	//棋局
	var timesData2 []int
	timesData2 = append(timesData2, int(t.QijvCount1Count), int(t.QijvCount2Count), int(t.QijvCount3Count), int(t.QijvCount4Count), int(t.QijvCount5Count), int(t.QijvCount6Count), int(t.QijvCount7Count), int(t.QijvCount8Count))
	t.dropRuleMap[viptypes.CostLevelRuleTypeChess] = timesData2

	//装备宝库
	var timesData3 []int
	baokuArr, err := coreutils.SplitAsIntArray(t.EquipBaoKuCount1Count)
	if err != nil {
		return template.NewTemplateFieldError("EquipBaoKuCount1Count", err)
	}
	for _, val := range baokuArr {
		timesData3 = append(timesData3, int(val))
	}
	t.dropRuleMap[viptypes.CostLevelRuleTypeEquipBaoKu] = timesData3

	// 宝库
	var timesData4 []int
	baokuArr2, err := coreutils.SplitAsIntArray(t.MaterialBaoKuCount1Count)
	if err != nil {
		return template.NewTemplateFieldError("MaterialBaoKuCount1Count", err)
	}
	for _, val := range baokuArr2 {
		timesData4 = append(timesData4, int(val))
	}
	t.dropRuleMap[viptypes.CostLevelRuleTypeMaterialBaoKu] = timesData4

	// 特戒宝库
	var timesData5 []int
	baokuArr3, err := coreutils.SplitAsIntArray(t.BaoKuBagCount1Count)
	if err != nil {
		return template.NewTemplateFieldError("BaoKuBagCount1Count", err)
	}
	for _, val := range baokuArr3 {
		timesData5 = append(timesData5, int(val))
	}
	t.dropRuleMap[viptypes.CostLevelRuleTypeRingBaoKu] = timesData5

	return nil
}

func (t *CostLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*CostLevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*CostLevelTemplate)
	}

	//验证 等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}
	//验证 升级条件
	err = validator.MinValidate(float64(t.NeedValue), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedValue)
		err = template.NewTemplateFieldError("NeedValue", err)
		return
	}

	//进阶系数规则
	for typ, data := range t.advancedRuleMap {
		err = validator.MinValidate(float64(data.GetMinRate()), float64(0), false)
		if err != nil {
			err = fmt.Errorf("功能模块:%d,min_count [%d] invalid", typ, data.GetMinRate())
			err = template.NewTemplateFieldError("min_count", err)
			return
		}
		err = validator.MinValidate(float64(data.GetMaxRate()), float64(0), false)
		if err != nil {
			err = fmt.Errorf("掉落类型:%d,max_count [%d] invalid", typ, data.GetMaxRate())
			err = template.NewTemplateFieldError("max_count", err)
			return
		}
	}

	//固定次数规则
	for typ, timesList := range t.dropRuleMap {
		for _, times := range timesList {
			err = validator.MinValidate(float64(times), float64(0), true)
			if err != nil {
				err = fmt.Errorf("掉落类型:%d,count_add [%d] invalid", typ, times)
				err = template.NewTemplateFieldError("count_add", err)
				return
			}
		}
	}

	return nil
}

func (t *CostLevelTemplate) FileName() string {
	return "tb_fufei_info.json"
}

func init() {
	template.Register((*CostLevelTemplate)(nil))
}
