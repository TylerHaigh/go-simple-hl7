package main

import (
	"com.tylerhaigh/go-simple-hl7/pkg/hl7"
)

func main() {
	c := hl7.Component{
		Data: []hl7.SubComponent{"sss"},
	}

	println(c.GetSubComponent(4))

}
