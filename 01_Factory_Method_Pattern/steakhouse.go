package main

import "errors"

func getFood(foodType string) (product, error) {
	if foodType == "steak" {
		return newSteak(), nil
	}
	if foodType == "pork chop" {
		return newPork_Chop(), nil
	}
	notFound := errors.New("the type of product you want is not exist.")
	return nil, notFound
}
