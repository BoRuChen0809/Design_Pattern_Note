package main

import "fmt"

type Cleaner struct {
}

func newCleaner() *Cleaner {
	return &Cleaner{}
}

func (c *Cleaner) clean() {
	fmt.Println("清理桌面、洗碗......")
}
