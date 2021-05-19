package utils

func PasswordFilter(str []byte) bool {
	rst := true
	for i := range str {
		if str[i] >= '0' && str[i] <= '9' ||
			str[i] >= 'a' && str[i] <= 'z' ||
			str[i] >= 'A' && str[i] <= 'Z' ||
			str[i] == '_' ||
			str[i] == 0x00 {
		} else {
			rst = false
			break
		}
	}
	return rst
}
