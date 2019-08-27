package player

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"time"
)

const (
	taskRefreshTime = time.Second * 20
)

//运营活动刷新数据
type RefreshActivityTask struct {
	pl player.Player
}

func (t *RefreshActivityTask) Run() {
	welfareManager := t.pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*PlayerWelfareManager)
	timeList := welfaretemplate.GetWelfareTemplateService().GetAllActivityTimeTemplate()
	for _, timeTemp := range timeList {
		welfareManager.RefreshActivityDataByGroupId(timeTemp.Group)
		// welfareManager.RefreshActivityData(timeTemp.GetOpenType(), timeTemp.GetOpenSubType())
	}

}

//间隔时间
func (t *RefreshActivityTask) ElapseTime() time.Duration {
	return taskRefreshTime
}

func CreateRefreshActivityTask(pl player.Player) *RefreshActivityTask {
	t := &RefreshActivityTask{
		pl: pl,
	}
	return t
}
