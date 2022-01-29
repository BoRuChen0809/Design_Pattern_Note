package main

type Nike struct {
}

func (n *Nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt{
			logo: "nike",
			size: "XXL",
		},
	}
}

func (n *Nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			logo: "nike",
			size: "US_12",
		},
	}
}
