/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*SystemScene)(nil))
}

/*系统场景日志*/
type SystemScene struct {
	SystemLogMsg `bson:",inline"`

	//内容
	Content string `json:"content"`
}

func (c *SystemScene) LogName() string {
	return "system_scene"
}
