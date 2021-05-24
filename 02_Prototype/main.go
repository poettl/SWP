package main

import (
	"fmt"
	"swp/prototype/prototypemanager"
)

func main() {
	prototypemanager.InitPrototype()
	mysword := prototypemanager.GetInstance("Longsword")
	fmt.Println(mysword)
}
