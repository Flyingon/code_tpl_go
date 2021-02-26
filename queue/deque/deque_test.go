package deque

import (
	"testing"
	"time"
)

func TestDeque_Run(t *testing.T) {
	deque := NewDeque(
		func(num int) (ret []string, err error) {
			for i := 0; i < num; i++ {
				ret = append(ret, "test")
			}
			return
		},
		func(s string) error {
			return nil
		},
		1, 1)
	deque.Run()
	time.Sleep(3 * time.Second)
	deque.RunSimple()
	time.Sleep(3 * time.Second)
}
