package merge

import (
	model "fgame/fgame/tools/dbperson/model"
	"fmt"
)

func getBackUpDbDump(p_dumpPath string, p_db *model.DBConfigInfo, p_savePath string) string {

	rst := fmt.Sprintf(p_dumpPath+" -h %s -P %d -u%s -p%s  --default-character-set=utf8mb4 --skip-add-locks -t %s > %s", p_db.Host, p_db.Port, p_db.UserName, p_db.PassWord, p_db.DBName, p_savePath)
	return rst
}

func getBackUpDbTableDump(p_dumpPath string, p_db *model.DBConfigInfo, p_tableName string, p_where string, p_savePath string) string {
	rst := fmt.Sprintf(p_dumpPath+" -h %s -P %d -u%s -p%s  --default-character-set=utf8mb4 --skip-add-locks -t %s %s  --where=\"%s\" > %s", p_db.Host, p_db.Port, p_db.UserName, p_db.PassWord, p_db.DBName, p_tableName, p_where, p_savePath)
	return rst
}

func getImportDump(p_mySqlPath string, p_db *model.DBConfigInfo, p_filePath string) string {
	rst := fmt.Sprintf(p_mySqlPath+" -h %s -P %d -u%s -p%s %s < %s", p_db.Host, p_db.Port, p_db.UserName, p_db.PassWord, p_db.DBName, p_filePath)
	return rst
}
