package main

import "fmt"

type VGA struct{}

func (vga *VGA) show() {
	fmt.Println("VGA monitor displaying......")
}
