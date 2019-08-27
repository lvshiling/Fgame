package export

import (
	model "fgame/fgame/tools/dbmerge/model"
	tool "fgame/fgame/tools/dbmerge/tool"
	"fmt"
)

// type allianceMySqlExport struct {
// }

// func (m *allianceMySqlExport) ExportTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error {

// }

func exportJieYi(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error {
	whereSql := fmt.Sprintf("EXISTS(SELECT 1 FROM t_jieyi WHERE jieYiId=t_jieyi.id and t_jieyi.serverId=%d)", p_serverid)
	return tool.BackUpTable(p_db, p_table, whereSql, p_savePath, p_dumpPath)
}

var (
	jieYiTableArray = []string{
		"t_jieyi_member",
	}
)

func init() {
	// export := &allianceMySqlExport{}
	for _, value := range jieYiTableArray {
		registerExport(value, IMysqlExportFunc(exportJieYi))
	}
}
