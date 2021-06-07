package main

import (
	"fmt"
)

func XORSwitch(lista []int) []int {
	//已知lista中有且只有两个元素 lista[0] , lista[1]
	//实现交换，并且不引入新的变量
	lista[0] = lista[0] ^ lista[1]
	lista[1] = lista[0] ^ lista[1]
	lista[0] = lista[0] ^ lista[1]
	return lista
}

func main() {
	List := []int{345, 678}
	fmt.Println(List)
	List = XORSwitch(List)
	fmt.Println("After using function XORSwitch : ")
	//fmt.Println("After using function XORSwitch : " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(List)), ""), "[]") )
	fmt.Println(List)
}
