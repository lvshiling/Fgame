package export

import (
	model "fgame/fgame/tools/dbperson/model"
	"fmt"
)

type IMySqlExport interface {
	ExportTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error
}

type IMysqlExportFunc func(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error

func (f IMysqlExportFunc) ExportTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error {
	return f(p_db, p_serverid, p_table, p_savePath, p_dumpPath)
}

var (
	tableExportMap map[string]IMySqlExport = make(map[string]IMySqlExport)
)

func GetExport(p_tableName string) IMySqlExport {
	if value, ok := tableExportMap[p_tableName]; ok {
		return value
	}
	return nil
}

func registerExport(p_tableName string, p_export IMySqlExport) error {
	if _, ok := tableExportMap[p_tableName]; ok {
		return fmt.Errorf("%s is already registered", p_tableName)
	}
	tableExportMap[p_tableName] = p_export
	return nil
}

func init() {
}
