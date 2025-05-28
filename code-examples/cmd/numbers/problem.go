package main

import (
	"fmt"
	"math"
)

func main() {
	var roomW, roomH, roomL, paintC float32 = 40.0, 3.4, 30.0, 5.1
	var paintD, paintH float32 = 0.15, 0.3
	var boxL, boxH, boxW float32 = 0.6, 0.3, 0.35

	cansRequired := paintWall(roomW, roomH, roomL, paintC)
	fmt.Printf("Number of cans required: %d\n", uint16(math.Ceil(float64(cansRequired))))

	cansInBox := cansInBoxes(paintD, paintH, boxL, boxH, boxW)
	fmt.Printf("Number of cans per box: %d\n", uint16(math.Floor(float64(cansInBox))))

	boxes := fullBoxes(paintD, paintH, boxL, boxH, boxW)
	fmt.Printf("Number of full boxes: %d\n", uint16(math.Floor(float64(boxes)))) //float64 is the type returned by ceil and floor functions, default for floating-point funcs

	cansUnpacked := uint16(math.Ceil(float64(cansRequired))) % uint16(math.Floor(float64(boxes)))
	fmt.Printf("Cans not packed in boxes %d\n", cansUnpacked)
}

func paintWall(roomW, roomH, roomL, paintC float32) float32 {
	roomWalls := (2 * roomW * roomH) + (2 * roomL * roomH)
	return roomWalls / paintC
}

func cansInBoxes(paintD, paintH, boxL, boxH, boxW float32) float32 {
	alongL := uint16(boxL / paintD)
	alongW := uint16(boxW / paintD)
	alongH := uint16(boxH / paintH)
	return float32(alongL * alongW * alongH)
}

func fullBoxes(paintD, paintH, boxL, boxH, boxW float32) float32 {
	radius := paintD / 2
	paintCanVolume := float32(math.Pi) * paintH * radius * radius
	boxArea := boxL * boxH * boxW
	return boxArea / paintCanVolume
}
