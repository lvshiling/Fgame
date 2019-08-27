/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*Exception)(nil))
}

/*异常*/
type Exception struct {
	SystemLogMsg `bson:",inline"`

	//异常内容
	Content string `json:"content"`
}

func (c *Exception) LogName() string {
	return "exception"
}
