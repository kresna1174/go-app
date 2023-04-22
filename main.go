package main

import (
	"fmt"

	go_module "github.com/kresna1174/go-module/v2"
)

type Person struct {
	Name string
}

func main() {
	person := Person{Name: "Kresna"}
	fmt.Println(go_module.SayHello(person.Name))
	fmt.Println(go_module.SayMyName(person.Name))
	fmt.Println(go_module.SayYourName(person.Name))
}
