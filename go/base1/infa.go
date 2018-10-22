package main

import (
	"fmt"
)

//定义接口
type person interface{
	run() string
	eat() string 
}

type person_method struct{
	mo 
}