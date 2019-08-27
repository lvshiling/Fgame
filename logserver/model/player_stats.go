/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerStats)(nil))
}

/*玩家点击统计*/
type PlayerStats struct {
	PlayerLogMsg `bson:",inline"`

	//统计内容
	Stats string `json:"stats"`
}

func (c *PlayerStats) LogName() string {
	return "player_stats"
}
