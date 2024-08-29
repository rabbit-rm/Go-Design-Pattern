package factoryPattern

type AdidasShoe struct {
	adidas
	size float32
}

type AdidasShirt struct {
	adidas
	size string
}

func (n *AdidasShoe) GetSize() float32 {
	return n.size
}

func (n *AdidasShirt) GetSize() string {
	return n.size
}

type adidas struct {
	LogoType
}

func (a *adidas) GetLogo() LogoType {
	return a.LogoType
}

func (a *adidas) CreateShoe(size float32) Shoe {
	return &AdidasShoe{
		adidas: *a,
		size:   size,
	}
}

func (a *adidas) CreateShirt(size string) Shirt {
	return &AdidasShirt{
		adidas: *a,
		size:   size,
	}
}
