package main

import (
	"fmt"

	"github.com/proway/exemplo/models"
)

func main() {
	var u models.User = models.User{
		ID:    1,
		Email: "usuario@gmail.com",
		Senha: "senha123"}
	fmt.Printf("%v", u.GetID())
}
