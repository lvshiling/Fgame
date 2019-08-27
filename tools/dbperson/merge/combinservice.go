package merge

import (
	"encoding/json"
	fdb "fgame/fgame/core/db"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	exp "fgame/fgame/tools/dbperson/export"

	log "github.com/Sirupsen/logrus"
)

type ICombinService interface {
	CombinService(p_fpid int64, p_fsid int, p_tpid int64, p_tsid int) error
}

func NewCombinService(centerDB fdb.DBService, p_ms IMergeService, p_dbms IDbConfigManage) ICombinService {
	rst := &combinService{
		centerDb: centerDB,
		ms:       p_ms,
		dbms:     p_dbms,
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

func (m *combinService) CombinService(p_fpid int64, p_fsid int, p_tpid int64, p_tsid int) error {

	toDB := m.dbms.GetDbConfigInfo(p_tpid, p_tsid)
	if toDB == nil {
		return fmt.Errorf("目标库不存在")
	}
	fromDB := m.dbms.GetDbConfigInfo(p_tpid, p_fsid)
	if fromDB == nil {
		return fmt.Errorf("源库不存在")
	}
	dbconfig := &fdb.DbConfig{
		Debug:       false,
		Dialect:     "mysql",
		User:        toDB.UserName,
		Password:    toDB.PassWord,
		Host:        fmt.Sprintf("%s:%d", toDB.Host, toDB.Port),
		DBName:      toDB.DBName,
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

	//导入前检查
	log.Info("检查...")
	//err = m.checkCombin(p_fpid, p_fsid, p_tpid, p_tsid)
	//if err != nil {
	//	return err
	//}

	//导入前的准备
	log.Info("准备...")
	err = m.prepareCombin(p_fpid, p_fsid)
	if err != nil {
		return err
	}
	//备份目标库
	log.Info("备份...")
	err = m.backupToDb(p_tpid, p_tsid)
	if err != nil {
		return err
	}
	//判断是否同一个库
	if toDB.Equal(fromDB) {
		log.Info("同一个数据库,不需要导出")
	} else {
		//导出备份
		log.Info("导出...")
		exportCount, err := m.exportTables(p_fpid, p_fsid)
		if err != nil {
			return err
		}
		log.Info("导出表", exportCount, "张")
		//导入备份
		//log.Info("导入...")
		inputCount, err := m.importTables(p_tpid, p_tsid)
		if err != nil {
			return err
		}
		log.Info("导入表", inputCount, "张")
	}

	log.Info("完成...")
	return nil
}

//导入前的检查
func (m *combinService) checkCombin(p_fpid int64, p_fsid int, p_tpid int64, p_tsid int) error {
	fromDB := m.dbms.GetDbConfigInfo(p_fpid, p_fsid)
	if fromDB == nil {
		return fmt.Errorf("源库不存在")
	}

	toDB := m.dbms.GetDbConfigInfo(p_tpid, p_tsid)
	if toDB == nil {
		return fmt.Errorf("目标库不存在")
	}
	err := m.checkTableUniformity(p_fpid, p_fsid, p_tpid, p_tsid)
	if err != nil {
		return err
	}
	return nil
}

//检查表结构一致性
func (m *combinService) checkTableUniformity(p_fpid int64, p_fsid int, p_tpid int64, p_tsid int) error {
	mesg := make([]string, 0)

	ftables, err := m.meta.GetAllTableName(p_fpid, p_fsid)
	if err != nil {
		return err
	}
	if ftables == nil || len(ftables) == 0 {
		return fmt.Errorf("from tables count is 0")
	}
	ttables, err := m.meta.GetAllTableName(p_tpid, p_tsid)
	if err != nil {
		return err
	}
	if ttables == nil || len(ttables) == 0 {
		return fmt.Errorf("to tables count is 0")
	}

	if len(ftables) != len(ttables) {
		mesg = append(mesg, fmt.Sprintf("两库表数量不一致"))
	}

	//检查表个数
	fNotExists := make([]string, 0)
	for _, fvalue := range ftables {
		exflag := false
		for _, tvalue := range ttables {
			if fvalue.TableName == tvalue.TableName {
				exflag = true
				break
			}
		}
		if !exflag {
			fNotExists = append(fNotExists, fvalue.TableName)
		}
	}

	tNotExists := make([]string, 0)
	for _, tvalue := range ttables {
		exflag := false
		for _, fvalue := range ftables {
			if fvalue.TableName == tvalue.TableName {
				exflag = true
				break
			}
		}
		if !exflag {
			tNotExists = append(tNotExists, tvalue.TableName)
		}
	}

	if len(fNotExists) > 0 {
		mesg = append(mesg, fmt.Sprintf("源库中存在表：%s未在目标库中", fNotExists))
	}
	if len(tNotExists) > 0 {
		mesg = append(mesg, fmt.Sprintf("目标库中存在表：%s未在源库中", tNotExists))
	}

	if len(mesg) > 0 { //返回表检查结果
		return fmt.Errorf("%s", mesg)
	}

	for _, fvalue := range ftables {
		export := exp.GetExport(fvalue.TableName)
		if export == nil {
			mesg = append(mesg, fmt.Sprintf("表%s未注册", fvalue.TableName))
		}
	}
	if len(mesg) > 0 { //返回表检查结果
		return fmt.Errorf("%s", mesg)
	}

	//检查列及列属性
	for _, tableinfo := range ftables {
		fcolumnList, err := m.meta.GetTableColumn(p_fpid, p_fsid, tableinfo.TableName)
		if err != nil {
			return err
		}
		if len(fcolumnList) == 0 {
			return fmt.Errorf("源库表%s字段为空", tableinfo.TableName)
		}
		tcolumnList, err := m.meta.GetTableColumn(p_tpid, p_tsid, tableinfo.TableName)
		if err != nil {
			return err
		}
		if len(tcolumnList) == 0 {
			return fmt.Errorf("目标库表%s字段为空", tableinfo.TableName)
		}

		if len(fcolumnList) != len(tcolumnList) {
			mesg = append(mesg, fmt.Sprintf("表%s字段数量不一致,源%d,目标:%d", tableinfo.TableName, len(fcolumnList), len(tcolumnList)))
		}

		fNotColumn := make([]string, 0)
		for _, fcValue := range fcolumnList {
			exflag := false
			for _, tcValue := range tcolumnList {
				if fcValue.ColumnName == tcValue.ColumnName && fcValue.ColumnType == tcValue.ColumnType {
					exflag = true
				}
			}
			if !exflag {
				fNotColumn = append(fNotColumn, fcValue.ColumnName)
			}
		}
		if len(fNotColumn) > 0 {
			mesg = append(mesg, fmt.Sprintf("表%s中的字段%s,源表中该字段，目标库的该表没有", tableinfo.TableName, fNotColumn))
		}

		tNotColumn := make([]string, 0)
		for _, tcValue := range tcolumnList {
			exflag := false
			for _, fcValue := range fcolumnList {
				if fcValue.ColumnName == tcValue.ColumnName && fcValue.ColumnType == tcValue.ColumnType {
					exflag = true
				}
			}
			if !exflag {
				tNotColumn = append(tNotColumn, tcValue.ColumnName)
			}
		}
		if len(tNotColumn) > 0 {
			mesg = append(mesg, fmt.Sprintf("表%s中的字段%s,目标表中该字段，源库的该表没有", tableinfo.TableName, tNotColumn))
		}
	}
	if len(mesg) > 0 {
		return fmt.Errorf("%s", strings.Replace(strings.Trim(fmt.Sprint(mesg), "[]"), " ", "\n", -1))
	}

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

//导出表到工作区
func (m *combinService) exportTables(p_fpid int64, p_fsid int) (int, error) {
	workFilePath, err := filepath.Abs(workPath)
	tdb := m.dbms.GetDbConfigInfo(p_fpid, p_fsid)
	if tdb == nil {
		return 0, fmt.Errorf("源库不存在")
	}
	beginNow := timeutils.TimeToMillisecond(time.Now())
	now := timeutils.TimeToMillisecond(time.Now())

	outPutCount := 0
	for _, table := range m.tableList {
		fileWorkTempPath := workFilePath + "/" + table.TableName + ".sql"

		whereName := fmt.Sprintf(table.TableWhere, p_fsid)
		log.Info("导出表:" + table.TableName)
		export := exp.GetExport(table.TableName)
		if export == nil {
			continue
			//	return outPutCount, fmt.Errorf("表%s未注册，导出失败", table.TableName)
		}
		// err = m.ms.BackUpTable(tdb, table.TableName, whereName, fileWorkTempPath)
		err = export.ExportTable(tdb, p_fsid, table.TableName, fileWorkTempPath, m.ms.DumpPath())
		newNow := timeutils.TimeToMillisecond(time.Now())
		spendSecond := newNow - now
		now = newNow
		log.Info("导出表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
		// log.Debug("导出表:"+table.TableName, spendSecond/1000, "秒")
		if err != nil {
			log.WithFields(log.Fields{
				"error":          err,
				"table":          table.TableName,
				"where":          whereName,
				"fromplatformid": p_fpid,
				"fromserverid":   p_fsid,
			}).Error("导出目标表异常")
			return 0, err
		}
		outPutCount++
	}
	endNow := timeutils.TimeToMillisecond(time.Now())
	totalSecond := endNow - beginNow
	log.Info("导出表:", outPutCount, "张,耗时：", totalSecond/1000, "秒")
	return outPutCount, nil
}

//导入目标表
func (m *combinService) importTables(p_tpid int64, p_tsid int) (int, error) {
	inPutCount := 0
	workFilePath, err := filepath.Abs(workPath)
	beginNow := timeutils.TimeToMillisecond(time.Now())
	now := timeutils.TimeToMillisecond(time.Now())
	tdb := m.dbms.GetDbConfigInfo(p_tpid, p_tsid)
	if tdb == nil {
		return 0, fmt.Errorf("目标库不存在")
	}
	for _, table := range m.tableList {
		fileWorkTempPath := workFilePath + "/" + table.TableName + ".sql"
		log.Info("目标表:" + table.TableName)
		_, err = os.Stat(fileWorkTempPath)
		if err != nil {
			if !os.IsNotExist(err) {
				return 0, err
			}
			err = nil
			continue
		}
		err = m.ms.ImportSqlPath(tdb, fileWorkTempPath)

		newNow := timeutils.TimeToMillisecond(time.Now())
		spendSecond := newNow - now
		now = newNow
		log.Info("目标表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
		// log.Debug("目标表:"+table.TableName, spendSecond/1000, "秒")

		if err != nil {
			return 0, err
		}
		inPutCount++
	}
	endNow := timeutils.TimeToMillisecond(time.Now())
	totalSecond := endNow - beginNow
	log.Info("导入表:", inPutCount, "张,耗时：", totalSecond/1000, "秒")
	return inPutCount, nil
}

//更改目标数据
func (m *combinService) updateServerData(p_fsid int, p_tsid int) error {
	beginNow := timeutils.TimeToMillisecond(time.Now())
	err := m.md.MergeChangeServerData(p_fsid, p_tsid)
	endNow := timeutils.TimeToMillisecond(time.Now())
	totalSecond := endNow - beginNow
	log.Info("修改表服务器表数据,", "耗时：", totalSecond/1000, "秒")
	if err != nil {
		return err
	}
	return nil
}

//更新中心服数据
func (m *combinService) updateCenterMergeData(p_fpid int32, p_fsid int32, p_tsid int32) (flag bool, err error) {
	flag, err = m.ifMerge(p_fpid, p_fsid, p_tsid)
	if err != nil {
		return
	}
	//是否合服过
	if flag {
		flag = false
		return
	}
	trans := m.centerDb.DB().Begin()
	defer func() {
		if err != nil {
			trans.Rollback()
		}
		return
	}()

	now := timeutils.TimeToMillisecond(time.Now())
	//添加合服记录
	e := &MergeRecordEntity{}
	e.Platform = p_fpid
	e.FromServerId = p_fsid
	e.ToServerId = p_tsid
	e.FinalServerId = p_tsid
	e.MergeTime = now
	e.CreateTime = now
	err = trans.Save(e).Error
	if err != nil {
		return
	}
	//更新所有最终服务器id是源服务器的
	//角色修改
	err = trans.Exec("UPDATE t_merge_record SET finalServerId=? WHERE platform=? and finalServerId = ?", p_tsid, p_fpid, p_fsid).Error
	if err != nil {
		return
	}
	err = trans.Commit().Error
	if err != nil {
		return
	}
	flag = true
	return
}

//是否合服过
func (m *combinService) ifMerge(pid int32, fsid int32, tsid int32) (flag bool, err error) {
	e := &MergeRecordEntity{}
	//判断源服务器有没有合过
	err = m.centerDb.DB().First(e, "platform=? and fromServerId=?", pid, fsid).Error
	if err == nil {
		return true, nil
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		err = nil
	}

	//判断目标服务器有没有合过
	err = m.centerDb.DB().First(e, "platform=? and fromServerId=?", pid, tsid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return
	}
	return true, nil
}

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
