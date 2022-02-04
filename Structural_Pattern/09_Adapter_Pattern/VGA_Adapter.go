package main

import "fmt"

type VGA_Adapter struct {
	vga *VGA
}

func (vgaa *VGA_Adapter) display() {
	fmt.Println("adapter converts VGA to DP.")
	vgaa.vga.show()
}
