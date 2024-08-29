package factoryPattern

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	nikeFactory, _ := GetLogoFactory(NikeLogo)
	adidasFactory, _ := GetLogoFactory(AdidasLogo)

	nikeShow := nikeFactory.CreateShoe(32)
	nikeShirt := nikeFactory.CreateShirt("175/180")

	adidasShow := adidasFactory.CreateShoe(42.5)
	adidasShirt := adidasFactory.CreateShirt("180/185")

	fmt.Printf("nike shoe:Logo:%s,size:%.2f\n", nikeShow.GetLogo(), nikeShow.GetSize())
	fmt.Printf("nike shirt:Logo:%s,size:%s\n", nikeShirt.GetLogo(), nikeShirt.GetSize())
	fmt.Printf("adiads shoe:Logo:%s,size:%.2f\n", adidasShow.GetLogo(), adidasShow.GetSize())
	fmt.Printf("adiads shirt:Logo:%s,size:%s\n", adidasShirt.GetLogo(), adidasShirt.GetSize())
}
