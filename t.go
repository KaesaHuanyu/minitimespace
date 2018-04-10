package main

import (
	"fmt"
	"minitimespace/models"
)

func main() {
	t, _ := models.GetTimespaceById(1)
	fmt.Println(t.GetUsers())
}
