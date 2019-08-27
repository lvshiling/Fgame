package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
)

type WelfareItemData struct {
	ItemId     int32
	Num        int32
	ExpireFlag int32
}

//开服活动配置
type OpenserverActivityTemplate struct {
	*OpenserverActivityTemplateVO
	openType             welfaretypes.OpenActivityType       //活动类型
	opneSubType          welfaretypes.OpenActivitySubType    //活动子类型
	rewItemMap           map[int32]int32                     //
	emailRewItemMap      map[int32]int32                     //
	emailRewItemDataList []*WelfareItemData                  //（指定物品非时效性）
	rewItemDataList      []*WelfareItemData                  //（指定物品非时效性）
	expireType           inventorytypes.NewItemLimitTimeType //过期类型
	expireTime           int64                               //过期时间
}

func (t *OpenserverActivityTemplate) TemplateId() int {
	return t.Id
}

func (t *OpenserverActivityTemplate) GetExpireType() inventorytypes.NewItemLimitTimeType {
	return t.expireType
}

func (t *OpenserverActivityTemplate) GetExpireTime() int64 {
	return t.expireTime
}

// func (t *OpenserverActivityTemplate) GetOpenType() welfaretypes.OpenActivityType {
// 	return t.openType
// }

// func (t *OpenserverActivityTemplate) GetOpenSubType() welfaretypes.OpenActivitySubType {
// 	return t.opneSubType
// }

func (t *OpenserverActivityTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *OpenserverActivityTemplate) GetEmailRewItemMap() map[int32]int32 {
	return t.emailRewItemMap
}

func (t *OpenserverActivityTemplate) GetEmailRewItemMapWithRatio(ratio int32) map[int32]int32 {
	newMap := make(map[int32]int32)
	for k, v := range t.emailRewItemMap {
		newMap[k] += v * ratio
	}
	return newMap
}

func (t *OpenserverActivityTemplate) GetRewItemDataList() []*WelfareItemData {
	return t.rewItemDataList
}

func (t *OpenserverActivityTemplate) GetEmailRewItemDataList() []*WelfareItemData {
	return t.emailRewItemDataList
}

func (t *OpenserverActivityTemplate) GetEmailRewItemDataListWithRatio(ratio int32) []*WelfareItemData {
	var newRewItemDataList []*WelfareItemData
	for _, data := range t.emailRewItemDataList {
		d := &WelfareItemData{
			ItemId:     data.ItemId,
			Num:        data.Num * ratio,
			ExpireFlag: data.ExpireFlag,
		}
		newRewItemDataList = append(newRewItemDataList, d)
	}

	return newRewItemDataList
}

func (t *OpenserverActivityTemplate) GetRewItemDataListWithRatio(ratio int32) []*WelfareItemData {
	var newRewItemDataList []*WelfareItemData
	for _, data := range t.rewItemDataList {
		d := &WelfareItemData{
			ItemId:     data.ItemId,
			Num:        data.Num * ratio,
			ExpireFlag: data.ExpireFlag,
		}
		newRewItemDataList = append(newRewItemDataList, d)
	}

	return newRewItemDataList
}

func (t *OpenserverActivityTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	t.emailRewItemMap = make(map[int32]int32)
	//验证 rew_item_id
	rewItemIdList, err := utils.SplitAsIntArray(t.AwardItemId)
	if err != nil {
		err = fmt.Errorf("[%s] split invalid", t.AwardItemId)
		return template.NewTemplateFieldError("AwardItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.AwardItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.AwardItemCount)
		return template.NewTemplateFieldError("AwardItemCount", err)
	}
	expireFlagList, err := utils.SplitAsIntArray(t.ItemExpireFlag)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemExpireFlag)
		return template.NewTemplateFieldError("ItemExpireFlag", err)
	}

	//过期类型
	t.expireType = inventorytypes.NewItemLimitTimeType(t.TimeType)
	if !t.expireType.Valid() {
		err = fmt.Errorf("[%d] invalid ", t.TimeType)
		err = template.NewTemplateFieldError("TimeType", err)
		return
	}

	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid ", t.AwardItemId, t.AwardItemCount)
		err = template.NewTemplateFieldError("AwardItemId or AwardItemCount", err)
		return
	}

	if t.expireType != inventorytypes.NewItemLimitTimeTypeNone && len(rewItemIdList) != len(expireFlagList) {
		err = fmt.Errorf("[%s][%s] invalid ", t.AwardItemId, t.ItemExpireFlag)
		err = template.NewTemplateFieldError("AwardItemId or ItemExpireFlag", err)
		return
	}

	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] nil invalid", t.AwardItemId)
				return template.NewTemplateFieldError("AwardItemId", err)
			}

			err = validator.MinValidate(float64(rewItemCountList[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("AwardItemCount", err)
			}

			_, ok := t.rewItemMap[itemId]
			if !ok {
				t.rewItemMap[itemId] = rewItemCountList[index]
			} else {
				t.rewItemMap[itemId] += rewItemCountList[index]
			}

			_, ok = t.emailRewItemMap[itemId]
			if !ok {
				t.emailRewItemMap[itemId] = rewItemCountList[index]
			} else {
				t.emailRewItemMap[itemId] += rewItemCountList[index]
			}

			//
			expireFlag := int32(0)
			if t.expireType != inventorytypes.NewItemLimitTimeTypeNone {
				expireFlag = expireFlagList[index]
			}
			d := &WelfareItemData{
				ItemId:     itemId,
				Num:        rewItemCountList[index],
				ExpireFlag: expireFlag,
			}
			t.emailRewItemDataList = append(t.emailRewItemDataList, d)
			t.rewItemDataList = append(t.rewItemDataList, d)
		}
	}

	return nil
}

func (t *OpenserverActivityTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//检查 Type
	t.openType = welfaretypes.OpenActivityType(t.Type)
	if !t.openType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("type", err)
	}
	//检查 SubType
	t.opneSubType = welfaretypes.CreateOpenActivitySubType(t.openType, t.SubType)
	if t.opneSubType == nil || !t.opneSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("subType", err)
	}

	//验证 value1
	err = validator.MinValidate(float64(t.Value1), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Value1)
		err = template.NewTemplateFieldError("Value1", err)
		return
	}

	//验证 value2
	err = validator.MinValidate(float64(t.Value2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Value2)
		err = template.NewTemplateFieldError("Value2", err)
		return
	}

	//验证 value3
	err = validator.MinValidate(float64(t.Value3), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Value3)
		err = template.NewTemplateFieldError("Value3", err)
		return
	}

	//验证 value4
	err = validator.MinValidate(float64(t.Value4), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Value4)
		err = template.NewTemplateFieldError("Value4", err)
		return
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewSilver)
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGold)
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(t.RewGoldBind), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGoldBind)
		err = template.NewTemplateFieldError("RewGoldBind", err)
		return
	}

	switch t.expireType {
	case inventorytypes.NewItemLimitTimeTypeExpiredToday:
		{
			expire, err := timeutils.ParseDayOfHHMMSS(strconv.FormatInt(int64(t.LimitTime), 10))
			if err != nil {
				err = fmt.Errorf("[%d] invalid ", t.LimitTime)
				err = template.NewTemplateFieldError("LimitTime", err)
				return err
			}
			t.expireTime = expire
		}
	case inventorytypes.NewItemLimitTimeTypeExpiredAfterTime:
		{
			t.expireTime = int64(t.LimitTime) * int64(common.SECOND)
		}
	case inventorytypes.NewItemLimitTimeTypeExpiredDate:
		{
			expire, err := timeutils.ParseDayOfYYYYMMDDHHMMSS(strconv.FormatInt(int64(t.LimitTime), 10))
			if err != nil {
				err = fmt.Errorf("[%d] invalid ", t.LimitTime)
				err = template.NewTemplateFieldError("LimitTime", err)
				return err
			}
			t.expireTime = expire
		}
	default:
		break
	}

	return nil
}

func (t *OpenserverActivityTemplate) PatchAfterCheck() {
	if t.RewSilver != 0 {
		t.emailRewItemMap[constanttypes.SilverItem] = t.RewSilver

		d := &WelfareItemData{
			ItemId:     constanttypes.SilverItem,
			Num:        t.RewSilver,
			ExpireFlag: 0,
		}
		t.emailRewItemDataList = append(t.emailRewItemDataList, d)
	}
	if t.RewGoldBind != 0 {
		t.emailRewItemMap[constanttypes.BindGoldItem] = t.RewGoldBind
		d := &WelfareItemData{
			ItemId:     constanttypes.BindGoldItem,
			Num:        t.RewGoldBind,
			ExpireFlag: 0,
		}
		t.emailRewItemDataList = append(t.emailRewItemDataList, d)
	}
	if t.RewGold != 0 {
		t.emailRewItemMap[constanttypes.GoldItem] = t.RewGold
		d := &WelfareItemData{
			ItemId:     constanttypes.GoldItem,
			Num:        t.RewGold,
			ExpireFlag: 0,
		}
		t.emailRewItemDataList = append(t.emailRewItemDataList, d)
	}

}

func (t *OpenserverActivityTemplate) FileName() string {
	return "tb_openserver_activity.json"
}

func init() {
	template.Register((*OpenserverActivityTemplate)(nil))
}
