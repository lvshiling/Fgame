package delete

import (
	model "fgame/fgame/tools/dbdelete/model"
	tool "fgame/fgame/tools/dbdelete/tool"
	"fmt"
)

// type allianceMySqlExport struct {
// }

// func (m *allianceMySqlExport) ExportTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error {

// }

type jieYiMySqlDelete struct {
}

func (d *jieYiMySqlDelete) DeleteTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_mysqlPath string) error {
	whereSql := fmt.Sprintf("EXISTS(SELECT 1 FROM t_jieyi WHERE jieYiId=t_jieyi.id and t_jieyi.serverId=%d)", p_serverid)
	return tool.DeleteTable(p_db, p_table, whereSql, p_mysqlPath)
}

func (d *jieYiMySqlDelete) IsServer() bool {
	return false
}

var (
	jieYiDel        = &jieYiMySqlDelete{}
	jieYiTableArray = []string{
		"t_jieyi_member",
	}
)

func init() {
	// export := &allianceMySqlExport{}
	for _, value := range jieYiTableArray {
		registerDelete(value, jieYiDel)
	}
}
