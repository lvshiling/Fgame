package delete

import (
	model "fgame/fgame/tools/dbdelete/model"
	"fmt"
)

type IMySqlDelete interface {
	DeleteTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_mysqlPath string) error
	IsServer() bool
}

// type IMysqlDeleteFunc func(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_mysqlPath string) error

// func (f IMysqlDeleteFunc) DeleteTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_mysqlPath string) error {
// 	return f(p_db, p_serverid, p_table, p_mysqlPath)
// }

var (
	tableDeleteMap map[string]IMySqlDelete = make(map[string]IMySqlDelete)
)

func GetDelete(p_tableName string) IMySqlDelete {
	if value, ok := tableDeleteMap[p_tableName]; ok {
		return value
	}
	return nil
}

func registerDelete(p_tableName string, p_export IMySqlDelete) error {
	if _, ok := tableDeleteMap[p_tableName]; ok {
		return fmt.Errorf("%s is already registered", p_tableName)
	}
	tableDeleteMap[p_tableName] = p_export
	return nil
}

func init() {
}
