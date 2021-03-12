package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 定义参数类型
type Param struct {
	Width,Height int

}

// 声明对象
type Rect struct {


}

// 定义公共方法  (本机以外的电脑也可以调用)
// error 是必须的    ret为指针的地址
func (r *Rect)Area(p Param,ret *int) error {

	*ret = p.Width * p.Width

	return nil
}


// 周长
func (r *Rect)Perimeter(p Param,ret *int) error {
	*ret = (p.Width + p.Height)*2
	return nil
}

var n int32

// func Add(server *http.Server)int32{
//
// 	n++
// 	return
// }

// rpc是为了解决两台电脑之间互相调用对方程序的

func main()  {

	//注册服务
	rect := new(Rect)
	rpc.Register(rect)
	// 开辟端口让别的机器可以调用
	rpc.HandleHTTP()
	err := http.ListenAndServe(":9090",nil)

	if err != nil {
        log.Fatal(err)
	}
}