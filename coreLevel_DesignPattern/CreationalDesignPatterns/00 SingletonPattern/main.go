package main

import (
	"fmt"
	"singleton-impl/singleton"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			obj := singleton.GetInstance()
			fmt.Printf("Goroutine %d -> Instance Address: %p\n", id, obj)
		}(i)
	}

	wg.Wait()
}
