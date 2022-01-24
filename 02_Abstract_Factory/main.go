package main

import "fmt"

func main() {
	nike_factory, err := getSportsFactory("Nike")
	if err != nil {
		fmt.Println(err)
	}
	nikeShoe := nike_factory.makeShoe()
	nikeShirt := nike_factory.makeShirt()
	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	ua_factory, err := getSportsFactory("UA")
	if err != nil {
		fmt.Println(err)
	}
	uaShirt := ua_factory.makeShirt()
	uaShoe := ua_factory.makeShoe()
	printShirtDetails(uaShirt)
	printShoeDetails(uaShoe)

}

func printShoeDetails(s iShoe) {
	fmt.Println("Shoe")
	fmt.Printf("Logo : %s\n", s.getLogo())
	fmt.Printf("Size : %s\n", s.getSize())
	fmt.Println("************************************************************************")
}

func printShirtDetails(s iShirt) {
	fmt.Println("Shirt")
	fmt.Printf("Logo : %s\n", s.getLogo())
	fmt.Printf("Size : %s\n", s.getSize())
	fmt.Println("************************************************************************")
}
