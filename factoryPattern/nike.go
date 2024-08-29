package factoryPattern

type NikeShoe struct {
	nike
	size float32
}

type NikeShirt struct {
	nike
	size string
}

func (n *NikeShoe) GetSize() float32 {
	return n.size
}

func (n *NikeShirt) GetSize() string {
	return n.size
}

type nike struct {
	LogoType
}

func (n *nike) GetLogo() LogoType {
	return n.LogoType
}

func (n *nike) CreateShoe(size float32) Shoe {
	return &NikeShoe{
		nike: *n,
		size: size,
	}
}

func (n *nike) CreateShirt(size string) Shirt {
	return &NikeShirt{
		nike: *n,
		size: size,
	}
}
