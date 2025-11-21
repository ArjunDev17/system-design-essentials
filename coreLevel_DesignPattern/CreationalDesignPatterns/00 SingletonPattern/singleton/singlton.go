package singleton

import (
	"fmt"
	"sync"
)

type MySingleton struct {
	Value int
}

var (
	instance *MySingleton
	once     sync.Once
)

func GetInstance() *MySingleton {
	once.Do(func() {
		fmt.Println(">>> Creating the Singleton Instance")
		instance = &MySingleton{Value: 10}
	})
	return instance
}
