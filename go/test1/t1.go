package main 
import (
	"fmt"
	"strings"
)

func main(){
	s := "hello,world"
	fmt.Println(strings.Contains(s,"hello"),strings.Contains(s,"?"))

}

