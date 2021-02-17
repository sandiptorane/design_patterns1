package main

type Adidas struct{
}

func (a *Adidas)MakeShoe() Shoe{  //implements SportFactory
	return &AdidasShoe{
		shoe : shoe{
			logo : "adidas",
			size : 13,
		},
	}
}

func (a *Adidas)MakeShort() Short{  //implements SportFactory
	return &AdidasShort{
		short : short{
			logo : "adidas",
			size : 14,
		},
	}
}
