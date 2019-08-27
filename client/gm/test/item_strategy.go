package test

import (
	"fgame/fgame/client/gm"
	"fgame/fgame/client/player"
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type testGmItemStrategy struct {
	p *player.Player
}

func (s *testGmItemStrategy) GetPlayer() *player.Player {
	return s.p
}

func (s *testGmItemStrategy) Run() {
	itemTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ItemTemplate)(nil))
	for _, item := range itemTemplateMap {
		gm.GmChangeItem(s.p, int32(item.TemplateId()), 10)
	}
}

func (s *testGmItemStrategy) OnError(code int32) {
	fmt.Printf("code:%d\n", code)
}

//物品变化了
func (s *testGmItemStrategy) OnItemChanged() {

}

func CreateTestGmItemStrategy(p *player.Player) player.Strategy {
	s := &testGmItemStrategy{
		p: p,
	}
	return s
}
