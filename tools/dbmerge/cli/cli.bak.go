package cli

// import (
// 	"encoding/json"
// 	merge "fgame/fgame/tools/dbmerge/merge"
// 	model "fgame/fgame/tools/dbmerge/model"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	fdb "fgame/fgame/core/db"
// 	timeutils "fgame/fgame/pkg/timeutils"

// 	log "github.com/Sirupsen/logrus"
// 	"github.com/codegangsta/cli"
// 	"github.com/rifflock/lfshook"
// )

// var (
// 	debug          = false
// 	configFile     = "./config/config.json"
// 	tebleFile      = "./config/table.json"
// 	bakUpPath      = "./backup"
// 	workPath       = "./worktemp"
// 	fromPlatFormId = int64(1)
// 	fromServerId   = 2
// 	toPlatFormId   = int64(1)
// 	toServerId     = 1
// 	notClearEnd    = false
// )

// type mergeConfig struct {
// 	DumpPath      string                `json:"dumpPath"`
// 	MySqlPath     string                `json:"mysqlPath"`
// 	DBServerArray []*model.DBConfigInfo `json:"dbMap"`
// }

// func readMergeConfig(p_filePath string) (config *mergeConfig, err error) {
// 	c, err := ioutil.ReadFile(p_filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sc := &mergeConfig{}
// 	err = json.Unmarshal(c, sc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sc, nil
// }

// type tebleConfigItem struct {
// 	TableName  string `json:"tableName"`
// 	TableWhere string `json:"tableWhere"`
// }

// type tableConfig struct {
// 	TableList []*tebleConfigItem `json:"tableList"`
// }

// func readTableList(p_filePath string) (config *tableConfig, err error) {
// 	c, err := ioutil.ReadFile(p_filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sc := &tableConfig{}
// 	err = json.Unmarshal(c, sc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sc, nil
// }

// func Start() {
// 	app := cli.NewApp()
// 	app.Name = "dbmerge"
// 	app.Usage = "dbmerge [global options]"

// 	app.Author = ""
// 	app.Email = ""

// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name:        "config,c",
// 			Value:       configFile,
// 			Usage:       "config file",
// 			Destination: &configFile,
// 		},
// 		cli.BoolFlag{
// 			Name:        "debug,d",
// 			Usage:       "debug",
// 			Destination: &debug,
// 		},
// 		cli.BoolFlag{
// 			Name:        "notClearEnd,nc",
// 			Usage:       "notClearEnd",
// 			Destination: &notClearEnd,
// 		},
// 		cli.Int64Flag{
// 			Name:        "fromPlatFormId,fp",
// 			Usage:       "fromPlatFormId",
// 			Destination: &fromPlatFormId,
// 		},
// 		cli.IntFlag{
// 			Name:        "fromServerId,fs",
// 			Usage:       "fromServerId",
// 			Destination: &fromServerId,
// 		},
// 		cli.Int64Flag{
// 			Name:        "toPlatFormId,tp",
// 			Usage:       "toPlatFormId",
// 			Destination: &toPlatFormId,
// 		},
// 		cli.IntFlag{
// 			Name:        "toServerId,ts",
// 			Usage:       "toServerId",
// 			Destination: &toServerId,
// 		},
// 	}
// 	app.Before = before
// 	app.Action = start
// 	app.Run(os.Args)
// }

// func before(ctx *cli.Context) error {
// 	log.AddHook(lfshook.NewHook(lfshook.PathMap{
// 		log.DebugLevel: "./logs/info.log",
// 		log.InfoLevel:  "./logs/info.log",
// 		log.WarnLevel:  "./logs/info.log",
// 		log.ErrorLevel: "./logs/error.log",
// 	}))

// 	if debug {
// 		log.SetLevel(log.DebugLevel)
// 	} else {
// 		log.SetLevel(log.WarnLevel)
// 	}
// 	return nil
// }

// func start(ctx *cli.Context) {
// 	// if fromPlatFormId == 0 || fromServerId == 0 || toPlatFormId == 0 || toServerId == 0 {
// 	// 	log.Fatalln("server err, failed: 没有足够的参数，缺失平台或者服务id")
// 	// }

// 	// fromPlatFormId = 1
// 	// fromServerId = 10
// 	// toPlatFormId = 1
// 	// toServerId = 2
// 	fmt.Println("fromPlatFormId:", fromPlatFormId)
// 	fmt.Println("fromServerId:", fromServerId)
// 	fmt.Println("toPlatFormId:", toPlatFormId)
// 	fmt.Println("toServerId:", toServerId)
// 	// return

// 	config, err := filepath.Abs(configFile)
// 	if err != nil {
// 		log.Fatalln("filepath abs failed:", err)
// 	}

// 	sc, err := readMergeConfig(config)
// 	if err != nil {
// 		log.Fatalln("read config file failed:", err)
// 	}

// 	tconfig, err := filepath.Abs(tebleFile)
// 	if err != nil {
// 		log.Fatalln("tebleFile abs failed:", err)
// 	}

// 	tableList, err := readTableList(tconfig)
// 	if err != nil {
// 		log.Fatalln("read teble file failed:", err)
// 	}

// 	if len(tableList.TableList) == 0 {
// 		log.Fatalln("read teble file is empty")
// 	}

// 	backFilePath, err := filepath.Abs(bakUpPath)
// 	if err != nil {
// 		log.Fatalln("bakUpPath abs failed:", err)
// 	}
// 	_, err = notExistCreate(backFilePath)
// 	if err != nil {
// 		log.Fatalln("workPath path failed:", err)
// 	}
// 	workFilePath, err := filepath.Abs(workPath)

// 	fmt.Println("清理工作区...")
// 	err = clearWorkPath(workFilePath)
// 	if err != nil {
// 		log.Fatalln("clear work temp files failed:", err)
// 	}

// 	_, err = notExistCreate(workFilePath)
// 	if err != nil {
// 		log.Fatalln("workPath path failed:", err)
// 	}

// 	dbconfigManage := merge.NewDbConfigManage()
// 	for _, item := range sc.DBServerArray {
// 		fmt.Println("resigter ", item.PlatformId, item.ServerId)
// 		dbconfigManage.RegisterDbConfigInfo(item.PlatformId, item.ServerId, item)
// 	}

// 	//检查输入的服务器是否都存在db
// 	fromDB := dbconfigManage.GetDbConfigInfo(int64(fromPlatFormId), fromServerId)
// 	if fromDB == nil {
// 		log.Fatalln("from db is not exists")
// 	}

// 	toDB := dbconfigManage.GetDbConfigInfo(int64(toPlatFormId), toServerId)
// 	if toDB == nil {
// 		log.Fatalln("to db is not exists")
// 	}

// 	fmt.Println("begin backup from db...")

// 	now := timeutils.TimeToMillisecond(time.Now())
// 	backFileName := backFilePath + "/" + fmt.Sprintf("%d_%d_%d.sql", fromPlatFormId, fromServerId, now)

// 	service := merge.NewMergeService(sc.DumpPath, sc.MySqlPath, dbconfigManage)
// 	allBegin := now
// 	//备份目标库
// 	log.Debug("开始备份目标库：", now)
// 	err = service.BackUpDb(int64(toPlatFormId), toServerId, backFileName)
// 	if err != nil {
// 		log.WithFields(log.Fields{
// 			"error": err,
// 		}).Error("备份目标表异常")
// 		log.Fatalln("backup To path failed:", err)
// 	}
// 	newNow := timeutils.TimeToMillisecond(time.Now())
// 	spendSecond := newNow - now
// 	now = newNow
// 	log.Debug("结束备份目标库，耗时：", spendSecond/1000, "秒")
// 	// return

// 	outPutCount := 0
// 	inPutCount := 0
// 	totalBegin := newNow
// 	for _, table := range tableList.TableList {
// 		fileWorkTempPath := workFilePath + "/" + table.TableName + ".sql"

// 		whereName := fmt.Sprintf(table.TableWhere, fromServerId)
// 		fmt.Println("导出表:" + table.TableName)
// 		err = service.BackUpTable(fromPlatFormId, fromServerId, table.TableName, whereName, fileWorkTempPath)
// 		newNow = timeutils.TimeToMillisecond(time.Now())
// 		spendSecond = newNow - now
// 		now = newNow
// 		fmt.Println("导出表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
// 		// log.Debug("导出表:"+table.TableName, spendSecond/1000, "秒")
// 		if err != nil {
// 			log.WithFields(log.Fields{
// 				"error":          err,
// 				"table":          table.TableName,
// 				"where":          whereName,
// 				"fromplatformid": fromPlatFormId,
// 				"fromserverid":   fromServerId,
// 			}).Error("导出目标表异常")
// 			log.Fatalln("backup table failed:", err)
// 		}
// 		outPutCount++
// 	}
// 	totalEnd := timeutils.TimeToMillisecond(time.Now())
// 	log.Debug("导出目标表总共耗时", (totalEnd-totalBegin)/1000, "秒")

// 	fmt.Println("开始导入目标库...")
// 	totalBegin = totalEnd
// 	for _, table := range tableList.TableList {
// 		fileWorkTempPath := workFilePath + "/" + table.TableName + ".sql"
// 		fmt.Println("目标表:" + table.TableName)
// 		err = service.ImportSqlPath(toPlatFormId, toServerId, fileWorkTempPath)

// 		newNow = timeutils.TimeToMillisecond(time.Now())
// 		spendSecond = newNow - now
// 		now = newNow
// 		fmt.Println("目标表:"+table.TableName+"耗时：", spendSecond/1000, "秒")
// 		// log.Debug("目标表:"+table.TableName, spendSecond/1000, "秒")

// 		if err != nil {
// 			log.Fatalln("import table failed:", err)
// 		}
// 		inPutCount++
// 	}
// 	fmt.Println("导入目标表完成!")
// 	totalEnd = timeutils.TimeToMillisecond(time.Now())
// 	log.Debug("入目标表总共耗时", (totalEnd-totalBegin)/1000, "秒")
// 	totalBegin = totalEnd

// 	dbconfig := &fdb.DbConfig{
// 		Debug:       false,
// 		Dialect:     "mysql",
// 		User:        toDB.UserName,
// 		Password:    toDB.PassWord,
// 		Host:        fmt.Sprintf("%s:%d", toDB.Host, toDB.Port),
// 		DBName:      toDB.DBName,
// 		ParseTime:   true,
// 		Charset:     "utf8mb4",
// 		MaxIdle:     50,
// 		MaxActive:   100,
// 		MaxLifeTime: 200,
// 	}

// 	db, err := fdb.NewDBService(dbconfig)
// 	if err != nil {
// 		log.Fatalln("to db init failed:", err)
// 	}

// 	fmt.Println("开始数据服合并...")
// 	mergeservice := merge.NewServerMergeOneDB(db)
// 	err = mergeservice.MergeChangeServerData(fromServerId, toServerId)
// 	if err != nil {
// 		log.Fatalln("merge change data failed:", err)
// 	}
// 	fmt.Println("结束数据服合并...")
// 	totalEnd = timeutils.TimeToMillisecond(time.Now())
// 	log.Debug("合并ServerID数据总共耗时", (totalEnd-totalBegin)/1000, "秒")
// 	allEnd := totalEnd
// 	log.Debug("全过程总共耗时:", (allEnd-allBegin)/1000, "秒", "导出:", outPutCount, "张表", "导入：", inPutCount, "张表")
// 	if !notClearEnd {
// 		fmt.Println("开始清除工作区...")
// 		err = clearWorkPath(workFilePath)
// 		if err != nil {
// 			log.Fatalln("clear work temp files failed:", err)
// 		}
// 	}

// 	// err = service.BackUpDb(1, 1, "E:/tempbak/mysqldata/db.sql")
// 	// if err != nil {
// 	// 	fmt.Println("error:", err)
// 	// }

// 	// err = service.BackUpTable(1, 1, "t_player", "", "E:/tempbak/mysqldata/t_player.sql")
// 	// if err != nil {
// 	// 	fmt.Println("error:", err)
// 	// }

// 	// err = service.BackUpTable(1, 1, "t_player", "id=4558226566359714 or id=4558244333392969", "E:/tempbak/mysqldata/t_player_one.sql")
// 	// if err != nil {
// 	// 	fmt.Println("error:", err)
// 	// }

// 	// err = service.ImportSqlPath(1, 2, "E:/tempbak/mysqldata/t_player_one.sql")
// 	// if err != nil {
// 	// 	fmt.Println("error:", err)
// 	// }
// }

// func notExistCreate(path string) (bool, error) {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true, nil
// 	}
// 	if os.IsNotExist(err) {
// 		err = os.Mkdir(path, os.ModePerm)
// 		if err != nil {
// 			return false, err
// 		}
// 		return true, nil
// 	}
// 	return false, err
// }

// func clearWorkPath(p_path string) error {
// 	// _, err := os.Stat(p_path)
// 	// if err == nil {
// 	// 	err = os.RemoveAll(p_path)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	// 	return nil
// 	// }
// 	// if os.IsNotExist(err) {
// 	// 	return nil
// 	// }

// 	// return err
// 	return nil
// }
