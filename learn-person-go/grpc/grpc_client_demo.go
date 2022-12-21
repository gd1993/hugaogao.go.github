package main

import (
	"fmt"
	"log"
	"net/rpc"
)

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

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}
	req := ArithRequest{9, 2}
	var res ArithResponse

	err2 := conn.Call("Arith.Multiply", req, &res)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err3 := conn.Call("Arith.Divide", req, &res)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("%d / %d 商 %d，余数 = %d\n", req.A, req.B, res.Quo, res.Rem)
}
