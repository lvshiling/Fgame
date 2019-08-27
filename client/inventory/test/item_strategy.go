package test

import (
	"fgame/fgame/client/gm"
	"fgame/fgame/client/inventory"
	"fgame/fgame/client/player"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

type testItemStrategy struct {
	p    *player.Player
	code chan int32
}

func (s *testItemStrategy) GetPlayer() *player.Player {
	return s.p
}

func (s *testItemStrategy) Run() {
	s.testUseWrong()
	s.testUseCorrect()
	s.testMerge()
	s.p.Close()
}

type itemUseTest struct {
	Index int32
	Num   int32
	Code  int32
}

//参数错误提示
var itemUseWrongTestSuits = []*itemUseTest{
	&itemUseTest{Index: -1, Num: -1, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
	&itemUseTest{Index: -1, Num: 0, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
	&itemUseTest{Index: -1, Num: 1, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
	&itemUseTest{Index: 0, Num: -1, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
	&itemUseTest{Index: 0, Num: 0, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
	&itemUseTest{Index: 1, Num: -1, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
	&itemUseTest{Index: 1, Num: 0, Code: int32(uipb.ErrorCode_InventoryArgumentInvalid)},
}

//测试错误
func (s *testItemStrategy) testUseWrong() {
	for _, itemUse := range itemUseWrongTestSuits {
		s.code = make(chan int32, 0)
		inventory.InventoryUse(s.p, itemUse.Index, itemUse.Num)
		code := <-s.code
		if code != itemUse.Code {
			log.WithFields(
				log.Fields{
					"index": itemUse.Index,
					"num":   itemUse.Num,
				}).Fatalf("inventory.test.testUseWrong:expect %d,but get %d", itemUse.Code, code)
		}
	}
}

//测试使用正确的
func (s *testItemStrategy) testUseCorrect() {

	itemTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ItemTemplate)(nil))
	var item *gametemplate.ItemTemplate
	for _, tItem := range itemTemplateMap {
		item = tItem.(*gametemplate.ItemTemplate)
		break
	}
	manager := s.p.GetManager(player.PlayerDataKeyInventory).(*inventory.PlayerInventoryDataManager)
	index := manager.FirstIndex(int32(item.TemplateId()))
	if index == -1 {
		s.code = make(chan int32, 0)
		gm.GmChangeItem(s.p, int32(item.TemplateId()), 1)
		code := <-s.code
		if code != 0 {
			log.WithFields(
				log.Fields{}).Fatalf("inventory.test.testUseCorrect:expect %d,but get %d", 0, code)
		}

	}
	index = manager.FirstIndex(int32(item.TemplateId()))
	//使用
	s.code = make(chan int32, 0)
	inventory.InventoryUse(s.p, index, 1)
	code := <-s.code
	if code != 0 {
		log.WithFields(
			log.Fields{}).Fatalf("inventory.test.testUseCorrect:expect %d,but get %d", 0, code)
	}
}

func (s *testItemStrategy) testMerge() {

	itemTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ItemTemplate)(nil))

	for _, tItem := range itemTemplateMap {
		item := tItem.(*gametemplate.ItemTemplate)
		s.code = make(chan int32, 0)
		gm.GmChangeItem(s.p, int32(item.TemplateId()), 1)
		code := <-s.code
		if code != 0 {
			break
		}
	}

	//合并
	s.code = make(chan int32, 0)
	inventory.InventoryMerge(s.p)
	code := <-s.code
	if code != 0 {
		log.WithFields(
			log.Fields{}).Fatalf("inventory.test.testMerge:expect %d,but get %d", 0, code)
	}
}

//错误代码
func (s *testItemStrategy) OnError(code int32) {
	log.WithFields(
		log.Fields{
			"playerId": s.p.Id(),
			"code":     code,
		}).Warn("inventory:获取错误提示")

	s.code <- code

}

//物品变化了
func (s *testItemStrategy) OnItemChanged() {
	s.code <- 0
}

func CreateTestItemStrategy(p *player.Player) player.Strategy {
	s := &testItemStrategy{
		p: p,
	}
	return s
}
