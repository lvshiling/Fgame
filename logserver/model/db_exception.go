/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*DBException)(nil))
}

/*数据库异常*/
type DBException struct {
	SystemLogMsg `bson:",inline"`

	//数据
	TableName string `json:"tableName"`

	//数据
	Data string `json:"data"`

	//错误
	Error string `json:"err"`
}

func (c *DBException) LogName() string {
	return "db_exception"
}
