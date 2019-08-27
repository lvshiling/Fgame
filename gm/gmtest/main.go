package main

import (
	"encoding/json"
	mongo "fgame/fgame/core/mongo"
	mglog "fgame/fgame/logserver/log"
	mgmodel "fgame/fgame/logserver/model"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type ITest interface {
	GetName() string
}

type Test struct {
	// TestA string `json:"testA"`
	// TestB int64  `json:"testB"`
	TestA string `testA`
	TestB int64  `testB`
}

func (m *Test) GetName() string {
	return "test"
}

func main() {
	// var intrst []int
	// intrst = append(intrst, 1)
	//192.168.1.13:27017
	mc := &mongo.MongoConfig{}
	mc.Addrs = []string{"192.168.1.13:27017"}
	mc.Database = "test"
	mc.FailFast = true
	tms, err := mongo.NewMongoService(mc)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	session := tms.Session().Clone()
	defer session.Close()

	info := &mgmodel.ChatContent{}
	// mglog.RegisterLogMsg(info)
	result := mglog.GetLogMsgList(info.LogName())

	c := session.DB("game_log").C(info.LogName())
	condition := bson.M{}
	// condition["platform"] = 1
	condition["logtime"] = bson.M{"$gte": 1538064000000}
	condition["logtime"] = bson.M{"$lte": 1538150400000, "$gte": 1538064000000}
	err = c.Find(condition).Sort("logtime").Limit(20).All(result)

	// c := session.DB("test").C("learyTest")
	// for i := 0; i < 100; i++ {
	// 	var test ITest
	// 	test = &Test{
	// 		TestA: fmt.Sprintf("key_%d", i),
	// 		TestB: int64(i),
	// 	}
	// 	err = c.Insert(test)
	// 	if err != nil {
	// 		fmt.Println("insert error:", err)
	// 		return
	// 	}
	// }

	// var test ITest
	// test = &Test{}

	// myinfo := &Test{}
	// err = c.Find(nil).Sort(`-testB`).One(test)

	// t := reflect.ValueOf(myinfo).Type()
	// // h := reflect.New(t).Interface()

	// result := reflect.MakeSlice(t, 0, 0)
	// // err = c.Find(bson.M{"testA": "test1"}).Sort("TestB")

	// // var valueString interface{}
	// err = c.Find(nil).Sort(`-testB`).All(&result)
	// // err = c.Find(nil).Sort(`-testB`).One(&valueString)
	// // err = c.Find(nil).Sort(`-testB`).Limit(10).Skip(5).All(result)
	if err != nil {
		fmt.Println("find error:", err)
		return
	}

	by, err := json.Marshal(result)
	fmt.Println(string(by))
	// by, err = json.Marshal(valueString)
	// fmt.Println(string(by))

	// newobj := valueString.(Test)
	// by, err = json.Marshal(&newobj)
	// fmt.Println(string(by))
}
