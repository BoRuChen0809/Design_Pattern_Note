package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type sun struct {
	description string
}

func (s *sun) Describe() string {
	return s.description
}

var sunInstance *sun

func getSun() *sun {
	if sunInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if sunInstance == nil {
			now := time.Now().Format("2006/01/02 15:04:05")
			fmt.Println("Create Sun.")
			sunInstance = &sun{fmt.Sprintf("created on : %s", now)}
			fmt.Println(sunInstance.Describe())
			return sunInstance
		} else {
			fmt.Println("Sun had created.")
			fmt.Println(sunInstance.Describe())
			return sunInstance
		}
	} else {
		fmt.Println("Sun had created.")
		fmt.Println(sunInstance.Describe())
		return sunInstance
	}
}
