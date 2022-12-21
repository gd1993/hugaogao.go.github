package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

//结构体   用于注册
type Arith struct {
}

//声明参数结构体
type ArithRequest struct {
	A, B int
}

//返回客户端的结果
type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

//乘积
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

//商和余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为零")
	}
	//除
	res.Quo = req.A / req.B
	//取模
	res.Rem = req.A % req.B
	return nil
}
func main() {
	//注册服务
	rect := new(Arith)
	//注册一个rect服务
	rpc.Register(rect)
	//服务处理绑定到http协议上
	rpc.HandleHTTP()
	//监听服务
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}
