package main

import (
	"com.tylerhaigh/go-simple-hl7/pkg/hl7"
)

func main() {
	c := hl7.Component{
		Data: []string{"sss"},
	}

	println(c.GetSubComponent(4))

}
