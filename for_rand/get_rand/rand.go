package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func main() {
	for i:=0; i < 1000;i ++ {
		sTime := time.Now()
		fmt.Println(GenerateRandomNumber(0, 3, 4))
		costTime := time.Now().Sub(sTime)
		fmt.Printf("cost time: %d ms\n", costTime.Milliseconds())
	}

}