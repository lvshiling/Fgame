package global

import (
	"sync"
)

type Account interface {
	EnablePC() bool
}

var (
	once sync.Once
	g    Account
)

func SetupAccount(tg Account) {
	once.Do(func() {
		g = tg
	})
}

func GetAccount() Account {
	return g
}
