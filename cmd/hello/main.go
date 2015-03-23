package main

import (
	"fmt"

	ct "github.com/gulyasm/go-meetup-2015/internal/github.com/daviddengcn/go-colortext"
)

func main() {
	ct.ChangeColor(ct.Green, true, ct.None, false)
	fmt.Println("Hello world!")
	ct.ResetColor()
}
