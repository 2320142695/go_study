package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

func (t T) String() {
	fmt.Println(strconv.Itoa(t.a) + " / " + strconv.FormatFloat(float64(t.b), 'f', 6, 64) + " / " + t.c)
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	t.String()
}
