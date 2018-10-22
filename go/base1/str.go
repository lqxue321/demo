package main

import "fmt"

type struct1 struct{
	x, y int
}


var (
	s1 = struct1{1,2}

	s2 = struct1{x:1}

	s3 = struct1{}

	p = &struct1{2,3}
)

func main(){
//	fmt.Println(struct1{1,5})
	v := struct1{1,2}
	fmt.Println(v)

	p :=&v
	fmt.Println(p)

	fmt.Println(s1,s2,s3,p)
	
}