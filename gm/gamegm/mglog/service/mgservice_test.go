package service_test

import (
	"encoding/json"
	mongo "fgame/fgame/core/mongo"
	rpservice "fgame/fgame/gm/gamegm/mglog/service"
	"fmt"
	"log"
	"testing"
	//  rpmodel "fgame/fgame/gm/gamegm/gm/center/staticreport/model"
)

func TestGetOnLineStatic(t *testing.T) {
	mongoConfig := &mongo.MongoConfig{}
	mongoConfig.Addrs = append(mongoConfig.Addrs, "192.168.1.13:27017")
	mongoConfig.Database = "game_log"
	mongoConfig.FailFast = true

	mgdb, err := mongo.NewMongoService(mongoConfig)
	if err != nil {
		log.Fatalln("init mongodb failed:", err)
	}
	service := rpservice.NewMgLogService(mgdb, mongoConfig)
	rst, err := service.GetChatLogMsg("chat_content", 1541779200000, 1541952000000, 0, 0, 0, 0, 0, 0, "")
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
