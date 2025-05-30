package main

import "fmt"

type camera struct {
	brand string
}

func (c *camera) trigger() string { //camera implements shooter interfaces becuase of trigger
	return fmt.Sprintf("has taken a photo with a %s\n", c.brand)
}
