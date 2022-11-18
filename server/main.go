package main

import (
	"github.com/lebrancconvas/Go-Proto-gRPC/pb/proto"  
	"fmt" 
)

func main() {
	a := proto.Mage{
		Name: "Lebranc",
		Level: 1,
		Element: "Fire",
		Health: 100,
		Mana: 50,
	}	
	fmt.Println(a.Name);  
}