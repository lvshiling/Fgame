package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//转生冲刺
type ZhuanShengInfo struct {
	ZhuanSheng int32   `json:zhuanSheng`  //转生数
	RewRecord  []int32 `json:"rewRecord"` //领取记录
	IsMail     bool    `json:"isMail"`    //是否邮件
}

func (info *ZhuanShengInfo) UpdateZhaunSheng(newZhuanSheng int32) {
	if info.ZhuanSheng >= newZhuanSheng {
		return
	}

	info.ZhuanSheng = newZhuanSheng
}

func (info *ZhuanShengInfo) IsCanReceiveRewards(needZhuanSheng int32) bool {
	//条件
	if info.ZhuanSheng < needZhuanSheng {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needZhuanSheng {
			return false
		}
	}

	return true
}

func (info *ZhuanShengInfo) AddRecord(needZhuanSheng int32) {
	info.RewRecord = append(info.RewRecord, needZhuanSheng)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeZhaunSheng, (*ZhuanShengInfo)(nil))
}
