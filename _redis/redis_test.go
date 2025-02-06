package _redis

import (
	"fmt"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	{
		machine := &Machine{
			Host:     `127.0.0.1`,
			Port:     `6379`,
			Database: `0`,
			Username: ``,
			Password: ``,
		}
		New().Machine(machine).Set("name", 0, 0)
		wg := sync.WaitGroup{}
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func(i int) {
				New().Machine(machine).Incr("name")
				wg.Done()
			}(i)
		}
		wg.Wait()
		res := New().Machine(machine).Get("name")
		fmt.Println(res)
	}
}
