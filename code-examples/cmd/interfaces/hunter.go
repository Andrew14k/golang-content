package main

import "fmt"

type hunter struct {
	person
	shooter // shooter type as anonymous - anything that implemented the trigger method
}

func (h *hunter) shoot() {
	fmt.Printf("Hunter %s %s", h.firstName, h.trigger())
}

func (h *hunter) setShooter(s shooter) {
	h.shooter = s
}
