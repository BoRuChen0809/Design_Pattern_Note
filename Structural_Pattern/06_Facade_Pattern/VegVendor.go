package main

import "fmt"

type VegVendor struct {
}

func newVegVendor() *VegVendor {
	return &VegVendor{}
}

func (v *VegVendor) purchase() {
	fmt.Println("δΎζθ¬θ......")
}
