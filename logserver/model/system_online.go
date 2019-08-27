/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*SystemOnline)(nil))
}

/*系统在线日志*/
type SystemOnline struct {
	SystemLogMsg `bson:",inline"`

	//在线人数
	OnlineNum int32 `json:"onlineNum"`
}

func (c *SystemOnline) LogName() string {
	return "system_online"
}
