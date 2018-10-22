package main

import (
	"fmt"
)

func main(){
//闭包要通过匿名函数来实现

	a := 10
	str := "Mike"

	func (){
		fmt.Println("a=",a)
		fmt.Println("str=",str)
	}()

	
	func (a,b int){
		fmt.Println("a=%d,b=%d",a,b)
	}(2,3)
}