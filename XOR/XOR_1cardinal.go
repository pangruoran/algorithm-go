package main

import (
	"fmt"
	"strconv"
)

func XOR(lista []int) int {
	//已知，数组中只有一种数出现了奇数次。
	xor := 0
	for _, v := range lista {
		xor ^= v
	}
	return xor
}

func main() {
	List := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 1, 2, 3, 1, 2, 1, 2, 1}
	Output := XOR(List)
	fmt.Println(List)
	fmt.Println("After using XOR function : " + strconv.Itoa(Output))
}
