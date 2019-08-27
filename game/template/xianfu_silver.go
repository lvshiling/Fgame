package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"fmt"
)

func init() {
	template.Register((*XianFuSilverTemplate)(nil))
}

//秘境仙府-银两副本数据配置
type XianFuSilverTemplate struct {
	*XianFuSilverTemplateVO
	nextTemplate         *XianFuSilverTemplate
	saodangItemMap       map[int32]int32
	saodangRewardItemMap map[int32]int32
	saodangRewardDropArr []int32
	challengeRewardsMap  map[int32]int32
	mapTemplate          *MapTemplate
}

func (t *XianFuSilverTemplate) TemplateId() int {
	return t.Id
}
func (t *XianFuSilverTemplate) FileName() string {
	return "tb_xianfu_silver.json"
}

//检查有效性
func (t *XianFuSilverTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证：BossID
	bossTmpObj := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
	if bossTmpObj == nil {
		return template.NewTemplateFieldError("BossId", fmt.Errorf("[%d] invalid", t.BossId))
	}
	//验证：升级所需银两
	err = validator.MinValidate(float64(t.UpgradeYinliang), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UpgradeYinliang", err)
	}
	//验证：升级所需元宝
	err = validator.MinValidate(float64(t.UpgradeGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UpgradeGold", err)
	}
	//验证：升级所需绑定元宝
	err = validator.MinValidate(float64(t.UpgradeBindGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UpgradeBindGold", err)
	}
	//验证：升级所需物品ID
	if t.UpgradeItemId != 0 {
		itemTmpObj := template.GetTemplateService().Get(int(t.UpgradeItemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("UpgradeItemId", fmt.Errorf("[%d] invalid", t.UpgradeItemId))
		}
		err = validator.MinValidate(float64(t.UpgradeItemId), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("UpgradeItemId", err)
		}
	}
	//验证：升级所需时间
	err = validator.MinValidate(float64(t.UpgradeTime), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UpgradeTime", err)
	}
	//验证：加速所需元宝
	err = validator.MinValidate(float64(t.UpgradeBindGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UpgradeBindGold", err)
	}
	//验证：扫荡所需元宝
	err = validator.MinValidate(float64(t.SaodangNeedGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("SaodangNeedGold", err)
	}

	//验证：扫荡奖励经验值
	err = validator.MinValidate(float64(t.RawExp), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawExp", err)
	}
	//验证：扫荡奖励经验点
	err = validator.MinValidate(float64(t.RawExpPoint), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawExpPoint", err)
	}
	//验证：扫荡奖励银两
	err = validator.MinValidate(float64(t.RawSilver), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawSilver", err)
	}
	//验证：扫荡奖励元宝
	err = validator.MinValidate(float64(t.RawGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawGold", err)
	}
	//验证：扫荡奖励绑定元宝
	err = validator.MinValidate(float64(t.RawBindGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("RawBindGold", err)
	}

	//验证：进入副本所需物品ID
	itemTmpObj := template.GetTemplateService().Get(int(t.NeedItemId), (*ItemTemplate)(nil))
	if itemTmpObj == nil {
		return template.NewTemplateFieldError("NeedItemId", fmt.Errorf("[%d] invalid", t.NeedItemId))
	}
	//验证：进入副本所需物品数量
	err = validator.MinValidate(float64(t.NeedItemCount), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedItemCount", err)
	}
	//验证：每日免费的次数
	err = validator.MinValidate(float64(t.Free), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("Free", err)
	}
	//验证：下一级ID
	if t.NextId != 0 {
		diff := t.NextId - t.Id
		if diff != 1 {
			return template.NewTemplateFieldError("NextId", fmt.Errorf("[%d] invalid", t.NextId))
		}
		xianfuSilverTmpObj := template.GetTemplateService().Get(t.NextId, (*XianFuSilverTemplate)(nil))
		if xianfuSilverTmpObj == nil {
			return template.NewTemplateFieldError("NextId", fmt.Errorf("[%d] invalid", t.NextId))
		}

		t.nextTemplate = xianfuSilverTmpObj.(*XianFuSilverTemplate)
	}

	return nil
}

//组合成需要的数据
func (t *XianFuSilverTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 map_id
	tempMapTemplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.mapTemplate = tempMapTemplate.(*MapTemplate)

	t.saodangItemMap = make(map[int32]int32)

	//验证: 扫荡所需的物品ID,逗号隔开
	//验证：扫荡所需的物品数量
	intSaodangNeedItemIdArr, err := utils.SplitAsIntArray(t.SaodangNeedItemId)
	if err != nil {
		return template.NewTemplateFieldError("SaodangNeedItemId", fmt.Errorf("[%s] invalid", t.SaodangNeedItemId))
	}
	intSaodangNeedItemCountArr, err := utils.SplitAsIntArray(t.SaodangNeedItemCount)
	if err != nil {
		return template.NewTemplateFieldError("SaodangNeedItemCount", fmt.Errorf("[%s] invalid", t.SaodangNeedItemCount))
	}
	if len(intSaodangNeedItemIdArr) != len(intSaodangNeedItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.SaodangNeedItemId, t.SaodangNeedItemCount)
		return template.NewTemplateFieldError("SaodangNeedItemId or SaodangNeedItemCount", err)
	}

	if len(intSaodangNeedItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intSaodangNeedItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("SaodangNeedItemId", fmt.Errorf("[%s] invalid", t.SaodangNeedItemId))
			}

			err = validator.MinValidate(float64(intSaodangNeedItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("SaodangNeedItemCount", err)
			}

			t.saodangItemMap[itemId] = intSaodangNeedItemCountArr[index]
		}
	}

	t.saodangRewardItemMap = make(map[int32]int32)

	//验证：扫荡奖励物品ID,逗号隔开
	//验证：扫荡奖励物品数量
	intRawItemIdArr, err := utils.SplitAsIntArray(t.RawItemId)
	if err != nil {
		return template.NewTemplateFieldError("RawItemId", fmt.Errorf("[%s] invalid", t.RawItemId))
	}
	intRawItemCountArr, err := utils.SplitAsIntArray(t.RawItemCount)
	if err != nil {
		return template.NewTemplateFieldError("RawItemCount", fmt.Errorf("[%s] invalid", t.RawItemCount))
	}
	if len(intRawItemIdArr) != len(intRawItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.RawItemId, t.RawItemCount)
		return template.NewTemplateFieldError("RawItemId or RawItemCount", err)
	}
	if len(intRawItemIdArr) > 0 {
		for index, itemId := range intRawItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("RawItemId", fmt.Errorf("[%s] invalid", t.RawItemId))
			}

			err = validator.MinValidate(float64(intRawItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RawItemCount", err)
			}

			//组合数据
			t.saodangRewardItemMap[itemId] = intRawItemCountArr[index]
		}
	}

	//验证：挑战通关奖励物品
	t.challengeRewardsMap = make(map[int32]int32)

	intGetRewItemIdArr, err := utils.SplitAsIntArray(t.GetItemId)
	if err != nil {
		return template.NewTemplateFieldError("GetItemId", fmt.Errorf("[%s] invalid", t.GetItemId))
	}
	intGetRawItemCountArr, err := utils.SplitAsIntArray(t.GetItemCount)
	if err != nil {
		return template.NewTemplateFieldError("GetItemCount", fmt.Errorf("[%s] invalid", t.GetItemCount))
	}
	if len(intGetRewItemIdArr) != len(intGetRawItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.GetItemId, t.GetItemCount)
		return template.NewTemplateFieldError("GetItemId or GetItemCount", err)
	}
	if len(intGetRewItemIdArr) > 0 {
		for index, itemId := range intGetRewItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("GetItemId", fmt.Errorf("[%s] invalid", t.GetItemId))
			}

			err = validator.MinValidate(float64(intGetRawItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("GetItemCount", err)
			}

			//组合数据
			t.challengeRewardsMap[itemId] = intGetRawItemCountArr[index]
		}
	}

	//验证：扫荡奖励掉落包ID，逗号隔开
	rawDropIdArr, err := utils.SplitAsIntArray(t.RawDropId)
	if err != nil {
		return template.NewTemplateFieldError("RawDropId", fmt.Errorf("[%s] invalid", t.RawDropId))
	}
	t.saodangRewardDropArr = make([]int32, 0, len(rawDropIdArr))
	if len(rawDropIdArr) > 0 {
		for _, dropId := range rawDropIdArr {

			itemTmpObj := template.GetTemplateService().Get(int(dropId), (*DropTemplate)(nil))
			if itemTmpObj == nil {
				return template.NewTemplateFieldError("RawDropId", fmt.Errorf("[%s] invalid", t.RawDropId))
			}

			t.saodangRewardDropArr = append(t.saodangRewardDropArr, dropId)
		}
	}

	return nil
}

//检验后组合
func (t *XianFuSilverTemplate) PatchAfterCheck() {
}

func (t *XianFuSilverTemplate) GetXianFuType() xianfutypes.XianfuType {
	return xianfutypes.XianfuTypeSilver
}

func (t *XianFuSilverTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *XianFuSilverTemplate) GetBossId() int32 {
	return t.BossId
}

//获取扫荡所需物品
func (t *XianFuSilverTemplate) GetSaodangItemMap(saoDangNum int32) map[int32]int32 {
	if saoDangNum > 1 {
		newMap := make(map[int32]int32)
		for itemId, num := range t.saodangItemMap {
			newMap[itemId] = num * saoDangNum
		}

		return newMap
	}
	return t.saodangItemMap
}

//获取扫荡奖励物品
func (t *XianFuSilverTemplate) GetSaodangRewardItemMap(saoDangNum int32) map[int32]int32 {
	if saoDangNum > 1 {
		newMap := make(map[int32]int32)
		for itemId, num := range t.saodangItemMap {
			newMap[itemId] = num * saoDangNum
		}

		return newMap
	}
	return t.saodangRewardItemMap
}

//获取扫荡掉落物品
func (t *XianFuSilverTemplate) GetSaodangRewardDropArr() []int32 {
	return t.saodangRewardDropArr
}

//获取通关奖励物品
func (t *XianFuSilverTemplate) GetChallengeRewardsItemMap() map[int32]int32 {
	return t.challengeRewardsMap
}

func (t *XianFuSilverTemplate) GetUpgradeTime() int64 {
	return int64(t.UpgradeTime) * common.MILL_METER
}

func (t *XianFuSilverTemplate) GetUpgradeGold() int32 {
	return t.UpgradeGold
}
func (t *XianFuSilverTemplate) GetUpgradeBindGold() int32 {
	return t.UpgradeBindGold
}
func (t *XianFuSilverTemplate) GetUpgradeYinliang() int64 {
	return t.UpgradeYinliang
}
func (t *XianFuSilverTemplate) GetUpgradeItemId() int32 {
	return t.UpgradeItemId
}
func (t *XianFuSilverTemplate) GetUpgradeItemNum() int32 {
	return t.UpgradeItemNum
}
func (t *XianFuSilverTemplate) GetSpeedUpNeedGold() float64 {
	return t.SpeedUpNeedGold
}
func (t *XianFuSilverTemplate) GetRawExp() int64 {
	return t.RawExp
}
func (t *XianFuSilverTemplate) GetRawExpPoint() int64 {
	return t.RawExpPoint
}
func (t *XianFuSilverTemplate) GetRawGold() int32 {
	return t.RawGold
}
func (t *XianFuSilverTemplate) GetRawBindGold() int32 {
	return t.RawBindGold
}
func (t *XianFuSilverTemplate) GetRawSilver() int64 {
	return t.RawSilver
}
func (t *XianFuSilverTemplate) GetNeedItemId() int32 {
	return t.NeedItemId
}
func (t *XianFuSilverTemplate) GetNeedItemCount() int32 {
	return t.NeedItemCount
}
func (t *XianFuSilverTemplate) GetNextId() int32 {
	return int32(t.NextId)
}
func (t *XianFuSilverTemplate) GetFree() int32 {
	return t.Free
}
func (t *XianFuSilverTemplate) GetSaodangNeedGold() int32 {
	return t.SaodangNeedGold
}

func (t *XianFuSilverTemplate) GetGroupLimit() int32 {
	return 0
}
