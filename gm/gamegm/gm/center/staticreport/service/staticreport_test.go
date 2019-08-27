package service_test

import (
	"encoding/json"
	mongo "fgame/fgame/core/mongo"
	rpservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	"fmt"
	"log"
	"testing"

	mongodb "gopkg.in/mgo.v2"
	//  rpmodel "fgame/fgame/gm/gamegm/gm/center/staticreport/model"
)

func TestGetOnLineStatic(t *testing.T) {
	logger := &fmtLog{}
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.13:27017")
	mongoConfig.Database = "game_log"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	mongodb.SetDebug(false)
	mongodb.SetLogger(logger)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewReportStatic(mgdb, mongoConfig)
	rst, err := service.GetOnLineStatic(1540828800000, 1540915200000, 1, 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	bytes, err := json.Marshal(rst)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(string(bytes))
	t.Fail()
}

func TestGetOnLinePlayerNumGroupSdk(t *testing.T) {
	logger := &fmtLog{}
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.13:27017")
	mongoConfig.Database = "game_log"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	mongodb.SetDebug(false)
	mongodb.SetLogger(logger)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewReportStatic(mgdb, mongoConfig)
	rst, err := service.GetOnLinePlayerNumGroupSdk(1539070075815, 1541779200000, 0, []int{1, 2})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	bytes, err := json.Marshal(rst)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(string(bytes))
	t.Fail()
}

func TestGetOnLinePlayerNum(t *testing.T) {
	logger := &fmtLog{}
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.13:27017")
	mongoConfig.Database = "game_log"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	mongodb.SetDebug(false)
	mongodb.SetLogger(logger)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewReportStatic(mgdb, mongoConfig)
	rst, err := service.GetOnLinePlayerNum(1539070075815, 1541779200000, 1, 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	bytes, err := json.Marshal(rst)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(string(bytes))
	t.Fail()
}

func TestGetOnLineServerFirstTime(t *testing.T) {
	logger := &fmtLog{}
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.13:27017")
	mongoConfig.Database = "game_log"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	mongodb.SetDebug(false)
	mongodb.SetLogger(logger)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewReportStatic(mgdb, mongoConfig)
	rst, err := service.GetOnLineServerFirstTime()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(rst)
	t.Fail()
}

func TestGetServerLoginPlayerNum(t *testing.T) {
	logger := &fmtLog{}
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.123:27017")
	mongoConfig.Database = "game_log_master"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	mongodb.SetDebug(false)
	mongodb.SetLogger(logger)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewReportStatic(mgdb, mongoConfig)
	rst, err := service.GetServerLoginPlayerNum(int64(1556899200000), int64(1559664000000))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	rstArray, _ := json.Marshal(rst)
	fmt.Println("result")
	fmt.Println(string(rstArray))
	t.Fail()
}

func TestGetOnLineStaticDaily(t *testing.T) {
	logger := &fmtLog{}
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.123:27017")
	mongoConfig.Database = "game_log_master"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	mongodb.SetDebug(false)
	mongodb.SetLogger(logger)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewReportStatic(mgdb, mongoConfig)
	rst, err := service.GetOnLineStaticDaily(int64(1556899200000), int64(1559664000000))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	rstArray, _ := json.Marshal(rst)
	fmt.Println("result")
	fmt.Println(string(rstArray))
	t.Fail()
}

func TestMapJson(t *testing.T) {
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	rstByte, _ := json.Marshal(myMap)
	fmt.Println(string(rstByte))
	t.Fail()
}

type fmtLog struct{}

func (m *fmtLog) Output(calldepth int, s string) error {
	fmt.Println(s)
	return nil
}
