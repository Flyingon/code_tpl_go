package main

import "fmt"

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLength := length - i
		oldArr := arrCopy(arr)
		genHeap(arr, lastLength)
		arr[0], arr[lastLength-1] = arr[lastLength-1], arr[0] // 最大元素移到最后
		fmt.Printf("第%d次结果: %v(%d) -> %v\n", i+1, oldArr, lastLength, arr)
	}
	return arr
}

// genHeap 构建大顶堆
func genHeap(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	noLeafStartIndex := length/2 - 1 // 从下往上，第一个非叶子节点的位置
	for i := noLeafStartIndex; i >= 0; i-- {
		rootChildIndex := i                                                       // 当前这个二叉树的跟(需要最大)
		leftChildIndex := 2*i + 1                                                 // 当前这个二叉树的左子树
		rightChildIndex := 2*i + 2                                                // 当前这个二叉树的右子树
		if leftChildIndex < length && arr[leftChildIndex] > arr[rootChildIndex] { // 这次只处理长度为length的数列，确定好范围
			rootChildIndex = leftChildIndex
		}
		if rightChildIndex < length && arr[rightChildIndex] > arr[rootChildIndex] {
			rootChildIndex = rightChildIndex
		}
		if rootChildIndex != i { // 需要调整
			arr[i], arr[rootChildIndex] = arr[rootChildIndex], arr[i]
		}
	}
	return arr
}

func arrCopy(arr []int) []int {
	ret := make([]int, len(arr))
	for index, elem := range arr {
		ret[index] = elem
	}
	return ret
}

func main() {
	arr := []int{
		22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, 70,
	}
	newArr := HeapSort(arr)
	fmt.Println(newArr)
}
