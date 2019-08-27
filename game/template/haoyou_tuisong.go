package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	friendtypes "fgame/fgame/game/friend/types"
	"fmt"
)

//好友推送配置
type FriendNoticeTemplate struct {
	*FriendNoticeTemplateVO
	noticeType friendtypes.FriendNoticeType
	rewItemMap map[int32]int32
}

func (t *FriendNoticeTemplate) TemplateId() int {
	return t.Id
}

func (t *FriendNoticeTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *FriendNoticeTemplate) GetFriendNoticeType() friendtypes.FriendNoticeType {
	return t.noticeType
}

func (t *FriendNoticeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	rewItemIdList, err := utils.SplitAsIntArray(t.ZhuheItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ZhuheItem)
		return template.NewTemplateFieldError("ZhuheItem", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.ZhuheItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ZhuheItemCount)
		return template.NewTemplateFieldError("ZhuheItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.ZhuheItem, t.ZhuheItemCount)
		return template.NewTemplateFieldError("ZhuheItem or ZhuheItemCount", err)
	}
	if len(rewItemIdList) > 0 {
		for index, itemId := range rewItemIdList {
			t.rewItemMap[itemId] = rewItemCountList[index]
		}
	}

	return nil
}

func (t *FriendNoticeTemplate) PatchAfterCheck() {
}

func (t *FriendNoticeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//推送类型
	t.noticeType = friendtypes.FriendNoticeType(t.Type)
	if !t.noticeType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//推送条件
	err = validator.MinValidate(float64(t.TiaoJian), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TiaoJian)
		return template.NewTemplateFieldError("TiaoJian", err)
	}

	// 反馈奖励1
	err = validator.MinValidate(float64(t.JiDanRewardExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiDanRewardExp)
		return template.NewTemplateFieldError("JiDanRewardExp", err)
	}

	// 反馈奖励2
	err = validator.MinValidate(float64(t.XianHuaRewardExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XianHuaRewardExp)
		return template.NewTemplateFieldError("XianHuaRewardExp", err)
	}

	// 反馈奖励1
	err = validator.MinValidate(float64(t.JiDanRewardExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JiDanRewardExpPoint)
		return template.NewTemplateFieldError("JiDanRewardExpPoint", err)
	}

	// 反馈奖励2
	err = validator.MinValidate(float64(t.XianHuaRewardExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XianHuaRewardExpPoint)
		return template.NewTemplateFieldError("XianHuaRewardExpPoint", err)
	}

	// 祝贺奖励
	for itemId, num := range t.rewItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			err = fmt.Errorf("[%s] invalid", t.ZhuheItem)
			return template.NewTemplateFieldError("ZhuheItem", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("ZhuheItemCount", err)
		}
	}

	return nil
}

func (edt *FriendNoticeTemplate) FileName() string {
	return "tb_haoyou_tuisong.json"
}

func init() {
	template.Register((*FriendNoticeTemplate)(nil))
}
