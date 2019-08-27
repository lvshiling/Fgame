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

func exportAlliance(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_savePath string, p_dumpPath string) error {
	whereSql := fmt.Sprintf("EXISTS(SELECT 1 FROM t_alliance WHERE allianceId=t_alliance.id and t_alliance.serverId=%d)", p_serverid)
	return tool.BackUpTable(p_db, p_table, whereSql, p_savePath, p_dumpPath)
}

var (
	allianceTableArray = []string{
		"t_alliance_invitation",
		"t_alliance_join_apply",
		"t_alliance_log",
		"t_alliance_member",
		"t_alliance_depot",
		"t_alliance_boss",
	}
)

func init() {
	// export := &allianceMySqlExport{}
	for _, value := range allianceTableArray {
		registerExport(value, IMysqlExportFunc(exportAlliance))
	}
}
