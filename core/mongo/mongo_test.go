package mongo_test

import (
	"sync"
	"testing"

	"gopkg.in/mgo.v2/bson"

	. "fgame/fgame/core/mongo"
)

var once sync.Once
var ms MongoService

func setupMongoService(t *testing.T) {
	once.Do(func() {
		mc := &MongoConfig{}
		mc.Addrs = []string{"192.168.99.100:32773", "192.168.99.100:32772", "192.168.99.100:32771"}
		mc.Database = "test"
		mc.FailFast = true
		tms, err := NewMongoService(mc)
		if err != nil {
			t.Fatalf("setup mongo service failed [%s]", err)
		}
		ms = tms
	})
}

type Test struct {
	TestA string `testA`
	TestB int64  `testB`
}

func TestInsert(t *testing.T) {
	setupMongoService(t)
	session := ms.Session().Clone()
	defer session.Close()
	c := session.DB("test").C("test")
	err := c.Insert(&Test{TestA: "test1", TestB: 1}, &Test{TestA: "test2", TestB: 2})
	if err != nil {
		t.Fatalf("mongo insert failed [%s]", err)
	}
	result := &Test{}
	err = c.Find(bson.M{"testA": "test1"}).One(result)
	if err != nil {
		t.Fatalf("mongo find failed [%s]", err)
	}

	if result.TestB != 1 {
		t.Fatal("mongo insert and query no same")
	}
}
