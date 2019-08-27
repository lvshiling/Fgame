package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
)

//开服活动时间配置
type OpenserverTimeTemplate struct {
	*OpenserverTimeTemplateVO
	openType     welfaretypes.OpenActivityType
	opneSubType  welfaretypes.OpenActivitySubType
	openTimeType welfaretypes.OpenTimeType
	beginTime    int64
	endTime      int64
	funcType     funcopentypes.FuncOpenType
	mailRewItems map[int32]int32
	groupList    []int32
}

func (t *OpenserverTimeTemplate) GetMailRewItems() map[int32]int32 {
	return t.mailRewItems
}

func (t *OpenserverTimeTemplate) IsChuanYin() bool {
	if t.IsChuanyin == int32(1) {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsRelationToGroup(groupId int32) bool {
	return coreutils.ContainInt32(t.groupList, groupId)
}

func (t *OpenserverTimeTemplate) IsMergeXunHuan() bool {
	return t.IsMergeCircle != 0
}

func (t *OpenserverTimeTemplate) IsXunHuan() bool {
	return t.IsCircle != 0
}

//
func (t *OpenserverTimeTemplate) IsTempRank() bool {
	if t.IsCircle != 0 {
		return true
	}
	if t.IsMergeCircle != 0 {
		return true
	}
	if t.openTimeType == welfaretypes.OpenTimeTypeOpenActivityNoMerge {
		return true
	}
	if t.openTimeType == welfaretypes.OpenTimeTypeWeek {
		return true
	}
	if t.openTimeType == welfaretypes.OpenTimeTypeMonth {
		return true
	}
	if t.CloseOpenDay > 0 {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) GetRelationToGroupList() []int32 {
	return t.groupList
}

func (t *OpenserverTimeTemplate) IsAllianceCheer() bool {
	if t.openType == welfaretypes.OpenActivityTypeAlliance && t.opneSubType == welfaretypes.OpenActivityAllianceSubTypeAlliance {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsArenaPvpCheer() bool {
	if t.openType == welfaretypes.OpenActivityTypeAlliance && t.opneSubType == welfaretypes.OpenActivityAllianceSubTypeNewWuLian {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsRankCharge() bool {
	if t.openType == welfaretypes.OpenActivityTypeRank && t.opneSubType == welfaretypes.OpenActivityRankSubTypeCharge {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsRankCost() bool {
	if t.openType == welfaretypes.OpenActivityTypeRank && t.opneSubType == welfaretypes.OpenActivityRankSubTypeCost {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsRankCharm() bool {
	if t.openType == welfaretypes.OpenActivityTypeRank && t.opneSubType == welfaretypes.OpenActivityRankSubTypeCharm {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsRankNunber() bool {
	if t.openType == welfaretypes.OpenActivityTypeRank && t.opneSubType == welfaretypes.OpenActivityRankSubTypeNumber {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) IsRankMarryDevelop() bool {
	if t.openType == welfaretypes.OpenActivityTypeRank && t.opneSubType == welfaretypes.OpenActivityRankSubTypeMarryDevelop {
		return true
	}
	return false
}

func (t *OpenserverTimeTemplate) GetOpenType() welfaretypes.OpenActivityType {
	return t.openType
}

func (t *OpenserverTimeTemplate) GetOpenSubType() welfaretypes.OpenActivitySubType {
	return t.opneSubType
}

func (t *OpenserverTimeTemplate) GetOpenFuncType() funcopentypes.FuncOpenType {
	return t.funcType
}

func (t *OpenserverTimeTemplate) TemplateId() int {
	return t.Id
}

func (t *OpenserverTimeTemplate) GetOpenTimeType() welfaretypes.OpenTimeType {
	return t.openTimeType
}

func (t *OpenserverTimeTemplate) GetBeginTime(now, openTime int64) (int64, error) {
	time := t.beginTime

	switch t.openTimeType {
	case welfaretypes.OpenTimeTypeOpenActivity,
		welfaretypes.OpenTimeTypeOpenActivityNoMerge,
		welfaretypes.OpenTimeTypeMerge:
		{
			openDayTime, err := timeutils.BeginOfNow(openTime)
			if err != nil {
				return 0, err
			}
			time += openDayTime
			break
		}
	case welfaretypes.OpenTimeTypeXunHuan,
		welfaretypes.OpenTimeTypeMergeXunHuan:
		{
			beginDay, err := timeutils.BeginOfNow(now)
			if err != nil {
				return 0, err
			}
			time += beginDay
			break
		}
	case welfaretypes.OpenTimeTypeWeek:
		{
			beginDay, err := timeutils.BeginOfWeekOfMillisecond(now)
			if err != nil {
				return 0, err
			}
			time += beginDay
			break
		}
	case welfaretypes.OpenTimeTypeMonth:
		{
			beginDay, err := timeutils.BeginOfMonthOfMillisecond(now)
			if err != nil {
				return 0, err
			}
			time += beginDay
			break
		}
	}

	return time, nil
}

func (t *OpenserverTimeTemplate) GetEndTime(now, openTime int64) (int64, error) {
	time := t.endTime

	switch t.openTimeType {
	case welfaretypes.OpenTimeTypeOpenActivity,
		welfaretypes.OpenTimeTypeOpenActivityNoMerge,
		welfaretypes.OpenTimeTypeMerge:
		{
			openDayTime, err := timeutils.BeginOfNow(openTime)
			if err != nil {
				return 0, err
			}
			time += openDayTime
			break
		}
	case welfaretypes.OpenTimeTypeXunHuan,
		welfaretypes.OpenTimeTypeMergeXunHuan:
		{
			beginDay, err := timeutils.BeginOfNow(now)
			if err != nil {
				return 0, err
			}
			time += beginDay
			break
		}
	case welfaretypes.OpenTimeTypeWeek:
		{
			beginDay, err := timeutils.BeginOfWeekOfMillisecond(now)
			if err != nil {
				return 0, err
			}
			time += beginDay
			break
		}
	case welfaretypes.OpenTimeTypeMonth:
		{
			beginDay, err := timeutils.BeginOfMonthOfMillisecond(now)
			if err != nil {
				return 0, err
			}
			time += beginDay
			break
		}
	}

	return time, nil
}

func (t *OpenserverTimeTemplate) IsOnTime(now, openTime int64) (flag bool, err error) {
	beginTime, err := t.GetBeginTime(now, openTime)
	if err != nil {
		return
	}
	endTime, err := t.GetEndTime(now, openTime)
	if err != nil {
		return
	}
	if beginTime <= now && now <= endTime {
		return true, nil
	}
	return false, nil
}

func (t *OpenserverTimeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 开启邮件奖励
	t.mailRewItems = make(map[int32]int32)
	intMailRewItemIdArr, err := coreutils.SplitAsIntArray(t.MailRewItem)
	if err != nil {
		return template.NewTemplateFieldError("MailRewItem", fmt.Errorf("[%s] invalid", t.MailRewItem))
	}
	intMailRewItemCountArr, err := coreutils.SplitAsIntArray(t.MailRewItemCount)
	if err != nil {
		return template.NewTemplateFieldError("MailRewItemCount", fmt.Errorf("[%s] invalid", t.MailRewItemCount))
	}
	if len(intMailRewItemIdArr) != len(intMailRewItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.MailRewItem, t.MailRewItemCount)
		return template.NewTemplateFieldError("MailRewItem or MailRewItemCount", err)
	}
	if len(intMailRewItemIdArr) > 0 {
		//组合数据
		for index, itemId := range intMailRewItemIdArr {
			t.mailRewItems[itemId] = intMailRewItemCountArr[index]
		}
	}

	return nil
}

func (t *OpenserverTimeTemplate) Check() (err error) {
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

	//检查 timeType
	t.openTimeType = welfaretypes.OpenTimeType(t.TimeType)
	if !t.openTimeType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.TimeType)
		return template.NewTemplateFieldError("TimeType", err)
	}

	//指定时间才能设置
	if t.CloseOpenDay != 0 && t.openTimeType != welfaretypes.OpenTimeTypeSchedule {
		err = fmt.Errorf("[%d] invalid", t.CloseOpenDay)
		return template.NewTemplateFieldError("CloseOpenDay", err)
	}

	switch t.openTimeType {
	case welfaretypes.OpenTimeTypeNotTimeliness:
		{
			// 无时效，时间配置为0
			//验证 value1
			if t.Value1 != 0 {
				err = fmt.Errorf("[%d] invalid", t.Value1)
				err = template.NewTemplateFieldError("Value1", err)
				return
			}

			//验证 value2
			if t.Value2 != 0 {
				err = fmt.Errorf("[%d] invalid", t.Value2)
				err = template.NewTemplateFieldError("Value2", err)
				return
			}

			// iscircle
			if t.IsXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsCircle)
				err = template.NewTemplateFieldError("IsCircle", err)
				return
			}
			if t.IsMergeXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsMergeCircle)
				err = template.NewTemplateFieldError("IsMergeCircle", err)
				return
			}
		}
	case welfaretypes.OpenTimeTypeOpenActivity,
		welfaretypes.OpenTimeTypeOpenActivityNoMerge:
		{
			//验证 value1
			err = validator.MinValidate(float64(t.Value1), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", t.Value1)
				err = template.NewTemplateFieldError("Value1", err)
				return
			}
			starDay := t.Value1 - 1
			t.beginTime = int64(common.DAY) * int64(starDay)

			//验证 value2
			err = validator.MinValidate(float64(t.Value2), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", t.Value2)
				err = template.NewTemplateFieldError("Value2", err)
				return
			}
			// iscircle
			if t.IsXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsCircle)
				err = template.NewTemplateFieldError("IsCircle", err)
				return
			}
			if t.IsMergeXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsMergeCircle)
				err = template.NewTemplateFieldError("IsMergeCircle", err)
				return
			}

			endDay := t.Value2 - 1
			t.endTime = int64(common.DAY) * int64(endDay)
		}
	case welfaretypes.OpenTimeTypeMerge:
		{
			//验证 value1
			err = validator.MinValidate(float64(t.Value1), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", t.Value1)
				err = template.NewTemplateFieldError("Value1", err)
				return
			}
			starDay := t.Value1 - 1
			t.beginTime = int64(common.DAY) * int64(starDay)

			//验证 value2
			err = validator.MinValidate(float64(t.Value2), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", t.Value2)
				err = template.NewTemplateFieldError("Value2", err)
				return
			}
			// iscircle
			if t.IsXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsCircle)
				err = template.NewTemplateFieldError("IsCircle", err)
				return
			}
			if t.IsMergeXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsMergeCircle)
				err = template.NewTemplateFieldError("IsMergeCircle", err)
				return
			}

			endDay := t.Value2 - 1
			t.endTime = int64(common.DAY) * int64(endDay)
		}
	case welfaretypes.OpenTimeTypeSchedule:
		{
			//活动开始时间 value1
			beginTime, err := timeutils.ParseDayOfYYYYMMDDHHMM(strconv.FormatInt(int64(t.Value1), 10))
			if err != nil {
				return template.NewTemplateFieldError("Value1", fmt.Errorf("[%d] invalid", t.Value1))
			}
			t.beginTime = beginTime

			//活动结束时间 value2
			endTime, err := timeutils.ParseDayOfYYYYMMDDHHMM(strconv.FormatInt(int64(t.Value2), 10))
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.Value2)
				return template.NewTemplateFieldError("Value2", err)
			}
			// iscircle
			if t.IsXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsCircle)
				err = template.NewTemplateFieldError("IsCircle", err)
				return err
			}
			if t.IsMergeXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsMergeCircle)
				err = template.NewTemplateFieldError("IsMergeCircle", err)
				return err
			}

			t.endTime = endTime
		}
	case welfaretypes.OpenTimeTypeXunHuan:
		{
			//验证 value1
			err = validator.MinValidate(float64(t.Value1), float64(1), true)
			if err != nil {
				err = fmt.Errorf("[%d] invalid", t.Value1)
				err = template.NewTemplateFieldError("Value1", err)
				return
			}

			// 只能配置一天
			if t.Value1 != 1 {
				err = fmt.Errorf("[%d] invalid", t.Value1)
				err = template.NewTemplateFieldError("Value1", err)
				return
			}

			// iscircle
			if !t.IsXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsCircle)
				err = template.NewTemplateFieldError("IsCircle", err)
				return
			}
			if t.IsMergeXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsMergeCircle)
				err = template.NewTemplateFieldError("IsMergeCircle", err)
				return
			}

			t.endTime = int64(common.DAY) * int64(t.Value1)
		}
	case welfaretypes.OpenTimeTypeMergeXunHuan:
		{

			// 只能配置一天
			if t.Value1 != 1 {
				err = fmt.Errorf("[%d] invalid", t.Value1)
				err = template.NewTemplateFieldError("Value1", err)
				return
			}

			// isMergeCircle
			if !t.IsMergeXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsMergeCircle)
				err = template.NewTemplateFieldError("IsMergeCircle", err)
				return
			}
			if t.IsXunHuan() {
				err = fmt.Errorf("[%d] invalid", t.IsCircle)
				err = template.NewTemplateFieldError("IsCircle", err)
				return
			}

			t.endTime = int64(common.DAY) * int64(t.Value1)
		}
	case welfaretypes.OpenTimeTypeWeek:
		{
			// 只能配置一天
			//TODO:zrc 验证周一到周日
			t.beginTime = int64(common.DAY) * int64(t.Value1-1)

			if t.Value2 != 1 {
				err = fmt.Errorf("[%d] invalid", t.Value2)
				err = template.NewTemplateFieldError("Value2", err)
				return
			}
			t.endTime = int64(common.DAY)*int64(t.Value2) + t.beginTime
		}
	case welfaretypes.OpenTimeTypeMonth:
		{
			// 只能配置一天
			//TODO:zrc 验证月初
			t.beginTime = int64(common.DAY) * int64(t.Value1-1)

			if t.Value2 != 1 {
				err = fmt.Errorf("[%d] invalid", t.Value2)
				err = template.NewTemplateFieldError("Value2", err)
				return
			}
			t.endTime = int64(common.DAY)*int64(t.Value2) + t.beginTime
		}
	default:
		{
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
		}
	}

	// 功能开启
	t.funcType = funcopentypes.FuncOpenType(t.OpenId)
	if !t.funcType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.OpenId)
		err = template.NewTemplateFieldError("OpenId", err)
		return
	}

	//开启邮件奖励
	for itemId, num := range t.mailRewItems {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("MailRewItem", fmt.Errorf("[%s] invalid", t.MailRewItem))
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("MailRewItemCount", err)
		}
	}

	//关联group
	t.groupList, err = coreutils.SplitAsIntArray(t.RelatedActivity)
	if err != nil {
		err = template.NewTemplateFieldError("RelatedActivity", err)
		return
	}

	return nil
}

func (t *OpenserverTimeTemplate) PatchAfterCheck() {

}

func (t *OpenserverTimeTemplate) FileName() string {
	return "tb_openserver_time.json"
}

func init() {
	template.Register((*OpenserverTimeTemplate)(nil))
}
