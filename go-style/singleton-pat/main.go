package main

import (
	"fmt"
	"sync"
)

type EntranceGate struct {
	numberPlate string
	entryTime   string
	vehchleType string
}

var instance *EntranceGate
var once sync.Once

// private costructor for initialization
func newEntranceGate() *EntranceGate {
	return &EntranceGate{}
}

func GetInstance() *EntranceGate {
	once.Do(func() {

		instance = newEntranceGate()
	})
	return instance
}

func main() {

	fmt.Println("welcome to LLD first mark")
	etrance1 := GetInstance()
	etrance1.entryTime = "22:36"
	etrance1.numberPlate = "jjsdcsd"
	etrance1.vehchleType = "2Wheeler"

}
