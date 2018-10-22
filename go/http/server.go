package main

 import (
	 "fmt"
     "net/http"
  //   "io/ioutil"
 )


func myHandler(w http.ResponseWriter,request *http.Request){
    fmt.Fprintf(w,"Hello.world %s!",request.URL.Path[1:])
}
/*
func te1(){
    resp,_ := http.Get("http:www.baidu.com")
    defer resp.Body.Close()
    body,_ :=ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}*/

func main() {
    http.HandleFunc("/", myHandler)
    http.ListenAndServe("localhost:8080", nil)  
 }
