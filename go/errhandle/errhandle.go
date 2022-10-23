package main

import "fmt"

type Order struct {
}

func (o *Order) Validate() error {
	return nil
}

func (o *Order) GetTotalPrice() error {
	return nil
}

func (o *Order) Summit() error {
	return nil
}

type Order2 struct {
	err error
}

func (o *Order2) Validate() {
	if o.err != nil {
		return
	}
	fmt.Println("do Validate")
}

func (o *Order2) GetTotalPrice() {
	if o.err != nil {
		return
	}
	fmt.Println("do GetTotalPrice")
}

func (o *Order2) Summit() {
	if o.err != nil {
		return
	}
	fmt.Println("do Summit")
}

func (o *Order2) createOrder() error {
	o.Validate()
	o.GetTotalPrice()
	o.Summit()
	if o.err != nil {
		return o.err
	}
	return nil
}

type Order3 struct {
	err error
}

func (o *Order3) Validate() *Order3 {
	if o.err != nil {
		fmt.Println("do Validate")
	}
	return o
}

func (o *Order3) GetTotalPrice() *Order3 {
	if o.err != nil {
		fmt.Println("do GetTotalPrice")
	}
	return o
}

func (o *Order3) Summit() *Order3 {
	if o.err != nil {
		fmt.Println("do Summit")
	}
	return o
}

func main() {
	// 1. Golang 經典 error handle
	order := Order{}
	err := order.Validate()
	if err != nil {
		return
	}
	err = order.GetTotalPrice()
	if err != nil {
		return
	}
	err = order.Summit()
	if err != nil {
		return
	}
	fmt.Println("成功送出訂單")

	// 2. 屏蔽過程中的 error 處理
	order2 := Order2{}
	err = order2.createOrder()
	if err != nil {
		return
	}
	fmt.Println("成功送出訂單")

	// 3. 優化上一種方法，加入鏈式概念
	order3 := Order3{}
	order3.Validate().GetTotalPrice().Summit()
	if order3.err != nil {
		return
	}
	fmt.Println("成功送出訂單")
}
