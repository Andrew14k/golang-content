package main

func main() {
	h := hunter{
		person: person{
			firstName: "Andrew",
			lastName:  "Smith",
			age:       23,
		},
		shooter: &camera{
			brand: "Minolta",
		}, //reference to camera
	}

	h.shoot()
}

// interfaces and runtime polymorphism come in to make changes to code by adding bits
