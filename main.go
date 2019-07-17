package main

import (
	"fmt"

	"github.com/malikpractice/GoPackage-Tutorial/config"
)

func main() {
	fmt.Println("Package tutorial")

	pc := config.GetPostgresConnection()

	fmt.Println(pc)
}
