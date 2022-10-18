package main

import (
	"fmt"
	"math"
)

const (
	maxInt8   = math.MaxInt8
	maxInt16  = math.MaxInt16
	maxInt32  = math.MaxInt32
	maxUint8  = math.MaxUint8
	maxUInt16 = math.MaxUint16
	maxUInt32 = math.MaxUint32
	minInt8   = math.MinInt8
	minInt16  = math.MinInt16
	minInt32  = math.MinInt32
	minUint   = 0
)

func main() {
	var a, b int16
	_, _ = fmt.Scan(&a, &b)
	a32 := int32(a)
	b32 := int32(b)

	switch multiply := a32 * b32; {
	case int32(multiply) < int32(minInt16):
		fmt.Println("int32")
	case int32(multiply) < int32(minInt8):
		fmt.Println("int16")
	case int32(multiply) < int32(minUint):
		fmt.Println("int18")
	case int32(multiply) == int32(0) && multiply < int32(maxUint8):
		fmt.Print("uint8")
	case int32(multiply) < int32(maxUInt16):
		fmt.Print("uint16")
	case int32(multiply) > int32(maxUInt16):
		fmt.Print("uint32")
	}

	// if multiply < int32(minInt16) {
	// 	fmt.Println("Сюда норм!")
	// } else if
}
