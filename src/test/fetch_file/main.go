package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func request(url string) ([]byte, error) {
    resp, _ := http.Get(url)
    defer (func() {
        resp.Body.Close()
    })()
    return ioutil.ReadAll(resp.Body)
}

func main(){
    body,_ := request("http://www.baidu.com");

    fmt.Println(string(body))  
}