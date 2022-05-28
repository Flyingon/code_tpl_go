package main

import (
	"fmt"
	"strconv"
)

// leetcode-752: https://leetcode.cn/problems/open-the-lock/

func openLockV1(deadends []string, target string) int {
	var step int
	deadMap := make(map[string]bool)
	for _, deadEnd := range deadends {
		deadMap[deadEnd] = true
	}
	q := []string{
		"0000",
	}
	visted := map[string]bool{
		"0000": true,
	}
	for len(q) != 0 {
		//fmt.Println("q: ", q)
		nextQueue := []string{}
		for i := 0; i < len(q); i++ {
			// 检查当前的队列
			cur := q[i]
			if deadMap[cur] {
				continue
			}
			if cur == target {
				return step
			}
			//fmt.Printf("当前：%s，", cur)

			for j := 0; j < 4; j++ {
				up := up(cur, j)
				if !visted[up] {
					nextQueue = append(nextQueue, up)
					visted[up] = true
				}
				down := down(cur, j)
				if !visted[down] {
					nextQueue = append(nextQueue, down)
					visted[down] = true
				}
			}
		}
		q = nextQueue
		// 在这里增加步数
		step++
		//fmt.Printf("第[%d]步: %+v\n", step, q)
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

// openLock 双向BFS遍历
func openLock(deadends []string, target string) int {
	var step int
	deadMap := make(map[string]bool)
	for _, deadEnd := range deadends {
		deadMap[deadEnd] = true
	}
	q := map[string]bool{
		"0000": true,
	}
	q2 := map[string]bool{
		target: true,
	}
	visted := map[string]bool{
		"0000": true,
	}
	for len(q) != 0 && len(q2) != 0 {
		temp := make(map[string]bool)
		for cur := range q {
			// 检查当前的队列
			if deadMap[cur] {
				continue
			}
			if q2[cur] {
				return step
			}

			visted[cur] = true
			for j := 0; j < 4; j++ {
				up := up(cur, j)
				if !visted[up] {
					temp[up] = true
				}
				down := down(cur, j)
				if !visted[down] {
					temp[down] = true
				}
			}
		}
		// 在这里增加步数
		step++
		//fmt.Printf("第[%d]步,temp: %+v\n", step,temp)
		// temp 相当于 q1
		// 这里交换 q1 q2，下一轮 循环 就是扩散 q2
		q = q2
		q2 = temp
		//fmt.Printf("第[%d]步: %+v, %+v\n", step, q, q2)
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

// 将 s[pos] 向上拨动一次
func up(s string, pos int) string {
	res := ""
	for index, numByte := range []byte(s) {
		if pos == index {
			num, _ := strconv.ParseInt(string(numByte), 10, 64)
			newNum := num + 1
			if newNum == 10 {
				newNum = 0
			}
			res += string(strconv.FormatInt(newNum, 10))
		} else {
			res += string(numByte)
		}
	}
	return res
}

// 将 s[pos] 向下拨动一次
func down(s string, pos int) string {
	res := ""
	for index, numByte := range []byte(s) {
		if pos == index {
			num, _ := strconv.ParseInt(string(numByte), 10, 64)
			newNum := num - 1
			if newNum == -1 {
				newNum = 9
			}
			res += string(strconv.FormatInt(newNum, 10))
		} else {
			res += string(numByte)
		}
	}
	return res
}

func main() {
	//cur := "0000"
	//for j := 0; j < 4; j++ {
	//	up := up(cur, j)
	//	fmt.Println(up)
	//	down := down(cur, j)
	//	fmt.Println(down)
	//}

	res := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	fmt.Println(res)
}
