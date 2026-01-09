package main

import (
	"fmt"

	"gitlab.com/lightspeed-b2b/shrimps/pkg.git/idsupport"
)

func main() {
	uuid := idsupport.UUID()
	fmt.Println(uuid)
}
