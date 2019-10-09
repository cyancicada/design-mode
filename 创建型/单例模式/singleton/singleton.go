package singleton

import "sync"

//	定义一个对象类型，该类型只有唯一的对象
type singleton struct {
	Name string
}

var (
	once sync.Once //	保证一个函数只运行一次

	instance *singleton //	被确保唯一的对象
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
