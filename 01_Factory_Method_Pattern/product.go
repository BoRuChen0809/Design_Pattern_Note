package main

type product interface {
	setName(name string)
	setDescribe(describe string)
	getName() string
	getDescribe() string
}
