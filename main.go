package main

import (
	"fmt"

	"github.com/malikpractice/GoPackage-Tutorial/config"
	"github.com/malikpractice/GoPackage-Tutorial/model"
)

func main() {
	fmt.Println("Golang package tutorial")

	pc := config.GetPostgresConnection()

	fmt.Println(pc)

	malik := model.Person{Id: "1", Name: "Malik"}

	fmt.Println(malik.Id, malik.Name)

}
