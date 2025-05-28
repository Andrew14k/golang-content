package main

import (
	"errors"
	"math"
)

func main() {
	// no implicit casting in Go, so this cannot work until var types are matched
	// var iNum int32 = 2
	// var fNum float64
	// fNum = iNum

	// var iSmall int8 = 127
	// //iSmall++ //causes overflow error so would retrun -128
	// fmt.Println(iSmall)

	// var uiSmall uint8 = 255
	// uiSmall += 5 //overlows so would return 4
	// fmt.Println(uiSmall)
}

// code that ensures whatever we are doing is "safe"
func safeAdd(a, b int64) (int64, error) {
	if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
		return 0, errors.New("overflow")
	}
	return a + b, nil
}
