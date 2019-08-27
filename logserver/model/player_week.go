/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerWeek)(nil))
}

/*周卡*/
type PlayerWeek struct {
	PlayerLogMsg `bson:",inline"`

	//上次过期时间
	LastExpireTime int64 `json:"lastExpireTime"`

	//周卡类型
	WeekType int32 `json:"weekType"`

	//原因编号
	Reason int32 `json:"reason"`

	//原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerWeek) LogName() string {
	return "player_week"
}
