// ● Create a new type: vehicle.
// 	 ○ The underlying type is a struct.
// 	 ○ The fields:
// 			■ doors
// 			■ color
// ● Create two new types: truck & sedan.
// 		○ The underlying type of each of these new types is a struct.
// 		○ Embed the “vehicle” type in both truck & sedan.
//		○ Give truck the field “fourWheel” which will be set to bool.
// 		○ Give sedan the field “luxury” which will be set to bool. solution
// ● Using the vehicle, truck, and sedan structs:
// 		○ using a composite literal, create a value of type truck and assign values to the
// 			fields;
// 		○ using a composite literal, create a value of type sedan and assign values to the
// 			fields.
// ● Print out each of these values.
// ● Print out a single field from each of these values.

package main

import "fmt"

type vehicule struct {
	doors int
	color string
}

type truck struct {
	vehicule
	fourWheel bool
}

type sedan struct {
	vehicule
	luxury bool
}

func main() {
	// truck
	t := truck{
		vehicule: vehicule{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	fmt.Println(t)

	// sedan
	s := sedan{
		vehicule: vehicule{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(s)

	fmt.Println(t.doors)
	fmt.Println(s.doors)
}
