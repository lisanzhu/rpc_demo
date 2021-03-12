package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 实现客户端

type Param struct {
	Width,Height int

}

func main()  {
	// 链接远程RPC
	rp,err := rpc.DialHTTP("tcp","127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}
    // 从服务器接收过来的值
	ret := 0
	wid:=100
	hei := 100
	for i:=1;i<10;i++{
		wid += 10
		hei+= 20
		e := rp.Call("Rect.Area",Param{wid,hei},&ret)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(ret)
	}
}