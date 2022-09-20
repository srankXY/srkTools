package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	//"io"
	//"log"
	"net/http"
)

//type test struct {
//}


func webroot(w http.ResponseWriter, req *http.Request){
	//io.WriteString(w, "hello, world!\n")
	w.Write([]byte("yes i can"))
	fmt.Println(req.RequestURI,"連接成功")

	 fmt.Println(req.Method)
}

//func (test) ServeHTTP(w http.ResponseWriter, req *http.Request){
//	w.Write([]byte("yes test"))
//}

func main() {
	client, err := rpc.Dial("https://rinkeby.infura.io/v3/a751414aa90e4970822056b85e48c037")
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}
	var networkid string
	client.Call(&networkid, "eth_hashrate")

	//var tg test
	//tg = test{}
	//http.Handle("/tg", tg)
	http.HandleFunc("/", webroot)

	go http.ListenAndServe(":6235", nil)


}

