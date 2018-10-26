package main

import (
	"fmt"
	"os"
)

func defer_call() {
	defer func() { fmt.Println("打印前") }()//3
	defer func() { fmt.Println("打印中") }()//2
	defer func() { fmt.Println("打印后") }()//1

	panic("触发异常")//4
}

func os_demo(){
	
}

func main() {
//	defer_call()
	list := os.Args
	n := len(list)
//	fmt.Println("n = ",n)

	for i:=0;i<n;i++{
		fmt.Printf("n = %d,value = %s\n",i,list[i])
	}


	for key,value := range list{
		fmt.Printf("list[%d] = %s\n",key,value)
	}
}

