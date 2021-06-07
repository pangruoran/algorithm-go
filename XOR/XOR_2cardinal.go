package main

import "fmt"

func XOR(lista []int) (int, int) {
	//已知两种数出现了奇数次。
	xor := 0
	for _, v := range lista {
		xor ^= v
	}
	//xor = 数1 ^ 数2

	//提取一个数最右侧的1
	xorr := xor & (^xor + 1)
	//此处变量xxor，即为第一个奇数数的变量。
	xxor := 0
	for _, v := range lista {
		if (v & xorr) == 1 {
			//此处还可以写成 (v & xorr) == 0
			xxor ^= v
		}
	}
	//xxor == 数1
	//xxor ^ xor => 数1 ^ (数1 ^ 数2) => 数2
	return xxor, xxor ^ xor
}

func main() {
	List := []int{1, 2, 3, 4, 1, 2, 3, 1, 2, 3, 1, 2}
	a, b := XOR(List)
	fmt.Println(List)
	fmt.Println("cardinal number is :", a, b)
}
