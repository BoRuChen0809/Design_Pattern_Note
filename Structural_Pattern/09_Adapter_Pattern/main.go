package main

func main() {
	c := &client{}

	dp := &Display{}
	c.watch(dp)

	vga := &VGA{}
	vgaa := &VGA_Adapter{vga: vga}
	c.watch(vgaa)
}
