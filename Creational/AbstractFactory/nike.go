package main

type Nike struct{
}

func (n *Nike)MakeShoe() Shoe{  //implements SportFactory
	return &NikeShoe{
		shoe :shoe{
			logo : "adidas",
			size : 13,
		},
	}
}

func (n *Nike)MakeShort() Short{  //implements SportFactory
	return &NikeShort{
		short : short{
			logo : "adidas",
			size : 14,
		},
	}
}

