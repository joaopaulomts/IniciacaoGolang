package main

import (
	"auction_golang/internal/entity/auction_entity"
	"fmt"
)

func main() {
	a := auction_entity.CreateAuction("Opala", auction_entity.New)

	fmt.Printf("Id %s", a.Id)
}
