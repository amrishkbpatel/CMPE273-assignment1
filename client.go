package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"

	"stocks"
)

var input string
var input2 string

var money float64
var per float64
var per2 float64
var opt float64
var tid float64

type Args struct {
	X, Y int
}

func GetStock() {
	fmt.Print("> ")
	fmt.Scanf("%s\n", &input)
}

func GetStock2() {
	fmt.Print("> ")
	fmt.Scanf("%s\n", &input2)
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{7, 8}
	var reply int
	c := jsonrpc.NewClient(client)
	err = c.Call("StockCal.Add", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Println("Enter 1st stock symbol:")
	GetStock()
	stock, err := stocks.GetQuote(input)

	fmt.Println("Enter %")
	fmt.Scanf("%f\n", &per)

	fmt.Println("Enter 2nd stock symbol:")
	GetStock2()
	stock2, err := stocks.GetQuote(input2)

	fmt.Println("Enter %")
	fmt.Scanf("%f\n", &per2)

	fmt.Println("Enter Budget:")
	fmt.Scanf("%f\n", &money)

	a := (per / 100) * money
	b := (per2 / 100) * money

	price, err := stock.GetPrice()
	if err != nil {
		fmt.Printf("Error getting price: %v", err)
	}
	price2, err := stock2.GetPrice()
	if err != nil {
		fmt.Printf("Error getting price: %v", err)
	}
	fmt.Println("Trade ID:112")
	fmt.Println(stock.GetName(), stock.GetSymbol(), price)
	fmt.Println(stock2.GetName(), stock2.GetSymbol(), price2)

	left := a - price
	left2 := b - price2
	fmt.Println("Unvested amount:")
	fmt.Println(left + left2) //unvested

	fmt.Println("Enter Trade id")
	fmt.Scanf("%f\n", &tid)

	fmt.Println(stock.GetName(), stock.GetSymbol(), price)
	fmt.Println(stock2.GetName(), stock2.GetSymbol(), price2)
	fmt.Println("Current Market Value:")
	fmt.Println(price + price2)
	fmt.Println("Unvested amount:")
	fmt.Println(left + left2)

}
