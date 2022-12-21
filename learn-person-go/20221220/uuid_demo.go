package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	u1, _ := uuid.NewUUID()
	fmt.Printf("uuid: %s\n", u1)

}
