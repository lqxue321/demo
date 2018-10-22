package main

import "fmt"
/*
var 数组名[数组长度]数组类型
*/

/*
向函数传递数组
*/

//参数设定数组长度
func test1(arr [5]int){
	for k,v := range arr{
		fmt.Println(k,"=>",v)
	}
}


//参数不设定数组长度
func test2(arr []int,size int)float32{
	
	var i int
    var avg, sum float32  

    for i = 0; i < size; ++i {
       sum += arr[i]
    }

   avg = sum / size

   return avg;
}

func main(){
/*	var arr1 = [5] int{1,2,3,45,71}

	test1(arr1)

	*/

var bal = []{5,8,12,14,36}
test2{bal,len(bal)}
}