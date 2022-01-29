package main

type food struct {
	name     string
	describe string
}

func (f *food) setName(name string) {
	f.name = name
}

func (f *food) getName() string {
	return f.name
}

func (f *food) setDescribe(describe string) {
	f.describe = describe
}

func (f *food) getDescribe() string {
	return f.describe
}
