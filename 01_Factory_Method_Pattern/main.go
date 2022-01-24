package main

import "fmt"

func main() {
	product_1, err := getFood("steak")
	if err != nil {
		fmt.Println(err)
	} else {
		product_detail(product_1)
	}
	product_2, err := getFood("pork chop")
	if err != nil {
		fmt.Println(err)
	} else {
		product_detail(product_2)
	}
}

func product_detail(p product) {
	fmt.Printf("Product Type : %s\n", p.getName())
	fmt.Printf("Product Describe : %s\n", p.getDescribe())
	fmt.Println("********************************************************************************************")
}
