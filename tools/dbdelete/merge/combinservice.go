package merge

import (
	"encoding/json"
	fdb "fgame/fgame/core/db"
	"fgame/fgame/pkg/timeutils"
	delete "fgame/fgame/tools/dbdelete/delete"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
)

type ICombinService interface {
	CombinService(p_fpid int64, p_fsid int) error
}

func NewCombinService(p_ms IMergeService, p_dbms IDbConfigManage) ICombinService {
	rst := &combinService{

		ms:   p_ms,
		dbms: p_dbms,
	}
	rst.meta = NewMySqlMetaService(p_dbms)
	return rst
}

var (
	bakUpPath = "./backup"
	workPath  = "./worktemp"
)

type combinService struct {
	md        IServerMergeOneDB
	ms        IMergeService
	dbms      IDbConfigManage
	meta      IMySqlMetaService
	tableList []*tebleConfigItem

	centerDb fdb.DBService
}

func (m *combinService) CombinService(p_fpid int64, p_fsid int) error {

	fromDB := m.dbms.GetDbConfigInfo(p_fpid, p_fsid)
	if fromDB == nil {
		return fmt.Errorf("源库不存在")
	}
	dbconfig := &fdb.DbConfig{
		Debug:       false,
		Dialect:     "mysql",
		User:        fromDB.UserName,
		Password:    fromDB.PassWord,
		Host:        fmt.Sprintf("%s:%d", fromDB.Host, fromDB.Port),
		DBName:      fromDB.DBName,
		ParseTime:   true,
		Charset:     "utf8mb4",
		MaxIdle:     50,
		MaxActive:   100,
		MaxLifeTime: 200,
	}

	db, err := fdb.NewDBService(dbconfig)
	if err != nil {
		return err
	}
	mergeservice := NewServerMergeOneDB(db)
	m.md = mergeservice

	//导入前的准备
	log.Info("准备...")
	err = m.prepareCombin(p_fpid, p_fsid)
	if err != nil {
		return err
	}
	//备份目标库
	log.Info("备份...")
	err = m.backupToDb(p_fpid, p_fsid)
	if err != nil {
		return err
	}

	//导出备份
	log.Info("删除...")
	err = m.deleteTables(p_fpid, p_fsid)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	log.Info("完成...")
	return nil
}

//导入前的准备
func (m *combinService) prepareCombin(p_fpid int64, p_fsid int) error {
	backFilePath, err := filepath.Abs(bakUpPath)
	if err != nil {
		return fmt.Errorf("备份路径错误:", err)
	}
	_, err = notExistCreate(backFilePath)
	if err != nil {
		return fmt.Errorf("备份路径错误:", err)
	}
	workFilePath, err := filepath.Abs(workPath)

	_, err = notExistCreate(workFilePath)
	if err != nil {
		return fmt.Errorf("导出路径错误:", err)
	}

	ftables, err := m.meta.GetAllTableName(p_fpid, p_fsid)
	if err != nil {
		return err
	}

	tableList := &tableConfig{}
	tableList.TableList = make([]*tebleConfigItem, 0)
	for _, value := range ftables {
		item := &tebleConfigItem{
			TableName: value.TableName,
		}
		tableList.TableList = append(tableList.TableList, item)
	}

	m.tableList = tableList.TableList

	return nil
}

//导入前的备份目标库
func (m *combinService) backupToDb(p_tpid int64, p_tsid int) error {
	backFilePath, err := filepath.Abs(bakUpPath)
	if err != nil {
		log.Fatalln("备份路径错误:", err)
	}
	tdb := m.dbms.GetDbConfigInfo(p_tpid, p_tsid)
	if tdb == nil {
		return fmt.Errorf("目标库不存在")
	}
	// now := timeutils.TimeToMillisecond(time.Now())
	date := time.Now().Format("20060102_150405")
	beginNow := timeutils.TimeToMillisecond(time.Now())
	backFileName := backFilePath + "/" + fmt.Sprintf("%d_%d_%s.sql", p_tpid, p_tsid, date)
	err = m.ms.BackUpDb(tdb, backFileName)
	if err != nil {
		return err
	}
	endNow := timeutils.TimeToMillisecond(time.Now())
	totalSecond := endNow - beginNow
	log.Info("备份目标库,耗时：", totalSecond/1000, "秒")
	return nil
}

//删除表
func (m *combinService) deleteTables(p_fpid int64, p_fsid int) error {
	tdb := m.dbms.GetDbConfigInfo(p_fpid, p_fsid)
	if tdb == nil {
		return fmt.Errorf("目标库不存在")
	}
	for _, table := range m.tableList {

		export := delete.GetDelete(table.TableName)
		if export == nil {
			return fmt.Errorf("表%s未注册，删除失败", table.TableName)
		}
		if export.IsServer() {
			continue
		}
		log.Info("删除表:" + table.TableName)
		now := timeutils.TimeToMillisecond(time.Now())
		err := export.DeleteTable(tdb, p_fsid, table.TableName, m.ms.MySqlPath())
		newNow := timeutils.TimeToMillisecond(time.Now())
		spendSecond := newNow - now
		now = newNow
		log.Info("删除表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
		if err != nil {
			log.WithFields(log.Fields{
				"error":          err,
				"table":          table.TableName,
				"fromplatformid": p_fpid,
				"fromserverid":   p_fsid,
			}).Error("删除目标表异常")
			return err
		}
	}
	//删除服务器表
	for _, table := range m.tableList {

		export := delete.GetDelete(table.TableName)
		if export == nil {
			return fmt.Errorf("表%s未注册，删除失败", table.TableName)
		}
		if !export.IsServer() {
			continue
		}
		log.Info("删除表:" + table.TableName)
		now := timeutils.TimeToMillisecond(time.Now())
		err := export.DeleteTable(tdb, p_fsid, table.TableName, m.ms.MySqlPath())
		newNow := timeutils.TimeToMillisecond(time.Now())
		spendSecond := newNow - now
		now = newNow
		log.Info("删除表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
		if err != nil {
			log.WithFields(log.Fields{
				"error":          err,
				"table":          table.TableName,
				"fromplatformid": p_fpid,
				"fromserverid":   p_fsid,
			}).Error("删除目标表异常")
			return err
		}
	}
	return nil
}

// //导出表到工作区
// func (m *combinService) exportTables(p_fpid int64, p_fsid int) (int, error) {
// 	workFilePath, err := filepath.Abs(workPath)
// 	tdb := m.dbms.GetDbConfigInfo(p_fpid, p_fsid)
// 	if tdb == nil {
// 		return 0, fmt.Errorf("源库不存在")
// 	}
// 	beginNow := timeutils.TimeToMillisecond(time.Now())
// 	now := timeutils.TimeToMillisecond(time.Now())

// 	outPutCount := 0
// 	for _, table := range m.tableList {
// 		fileWorkTempPath := workFilePath + "/" + table.TableName + ".sql"

// 		whereName := fmt.Sprintf(table.TableWhere, p_fsid)
// 		log.Info("导出表:" + table.TableName)
// 		export := exp.GetExport(table.TableName)
// 		if export == nil {
// 			return outPutCount, fmt.Errorf("表%s未注册，导出失败", table.TableName)
// 		}
// 		// err = m.ms.BackUpTable(tdb, table.TableName, whereName, fileWorkTempPath)
// 		err = export.ExportTable(tdb, p_fsid, table.TableName, fileWorkTempPath, m.ms.DumpPath())
// 		newNow := timeutils.TimeToMillisecond(time.Now())
// 		spendSecond := newNow - now
// 		now = newNow
// 		log.Info("导出表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
// 		// log.Debug("导出表:"+table.TableName, spendSecond/1000, "秒")
// 		if err != nil {
// 			log.WithFields(log.Fields{
// 				"error":          err,
// 				"table":          table.TableName,
// 				"where":          whereName,
// 				"fromplatformid": p_fpid,
// 				"fromserverid":   p_fsid,
// 			}).Error("导出目标表异常")
// 			return 0, err
// 		}
// 		outPutCount++
// 	}
// 	endNow := timeutils.TimeToMillisecond(time.Now())
// 	totalSecond := endNow - beginNow
// 	log.Info("导出表:", outPutCount, "张,耗时：", totalSecond/1000, "秒")
// 	return outPutCount, nil
// }

// //导入目标表
// func (m *combinService) importTables(p_tpid int64, p_tsid int) (int, error) {
// 	inPutCount := 0
// 	workFilePath, err := filepath.Abs(workPath)
// 	beginNow := timeutils.TimeToMillisecond(time.Now())
// 	now := timeutils.TimeToMillisecond(time.Now())
// 	tdb := m.dbms.GetDbConfigInfo(p_tpid, p_tsid)
// 	if tdb == nil {
// 		return 0, fmt.Errorf("目标库不存在")
// 	}
// 	for _, table := range m.tableList {
// 		fileWorkTempPath := workFilePath + "/" + table.TableName + ".sql"
// 		log.Info("目标表:" + table.TableName)
// 		err = m.ms.ImportSqlPath(tdb, fileWorkTempPath)

// 		newNow := timeutils.TimeToMillisecond(time.Now())
// 		spendSecond := newNow - now
// 		now = newNow
// 		log.Info("目标表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
// 		// log.Debug("目标表:"+table.TableName, spendSecond/1000, "秒")

// 		if err != nil {
// 			return 0, err
// 		}
// 		inPutCount++
// 	}
// 	endNow := timeutils.TimeToMillisecond(time.Now())
// 	totalSecond := endNow - beginNow
// 	log.Info("导入表:", inPutCount, "张,耗时：", totalSecond/1000, "秒")
// 	return inPutCount, nil
// }

type tebleConfigItem struct {
	TableName  string `json:"tableName"`
	TableWhere string `json:"tableWhere"`
}

type tableConfig struct {
	TableList []*tebleConfigItem `json:"tableList"`
}

func readTableList(p_filePath string) (config *tableConfig, err error) {
	c, err := ioutil.ReadFile(p_filePath)
	if err != nil {
		return nil, err
	}
	sc := &tableConfig{}
	err = json.Unmarshal(c, sc)
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func notExistCreate(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, err
}
