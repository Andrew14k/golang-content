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

	canon := &camera{brand: "Canon"}
	h.setShooter(canon)
	h.shoot()

	samsung := &laser{
		brand: "Samsung Laser",
		sound: "pew! pew!",
	}
	h.setShooter(samsung)
	h.shoot()

	pistol := &laser{
		brand: "Mattel water pistol",
		sound: "shoosh! splash!",
	}
	h.setShooter(pistol)
	h.shoot()
}

// interfaces and runtime polymorphism come in to make changes without modifying existsing code
// go is a compositional language, so it favours composition over inheritance

// used for setting up MySql database, and/or swapping that out for a different database
// interfaces can be used for logging to screen or file (and swap out for different circumstances)
