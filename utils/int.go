package utils

func ConvertUint32sToInt32s(src []uint32) []int32 {
	int32s := make([]int32, len(src))
	for i := range src {
		int32s[i] = int32(src[i])
	}
	return int32s
}

func Convertint32sToUint32s(src []int32) []uint32 {
	uint32s := make([]uint32, len(src))
	for i := range src {
		uint32s[i] = uint32(src[i])
	}
	return uint32s
}

func ConvertUint16sToInt16s(src []uint16) []int16 {
	int16s := make([]int16, len(src))
	for i := range src {
		int16s[i] = int16(src[i])
	}
	return int16s
}

func Convertint16sToUint16s(src []int16) []uint16 {
	uint16s := make([]uint16, len(src))
	for i := range src {
		uint16s[i] = uint16(src[i])
	}
	return uint16s
}

func Convert9Uint16sToInt16s(src [9]uint16) []int16 {
	int16s := make([]int16, 9)
	for i := range src {
		int16s[i] = int16(src[i])
	}
	return int16s
}

func Convertint16sTo9Uint16s(src []int16) [9]uint16 {
	uint16s := [9]uint16{}
	for i := range src {
		uint16s[i] = uint16(src[i])
	}
	return uint16s
}
