package api

import (
	"context"
	"fgame/fgame/cross/player/player"
	settimepb "fgame/fgame/cross/settime/pb"
	"fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/timeutils"

	"time"
)

//服务器服务
type SetTimeServer struct {
}

//设置时间
func (ss *SetTimeServer) SetTime(ctx context.Context, req *settimepb.SetTimeRequest) (res *settimepb.SetTimeResponse, err error) {
	currentTime := req.GetTime()
	t := timeutils.MillisecondToTime(currentTime)
	now := time.Now()
	offTime := int64(t.Sub(now)) / int64(time.Millisecond)
	global.GetGame().GetTimeService().SetOffTime(offTime)
	nowInt := global.GetGame().GetTimeService().Now()
	res = &settimepb.SetTimeResponse{}
	res.Time = nowInt

	scGetTime := pbutil.BuildSCGetTime(nowInt)
	player.GetOnlinePlayerManager().BroadcastMsg(scGetTime)
	return
}

func NewSetTimeServer() *SetTimeServer {
	ss := &SetTimeServer{}
	return ss
}
