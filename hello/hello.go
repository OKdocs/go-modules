package hello

import (
	"fmt"

	"github.com/petrkotek/go-vector/2d/32bit/vec"
)

func Hello() string {
	v1 := vec.New(1, 2)
	v2 := vec.New(3, 4)
	fmt.Println(v1.Add(v2))
	return "Hello, World"
}
