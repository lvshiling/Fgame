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

type allianceMySqlDelete struct {
}

func (d *allianceMySqlDelete) DeleteTable(p_db *model.DBConfigInfo, p_serverid int, p_table string, p_mysqlPath string) error {
	whereSql := fmt.Sprintf("EXISTS(SELECT 1 FROM t_alliance WHERE allianceId=t_alliance.id and t_alliance.serverId=%d)", p_serverid)
	return tool.DeleteTable(p_db, p_table, whereSql, p_mysqlPath)
}

func (d *allianceMySqlDelete) IsServer() bool {
	return false
}

var (
	allianceDel        = &allianceMySqlDelete{}
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
		registerDelete(value, allianceDel)
	}
}
