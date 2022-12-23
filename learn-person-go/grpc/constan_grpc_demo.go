package main

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
