package factoryPattern

import "errors"

func GetLogoFactory(logoType LogoType) (Factory, error) {
	switch logoType {
	case AdidasLogo:
		return &adidas{logoType}, nil
	case NikeLogo:
		return &nike{logoType}, nil
	default:
		return nil, errors.New("product type not supported")
	}
}

type LogoType string

const (
	NikeLogo   LogoType = "Nike"
	AdidasLogo          = "Adidas"
)

type Shoe interface {
	GetLogo() LogoType
	GetSize() float32
}

type Shirt interface {
	GetLogo() LogoType
	GetSize() string
}

type Factory interface {
	CreateShoe(float32) Shoe
	CreateShirt(string) Shirt
}
