package main

import (
	"fmt"

	"github.com/ysugiyama12/go-api/api"
)

func main() {
	fmt.Println("main")
	res := Hello("world")
	fmt.Println(res)
	api.Execute()
}
