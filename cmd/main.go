package main

import (
	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7"
)

func main() {
	c := hl7.Component{
		Data: []hl7.SubComponent{"sss"},
	}

	println(c.GetSubComponent(4))

	s := hl7.SegmentFromComponentString(
		"MSH",
		[
			// ["1"]
			// [],
			// [
			// 	[ "111^aa^aa" ],
			// 	[ "111^aa^aa" ],
			// ]
		]
	)

}
