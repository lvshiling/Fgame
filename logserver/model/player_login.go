/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerLogin)(nil))
}

/*玩家登陆日志*/
type PlayerLogin struct {
	PlayerLogMsg `bson:",inline"`

	//ip1
	Ip1 string `json:"ip1"`
}

func (c *PlayerLogin) LogName() string {
	return "player_login"
}
