package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main(){
	resp,err := http.Get("127.0.0.1:8080/server")
	if err != nil{
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("header = ",resp.Header)
	fmt.Println("resp status %s\nstatusCode %d\n",resp.Status,resp.StatusCode)
	fmt.Printf("boby type = %T\n",resp.Body)

	buf := make([]byte,2048)
	var tmp string
	
	for{
		n,err := resp.Body.Read(buf)
		if err != nil && err != io.EOF{
			fmt.Println(err)
			return
		}
		if n == 0{
			fmt.Println("读取内容结束")
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("buf = ",string(tmp))
}