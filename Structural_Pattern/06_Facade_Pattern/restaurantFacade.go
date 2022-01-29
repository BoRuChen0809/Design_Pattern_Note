package main

type restaurant struct {
	waiter    *Waiter
	cleaner   *Cleaner
	chef      *Chef
	vegvendor *VegVendor
}

func newRestaurant() *restaurant {
	return &restaurant{
		waiter:    newWaiter(),
		cleaner:   newCleaner(),
		chef:      newChef(),
		vegvendor: newVegVendor(),
	}
}

func (r *restaurant) order_dishes() {
	r.vegvendor.purchase()
	r.waiter.order()
	r.chef.cook()
	r.waiter.service()
	r.cleaner.clean()
}
