package main

import (
	"fmt"
	"strconv"
)

func BbGame(lista []int) (result bool) {
	/*黑板游戏逻辑函数：
	      实现初始数组各元素异或为0，返回true；
	          初始数组各元素异或不为0，数组元素个数为偶数时，返回true；
			  其他情况返回false。*/
	xor := 0
	for _, v := range lista {
		xor ^= v
	}
	if (xor == 0) || (len(lista)%2 == 0) {
		result = true
	}
	return
}

func main() {
	List1 := []int{1, 1, 2, 2, 3, 4}
	fmt.Println(List1)
	fmt.Println("Result is : " + strconv.FormatBool(BbGame(List1)))
	List2 := []int{1, 1, 2, 2, 3}
	fmt.Println(List2)
	fmt.Println("Result is : " + strconv.FormatBool(BbGame(List2)))
}
