/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*SystemNeiguaOnline)(nil))
}

/*系统内挂在线日志*/
type SystemNeiguaOnline struct {
	SystemLogMsg `bson:",inline"`

	//在线人数
	OnlineNum int32 `json:"onlineNum"`
}

func (c *SystemNeiguaOnline) LogName() string {
	return "system_neigua_online"
}
