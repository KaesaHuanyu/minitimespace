package main

import (
	"fmt"
	"minitimespace/models"
)

func main() {
	u1, _ := models.GetUserById(1)
	u2, _ := models.GetUserById(2)
	t, _ := models.GetTimespaceById(1)
	fmt.Println(t.AddUser(u1, u2))
}
