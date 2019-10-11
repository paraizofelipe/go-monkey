package kong

import "sync"

var (
	once     sync.Once
	instance *Kong
)

type Kong struct {
	BaseURL        string
	MonkeyCommands map[string]interface{}
}

func GetInstance() *Kong {
	once.Do(func() {
		instance = &Kong{
			MonkeyCommands: make(map[string]interface{}),
		}
	})
	return instance
}

func (k *Kong) GetBaseURL() string {
	return k.BaseURL
}

func (k *Kong) GetMonkeyCommands() map[string]interface{} {
	return k.MonkeyCommands
}

func (k *Kong) AddMonkeyCommands(name string, mkCommand interface{}) {
	k.MonkeyCommands[name] = mkCommand
}
