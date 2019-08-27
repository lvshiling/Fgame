package merge

import (
	model "fgame/fgame/tools/dbmerge/model"
	tool "fgame/fgame/tools/dbmerge/tool"
)

type IMergeService interface {
	BackUpDb(db *model.DBConfigInfo, p_savePath string) error
	BackUpTable(db *model.DBConfigInfo, p_table string, p_where string, p_savePath string) error

	ImportSqlPath(db *model.DBConfigInfo, p_sqlPath string) error

	DumpPath() string
	MySqlPath() string
}

type mergeService struct {
	dumpPath  string
	mysqlPath string
}

func (m *mergeService) BackUpDb(db *model.DBConfigInfo, p_savePath string) error {

	return tool.BackUpDb(db, p_savePath, m.dumpPath)
}

func (m *mergeService) BackUpTable(db *model.DBConfigInfo, p_table string, p_where string, p_savePath string) error {

	return tool.BackUpTable(db, p_table, p_where, p_savePath, m.dumpPath)
}

func (m *mergeService) ImportSqlPath(db *model.DBConfigInfo, p_sqlPath string) error {
	return tool.ImportSqlPath(db, p_sqlPath, m.mysqlPath)
}

func (m *mergeService) DumpPath() string {
	return m.dumpPath
}

func (m *mergeService) MySqlPath() string {
	return m.mysqlPath
}

func NewMergeService(p_dumpPath string, p_mysqlPath string) IMergeService {
	rst := &mergeService{
		dumpPath:  p_dumpPath,
		mysqlPath: p_mysqlPath,
	}
	return rst
}
