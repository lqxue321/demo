package main

import (
	"fmt"
)

//type的几种用法

//1.定义结构体
type Person struct{
	name string
	age int
}


//2.类型等价定义，相当于类型重命名
type name1 string

type name2 string

func (name name2)len() int{
	return len(name)
}

//3.结构体内嵌匿名成员
type Animal struct{
	string
	isPredator bool
}

//4.定义接口类型,补充  实现接口1.结构体实现  2.函数对象
type Personer interface{
	Run()
	Name() string
} 

type Person1 struct{
	name string 
	age int
}

func (Person1)Run(){
	fmt.Println("running.....")
}

func (person1 Person1)Name() string{
	return person1.Name
}

func  main(){
	/*
	p := Person{"",18}
	fmt.Println(p.age)

	

	var myname name = "Rose"
	i := []byte(myname)
	fmt.Println(len(i))
	
	
	animal1 := Animal{"name",false}
	fmt.Println(animal1)

	*/

	var p Person1
	fmt.Println(p)
}