package main

import "fmt"

type laser struct {
	brand string
	sound string
}

func (l *laser) trigger() string {
	return fmt.Sprintf("has taken a shot with their %s - %s\n", l.brand, l.sound)
}
