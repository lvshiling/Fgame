package player

//策略
type Strategy interface {
	//玩家
	GetPlayer() *Player
	//执行
	Run()
	//错误提示
	OnError(code int32)
	//物品改变
	OnItemChanged()
}

type StrategyFactory interface {
	CreateStrategy(p *Player) Strategy
}

type StrategyFactoryFunc func(p *Player) Strategy

func (sff StrategyFactoryFunc) CreateStrategy(p *Player) Strategy {
	return sff(p)
}

type StrategyKey string

var (
	strategyFactoryMap map[StrategyKey]StrategyFactory = make(map[StrategyKey]StrategyFactory)
)
