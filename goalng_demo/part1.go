package main

import (
	"fmt"
)


func struct_func(){
	type Car struct{
		color string
		displacement string
	}
	
	//new 的结构体 返回结构体的指针
	//new 's struct return a pointer to the struct
	car := new (Car)
	car.color = "black"
	car.displacement = "2.5T"

	
	//用var 则是结构体自身
	// use var is the struct itselt
	var car2 Car
	car2.color = "red"
	car2.displacement = "3.0T"

	fmt.Println(car)
	fmt.Println(car2)
}

func array_demo(){
	arr1 := [2]int{2,5}
	arr2 := []int {3,4}
	arr3 := [...] int {5,8}

	fmt.Println(arr1,arr2,arr3)
}

func array_demo2(){
	array := [5]int {1:20,4:35}
	for i,j :=range array{
		fmt.Println(i,j)
	}
}

func slice_demo1(){
	arr := [5]int{1,2,3,4,5}
	slice := arr[0:3]
	fmt.Println(slice)
}

func map_Demo(){
	arr := map[int] string{
		1:"Tom",
		2:"Jack",
	}
	for k,v := range arr{
		fmt.Println(k,v)
	}
}

//if we modify second_slice,the first_slice was been modified
func modify_slice(){
	//create a slice
	slice := []int {3,47,49,53,98}

	//create a new slice 
	new_slice := slice[1:4]

	//
	fmt.Println("print slice's value")
	fmt.Println(slice)
	fmt.Println(new_slice)
	//updata new_slcie's value,then print again
	new_slice[1] = 35
	fmt.Println("updata new_slice's value,then print again")
	fmt.Println(slice)
	fmt.Println(new_slice)
}

func add_element_for_slice(){
	//create slice1
	slice := []int {1,2,3,4}

	//create slice2
	slice = append(slice,5)
	fmt.Println(slice)
}

func new_slice(){
	slice1 := []int {1,2}
	slice2 := []int {3,4}
	slice2 = append(slice1)
	fmt.Println(slice2)
}

func map_demo1(){
	dict := map[int]string {
		1:"red",
		2:"green",
	}

	for k,v := range dict{
		fmt.Println(k,v)
	}
}

/*
func method_demo(){

	//定义结构体
	type user struct{
		name string 
	}


	//定义方法  使用值接受者实现方法
	func (u user) value_method1_introduce() {
		fmt.Println("welcome back! %s",u.name)
	}

	//定义方法  使用指正接受者实现一个方法
	func  (u *user) pointer_method2_changeName(name string){
		u.name = name
	}

	user1 := user{name:"Susan"} 
	user1.value_method1_introduce()
}*/

type user struct{
	name string 
}


//定义方法  使用值接受者实现方法
func (u user) value_method1_introduce() {
	fmt.Println("welcome back! ",u.name)
}

//定义方法  使用指正接受者实现一个方法
func  (u *user) pointer_method2_changeName(name string){
	u.name = name
}



func test1() int {
	var x int
	x++
	return x*x
}



func test2() func() int{
	var x int 
	return func() int{
		x++
		return x*x
	}
	
}




func main(){

	f:=test2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	//闭包的地址地址引用
	a := 10
	str := "mike"

	func(){
		a = 666
		str = "go"
		fmt.Printf("func 的a= %d,str= %s\n",a,str)
	}()
	fmt.Printf("非func 的a= %d,str= %s\n",a,str)

//闭包的特点
	fmt.Println("直接调用方法")
	fmt.Println(test1())
	fmt.Println(test1())



	user1 := user{name:"Susan"} 
	user1.value_method1_introduce()
	user1.pointer_method2_changeName("Jack")
	user1.value_method1_introduce()


	//	struct_func()
//	array_demo()
//	map_Demo()
//	array_demo2()

/*
	x,y := 3,6
	arr := [...]int{2:2}
	var pf *[3]int = &arr
	fmt.Println(pf)

	pfarr := [...] *int {&x,&y}
	fmt.Println(pfarr)

	//slice
	slice1 := []string {"green","red"}
	fmt.Println(slice1)

	//create null slice
	slice2 := []int {}
	fmt.Println(slice2)

	//declare null slice
	var slice3 []int
	fmt.Println(slice3)
*/

//	modify_slice()
//	add_element_for_slice()	

//	new_slice()

//	map_demo1()




}
