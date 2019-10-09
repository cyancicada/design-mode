package main

import (
	"fmt"
	"strconv"
	"sync"

	pool "design-mode/创建型/对象池模式/objectpool"
)

func main() {
	p := pool.NewPool(5)
	wait := sync.WaitGroup{}
	for i := 0; i < 7; i++ {
		index := i
		wait.Add(1)
		go func(pool pool.Pool, ind int) {
			select {
			case Obj := <-pool:
				Obj.Do(ind)
				pool <- Obj
			default:
				fmt.Println("No Object:" + strconv.Itoa(ind))
			}
			wait.Done()
		}(*p, index)
	}
	wait.Wait()
}
