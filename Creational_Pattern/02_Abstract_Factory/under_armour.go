package main

type UA struct {
}

func (ua *UA) makeShirt() iShirt {
	return &uaShirt{
		shirt{
			logo: "UA",
			size: "XL",
		},
	}
}

func (ua *UA) makeShoe() iShoe {
	return &uaShoe{
		shoe{
			logo: "UA",
			size: "US_11",
		},
	}
}
