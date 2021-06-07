package utils

import "math"

// BigByteToInt convert
func BigByteToInt(data []byte) int64 {
	ln := len(data)
	rt := int64(0)
	for i := 0; i < ln; i++ {
		rt |= int64(data[ln-1-i]) << (i * 8)
	}
	return rt
}

// BigIntToByte convert
func BigIntToByte(data int64, ln int) []byte {
	rt := make([]byte, ln)
	for i := 0; i < ln; i++ {
		rt[ln-1-i] = byte(data >> (i * 8))
	}
	return rt
}

// LitByteToInt convert
func LitByteToInt(data []byte) int64 {
	ln := len(data)
	rt := int64(0)
	for i := 0; i < ln; i++ {
		rt |= int64(data[i]) << (i * 8)
	}
	return rt
}

// LitIntToByte convert
func LitIntToByte(data int64, ln int) []byte {
	rt := make([]byte, ln)
	for i := 0; i < ln; i++ {
		rt[i] = byte(data >> (i * 8))
	}
	return rt
}

// BigByteToFloat32 convert
func BigByteToFloat32(data []byte) float32 {
	v := BigByteToInt(data)
	return math.Float32frombits(uint32(v))
}

// BigFloatToByte32 convert
func BigFloatToByte32(data float32) []byte {
	v := math.Float32bits(data)
	return BigIntToByte(int64(v), 4)
}

// BigByteToFloat64 convert
func BigByteToFloat64(data []byte) float64 {
	v := BigByteToInt(data)
	return math.Float64frombits(uint64(v))
}

// BigFloatToByte64 convert
func BigFloatToByte64(data float64) []byte {
	v := math.Float64bits(data)
	return BigIntToByte(int64(v), 8)
}

// LitByteToFloat32 convert
func LitByteToFloat32(data []byte) float32 {
	v := LitByteToInt(data)
	return math.Float32frombits(uint32(v))
}

// LitFloatToByte32 convert
func LitFloatToByte32(data float32) []byte {
	v := math.Float32bits(data)
	return LitIntToByte(int64(v), 4)
}

// LitByteToFloat64 convert
func LitByteToFloat64(data []byte) float64 {
	v := LitByteToInt(data)
	return math.Float64frombits(uint64(v))
}

// LitFloatToByte64 convert
func LitFloatToByte64(data float64) []byte {
	v := math.Float64bits(data)
	return LitIntToByte(int64(v), 8)
}
