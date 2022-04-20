package main

import "fmt"

func shoppingPrice() int {
	var totalItem, totalPrice int
	fmt.Print("Total item: ")
	fmt.Scanf("%d", &totalItem)

	prices := make([]int, totalItem)
	fmt.Print("Harga: ")
	for i := 0; i < totalItem; i++ {
		fmt.Scanf("%d", &prices[i])
	}

	for i := 0; i < len(prices); i++ {
		totalPrice += prices[i]
	}

	return totalPrice
}

func shoppingPriceSrp() int {
	totalItem := getTotalItem()
	prices := getPrices(totalItem)
	totalPrice := countTotalPrice(prices)

	return totalPrice
}

func getTotalItem() int {
	var totalItem int
	fmt.Print("Total item: ")
	fmt.Scanf("%d", &totalItem)

	return totalItem
}

func getPrices(totalItem int) []int {
	prices := make([]int, totalItem)
	fmt.Print("Harga: ")
	for i := 0; i < totalItem; i++ {
		fmt.Scanf("%d", &prices[i])
	}

	return prices
}

func countTotalPrice(prices []int) int {
	totalPrice := 0
	for i := 0; i < len(prices); i++ {
		totalPrice += prices[i]
	}

	return totalPrice
}

func main() {
	fmt.Println(shoppingPriceSrp())
}
